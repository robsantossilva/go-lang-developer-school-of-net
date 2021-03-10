package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", body)
}
