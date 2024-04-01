package adduser

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GAddUserStruct struct {
	LoginId     int    `json:"loginid"`
	UserId      string `json:"userid"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	CurrentUser string `json:"currentUser"`
}

type AddUserRespStruct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func AddUserApi(w http.ResponseWriter, r *http.Request) {
	log.Println("AddUserApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lAddUserReq GAddUserStruct
		var lAddUserResp AddUserRespStruct

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error in ADDUSERAPI request :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lAddUserReq)
			if err != nil {
				log.Println("AddUser_Unmarshal :", err)
				fmt.Println(err.Error())
			} else {
				// check user exist before adding new user
				lNewUser := CheckUserExist(lAddUserReq.UserId)
				if lNewUser {
					if AddNewUser(lAddUserReq) == 1 {
						lAddUserResp.Status = "S"
						lAddUserResp.ErrMsg = ""
					} else {
						lAddUserResp.Status = "E"
						lAddUserResp.ErrMsg = "error"
					}
				} else {
					lAddUserResp.Status = "User Already Exist"
					lAddUserResp.ErrMsg = "error"
				}
			}
		}

		// marshal the Response struct
		data, err := json.Marshal(lAddUserResp)
		if err != nil {
			log.Println("Error in AddUserApi Marshal :", err)
			fmt.Fprintf(w, "Error in AddUserApi Marshal "+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("AddUserApi (-)")
	}
}

func CheckUserExist(userId string) (newUser bool) {

	var lNewUser bool
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		//QUERY FOR RESPECTIVE USER ID PERSON IS ALREADY EXIST CHECKING
		queryString := ` select * from medapp_login where user_id = ? `

		sqlResult, err := db.Query(queryString, userId)
		if err != nil {
			log.Println("Error in CheckUsers()exist :", err)
			fmt.Println(err.Error())
		} else {
			if sqlResult.Next() {
				lNewUser = false
			} else {
				lNewUser = true
			}
		}
	}
	return lNewUser
}

func AddNewUser(user GAddUserStruct) int {
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
		insert into medapp_login(user_id, password, role, created_by, created_date, updated_by, updated_date)
		values( ?, ?, ?, ?, now(), ?, now() )
		`
		sqlResult, err := db.Exec(queryString, user.UserId, user.Password, user.Role, user.CurrentUser, user.CurrentUser)
		if err != nil {
			log.Println("Error in AddNewUser() :", err)
			fmt.Println(err.Error())
		} else {
			num, _ := sqlResult.RowsAffected()
			lRecordsInserted = int(num)
		}
	}
	return lRecordsInserted
}
