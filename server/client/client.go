package client

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jbrre/workshop-docker/client/models"
)

var db *sql.DB

type RequestStatus struct {
	StatusCode int

	Err error
}

func InitDb() error {
	log.Println("Connecting to db...")
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBADRESS"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Printf("SQL database open error, %v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("SQL database ping error, %v", err)
		return err
	}
	log.Println("Connected!")
	return nil
}

func GetUserList() ([]models.User, error) {
	userList := make([]models.User, 0)
	rows, err := db.Query("SELECT username, email_adress FROM users")
	if err != nil {
		log.Printf("SQL database query error, %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var curService models.User
		err := rows.Scan(&curService.Username, &curService.EmailAdress)
		if err != nil {
			log.Printf("SQL database reading error, %s", err.Error())
			return nil, err
		}
		userList = append(userList, curService)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("SQL rows error, %s", err.Error())
		return nil, err
	}
	return userList, nil
}
