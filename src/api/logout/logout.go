package logout

import (
	dbconnect "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LogoutReqStruct struct {
	UserId string `json:"userId"`
}

type LogoutRespStruct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func LogoutApi(w http.ResponseWriter, r *http.Request) {

	log.Println("LogoutApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {

		var lReqStruct LogoutReqStruct
		var lRespStruct LogoutRespStruct

		// reading request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("LogoutApi_ioutil.ReadAll :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lReqStruct)
			if err != nil {
				log.Println("LogoutApi_Unmarshal :", err)
				fmt.Println(err.Error())
			} else {
				// code
				if UserLogout(lReqStruct.UserId) == 0 {
					lRespStruct.Status = "E"
					lRespStruct.ErrMsg = "error"
				} else {
					lRespStruct.Status = "S"
					lRespStruct.ErrMsg = ""
				}
			}
		}

		// Marshall the Response Struct
		data, err := json.Marshal(lRespStruct)
		if err != nil {
			log.Println("LogoutApi_Unmarshal :", err)
			fmt.Fprintf(w, "Error in LogoutApi "+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}

	}
	log.Println("LogoutApi (-)")
}

func UserLogout(userId string) int {

	var lRecordsAffected int
	// OPEN DB connection
	db, err := dbconnect.LocalDbConnect()
	if err != nil {
		log.Println("MethodName()_LocalDbConnect :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		queryString := ` 
		update medapp_login_history set logout_date=NOW(),logout_time=curtime()
		where login_history_id=(select max(login_history_id) from medapp_login_history where created_by = ?)
		`
		sqlResult, err := db.Exec(queryString, userId)
		if err != nil {
			log.Println("MethodName()_db.Query :", err)
			fmt.Println(err.Error())
		} else {
			num, _ := sqlResult.RowsAffected()
			lRecordsAffected = int(num)
		}
	}
	return lRecordsAffected
}
