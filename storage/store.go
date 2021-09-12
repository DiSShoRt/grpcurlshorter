package storage

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"math/rand"
	u "net/url"
	"strings"
	"time"
)

type UserModele struct {
	Surl string
	Lurl string
}

const connStr = "user=anton password=123 dbname=postgres sslmode=disable"

func AddUrlToDb(s string) (string, error) {
	err := Check(s)
	if err != nil {
		log.Fatalf("this url %s is in the database\n", s)
		return s, fmt.Errorf("this url %s is in the database\n", s)
		panic(err)
	}

	if s == "" {
		return "", fmt.Errorf("empty string in request\n")
	}

	if !IsUrl(s) {
		return s, fmt.Errorf("no valid url\n")
	}

	fmt.Println(err)

	conf, err := pgx.ParseConnectionString(connStr)

	//connect to db
	db, err := pgx.Connect(conf)
	if err != nil {
		log.Printf("cant connect to db %s", err)
	}
	// close connection
	defer db.Close()
	var shortUrl string
	shortUrl = GetHash()

	var surl string

	if err != nil {
		err = db.QueryRow("INSERT INTO url VALUES ( $1, $2) returning short_url", shortUrl, s).Scan(&surl)
	}
	if err != nil {
		log.Printf("cant connect to db %s", err)
	}
	return surl, nil

}

func GetUrlFromDb(s string) (string, error) {

	if s == "" {
		log.Println("empty string")
		return s, fmt.Errorf("empty string in request")
	}

	if len(s) != 10 {
		log.Printf("no valid url %s", s)
		return s, fmt.Errorf("no valid short url %s", s)
	}

	conf, err := pgx.ParseConnectionString(connStr)

	//connect to db
	db, err := pgx.Connect(conf)
	if err != nil {
		log.Println("cant connect to db %s", err)
	}

	defer db.Close()

	var url string

	err = db.QueryRow("SELECT long_url FROM url WHERE short_url = $1", s).Scan(&url)
	if err != nil {
		log.Println(err)
		return url, fmt.Errorf("connect to db")
	}

	return url, nil
}

func GetHash() string {
	rand.Seed(time.Now().Unix())
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_1234567890")

	var b strings.Builder
	for i := 0; i < 10; i++ {
		b.WriteRune(letter[rand.Intn(len(letter))])
	}
	return b.String()
}

func IsUrl(url string) bool {
	_, err := u.ParseRequestURI(url)

	if err != nil {
		log.Println("no valid url", err)
		return false
	}
	return true
}

func Check(s string) error {

	//create connect config
	conf, err := pgx.ParseConnectionString(connStr)

	//connect to db
	db, err := pgx.Connect(conf)
	if err != nil {
		log.Printf("cant connect to db %s\n", err)
	}

	defer db.Close()
	//one row

	rows, err := db.Query("SELECT * FROM url WHERE long_url = $1", s)

	for rows.Next() {
		p := UserModele{}
		err := rows.Scan(&p.Lurl, &p.Surl)
		fmt.Println(p)
		if s == p.Lurl || s == p.Surl {
			fmt.Println(err)
			fmt.Printf("this url %s is in database\n", s)
			return fmt.Errorf("this url %s is in database\n", s)

		}

	}

	return nil
}
