package billentry

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SaveBillReqStruct struct {
	MedicineName string  `json:"medicinename"`
	Quantity     int     `json:"quantity"`
	UnitPrice    int     `json:"unitprice"`
	BillNo       int     `json:"billNo"`
	BillDate     string  `json:"billDate"`
	BillAmount   float64 `json:"billAmount"`
	BillGst      float64 `json:"billGst"`
	NetPrice     float64 `json:"netPrice"`
	CurrentUser  string  `json:"currentUser"`
}
type SaveBillRespStruct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func SaveBillApi(w http.ResponseWriter, r *http.Request) {
	log.Println("SaveBill (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lSaveBillReq SaveBillReqStruct
		var lSaveBillResp SaveBillRespStruct
		var bill bool
		var quan bool

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in SaveBill request :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lSaveBillReq)
			if err != nil {
				log.Println("SaveBill_Unmarshal :", err)
				fmt.Println(err.Error())
			} else {
				bill = MasterBillEntry(lSaveBillReq, lSaveBillResp)
				if bill {
					quan = QuantityUpdate(lSaveBillReq, lSaveBillResp, quan)
					if quan {
						lSaveBillResp.Status = "S"
						lSaveBillResp.ErrMsg = ""
					} else {
						lSaveBillResp.Status = "E"
						lSaveBillResp.ErrMsg = "insertion failed"
					}
				} else {
					log.Println("Error the data's not inserted", err.Error())
				}
			}
		}
		// marshal the Response struct
		data, err := json.Marshal(lSaveBillResp)
		if err != nil {
			log.Println("Error in SaveBill Marshal :", err)
			fmt.Fprintf(w, "Error in SaveBill Marshal "+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("SaveBill (-)")
	}
}
func QuantityUpdate(lSaveBillReq SaveBillReqStruct, lSaveBillResp SaveBillRespStruct, quan bool) bool {
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := `Update medapp_stock
						set quantity=quantity-?
						where medicine_master_id=(select medicine_master_id
							from medapp_medicine_master
							where medicine_name=?)`
		_, err := db.Exec(queryString, lSaveBillReq.Quantity, lSaveBillReq.MedicineName)
		if err != nil {
			lSaveBillResp.Status = "E"
			lSaveBillResp.ErrMsg = "Error" + err.Error()
			log.Println("Error in Inserting Stock table :", err)
			fmt.Println(err.Error())
		} else {
			lSaveBillResp.Status = "S"
			lSaveBillResp.ErrMsg = ""
			quan = true
			log.Println("Inserted Succesfully on Stock table")
		}
	}
	return quan
}

func MasterBillEntry(lSaveBillReq SaveBillReqStruct, lSaveBillResp SaveBillRespStruct) bool {
	// OPEN DB connection
	var lMaster bool
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := ` insert into medapp_bill_master(bill_no, bill_date, bill_amount, bill_gst, 
			net_price, login_id, created_by, created_date,
						updated_by, updated_date)
						Values(?, ?, ?, ?, ?,
						(select login_id 
						from medapp_login
						where user_id= ?),?, NOW(), ?, NOW())`

		_, err := db.Exec(queryString, lSaveBillReq.BillNo, lSaveBillReq.BillDate, lSaveBillReq.BillAmount,
			lSaveBillReq.BillGst, lSaveBillReq.NetPrice, lSaveBillReq.CurrentUser, lSaveBillReq.CurrentUser, lSaveBillReq.CurrentUser)
		if err != nil {
			lSaveBillResp.Status = "E"
			lSaveBillResp.ErrMsg = "Error" + err.Error()
			log.Println("Error in Inserting medapp_bill_master table :", err)
			fmt.Println(err.Error())
		} else {
			lSaveBillResp.Status = "S"
			lSaveBillResp.ErrMsg = ""
			log.Println("Inserted Succesfully on medapp_bill_master")
			lMaster = DetailBillEntry(lSaveBillReq, lSaveBillResp)
		}
	}
	return lMaster
}

func DetailBillEntry(lSaveBillReq SaveBillReqStruct, lSaveBillResp SaveBillRespStruct) bool {
	var lDetail bool
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := `insert into medapp_bill_details(bill_no, 
			medicine_master_id, quantity,unit_price,amount,
			created_by, created_date,updated_by, updated_date)
			values(?,(select medicine_master_id
				from medapp_medicine_master
				where medicine_name=?),?,?,?,?,now(),?,now())`

		_, err := db.Exec(queryString, lSaveBillReq.BillNo, lSaveBillReq.MedicineName, lSaveBillReq.Quantity,
			lSaveBillReq.UnitPrice, lSaveBillReq.BillAmount, lSaveBillReq.CurrentUser, lSaveBillReq.CurrentUser)
		if err != nil {
			lSaveBillResp.Status = "E"
			lSaveBillResp.ErrMsg = "Error" + err.Error()
			log.Println("Error in Inserting medapp_bill_details table :", err)
			fmt.Println(err.Error())
		} else {
			lSaveBillResp.Status = "S"
			lSaveBillResp.ErrMsg = ""
			lDetail = true
			log.Println("Inserted Succesfully on medapp_bill_details")
			// lSaveBillResp.SaveBillArr = append(lSaveBillResp.SaveBillArr, lSaveBillReq)
		}
	}
	return lDetail
}
