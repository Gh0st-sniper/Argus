package main

import (
	"fmt"

	"bufio"

	"os"

	"log"

	"net/http"
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

		//fmt.Println(text)

		full_domain := protocol + text + "." + domain

		//fmt.Println(full_domain)

		resp, err := http.Get(full_domain)

		if err != nil {

			log.Fatal(err)

		} else {

			if resp.StatusCode == 200 {

				fmt.Printf("Found subdomain :: %s\n", full_domain)

			}

			resp.Body.Close()

		}

	}
}
