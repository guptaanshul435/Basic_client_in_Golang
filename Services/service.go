package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/anshulg/concurrentclient/models"
)

func AddUser(user *models.User, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return
	}

	// Read and process the response body

	fmt.Println("POST request successful!")

	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Process the responseData

}

func FetchUser(url string, wg *sync.WaitGroup) {
	wg.Done()

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	json.NewDecoder(resp.Body).Decode(&users)

	for _, user := range users {
		fmt.Println(user.Name)
	}

	time.Sleep(time.Second * 3)
}
