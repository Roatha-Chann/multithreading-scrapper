package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dchest/uniuri"
)

func main() {

	start := time.Now()

	rep := 20
	results := make(chan string)

	// Use goroutine to send multiple time-consuming jobs to the channel.
	for i := 0; i < rep; i++ {
		go func(num int) {
			results <- getContent(num)
		}(i)
	}

	// Receive results from the channel and use them.
	for i := 0; i < rep; i++ {
		writeFile(<-results)
	}

	elapsed := time.Since(start)
	log.Printf("Processing time: %s", elapsed)

}

func getContent(num int) string {

	var url [20]string
	url[0] = "https://jsonplaceholder.typicode.com/posts/1"
	url[1] = "https://jsonplaceholder.typicode.com/posts/2"
	url[2] = "https://jsonplaceholder.typicode.com/posts/3"
	url[3] = "https://jsonplaceholder.typicode.com/posts/4"
	url[4] = "https://jsonplaceholder.typicode.com/posts/5"
	url[5] = "https://jsonplaceholder.typicode.com/posts/6"
	url[6] = "https://jsonplaceholder.typicode.com/posts/7"
	url[7] = "https://jsonplaceholder.typicode.com/posts/8"
	url[8] = "https://jsonplaceholder.typicode.com/posts/9"
	url[9] = "https://jsonplaceholder.typicode.com/posts/10"
	url[10] = "https://jsonplaceholder.typicode.com/posts/11"
	url[11] = "https://jsonplaceholder.typicode.com/posts/12"
	url[12] = "https://jsonplaceholder.typicode.com/posts/13"
	url[13] = "https://jsonplaceholder.typicode.com/posts/14"
	url[14] = "https://jsonplaceholder.typicode.com/posts/15"
	url[15] = "https://jsonplaceholder.typicode.com/posts/16"
	url[16] = "https://jsonplaceholder.typicode.com/posts/17"
	url[17] = "https://jsonplaceholder.typicode.com/posts/18"
	url[18] = "https://jsonplaceholder.typicode.com/posts/19"
	url[19] = "https://jsonplaceholder.typicode.com/posts/20"

	resp, err := http.Get(url[num])

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	email_source := string(body)

	return email_source
}

func writeFile(content string) bool {

	random_string := uniuri.New()

	filename := random_string + ".html"

	ioutil.WriteFile("pages/"+filename, []byte(string(content)), 0755)

	return true
}
