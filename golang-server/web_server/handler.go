package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde el handler")
}
