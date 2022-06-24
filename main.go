package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.Info(r.URL.Path[1:])
		fmt.Fprintf(w, "resp: %q\n", r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
