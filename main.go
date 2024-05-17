package main

import (
	"fmt"
	"github.com/Montheankul-K/bank/handler"
	"github.com/Montheankul-K/bank/repository"
	"github.com/Montheankul-K/bank/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"time"
)

// go run . run go if have one main.go
func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	_ = customerRepository
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	// [0-9]+ is regx if param not a number return 404

	log.Printf("Banking service started at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // if have env variable use env
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// APP_PORT in env is equal app.port

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// should config time zone to prevent problem when run in container
func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	// location refer to time zone in IANA time zone database
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	// add ?parseTime=true when want to use time.Time

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute) // set timeout
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
