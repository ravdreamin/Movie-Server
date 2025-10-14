package main

type Movies struct {
	id       string    `json: "id"`
	isbn     string    `json: "isbn"`
	title    string    `json: "title"`
	director *Director `json: "director"`
}

type Director struct {
	firstname string `json: "firstname"`
	lastname  string `json: "lastname"`
}

var movie []Movies

func main() {

}
