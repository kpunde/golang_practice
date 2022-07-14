package requestRest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

type PostRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

type PostResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func hitGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	handleError(err)

	body := res.Body
	bytes, err := ioutil.ReadAll(body)
	handleError(err)

	var people []Person
	err = json.Unmarshal(bytes, &people)
	handleError(err)

	for _, item := range people {
		fmt.Printf("%v	%v	%v\n", item.Id, item.Username, item.Email)
	}
}

func hitPostRequest() {
	post := PostRequest{
		Title:  "Test PostRequest",
		Body:   "This is Test PostRequest",
		UserId: 1,
	}

	jsonBody, err := json.Marshal(post)
	handleError(err)

	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonBody))
	handleError(err)

	body := res.Body
	bytes, err := ioutil.ReadAll(body)
	handleError(err)

	var postResponse PostResponse
	err = json.Unmarshal(bytes, &postResponse)
	handleError(err)

	fmt.Println(postResponse)
}
