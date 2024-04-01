package dashboard

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TodaySalesStruct struct {
	Ysales int `json:"ysales"`
	Tsales int `json:"tsales"`
}

type TodaySalesRespStruct struct {
	// DbSalesArr []TodaySalesStruct `json:"dbSalesArr"`
	Ysales   int    `json:"ysales"`
	Tsales   int    `json:"tsales"`
	CurValue int    `json:"curvalue"`
	Symbol   bool   `json:"symbol"`
	Status   string `json:"status"`
	ErrMsg   string `json:"errMsg"`
}

func SalesApi(w http.ResponseWriter, r *http.Request) {
	log.Println("SalesApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {

		// var lSaleData TodaySalesStruct
		var lSaleBillResp TodaySalesRespStruct
		lSaleBillResp = TodaySale(lSaleBillResp)
		lSaleBillResp = YesterdaySale(lSaleBillResp)
		lSaleBillResp = CurStockValue(lSaleBillResp)
		if lSaleBillResp.Status == "S" {
			if lSaleBillResp.Ysales > lSaleBillResp.Tsales {
				lSaleBillResp.Symbol = true
				log.Println("Dashboard value Fetched")
			} else if lSaleBillResp.Ysales < lSaleBillResp.Tsales {
				lSaleBillResp.Symbol = false
			} else {
				lSaleBillResp.Symbol = true
			}
		} else {
			log.Println("Dashboard value failed")
		}
		// marshal the Response struct
		data, err := json.Marshal(lSaleBillResp)
		if err != nil {
			log.Println("Error in AddBill Marshal :", err)
			fmt.Fprintf(w, "Error in AddBill Marshal "+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("SalesApi (-)")
	}
}
func CurStockValue(lSaleBillResp TodaySalesRespStruct) TodaySalesRespStruct {
	log.Println("CurStockValue (+)")
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		//QUERY FOR RESPECTIVE MEdicine qunatity is already EXIST CHECKING
		queryString := `select SUM(quantity) from medapp_stock`
		sqlResult, err := db.Query(queryString)
		if err != nil {
			lSaleBillResp.Status = "E"
			lSaleBillResp.ErrMsg = ""
			log.Println("Error in CheckSTOCK VALUE:", err)
			fmt.Println(err.Error())
		} else {
			for sqlResult.Next() {
				err := sqlResult.Scan(&lSaleBillResp.CurValue)
				if err != nil {
					lSaleBillResp.Status = "E"
					lSaleBillResp.ErrMsg = "Fetching place error"
				} else {
					lSaleBillResp.Status = "S"
					lSaleBillResp.ErrMsg = ""
				}
			}
		}
	}
	log.Println("CurStockValue (-)")
	return lSaleBillResp

}
func YesterdaySale(lSaleBillResp TodaySalesRespStruct) TodaySalesRespStruct {
	log.Println("YesterdaySale (+)")
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		//QUERY FOR RESPECTIVE MEdicine qunatity is already EXIST CHECKING
		queryString := `select SUM(bill_amount) from medapp_bill_master
						where bill_date=curDate()-1`

		sqlResult, err := db.Query(queryString)
		if err != nil {
			lSaleBillResp.Status = "E"
			lSaleBillResp.ErrMsg = "Quantity is not enough"
			log.Println("Error in YesterdaySale :", err)
			fmt.Println(err.Error())
		} else {
			for sqlResult.Next() {
				err := sqlResult.Scan(&lSaleBillResp.Ysales)
				if err != nil {
					lSaleBillResp.Status = "E"
					lSaleBillResp.ErrMsg = "Fetching place error"
				} else {
					lSaleBillResp.Status = "S"
					lSaleBillResp.ErrMsg = ""
				}
			}
		}
	}
	log.Println("YesterdaySale (-)")
	return lSaleBillResp
}

func TodaySale(lSaleBillResp TodaySalesRespStruct) TodaySalesRespStruct {
	log.Println("TodaySale (+)")
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Establishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		//QUERY FOR RESPECTIVE MEdicine qunatity is already EXIST CHECKING
		queryString := `select SUM(bill_amount) from medapp_bill_master
						where bill_date=curDate()`
		// `select SUM(bill_amount) from medapp_bill_master
		// 				where bill_date=curDate()-1`

		sqlResult, err := db.Query(queryString)
		if err != nil {
			lSaleBillResp.Status = "E"
			lSaleBillResp.ErrMsg = "Quantity is not enough"
			log.Println("Error in TodaySale:", err)
			fmt.Println(err.Error())
		} else {
			for sqlResult.Next() {
				err := sqlResult.Scan(&lSaleBillResp.Tsales)
				if err != nil {
					lSaleBillResp.Status = "E"
					lSaleBillResp.ErrMsg = "Fetching place error"
				} else {
					lSaleBillResp.Status = "S"
					lSaleBillResp.ErrMsg = ""
				}
			}
		}
	}
	log.Println("TodaySale (-)")
	return lSaleBillResp
}
