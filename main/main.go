package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

var accesskey string = "708db701d123a1c53f2dd121e6b49c82cef10206c5a392bdce64c81b98f530cb"

type images struct {
	Full  string
	Thumb string
}

type HomeAPI struct {
	Id    string
	Urls  images
	Likes int
}

func getHome() []HomeAPI {
	url := "https://api.unsplash.com/photos?client_id=" + accesskey
	start := time.Now()
	resp, err := http.Get(url)
	fmt.Println(time.Since(start).Seconds())
	if err != nil {
		fmt.Print(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	test := []HomeAPI{}
	_ = json.Unmarshal(body, &test)
	return test
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, getHome())
	})
	http.ListenAndServe(":8080", nil)

}
