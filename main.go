package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var err error

type People struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

func main() {
	// JSON dosyasını açalım.
	jsonDosya, err := os.Open("people.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonDosya.Close()

	normalMethod(jsonDosya)
	fmt.Println("Normal Method Bitti ***********************")

	// Dosyanın başına git
	jsonDosya.Seek(0, 0)
	dataStreaming(jsonDosya)
	fmt.Println("Data Streaming Bitti ***********************")
}

func normalMethod(jsonDosya *os.File) {
	// Dosya içeriğini okuyalım.
	icerik, err := ioutil.ReadAll(jsonDosya)
	if err != nil {
		fmt.Println(err)
	}

	// Dosyadaki veriyi ayrıştırıp değişkenimize aktaralım.
	var people []People
	json.Unmarshal(icerik, &people)

	// JSON verisindeki isimleri ekrana yazdıralım.
	for _, person := range people {
		fmt.Println(person.Name)
	}
}

func dataStreaming(jsonDosya *os.File) {
	dec := json.NewDecoder(jsonDosya)

	// Başlangıç verisini ([) okuyalım
	_, err = dec.Token()
	if err != nil {
		log.Println(err)
	}

	var person People
	// Bu döngü okunacak bir sonraki obje var ise devam eder.
	for dec.More() {
		err = dec.Decode(&person)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(person.Name)
	}

	// Kapanış verisini (]) okuyalım
	_, err = dec.Token()
	if err != nil {
		log.Println(err)
	}
}
