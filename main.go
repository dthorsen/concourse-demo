package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(GetGithubZen())
}

func Multiply(a, b int) int {
	return a * b
}

func GetGithubZen() string {
	resp, err := http.Get("https://api.github.com/zen")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
