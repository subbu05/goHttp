package main


import (
	"net/http"
	"log"
	"os"
	"fmt"
)

func main() {
	
	if len(os.Args) > 1 {
	      port := ":"+os.Args[1]
	}else{
		port := ":9000"	
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Golang server Welcomes you!")
	})

	log.Fatal(http.ListenAndServe(port, nil))

}
