package stockentry

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AddstockReqStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	CurrentUser  string `json:"currentUser"`
}

type AddstockRespStruct struct {
	MedProductsArr []AddstockReqStruct `json:"medProductsArr"`
	Status         string              `json:"status"`
	ErrMsg         string              `json:"errMsg"`
}

func AddStockApi(w http.ResponseWriter, r *http.Request) {
	log.Println("AddStockApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {

		var lAddStockReq AddstockReqStruct
		var lAddStockResp AddstockRespStruct

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in ADDSTOCKAPI request :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lAddStockReq)
			if err != nil {
				log.Println("Addstock_Unmarshal :", err)
				fmt.Println(err.Error())
			} else {
				// check user exist before adding new user
				lNewStock := CheckStockExist(lAddStockReq.MedicineName, lAddStockReq.Brand)
				if lNewStock {
					if AddStock(lAddStockReq) == 1 {
						lAddStockResp.Status = "S"
						lAddStockResp.ErrMsg = ""
					} else {
						lAddStockResp.Status = "E"
						lAddStockResp.ErrMsg = "error"
					}
				} else {
					lAddStockResp.Status = "Medicine Already Exist"
					lAddStockResp.ErrMsg = "error"
				}
			}
		}

		// marshal the Response struct
		data, err := json.Marshal(lAddStockResp)
		if err != nil {
			log.Println("Error in AddStockApi Marshal :", err)
			fmt.Fprintf(w, "Error in AddStockApi Marshal "+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("AddStockApi (-)")
	}
}

// 1st method for checking stock is exist or not in medaapp_medicine_master table
func CheckStockExist(MedicineName string, Brand string) (newStock bool) {
	log.Println("CHeck stock exist (+)")
	var lNewStock bool
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		//QUERY FOR RESPECTIVE MedName & brand IS ALREADY EXIST CHECKING
		queryString := `select medicine_master_id,medicine_name,brand
						from medapp_medicine_master 
						where medicine_name=? and brand=?`

		sqlResult, err := db.Query(queryString, MedicineName, Brand)
		if err != nil {
			log.Println("Error in CheckStockExist() :", err)
			fmt.Println(err.Error())
		} else {
			if sqlResult.Next() {
				lNewStock = false
			} else {
				lNewStock = true
			}
		}
	}
	log.Println("CHeck stock exist (-)")
	return lNewStock
}

// 2nd method for Adding stock in medaapp_medicine_master table
func AddStock(stock AddstockReqStruct) int {
	log.Println("Add stock (+)")
	var lRecordsInserted int

	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := ` 
		insert into medapp_medicine_master(medicine_name, brand,created_by,created_date)
		values (?,?,?,NOW())`

		sqlResult, err := db.Exec(queryString, stock.MedicineName, stock.Brand, stock.CurrentUser)
		if err != nil {
			log.Println("Error in AddStockMedmaster() :", err)
			fmt.Println(err.Error())
		} else {
			num, _ := sqlResult.RowsAffected()
			Temp := int(num)
			if Temp == 1 {
				//3rd  method calling
				lRecordsInserted = InsertingStock(stock)
			} else {
				lRecordsInserted = 0
				log.Println("Error in InsertingStock Error :", err)
			}
		}
	}
	log.Println("Add stock (-)")
	return lRecordsInserted
}

// 3rd added Stock is succesful means im inserting the id & user details in stock table
func InsertingStock(stock AddstockReqStruct) int {

	log.Println("Inserting stock (+)")
	var lRecordsInserted int
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := ` insert into medapp_stock (medicine_master_id ,quantity,
							unit_price,created_by,created_date,updated_by,updated_date )
							Values ((select max(medicine_master_id)
							from medapp_medicine_master mmm),0,0,?,NOW(),?,NOW())`

		sqlResult, err := db.Exec(queryString, stock.CurrentUser, stock.CurrentUser)
		if err != nil {
			log.Println("Error in InsertingStocktable() :", err)
			fmt.Println(err.Error())
		} else {
			num, _ := sqlResult.RowsAffected()
			lRecordsInserted = int(num)
		}
	}
	log.Println("Inserting stock (-)")
	return lRecordsInserted
}
