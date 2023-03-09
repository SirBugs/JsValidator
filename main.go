package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func rzlt_file() { // This func is created for checking if the results file if exist or not

	if _, err := os.Stat("valid_js_files.txt"); err != nil {
		f, err := os.Create("valid_js_files.txt") // Creating the file if it's not existing
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		fmt.Printf("") // output_msg
	} else { // Already rzlts file is found, No need to create or replace it.
		fmt.Printf("") // output_msg
	}

}

func request(_url_ string) {
	client := http.Client{
		Timeout: 7 * time.Second,
	}
	req, err := http.NewRequest("GET", _url_, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0")
	resp, err := client.Do(req)
	if resp == nil {
		fmt.Printf("[ - ] :: %v\n", _url_)
	} else {
		if resp.StatusCode == 200 {
			fmt.Printf("[ %v ] :: %v\n", resp.StatusCode, _url_)
			f, err := os.OpenFile("valid_js_files.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			f.WriteString(_url_ + "\n")
			f.Close()
		} else {
			fmt.Printf("[ %v ] :: %v\n", resp.StatusCode, _url_)
		}
	}
}

func main() {
	rzlt_file()

	arg := os.Args[1]
	myFile, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	myScanner := bufio.NewScanner(myFile)

	for myScanner.Scan() {
		my_url_ := myScanner.Text()
		go request(my_url_)
		time.Sleep(75 * time.Millisecond)
	}
	time.Sleep(10 * time.Second)

}
