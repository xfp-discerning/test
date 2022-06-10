package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const baseURL = "https://jsonplaceholder.typicode.com/posts/"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(172.31.160.1:3306)/test")
	if err != nil {
		panic(err)
	}
}

type Result struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	m := 100
	in := make(chan string, m)
	out := make(chan *Result, m)
	for i := 1; i <= m; i++ {
		in <- fmt.Sprintf("%s%d", baseURL, i)
	}

	Fetch(in, out, 10)
	Save(out)
}

func Fetch(in <-chan string, out chan<- *Result, n int) {
	for i := 0; i < n; i++ {
		go func(i int) {
			for {
				url := <-in
				res, err := http.Get(url)
				if err != nil {
					log.Println(err)
					continue
				}
				if res.StatusCode != http.StatusOK {
					log.Println("status not ok")
					continue
				}
				var r Result
				err = json.NewDecoder(res.Body).Decode(&r)
				if err != nil {
					continue
				}
				out <- &r
			}
		}(i)
	}
}

func Save(ch <-chan *Result) {
	for r := range ch {
		// TODO: save  to database
		_, err := db.Exec("insert into post (id,user_id,title,body) values(?,?,?,?)", r.ID, r.UserID, r.Title, r.Body)
		if err != nil {
			log.Printf("error saving post id=%d, %v", r.ID, err)
			continue
		}
	}
}
