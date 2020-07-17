package config

import (
	"database/sql"
	"fmt"
	"log"

	//driver for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

//GetCustomConf GetCustomConf
func GetCustomConf(key, defVal string) string {
	viper.SetConfigName("configfile")
	viper.AddConfigPath("../master/config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", key, defVal)
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

//InitDB InitDB
func Connection() *sql.DB {
	dbUser := GetCustomConf("DB_USER", "root")
	dbPass := GetCustomConf("DB_PASSWORD", "password")
	dbHost := GetCustomConf("DB_HOST", "localhost")
	dbPort := GetCustomConf("DB_PORT", "3306")
	schemaName := GetCustomConf("DB_SCHEMA", "dbName")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, _ := sql.Open("mysql", dbPath)
	err := dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return dbConn
}
