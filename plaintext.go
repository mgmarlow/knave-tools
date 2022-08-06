package main

import "fmt"

func PlainTextKnave(k Knave) string {
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
		abilityString(k.Strength),
		abilityString(k.Dexterity),
		abilityString(k.Constitution),
		abilityString(k.Intelligence),
		abilityString(k.Wisdom),
		abilityString(k.Charisma),
		abilityString(k.Armor),
		plainTextInventory(k.Inventory),
	)
}

func abilityString(defense int) string {
	return fmt.Sprintf("+%d %d", Bonus(defense), defense)
}

func plainTextInventory(items []Item) string {
	str := ""

	for _, item := range items {
		str += fmt.Sprintf("* %s\n", item)
	}

	return str
}
