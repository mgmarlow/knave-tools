package main

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

func Bonus(defense int) int {
	return defense - 10
}

func rollAbilityDefense() int {
	rolls := []int{Roll(6), Roll(6), Roll(6)}
	return Min(rolls) + 10
}
