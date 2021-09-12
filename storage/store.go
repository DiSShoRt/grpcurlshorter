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


//connStr := "user=anton password=123 dbname=postgres sslmode=disable"

func AddUrlToDb(s string) (string, error) {

	if s == "" {
		return "", fmt.Errorf("empty string in request")
	}

	if !IsUrl(s) {
		return s, fmt.Errorf("no valid url")
	}

	connStr := "user=anton password=123 dbname=postgres sslmode=disable"

	conf, err := pgx.ParseConnectionString(connStr)

	//connect to db
	db, err := pgx.Connect(conf)
	if err != nil {
		log.Println("cant connect to db %s", err)
	}
	// close connection
	defer db.Close()
	var shortUrl string
	shortUrl = GetHash()
	//fmt.Println(shortUrl)
	var surl string
	err = db.QueryRow("INSERT INTO url VALUES ( $1, $2) returning short_url", shortUrl, s).Scan(&surl)

	if err != nil {
		log.Println("cant connect to db %s", err)
	}
	return surl, nil

}

type UserModele struct {
	Surl string
	Lurl string
}

func GetUrlFromDb(s string) (string, error) {

	if s == "" {
		log.Println("empty string")
		return s, fmt.Errorf("empty string in request")
	}

	if len(s) != 10 {
		log.Println("no valid")
		return s, fmt.Errorf("no valid short url")
	}

	//create config string
	connStr := "user=anton password=123 dbname=postgres sslmode=disable"
	//create connect config
	conf, err := pgx.ParseConnectionString(connStr)

	//connect to db
	db, err := pgx.Connect(conf)
	if err != nil {
		log.Println("cant connect to db %s", err)
	}

	defer db.Close()
	//one row
	var url string
	//var rez grpcResp
	err = db.QueryRow("SELECT long_url FROM url WHERE short_url = $1", s).Scan(&url)

	//defer post.Close()
	//walk to posts

	//fmt.Println(url)
	if err != nil {
		log.Println(err)
		return url, fmt.Errorf("connect to db")
	}
	//fmt.Println(p)
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
		log.Println("url", err)
		return false
	}
	return true
}
