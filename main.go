package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for {
		msg := getAttack("http://test.sinmetal.jp")
		fmt.Print(msg)
	}
}

func getAttack(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("feiled err=%+v\n", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("feiled read from responce body. err=%+v\n", err)
	}
	defer resp.Body.Close()
	return fmt.Sprintf("Status : %s, %s\n", resp.Status, string(body))
}
