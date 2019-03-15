package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	Credit9service "gitlab.com/c9/Credit9"
	mysqldbs "gitlab.com/c9/Mysqldb"
)

var (
	dbFile      = "hostdb"
	sqlFile     = "paybox.db"
	mode        = "dev"
	Version     = "undefined"
	BuildTime   = "undefined"
	GitHash     = "undefined"
	logFlag     = flag.String("l", "debug", "กำหนดระดับ log -> info, warn, error, fatal, panic")
	proFlag     = flag.Bool("p", false, "รันในโหมดโปรดักชั่น ใช้งานจริง ถ้าไม่ใส่โปรแกรมจะไม่เปิดอุปกรณ์รับเงิน")
	versionFlag = flag.Bool("v", false, "show version info")
)

var mysql_np *sqlx.DB
var mysql_dbc *sqlx.DB
var sql_dbc *sqlx.DB
var nebula_dbc *sqlx.DB
var (
	pgEnv     = "development" //default
	pgSSLMode = "disable"
	pgDbHost  = "192.168.0.163"
	pgDbUser  = "postgres"
	pgDbPass  = "postgres"
	pgDbName  = "backup"
	pgDbPort  = "5432"
)

func ConnectMySqlDB(dbName string) (db *sqlx.DB, err error) {
	fmt.Println("Connect MySql")
	//dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	dsn := "it:[ibdkifu@tcp(192.168.0.89:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sql error =", err)
		return nil, err
	}
	return db, err
}

func ConnectMysqlNP(dbName string) (db *sqlx.DB, err error) {
	fmt.Println("Connect MySql")
	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	//fmt.Println(dsn,"DBName =", dbName)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sql error =", err)
		return nil, err
	}
	return db, err
}

func ConnectSqlDB() (msdb *sqlx.DB, err error) {
	db_host := "192.168.0.7"
	db_name := "expertshop"
	db_user := "sa"
	db_pass := "[ibdkifu"
	port := "1433"
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql", dsn)
	if msdb.Ping() != nil {
		fmt.Println("Error ")
	}

	return msdb, nil
}

func ConnectNebula() (msdb *sqlx.DB, err error) {
	db_host := "192.168.0.7"
	db_name := "bcnp"
	db_user := "sa"
	db_pass := "[ibdkifu"
	port := "1433"
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql", dsn)
	if msdb.Ping() != nil {
		fmt.Println("Error ")
	}

	return msdb, nil
}
func init() {

	mysql_nopadol, err := ConnectMysqlNP("npdl")
	if err != nil {
		fmt.Println(err.Error())
	}
	mysql_np = mysql_nopadol

}

func main() {
	flag.Parse()
	log.Printf("#### Version: %s", Version)
	log.Printf("#### Build Time: %s", BuildTime)
	log.Printf("#### Git Hash: %s", GitHash)

	switch {
	case *versionFlag:
		log.Printf("App Version: %s", Version)
		log.Printf("Build Time: %s", BuildTime)
		log.Printf("Git Hash: %s", GitHash)
		return
	case *proFlag:
		log.Println("### APP Mode = Production ###")
		mode = "pro"
	}
	salesRepo := mysqldbs.NewCredit9Repository(mysql_np)
	salesService := Credit9service.New(salesRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheckHandler)
	mux.HandleFunc("/version", apiVersionHandler)

	fmt.Println("Waiting for Accept Connection : 9999")

	mux.Handle("/credit9/", http.StripPrefix("/credit9/v1", Credit9service.MakeHandler(salesService)))
	http.ListenAndServe(":9999", mux)
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Success bool `json:"api success"`
	}{true})
}

func apiVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	//t := time.Now()
	json.NewEncoder(w).Encode(struct {
		Version     string `json:"version"`
		Description string `json:"description"`
		Creator     string `json:"creator"`
		LastUpdate  string `json:"lastupdate"`
	}{
		"0.1.0 BETA",
		"ERP Cloud Client Service",
		"ERP dev team 2019",
		"2019-01-01",
	})
}
