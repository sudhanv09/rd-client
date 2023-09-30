package main

import "fmt"

func main() {
	c := httpClient(getApiKey())
	t, err := c.rdGetUser()
	if err != nil {
		fmt.Println("user fetch failed")
	}
	fmt.Println(t)
}
