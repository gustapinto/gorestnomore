package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ConnectPostgres() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	if db, err := sql.Open("pgx", connString); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return db
	}
}

func AddBook(book *Book) {
	db := ConnectPostgres()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	query := `insert into "books" ("title", "author") values ($1, $2)`

	if _, err := db.Exec(query, book.Title, book.Author); err != nil {
		log.Fatal(err)
	}
}

func GetAllBooks() []Book {
	var books []Book

	db := ConnectPostgres()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	rows, err := db.Query("select * from books order by id")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var title, author string

		if err := rows.Scan(&id, &title, &author); err != nil {
			log.Fatal(err)
		}

		books = append(books, Book{id, title, author})
	}

	return books
}

func GetBookById(bookId string) (*Book, error) {
	db := ConnectPostgres()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	row := db.QueryRow("select * from books where id=$1", bookId)

	book := &Book{}

	if err := row.Scan(&book.Id, &book.Title, &book.Author); err != nil {
		return nil, err
	} else {
		return book, nil
	}
}

func UpdateBook(bookId string, newBook *Book) error {
	db := ConnectPostgres()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	if _, err := db.Exec("update books set title=$1, author=$2 where id=$3",
		newBook.Title, newBook.Author, bookId); err != nil {
		return err
	}

	return nil
}

func DeleteBook(bookId string) error {
	db := ConnectPostgres()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	if _, err := db.Exec("delete from books where id=$1", bookId); err != nil {
		return err
	}

	return nil
}
