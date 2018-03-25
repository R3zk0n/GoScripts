package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/fatih/color"
)

func main() {

	var website string
	fmt.Scan(&website)

	fmt.Println("Target Website:", website)

	req, err := http.NewRequest("GET", website, nil)
	response, err := http.DefaultClient.Do(req)

	//error
	if err != nil {
		fmt.Println("Invaild URL or missing HTTPS")
		os.Exit(0)
	}

	defer response.Body.Close() // Close body
	//body, err := ioutil.ReadAll(response.Body)
	// Header work, Filtering for strict transport etc etc
	if response.StatusCode == http.StatusOK {
		color.Yellow("Status: 200")
		Server := response.Header.Get("Server")
		XSS := response.Header.Get("X-Xss-Protection")
		Frame := response.Header.Get("X-Frame-Options")

		//fmt.Println(response)
		fmt.Println("Server type:", Server)
		if len(XSS) < 2 {
			fmt.Println(website, "Missing XSS-Protection URL")
		}
		if len(Frame) < 2 {
			fmt.Println(website, "Missing X-Frame-Options")
		}

	}

	//fmt.Println(string(body))
}


