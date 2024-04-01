package salesreport

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SaleReportReqStruct struct {
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	BillNo      int    `json:"billNo"`
	BillDate    string `json:"billDate"`
	MedicineId  string `json:"medicineid"`
	Quantity    int    `json:"quantity"`
	Amount      int    `json:"amount"`
	CurrentUser string `json:"currentUser"`
}

type SalesRespStruct struct {
	SalesArr []SaleReportReqStruct `json:"salesArr"`
	Status   string                `json:"status"`
	ErrMsg   string                `json:"errMsg"`
}

func SalesReportApi(w http.ResponseWriter, r *http.Request) {
	log.Println("SalesReportApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {

		var lSaleReportReq SaleReportReqStruct
		var lSaleReportResp SalesRespStruct

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in SalesReportApi request :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lSaleReportReq)
			if err != nil {
				log.Println("SaleReport_Unmarshal :", err)
				fmt.Println(err.Error())
			} else {
				log.Println("api1", lSaleReportReq.StartDate, lSaleReportReq.EndDate)
				lSaleReportResp = SaleReport(lSaleReportReq, lSaleReportResp)
			}
		}
		// marshal the Response struct
		data, err := json.Marshal(lSaleReportResp)
		if err != nil {
			log.Println("Error in SalesReportApi Marshal :", err)
			fmt.Fprintf(w, "Error in SalesReportApi Marshal"+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("SalesReportApi (-)")
	}
}

// method to update stock in the stock table
func SaleReport(stock SaleReportReqStruct, lSaleReportResp SalesRespStruct) SalesRespStruct {
	log.Println("SaleReport (+)")
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		log.Println("scccvcc", stock.StartDate, stock.EndDate)
		defer db.Close()
		queryString := `select bd.bill_no,bd.created_date,bd.medicine_master_id, bd.quantity, bd.amount
						from  medapp_bill_details bd
						where bd.created_date between '` + stock.StartDate + `' and  '` + stock.EndDate + `'`

		lResult, err := db.Query(queryString)
		if err != nil {
			log.Println("Error in SalesReport Fetching :", err)
			fmt.Println(err.Error())
		} else {

			for lResult.Next() {
				err := lResult.Scan(&stock.BillNo, &stock.BillDate, &stock.MedicineId, &stock.Quantity,
					&stock.Amount)
				if err != nil {
					log.Println("Error in SalesReport Fetching :", err)
					lSaleReportResp.Status = "E"
					lSaleReportResp.ErrMsg = ""
					fmt.Println(err.Error())
				} else {
					lSaleReportResp.Status = "S"
					lSaleReportResp.ErrMsg = ""
					log.Println("SUccess")
				}
				lSaleReportResp.SalesArr = append(lSaleReportResp.SalesArr, stock)
			}

			log.Println("ddd", lSaleReportResp.SalesArr)
			log.Println("SUccess")
		}
	}
	log.Println("SalesReportApi (-)")
	return lSaleReportResp
}

// for lResult.Next() {
// 	err := lResult.Scan(&stock.BillNo, &stock.BillDate, &stock.MedicineId, &stock.Quantity,
// 		&stock.Amount)
// 	if err != nil {
// 		log.Println("Error in SalesReport Fetching :", err)
// 		lSaleReportResp.Status = "E"
// 		lSaleReportResp.ErrMsg = ""
// 		fmt.Println(err.Error())
// 	} else {
// 		lSaleReportResp.Status = "S"
// 		lSaleReportResp.ErrMsg = ""
// 		log.Println("SUccess")
// 	}
// 	lSaleReportResp.SalesArr = append(lSaleReportResp.SalesArr, stock)
// }
