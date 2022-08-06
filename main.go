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
		return defense - 10
	},
	"defense": func(bonus int) int {
		return bonus + 10
	},
	"formatTraits": func(t Traits) string {
		return t.toString()
	},
	"slots": func(inv Inventory) int {
		return Slots(inv)
	},
	"formatItem": func(item Item) string {
		if item.ArmorBonus != 0 {
			return fmt.Sprintf("%s (%d slot(s), %d quality)", item.Name, item.Slots, item.Quality)
		}

		if item.Damage != 0 {
			return fmt.Sprintf("%s (d%d damage, %d slot(s) %d hand(s), %d quality)", item.Name, item.Damage, item.Slots, item.Hands, item.Quality)
		}

		return item.Name
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
