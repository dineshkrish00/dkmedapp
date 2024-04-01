package stockview

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CurrentStockStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	Quantity     int    `json:"quantity"`
	UnitPrice    int    `json:"unitprice"`
}

type CurrentStockRespStruct struct {
	StockViewArr []CurrentStockStruct `json:"stockViewArr"`
	Status       string               `json:"status"`
	ErrMsg       string               `json:"errMsg"`
}

func CurrentStockApi(w http.ResponseWriter, r *http.Request) {
	log.Println("CurrentStockApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "GET" {
		var lCurrentStockReq CurrentStockStruct
		var lCurrentStockResp CurrentStockRespStruct
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in UPDATESTOCKAPI request :", err)
			fmt.Println(err.Error())
		} else {
			lCurrentStockResp = FetchStock(lCurrentStockReq, lCurrentStockResp)
		}

		// marshal the Response struct
		data, err := json.Marshal(lCurrentStockResp)
		if err != nil {
			log.Println("Error in UpdateStockApi Marshal :", err)
			fmt.Fprintf(w, "Error in UpdateStockApi Marshal"+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("CurrentStockApi (-)")
	}
}

// method to update stock in the stock table
func FetchStock(lCurrentStockReq CurrentStockStruct, lFetchArr CurrentStockRespStruct) CurrentStockRespStruct {
	log.Println("Current stock (+)")
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
		select m.medicine_name, m.brand, s.quantity,s.unit_price
		from medapp_medicine_master m, medapp_stock s
		where m.medicine_master_id =s.medicine_master_id;`
		sqlResult, err := db.Query(queryString)
		if err != nil {
			log.Println("Error in Fetchstock() :", err)
			fmt.Println(err.Error())
		} else {
			for sqlResult.Next() {
				// query returns user details
				err := sqlResult.Scan(&lCurrentStockReq.MedicineName, &lCurrentStockReq.Brand,
					&lCurrentStockReq.Quantity, &lCurrentStockReq.UnitPrice)
				if err != nil {
					lFetchArr.Status = "E"
					lFetchArr.ErrMsg = err.Error()
					log.Println("Error in Scanning StockView_Query :", err.Error())
				} else {
					lFetchArr.StockViewArr = append(lFetchArr.StockViewArr, lCurrentStockReq)
					lFetchArr.Status = "S"
					lFetchArr.ErrMsg = ""
				}
			}
		}
	}
	log.Println("Current stock (-)")
	return lFetchArr
}
