package main

import (
	"api/adduser"
	"api/billentry"
	"api/dashboard"
	login "api/login"
	"api/loginhistory"
	"api/logout"
	"api/salesreport"
	"api/stockentry"
	"api/stockview"
	"log"
	"net/http"
)

func main() {

	log.Println("Server Started (+)")

	http.HandleFunc("/logdownload", download.LogFileDownload)

	http.HandleFunc("/getLogin", login.GetuserActivityApi)

	http.HandleFunc("/getLogout", logout.LogoutApi)

	http.HandleFunc("/getLoginHistory", loginhistory.LoginHistoryApi)

	http.HandleFunc("/addUser", adduser.AddUserApi)

	http.HandleFunc("/addStock", stockentry.AddStockApi)
	http.HandleFunc("/updateStock", stockentry.UpdateStockApi)
	http.HandleFunc("/fetchStock", stockentry.FetchMedApi)

	http.HandleFunc("/viewStock", stockview.CurrentStockApi)

	http.HandleFunc("/fetchmedData", billentry.FetchMednameApi)
	http.HandleFunc("/addBillData", billentry.AddBillApi)
	http.HandleFunc("/saveBillData", billentry.SaveBillApi)

	http.HandleFunc("/salesReport", salesreport.SalesReportApi)

	http.HandleFunc("/salesApi", dashboard.SalesApi)

	http.ListenAndServe(":29091", nil)

	log.Println("Server Terminated (-)")

}
