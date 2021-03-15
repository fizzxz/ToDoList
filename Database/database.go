package Database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func OpenMySQLDB(dbUser, dbPass, ipAddr, portNo, dbName string) (*sql.DB, error) {
	return sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+ipAddr+":"+portNo+")/"+dbName)
}

func GetDatabaseCredentials(jsonCredFile string) (string, string) {
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
