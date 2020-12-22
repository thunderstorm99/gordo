package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type winner struct {
	Name   string
	Number int
	Price  int
}

type apiResponse struct {
	Number int `json:"numero"`
	Price  int `json:"premio"`
}

type lose struct {
	Owners []struct {
		Name    string `json:"name"`
		Tickets []int  `json:"tickets"`
	} `json:"owners"`
}

const (
	siteURL = "https://api.elpais.com/ws/LoteriaNavidadPremiados?n="
)

// global variables
var winners []winner

func main() {
	l := readJSON()
	getResults(l)

}

func readJSON() lose {
	bytes, err := ioutil.ReadFile("tickets.json")
	if err != nil {
		panic(err)
	}

	var j lose

	json.Unmarshal(bytes, &j)
	return j
}

func getResults(l lose) {
	// loop forever
	for {
		// loop through each owner
		for i := range l.Owners {
			// loop through their tickets
			for j := range l.Owners[i].Tickets {
				url := siteURL + strconv.Itoa(l.Owners[i].Tickets[j])
				response, err := http.Get(url)
				if err != nil {
					panic(err)
				}
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					panic(err)
				}

				// remove non standard JSON
				body = bytes.Replace(body, []byte("busqueda="), []byte(""), -1)

				// variable to hold body
				var b apiResponse

				json.Unmarshal(body, &b)
				if b.Price != 0 {
					// a price was won, append to winners
					appendWinner(l.Owners[i].Name, b.Number, b.Price)
				}
			}
		}
		// announce winners
		announce()
		time.Sleep(10 * time.Second)
	}
}

func appendWinner(name string, number int, price int) {
	// check if price is already in winners list
	for i := range winners {
		if winners[i].Number == number {
			// this number is already in winners list
			if winners[i].Price == price {
				// this is exactly the price we already have
				return
			}
			continue
		}
	}
	// append winner to the array
	winners = append(winners, winner{Name: name, Number: number, Price: price})
}

func announce() {
	// clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Println("Lotería de Navidad", time.Now().Year())

	// loop through winners
	for i := range winners {
		fmt.Println(winners[i].Name, "won", winners[i].Price, "Euro with Ticket", winners[i].Number)
	}
}
