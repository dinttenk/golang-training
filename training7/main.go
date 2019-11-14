package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	Id      int
	Title   string
	Content string
}

var data = []article{
	article{1, "judul pertama", "isi pertama"},
	article{2, "judul kedua", "isi kedua"},
}

func articles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	}
	http.Error(w, "", http.StatusBadRequest)
	return
}
func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "hello", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
	return
}

func main() {
	http.HandleFunc("/articles", articles)
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// type StarWarsPeople struct {
// 	Name      string   `json: "name"`
// 	Height    string   `json: "height" `
// 	Mass      string   `json:"mass"`
// 	HairColor string   `json: "hair_color"`
// 	SkinColor string   `json:"skin_color"`
// 	EyeColor  string   `json: "eye_color"`
// 	BirthYear string   `json: "birth_year"`
// 	Gender    string   `json: "gender"`
// 	Films     []string `json: "films"`
// }

// func main() {
// 	response, _ := http.Get("http://swapi.co/api/people/1")

// 	responseData, _ := ioutil.ReadAll(response.Body)
// 	defer response.Body.Close()

// 	fmt.Println(string(responseData))
// 	var People StarWarsPeople
// 	json.Unmarshal(responseData, &People)
// 	fmt.Println("----- Print failed -----")
// 	fmt.Println("Name:", People.Name)
// 	fmt.Println("Heigth:", People.Height)
// 	fmt.Println("Hair Color:", People.HairColor)
// 	fmt.Println("Mass:", People.Mass)
// 	fmt.Println("films:", People.Films[0])

// }
