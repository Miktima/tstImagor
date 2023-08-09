package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func checkStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Print("\n")
		fmt.Println("Link: ", url, " | Status:", resp.Status, "!!!!!!")
	} else {
		fmt.Print(".")
	}
}

func main() {
	// initial values
	type initPar struct {
		URLImagor   string
		Image       string
		StartHeight int
		StartWidth  int
		Step        int
		QPoint      int
	}
	var par initPar
	// Open our jsonFile
	byteValue, err := os.ReadFile("conf.json")
	// if we os.ReadFile returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	errUn := json.Unmarshal(byteValue, &par)
	if err != nil {
		fmt.Println(errUn)
	}
	fmt.Println("URL to Imagor: ", par.URLImagor)
	fmt.Println("Image: ", par.Image)
	fmt.Println("Start height (px): ", par.StartHeight)
	fmt.Println("Start width (px): ", par.StartWidth)
	fmt.Println("Step (px): ", par.Step)
	fmt.Println("Quantity of points: ", par.QPoint)

	var input2 string
	fmt.Scanln(&input2)

	var link string
	for i := 0; i < par.QPoint; i++ {
		w := strconv.Itoa(i*par.Step + par.StartWidth)
		h := strconv.Itoa(i*par.Step + par.StartHeight)
		link = par.URLImagor + "unsafe/fit-in/" + w + "x" + h + "/" + par.Image
		// link = par.URLImagor + "unsafe/" + strconv.Itoa(i*par.Step) + "x" + strconv.Itoa(i*par.Step) + ":" + w + "x" + h + "/" + par.Image
		go checkStatus(link)
	}
	var input string
	fmt.Scanln(&input)
}
