package login

import (
	db "api/db_connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UserDetailStruct struct {
	LoginId  int    `json:"loginId"`
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRespStruct struct {
	UserName string `json:"userName"`
	UserRole string `json:"userRole"`
	Status   string `json:"status"`
	ErrMsg   string `json:"errMsg"`
}

type LoginReqStruct struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

func GetuserActivityApi(w http.ResponseWriter, r *http.Request) {
	log.Println("GetuserActivityApi (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lloginReq LoginReqStruct
		var lloginResp LoginRespStruct

		// reading request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error at reading login username & pass :", err)
			fmt.Println(err.Error())
		} else {
			// Unmarshall request body
			err := json.Unmarshal(body, &lloginReq)
			if err != nil {
				log.Println("logindetails_Unmarshal :", err)
				fmt.Println(err.Error())
				lloginResp.Status = "E"
				lloginResp.ErrMsg = "Error:" + err.Error()
			} else {
				resp := GetUserLogin(lloginReq.UserId, lloginReq.Password)
				lloginResp = resp

			}
		}
		// marshal the Response struct
		data, err := json.Marshal(lloginResp)
		if err != nil {
			log.Println("Error in loginApi Marshal :", err)
			fmt.Fprintf(w, "Error in loginApi Marshal "+err.Error())
		} else {
			log.Println("Marshal data :", string(data))
			fmt.Fprint(w, string(data))
		}
		log.Println("loginApi (-)")
	}
	log.Println("GetuserActivityApi(-)")
}

// fetching 1st method
func GetUserLogin(userId, password string) (lloginResp LoginRespStruct) {
	var lrespArr LoginRespStruct
	var lUserDetails []UserDetailStruct
	// OPEN DB connection
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		//fetching the datas from the login table
		queryString := `select user_id, password, role from medapp_login`
		sqlResult, err := db.Query(queryString)
		if err != nil {
			log.Println("Error in GetUsers() :", err)
			fmt.Println(err.Error())
		} else {

			// query returns user details
			for sqlResult.Next() {
				var lUserDetail UserDetailStruct
				err := sqlResult.Scan(&lUserDetail.UserId, &lUserDetail.Password, &lUserDetail.Role)
				if err != nil {
					log.Println("Error in Scanning User Details :", err.Error())
				} else {
					lUserDetails = append(lUserDetails, lUserDetail)
					lrespArr = CheckUserLogin(userId, password, lUserDetails)
				}
			}
			if lrespArr.Status == "S" {
				LoginHistory(userId)
			} else {
				log.Println("LoginHisotry method failed to execute")
			}
		}
	}
	return lrespArr
}

// 2nd method
func CheckUserLogin(userId, password string, lUserDetails []UserDetailStruct) LoginRespStruct {
	var lloginResp LoginRespStruct
	for _, v := range lUserDetails {
		if userId == v.UserId && password == v.Password {
			lloginResp.Status = "S"
			lloginResp.ErrMsg = ""
			lloginResp.UserName = v.UserId
			lloginResp.UserRole = v.Role
			break
		} else {
			lloginResp.Status = "E"
			lloginResp.ErrMsg = "Invalid UserId or Password"
		}
	}
	return lloginResp
}

// 3rd method
// login history inserting
func LoginHistory(userId string) {
	db, err := db.LocalDbConnect()
	if err != nil {
		log.Println("Error in Eastablishing DB-CONNECTION :", err)
		fmt.Println(err.Error())
	} else {
		// CLOSE DB connection after executing QUERY
		defer db.Close()
		logHistory := `
			insert into medapp_login_history
					(login_id,login_date,login_time,created_by,
					created_date,updated_by,updated_date)
			values((select login_id from medapp_login where user_id=?),CurDate(), CURTIME(),?,CurDate(),?,CurDate())`
		_, err := db.Exec(logHistory, userId, userId, userId)
		if err != nil {
			log.Println("Executing the loghistory :", err)
			fmt.Println(err.Error())
		} else {
			log.Println("inserted succesfully in login history ")
		}
	}
}
