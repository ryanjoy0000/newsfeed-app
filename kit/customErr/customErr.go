package customErr

import (
	"log"
	"net/http"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func HandleHttpErr(respW http.ResponseWriter, err error) {
	if err != nil {
		http.Error(respW, err.Error(), http.StatusInternalServerError)
		HandleErr(err)
		return
	}
}
