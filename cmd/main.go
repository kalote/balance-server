package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := newRouter()
	fmt.Println("Server running on port 3000")
	panic(http.ListenAndServe(":3000", r))
}
