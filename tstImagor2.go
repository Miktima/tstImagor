package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func checkStatus(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Print("\n")
		fmt.Println("Elapsed time", elapsed, "| Link: ", url, " | Status:", resp.Status)
	} else {
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Print(elapsed, "|")
	}
}

func main() {
	// Open our txt file with links
	readFile, err := os.Open("lat_links56.txt")
	// if we os.ReadFile returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	// link_prefix := "https://cdnpl1.img.sputniknews.com"
	// link_prefix := "https://cdn1.sputniknewsbr.com.br"
	// link_prefix := "https://img.pl.sputniknews.com"
	// link_prefix := "https://img.sputniknewsbr.com.br"
	link_prefix := "https://img.sputniknews.lat"

	i := 0
	var link string
	for fileScanner.Scan() {
		link = link_prefix + fileScanner.Text()
		if i == 0 {
			fmt.Println(link)
		}
		go checkStatus(link)
		i += 1
		// if i == 8000 {
		// 	break
		// }
	}
	var input string
	fmt.Scanln(&input)
}
