package billentry

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FetchDataStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	CurrentUser  string `json:"currentUser"`
}

type FetchDataRespStruct struct {
	ProductsArr []FetchDataStruct `json:"productsArr"`
	Status      string            `json:"status"`
	BillNo      int               `json:"billNo"`
	ErrMsg      string            `json:"errMsg"`
}

func FetchMednameApi(w http.ResponseWriter, r *http.Request) {
	log.Println("FetchMednameApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "GET" {

		var lFetchStockReq FetchDataStruct
		var lFetchStockResp FetchDataRespStruct
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in UPDATESTOCKAPI request :", err)
			fmt.Println(err.Error())
		} else {
			lFetchStockResp = FetchStock(lFetchStockReq, lFetchStockResp)
			lFetchStockResp = BillNO(lFetchStockResp)
		}

		// marshal the Response struct
		data, err := json.Marshal(lFetchStockResp)
		if err != nil {
			log.Println("Error in UpdateStockApi Marshal :", err)
			fmt.Fprintf(w, "Error in UpdateStockApi Marshal"+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("FetchMednameApi (-)")
	}
}

func BillNO(lFetchStockResp FetchDataRespStruct) FetchDataRespStruct {
	log.Println("BillNO (+)")
	// var lRecordsInserted int

	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := `select max(bill_no)+1 from medapp_bill_master `

		sqlResult, err := db.Query(queryString)
		if err != nil {
			log.Println("Error in FetchStock :", err)
			fmt.Println(err.Error())
		} else {
			for sqlResult.Next() {
				// query returns user details
				err := sqlResult.Scan(&lFetchStockResp.BillNo)
				if err != nil {
					lFetchStockResp.Status = "E"
					lFetchStockResp.ErrMsg = err.Error()
					log.Println("Error in Scanning GetMedData_Query :", err.Error())
				} else {
					// lFetchStockResp.ProductsArr = append(lFetchStockResp.ProductsArr, lFetchStockReq)
					lFetchStockResp.Status = "S"
					lFetchStockResp.ErrMsg = ""
				}
			}
		}
	}
	log.Println("BillNO (-)")
	return lFetchStockResp
}

// method to update stock in the stock table
func FetchStock(lFetchStockReq FetchDataStruct, lFetchArr FetchDataRespStruct) FetchDataRespStruct {
	log.Println("Fetch Stock (+)")
	// var lRecordsInserted int

	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := ` 
			select medicine_name, brand 
			from medapp_medicine_master`

		sqlResult, err := db.Query(queryString)
		if err != nil {
			log.Println("Error in FetchStock :", err)
			fmt.Println(err.Error())
		} else {
			for sqlResult.Next() {
				// query returns user details
				err := sqlResult.Scan(&lFetchStockReq.MedicineName, &lFetchStockReq.Brand)
				if err != nil {
					lFetchArr.Status = "E"
					lFetchArr.ErrMsg = err.Error()
					log.Println("Error in Scanning GetMedData_Query :", err.Error())
				} else {
					lFetchArr.ProductsArr = append(lFetchArr.ProductsArr, lFetchStockReq)
					lFetchArr.Status = "S"
					lFetchArr.ErrMsg = ""
				}
			}
		}
	}
	log.Println("Fetch Stock (-)")
	return lFetchArr
}
