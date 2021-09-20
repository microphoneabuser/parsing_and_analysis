package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
)

const domain = "https://xkcd.com"
const jsonName = "info.0.json"
const maxCount = 2491

func main() {
	var comics Comics
	ch := make(chan string)
	for i := 1; i <= maxCount; i++ {
		query := fmt.Sprintf("%s/%d/%s", domain, i, jsonName)
		go ReadComic(query, &comics, ch)
	}
	for i := 1; i <= maxCount; i++ {
		fmt.Printf("%s", <-ch)
	}
	if err := WriteComics("comics.json", &comics); err != nil {
		fmt.Printf("writing json file failed. %s", err)
	} else {
		fmt.Println("Done.")
	}
}

func ReadComic(query string, comics *Comics, ch chan<- string) {
	var comic Comic
	resp, err := http.Get(query)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		ch <- fmt.Sprintln(err)
		return
	}
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		ch <- fmt.Sprintln(err)
		return
	}
	resp.Body.Close()
	comics.Comics = append(comics.Comics, &comic)
	ch <- fmt.Sprintf("%d...", comic.Num)
}

func WriteComics(path string, comics *Comics) error {
	sort.Slice(comics.Comics, func(i, j int) bool {
		return comics.Comics[i].Num < comics.Comics[j].Num
	})
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(file).Encode(comics); err != nil {
		return err
	}
	file.Close()
	return nil
}

type Comics struct {
	Comics []*Comic `json:"comics"`
}

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"self_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
