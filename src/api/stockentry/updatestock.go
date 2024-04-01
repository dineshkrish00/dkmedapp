package stockentry

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type StockUpdateReqStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	Quantity     string `json:"quantity"`
	UnitPrice    string `json:"unitprice"`
	CurrentUser  string `json:"currentUser"`
}

type StockUpdateRespStruct struct {
	ProductsArr []StockUpdateReqStruct `json:"productsArr"`
	Status      string                 `json:"status"`
	ErrMsg      string                 `json:"errMsg"`
}

func UpdateStockApi(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateStockApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {

		var lUpdateStockReq StockUpdateReqStruct
		var lUpdateStockResp StockUpdateRespStruct

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in UPDATESTOCKAPI request :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lUpdateStockReq)
			if err != nil {
				log.Println("Updatestock_Unmarshal :", err)
				fmt.Println(err.Error())
			} else {
				numRowsAffected := UpdateStock(lUpdateStockReq)
				// check medicine exist before adding new user
				if numRowsAffected == 0 {
					lUpdateStockResp.Status = "E"
					lUpdateStockResp.ErrMsg = "Error in updating stock"
				} else {
					lUpdateStockResp.Status = "S"
					lUpdateStockResp.ErrMsg = ""
				}
			}
		}
		// marshal the Response struct
		data, err := json.Marshal(lUpdateStockResp)
		if err != nil {
			log.Println("Error in UpdateStockApi Marshal :", err)
			fmt.Fprintf(w, "Error in UpdateStockApi Marshal"+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("UpdateStockApi (-)")
	}
}

// method to update stock in the stock table
func UpdateStock(stock StockUpdateReqStruct) int {
	log.Println("UpdateStock (+)")
	var lRecordsInserted int
	lquantity, err := strconv.Atoi(stock.Quantity)
	if err != nil {
		log.Println("error handled failed", err)
	} else {
		log.Println("Succesfully changed quantity")
	}
	lunitpirce, err := strconv.Atoi(stock.UnitPrice)
	if err != nil {
		log.Println("error handled failed", err)
	} else {
		log.Println("Succesfully changed quantity")
	}
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := `update medapp_stock
						set quantity = quantity+?, unit_price = ?, updated_by= ?,updated_date= now() 
						where medicine_master_id=(select medicine_master_id
						from medapp_medicine_master
						where medicine_name=? and brand=?)`

		lResult, err := db.Exec(queryString, lquantity, lunitpirce, stock.CurrentUser, stock.MedicineName, stock.Brand)
		if err != nil {
			log.Println("Error in UpdatedStocktable() :", err)
			fmt.Println(err.Error())
			return 0
		} else {
			num, _ := lResult.RowsAffected()
			lRecordsInserted = int(num)
		}
	}
	log.Println("UpdateStock (-)")
	return lRecordsInserted
}
