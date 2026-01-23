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

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/tools/go/analysis/passes/defers"
)

type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message. omitempyt"`
}

func createConnectionDb() *sql.DB {
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
	stocks, err:= getAllStocks()
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

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r http.Response){
	params:= mux.Vars(r)
	id, err:= strconv.ParseInt(params["id"])
	if err != nil {
		log.Fatal("Unable to convet the string into int %v",err)
	}

	deleteRows:= deleteStock(int64(id))

	msg:= fmt.Sprintf("Stock deleted successfully. Total rows/records %v", deleteRows)

	res:= response{
		ID: int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}


func insertStock(stock model.Stock) int64{
	db:= createConnectionDb()
	defer db.Close()

	sqlStatement:= `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int64
	err:= db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil{
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("Inserted a single query %v", id)
	return  id

}

func getStock (id int64) (model.Stock, error){
	db:= createConnectionDb()
	defer db.Close()

	var stock model.Stock
	sqlStatement:= `SELECT * FROM stocks WHERE stockid= $1`

	row:= db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Company, &stock.Price)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned ")
		return  stock, nil
	case nil :
		return stock, nil
	default:
		log.Fatalf("Unable to scan rows %v", err)

	}

	return stock, err


}





