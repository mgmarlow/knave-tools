package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

//go:embed views/*
var files embed.FS

// Template functions
var funcs = template.FuncMap{
	"bonus": func(defense int) int {
		return Bonus(defense)
	},
	"slots": func(inv Inventory) int {
		return Slots(inv)
	},
	"plaintext": func(k Knave) string {
		return PlainTextKnave(k)
	},
}

type IndexContent struct {
	Knave *Knave
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(
		template.New("layout.html").Funcs(funcs).ParseFS(files, "views/layout.html", "views/index.html"))
	content := &IndexContent{
		Knave: NewKnave(),
	}
	t.Execute(w, content)
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
