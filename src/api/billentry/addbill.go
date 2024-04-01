package billentry

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BillDetailReqStruct struct {
	MedicineName string `json:"medicinename"`
	Quantity     int    `json:"quantity"`
	CurrentUser  string `json:"currentUser"`
}
type BillFetchingStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	Quantity     int    `json:"quantity"`
	UnitPrice    int    `json:"unitprice"`
	MedId        string `json:"medId"`
}

type BillFetchRespStruct struct {
	AddBillArr []BillFetchingStruct `json:"addBillArr"`
	Status     string               `json:"status"`
	ErrMsg     string               `json:"errMsg"`
}

func AddBillApi(w http.ResponseWriter, r *http.Request) {
	log.Println("AddBill (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lAddBillReq BillDetailReqStruct
		var lBillData BillFetchingStruct
		var lAddBillResp BillFetchRespStruct

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in AddBill request :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lAddBillReq)
			if err != nil {
				log.Println("AddUser_Unmarshal :", err)
				fmt.Println(err.Error())
			} else {
				lAddBillResp = CheckMedExist(lAddBillReq.MedicineName, lAddBillReq.Quantity, lBillData, lAddBillResp)
				log.Println("dsdsd", lAddBillResp)
			}
		}
		// marshal the Response struct
		data, err := json.Marshal(lAddBillResp)
		if err != nil {
			log.Println("Error in AddBill Marshal :", err)
			fmt.Fprintf(w, "Error in AddBill Marshal "+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("AddBill (-)")
	}
}

func CheckMedExist(medname string, quantity int, lBillData BillFetchingStruct, lAddBillResp BillFetchRespStruct) BillFetchRespStruct {
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		//QUERY FOR RESPECTIVE MEdicine qunatity is already EXIST CHECKING
		queryString := ` select mm.medicine_name, mm.medicine_master_id, ms.quantity, ms.unit_price, mm.brand
							from medapp_stock ms, medapp_medicine_master mm
							where ms.medicine_master_id =mm.medicine_master_id
							and  mm.medicine_name =? and ms.quantity>=?`

		sqlResult, err := db.Query(queryString, medname, quantity)
		if err != nil {
			lAddBillResp.Status = "E"
			lAddBillResp.ErrMsg = "Quantity is not enough"
			log.Println("Error in CheckMedExist(Quantity) :", err)
			fmt.Println(err.Error())
		} else {
			for sqlResult.Next() {
				err := sqlResult.Scan(&lBillData.MedicineName, &lBillData.MedId, &lBillData.Quantity, &lBillData.UnitPrice, &lBillData.Brand)
				if err != nil {
					lAddBillResp.Status = "E"
					lAddBillResp.ErrMsg = "Fetching place error"
				} else {
					lAddBillResp.Status = "S"
					lAddBillResp.ErrMsg = ""
					lAddBillResp.AddBillArr = append(lAddBillResp.AddBillArr, lBillData)
				}
			}
		}
	}
	return lAddBillResp
}
