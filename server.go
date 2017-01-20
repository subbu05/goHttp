package main

import "fmt"
import (
	"net/http"
	"log"
)

func main() {




	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Golang server Welcomes you!")
	})

	log.Fatal(http.ListenAndServe(":9000", nil))

}
