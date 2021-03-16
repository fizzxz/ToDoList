package Database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ConnectToDB(ipAddr, portNo, dbName string) (*sql.DB, error) {
	dbUser, dbPass := getDatabaseCredentials("credentials.json")
	db, err := openMySQLDB(dbUser, dbPass, ipAddr, portNo, dbName)
	if err != nil {
		fmt.Println("issa no openin'") // do something here
		log.Fatal(err)
	}
	return db, err
}

func openMySQLDB(dbUser, dbPass, ipAddr, portNo, dbName string) (*sql.DB, error) {
	return sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+ipAddr+":"+portNo+")/"+dbName)
}

func getDatabaseCredentials(jsonCredFile string) (string, string) {
	// Open our jsonFile
	jsonFile, err := os.Open(jsonCredFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened credentials.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	// initialize our Users array
	var credentials Credentials

	// unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &credentials)
	if err != nil {
		fmt.Println(err)
	}
	return credentials.Username, credentials.Password
}
