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
var funcs = template.FuncMap{
	"bonus": func(defense int) int {
		return defense - 10
	},
	"armorName": func(kind int) string {
		return armorDisplayName[kind]
	},
}

const (
	NoArmor   int = 11
	Gambeson      = 12
	Brigadine     = 13
	Chain         = 14
	HalfPlate     = 15
	FullPlate     = 16
)

var armorDisplayName = map[int]string{
	FullPlate: "Full Plate",
	HalfPlate: "Half Plate",
	Chain:     "Chain",
	Brigadine: "Brigadine",
	Gambeson:  "Gambeson",
	NoArmor:   "No armor",
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
	Armor        int
}

func NewKnave() *Knave {
	return &Knave{
		Name:         "foo bar",
		Strength:     rollAbilityDefense(),
		Dexterity:    rollAbilityDefense(),
		Constitution: rollAbilityDefense(),
		Intelligence: rollAbilityDefense(),
		Wisdom:       rollAbilityDefense(),
		Charisma:     rollAbilityDefense(),
		Armor:        rollArmor(),
	}
}

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func rollAbilityDefense() int {
	rolls := []int{roll(6), roll(6), roll(6)}
	return min(rolls) + 10
}

func min(arr []int) int {
	cur := arr[0]

	for _, v := range arr {
		if cur > v {
			cur = v
		}
	}

	return cur
}

func rollArmor() int {
	roll := roll(20)
	if roll == 20 {
		return Chain
	}
	if roll >= 15 {
		return Brigadine
	}
	if roll >= 4 {
		return Gambeson
	}
	return NoArmor
}
