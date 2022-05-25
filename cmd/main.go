package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ryanjoy0000/newsfeed-app/kit/config"
	"github.com/ryanjoy0000/newsfeed-app/kit/customErr"
	"github.com/ryanjoy0000/newsfeed-app/platform/newsFeed"
	"github.com/ryanjoy0000/newsfeed-app/platform/newsFeed/models"
)

var feedSys *newsFeed.FeedSystem

const addr = "localhost:4400"

func main() {

	db := config.InitMySQLDB()
	feedSys = newsFeed.CreateNewFeedSys(db)

	chiMux := chi.NewRouter()

	// GET
	chiMux.Get("/newsfeed", listPage)

	// POST
	chiMux.Post("/newsfeed", newPage)
	fmt.Println("Starting server...")
	err := http.ListenAndServe(addr, chiMux)
	customErr.HandleErr(err)

	defer config.CloseDBConn(db)
}

func listPage(respW http.ResponseWriter, req *http.Request) {
	setHTMLContentType(respW)
	receivedList := feedSys.Get()
	fmt.Println("Received from DB: ", receivedList)
	json.NewEncoder(respW).Encode(receivedList)
}

func newPage(respW http.ResponseWriter, req *http.Request) {
	setHTMLContentType(respW)
	if req.Method == http.MethodPost {

		// fItem := models.FeedItem{
		// 	Content: "Some more test content",
		// }
		// feedSys.Add(fItem)

		var item models.FeedItem
		json.NewDecoder(req.Body).Decode(&item)
		feedSys.Add(item)
		respW.Write([]byte("Saved successfully"))
	}
}

func setHTMLContentType(respW http.ResponseWriter) {
	respW.Header().Set("Content-Type", "text/html; charset=utf-8")
}
