package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	c := http.Client{}
	req, error := http.NewRequest("GET", "http://google.com", nil)
	if error != nil {
		panic(error)
	}
	req.Header.Set("Accept", "application/json")
	resp, error := c.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	body, err  := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}