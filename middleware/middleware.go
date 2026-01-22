package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	model "go-postgrsql/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message. omitempyt"`
}

func CreateConnectionDb() *sql.DB {
	err:= godotenv.Load(".env")

	if err != nil{
		log.Fatal("Error failed to load env...")
	}

	db, err :=  sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err= db.Ping()

	if err != nil{
		panic(err)
	}

	fmt.Println("Successfully Connected to db - postgres")
	return  db

}

func CreateStock(w http.ResponseWriter, r *http.Request){
	var stock model.Stock

	err:= json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Unable ot decode the request body %v", err)
	}

	insertID:= insertStock(stock)
	res:= response{
		ID: insertID,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}


func GetStock(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil  {
		log.Fatal("Unable to convert the string into int. %v", err)
	}

	stock, err:= getStock(int64(id))

	if err != nil{
		log.Fatal("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request){
	stocks, err:= egtAllStocks()
	if err != nil{
		log.Fatal("Unable to get all the stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Response){
	params:= mux.Vars(r)

	id, err:= strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Unable to convert the string into int %v", err)
	}

	var stock model.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)


	if err != nil{
		log.Fatalf("Unable to deocde the request body %v", err)
	}

	updateRows := updateStock(int64(id), stock)

	msg:= fmt.Sprintf("Stock updated successsfully. Total rows/records affected %v", updateRows)
	res:= response{
		ID: int64(id),
		Message: msg,
	}

	json.NewDecoder(w).Decode(res)
}





