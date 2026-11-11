package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/flvsantos15/code-bank-golang/domain"
	"github.com/flvsantos15/code-bank-golang/infrastructure/repository"
	"github.com/flvsantos15/code-bank-golang/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Flavio"
	cc.ExpirationYear = 2026
	cc.ExpirationMonth = 4
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	repo.CreateCreditCard(*cc)
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		// log.Fatal("error connection to database")
	}
	return db
}
