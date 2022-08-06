package main

type Knave struct {
	Name         string
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
		Name:         "foo bar",
		Strength:     rollAbilityDefense(),
		Dexterity:    rollAbilityDefense(),
		Constitution: rollAbilityDefense(),
		Intelligence: rollAbilityDefense(),
		Wisdom:       rollAbilityDefense(),
		Charisma:     rollAbilityDefense(),
		Armor:        getArmorBonus(inventory),
		Inventory:    inventory,
		Traits:       NewTraits(),
	}
}

func getArmorBonus(items []Item) int {
	count := 0

	for _, item := range items {
		count += item.ArmorBonus
	}

	if count > 0 {
		return count
	}

	return 1
}

func rollAbilityDefense() int {
	rolls := []int{Roll(6), Roll(6), Roll(6)}
	return Min(rolls) + 10
}
