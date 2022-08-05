package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
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

func ArmorToDisplayName(kind int) string {
	return armorDisplayName[kind]
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
