package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Name  string `json:"name"`
	Add   string `json:"add"`
	Class string `json:"class"`
}
type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{Name: "Chandu kumar", Add: "Obra", Class: "BCA 3rd"},
		Article{Name: "sonal kumar", Add: "bumaru", Class: "BCA 2nd"},
		Article{Name: "manish kumar", Add: "basdiha", Class: "MCA 2nd"},
	}
	fmt.Println("All Article Hit Endpoint :", "\n", article, "\n")
	json.NewEncoder(w).Encode(article)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage hit ")
	//fmt.Fprintf(w, "all articles")
	//fmt.Fprintf(w, "Enter the first number -")
	//fmt.Fprintf(w, "Enter the second number -")
	//fmt.Fprintf(w, strconv.Itoa(e+u), "sum of the number = ")

}
func handle() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aticle", allArticle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handle()

}
