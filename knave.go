package main

import "fmt"

type Knave struct {
	Hp           int
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Armor        int
	Inventory    Inventory
	Traits       Traits
}

func NewKnave() *Knave {
	inventory := NewInventory()

	return &Knave{
		Hp:           Roll(8),
		Strength:     rollAbilityDefense(),
		Dexterity:    rollAbilityDefense(),
		Constitution: rollAbilityDefense(),
		Intelligence: rollAbilityDefense(),
		Wisdom:       rollAbilityDefense(),
		Charisma:     rollAbilityDefense(),
		Armor:        ArmorDefense(inventory),
		Inventory:    inventory,
		Traits:       NewTraits(),
	}
}

func (k Knave) String() string {
	str := `
HP: %d

Traits
====================
%s

Attributes
====================
STR: %s
DEX: %s
CON: %s
INT: %s
WIS: %s
CHA: %s
ARM: %s
====================

Inventory
====================
%s`

	return fmt.Sprintf(str,
		k.Hp,
		k.Traits,
		AbilityString(k.Strength),
		AbilityString(k.Dexterity),
		AbilityString(k.Constitution),
		AbilityString(k.Intelligence),
		AbilityString(k.Wisdom),
		AbilityString(k.Charisma),
		AbilityString(k.Armor),
		PlainTextInventory(k.Inventory),
	)
}

func Bonus(defense int) int {
	return defense - 10
}

func AbilityString(defense int) string {
	return fmt.Sprintf("+%d %d", Bonus(defense), defense)
}

func rollAbilityDefense() int {
	rolls := []int{Roll(6), Roll(6), Roll(6)}
	return Min(rolls) + 10
}
