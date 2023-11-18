package main

import (
	"sync"

	services "github.com/anshulg/concurrentclient/Services"
	"github.com/anshulg/concurrentclient/models"
)

const url string = "http://localhost:4000/users"

var wg sync.WaitGroup

func main() {

	user := models.User{
		// UserId:  13,
		Name:    "Keshav",
		Address: "noida",
		GmailId: "guptaanshul435@gmail.com",
		PhoneNo: "9557876138",
	}

	wg.Add(2)
	// ch := make(chan string)
	// defer close(ch)
	go services.AddUser(&user, url, &wg)

	go services.FetchUser(url, &wg)

	// time.Sleep(time.Second * 3)
	wg.Wait()
}
