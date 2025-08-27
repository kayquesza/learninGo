package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	Id_user    uint32 `json:"id_user"`
	Name_user  string `json:"name_user"`
	Email_user string `json:"email_user"`
}

// CreateUser - Inserts a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Failed to read request body!"))
		return
	}

	// var <variable name> <struct type>
	var user user
	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("Error converting user to struct."))
		return
	}

	// Creating database connection
	db, error := database.Connect()
	if error != nil {
		w.Write([]byte(fmt.Sprintf("Error connecting to database: %v", error)))
		return
	}
	defer db.Close()

	// Prepare Statement - Creates an insert command to prevent SQL Injection attacks
	statement, error := db.Prepare("insert into users (name_user, email_user) values (?, ?)")
	if error != nil {
		w.Write([]byte("Error creating statement."))
		return
	}
	defer statement.Close()

	insert, error := statement.Exec(user.Name_user, user.Email_user)
	if error != nil {
		w.Write([]byte("Error executing statement."))
		return
	}

	idInsert, error := insert.LastInsertId()
	if error != nil {
		w.Write([]byte("Error getting inserted id."))
		return
	}

	// Status Codes

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted successfully! ID: %d", idInsert)))

}

// FoundUsers - Searches for all users saved in the database
func FoundUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to database!"))
		return
	}
	defer db.Close()

	// SQL: SELECT * FROM USERS
	rows, error := db.Query("select * from users")
	if error != nil {
		w.Write([]byte("Error searching for users!"))
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user

		if error := rows.Scan(&user.Id_user, &user.Name_user, &user.Email_user); error != nil {
			w.Write([]byte("Error scanning user"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(users); error != nil {
		w.Write([]byte("Error converting users to JSON."))
		return
	}
}

// FoundUser - Searches for a specific user saved in the database
func FoundUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	// Converts ID to integer
	Id_user, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error converting parameter to integer."))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to database."))
		return
	}

	rows, error := db.Query("select * from users where id_user = ?", Id_user)
	if error != nil {
		w.Write([]byte("Error searching for user."))
		return
	}

	var user user
	if rows.Next() {
		if error := rows.Scan(&user.Id_user, &user.Name_user, &user.Email_user); error != nil {
			w.Write([]byte("Error scanning user."))
			return
		}
	}

	if error := json.NewEncoder(w).Encode(user); error != nil {
		w.Write([]byte("Error converting user to JSON."))
		return
	}
}

// UpdateUser - Updates user data in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	Id_user, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error converting parameter to integer"))
		return
	}

	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Error reading request body!"))
		return
	}

	var user user
	if error := json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("Error converting user to struct"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to database."))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("update users set name_user = ?, email_user = ? where id_user = ?")
	if error != nil {
		w.Write([]byte("Error creating statement"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(user.Name_user, user.Email_user, Id_user); error != nil {
		w.Write([]byte("Error updating user."))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeleteUser - Removes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id_user, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error converting parameter to integer."))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to database."))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("delete from users where id_user = ?")
	if error != nil {
		w.Write([]byte("Error creating statement."))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(Id_user); error != nil {
		w.Write([]byte("Error deleting user."))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
