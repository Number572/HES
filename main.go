package main 

import (
	"io"
	"net/http"
	"os"
	gp "github.com/number571/gopeer"
)

func hello(w http.ResponseWriter, r *http.Request) {
	gp.GenerateKey(1024)
	io.WriteString(w, "hello, world!")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
