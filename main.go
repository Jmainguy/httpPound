package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GET(URL string, i int) {
	var netClient = &http.Client{
		Timeout: time.Second * 120,
	}
	response, err := netClient.Get(URL)
	if err != nil {
		fmt.Println(err)
	} else {
        	fmt.Println("Response status:", response.Status)
                response.Body.Close()
        }
	fmt.Println(i)
}

func main() {

	URL := os.Getenv("URL")
	COUNTSTRING := os.Getenv("COUNT")

	COUNT, _ := strconv.Atoi(COUNTSTRING)

	fmt.Println(COUNT)
	fmt.Println(URL)

	i := 0

	for i < COUNT {
		go GET(URL, i)
		i++
	}

	time.Sleep(120 * time.Second)

	os.Exit(0)
}
