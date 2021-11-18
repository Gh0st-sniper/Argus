package main

import (
	"fmt"

	"bufio"

	"os"

	"log"

	"net/http"

	"time"
)

func main() {

	fmt.Println("Enter domain to enumerate : ")

	var wordlist string
	var domain string

	fmt.Scanf("%s", &domain)

	fmt.Println("Enter wordlist")

	fmt.Scanf("%s", &wordlist)

	enumerate(wordlist, domain)

}

func enumerate(filename, domain string) {

	protocol := "https://"

	file, err := os.Open(filename)

	if err != nil {

		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		text := scanner.Text()

		full_domain := protocol + text + "." + domain

		resp, err := http.Get(full_domain)

		if err != nil {

			fmt.Printf("Missed subdomain :: %s\n", full_domain)

		} else {

			if resp.StatusCode == 200 {

				fmt.Printf("Found subdomain :: %s\n", full_domain)
				time.Sleep(5 * time.Second)

			}

			resp.Body.Close()

		}

	}
}
