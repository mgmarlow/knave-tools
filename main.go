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

//go:embed views/*
var files embed.FS

// Template functions
var funcs = template.FuncMap{}

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
	rand.Seed(time.Now().UnixMicro())

	http.HandleFunc("/", indexHandler)

	fmt.Println("Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Knave struct {
	Name         string
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

func NewKnave() *Knave {
	return &Knave{
		Name:         "foo bar",
		Strength:     abilityRoll(),
		Dexterity:    abilityRoll(),
		Constitution: abilityRoll(),
		Intelligence: abilityRoll(),
		Wisdom:       abilityRoll(),
		Charisma:     abilityRoll(),
	}
}

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func abilityRoll() int {
	rolls := []int{roll(6), roll(6), roll(6)}

	min := rolls[0]
	minIndex := 0
	for i, roll := range rolls {
		if min > roll {
			min = roll
			minIndex = i
		}
	}

	return sum(removeIndex(rolls, minIndex)...)
}

func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func sum(n ...int) int {
	count := 0
	for _, v := range n {
		count += v
	}
	return count
}
