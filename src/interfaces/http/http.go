package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting HTTP Application ...")
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Failed to load")
		os.Exit(1)
	}

	fmt.Println("Status", resp.Status)
	fmt.Println("Status Code", resp.StatusCode)
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
