package loginhistory

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginUserStruct struct {
	UserId string
	Date   string
	Time   string
	Type   string
}

type LoginHistoryRespStruct struct {
	UserDetailArr []LoginUserStruct `json:"userDetailArr"`
	Status        string            `json:"status"`
	ErrMsg        string            `json:"errMsg"`
}

func LoginHistoryApi(w http.ResponseWriter, r *http.Request) {

	log.Println("LoginHistoryApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {

		var lLoginHistoryResp LoginHistoryRespStruct

		lLoginMethod := GetUserLogin()
		lLogoutMethod := GetUserLogout()
		lLoginHistoryResp.UserDetailArr = append(lLoginHistoryResp.UserDetailArr, lLoginMethod...)
		log.Println("login get informatin", lLoginHistoryResp.UserDetailArr)
		lLoginHistoryResp.UserDetailArr = append(lLoginHistoryResp.UserDetailArr, lLogoutMethod...)
		log.Println("logout get informatin", lLoginHistoryResp.UserDetailArr)

		lLoginHistoryResp.Status = "S"
		lLoginHistoryResp.ErrMsg = ""

		// marshal the Response struct
		data, err := json.Marshal(lLoginHistoryResp)
		if err != nil {
			log.Println("Error in GetuserActivityApi Marshal :", err)
			fmt.Fprintf(w, "Error in GetuserActivityApi Marshal "+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("LoginHistoryApi (-)")
	}
}

func GetUserLogin() (UserDetailArr []LoginUserStruct) {
	var lUserActivity LoginUserStruct
	var lActivityArr []LoginUserStruct
	lUserActivity.Type = "Login"
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()

		queryString := ` 
		select created_by , login_date , login_time 
		from medapp_login_history
		where login_date is not null and login_time is not null 
		`
		sqlResult, err := db.Query(queryString)
		if err != nil {
			log.Println("Error in GetUserslogin() :", err)
			fmt.Println(err.Error())
		} else {
			// query returns user details
			for sqlResult.Next() {
				err := sqlResult.Scan(&lUserActivity.UserId, &lUserActivity.Date, &lUserActivity.Time)
				if err != nil {
					log.Println("Error in Scanning GetLoginActivity_Query :", err.Error())
				} else {
					lActivityArr = append(lActivityArr, lUserActivity)
				}
			}
		}
	}
	return lActivityArr
}

func GetUserLogout() (UserDetailArr []LoginUserStruct) {

	var lUserActivity LoginUserStruct
	var lActivityArr []LoginUserStruct
	lUserActivity.Type = "Logout"

	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()

		queryString := ` 
		select created_by , logout_date , logout_time 
		from medapp_login_history
		where logout_date is not null and logout_time is not null
		`
		sqlResult, err := db.Query(queryString)
		if err != nil {
			log.Println("Error in GetUserslogout() :", err)
			fmt.Println(err.Error())
		} else {
			// query returns user details
			for sqlResult.Next() {
				err := sqlResult.Scan(&lUserActivity.UserId, &lUserActivity.Date, &lUserActivity.Time)
				if err != nil {
					log.Println("Error in Scanning GetLoginActivity_Query :", err.Error())
				} else {
					lActivityArr = append(lActivityArr, lUserActivity)
				}
			}
		}
	}
	return lActivityArr
}
