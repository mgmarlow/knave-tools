package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

type Item struct {
	Name       string
	ArmorBonus int
	Slots      int
	Quality    int
	Hands      int
	Damage     int
}

var weapons = []string{
	"Dagger", "Cudgel", "Sickle", "Staff",
	"Spear", "Sword", "Mace", "Axe", "Flail",
	"Halberd", "War Hammer", "Long Sword", "Battle Axe",
	"Sling", "Bow", "Crossbow",
}

var dungeoneeringGear = []string{
	"Rope, 50ft", "Crowbar", "Lantern 2", "Pulleys",
	"Tinderbox", "Lamp oil", "Candles, 5", "Grap. hook",
	"Padlock", "Chain, 10ft", "Hammer", "Manacles",
	"Chalk, 10", "Waterskin", "Mirror", "Pole, 10ft",
	"Sack", "Tent", "Spikes, 5", "Torches",
}

var generalGear1 = []string{
	"Air bladder", "Bear trap", "Shovel", "Bellows", "Grease",
	"Saw", "Bucket", "Caltrops", "Chisel", "Drill",
	"Fish. rod", "Marbles", "Glue", "Pick", "Hourglass",
	"Net", "Tongs", "Lockpicks", "Metal file", "Nails",
}

var generalGear2 = []string{
	"Incense", "Sponge", "Lens", "Perfume", "Horn",
	"Bottle", "Soap", "Spyglass", "Tar pot", "Twine",
	"Fake jewels", "Blank book", "Card deck", "Dice set", "Cook pots",
	"Face paint", "Whistle", "Instrument", "Quill & Ink", "Small bell",
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
	Inventory    []Item
}

func NewKnave() *Knave {
	inventory := buildInventory()
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

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func rollAbilityDefense() int {
	rolls := []int{roll(6), roll(6), roll(6)}
	return Min(rolls) + 10
}

// TODO: Maybe just sample these and repeat names for probability
func rollArmor() []string {
	roll := roll(20)
	if roll == 20 {
		return []string{"Chain"}
	}
	if roll >= 15 {
		return []string{"Brigadine"}
	}
	if roll >= 4 {
		return []string{"Gambeson"}
	}
	return []string{}
}

func rollHelmetsAndShield() []string {
	roll := roll(20)
	if roll == 20 {
		return []string{"Helmet", "Shield"}
	}
	if roll >= 17 {
		return []string{"Shield"}
	}
	if roll >= 14 {
		return []string{"Helmet"}
	}
	return []string{}
}

func buildInventory() []Item {
	items := []string{}

	items = append(items, Sample(weapons))
	items = append(items, rollArmor()...)
	items = append(items, rollHelmetsAndShield()...)
	items = append(items, []string{
		Sample(dungeoneeringGear),
		Sample(dungeoneeringGear),
		Sample(generalGear1),
		Sample(generalGear2),
		"Travel rations (2 days)",
	}...)

	return Map(items, func(s string, _ int) Item {
		return getItem(s)
	})
}

func getItem(name string) Item {
	withName := func(i Item) Item {
		if i.Name == "" {
			i.Name = name
		}
		return i
	}

	items := map[string]Item{
		"Full Plate": {
			ArmorBonus: 6,
			Slots:      5,
			Quality:    7,
		},
		"Half Plate": {
			ArmorBonus: 5,
			Slots:      4,
			Quality:    6,
		},
		"Chain": {
			ArmorBonus: 4,
			Slots:      3,
			Quality:    5,
		},
		"Brigadine": {
			ArmorBonus: 3,
			Slots:      2,
			Quality:    4,
		},
		"Gambeson": {
			ArmorBonus: 2,
			Slots:      1,
			Quality:    3,
		},
		"Helmet": {
			ArmorBonus: 1,
			Slots:      1,
			Quality:    1,
		},
		"Shield": {
			ArmorBonus: 1,
			Slots:      1,
			Quality:    1,
		},
		"Sling": {
			Damage:  4,
			Slots:   1,
			Hands:   1,
			Quality: 3,
		},
		"Bow": {
			Damage:  6,
			Slots:   2,
			Hands:   2,
			Quality: 3,
		},
		"Crossbow": {
			Damage:  8,
			Slots:   3,
			Hands:   2,
			Quality: 3,
		},
	}

	if val, ok := items[name]; ok {
		return withName(val)
	}

	// Small weapon
	if Includes([]string{"Dagger", "Cudgel", "Sickle", "Staff"}, name) {
		return withName(Item{
			Damage:  6,
			Slots:   1,
			Hands:   1,
			Quality: 3,
		})
	}

	// Medium weapon
	if Includes([]string{"Spear", "Sword", "Mace", "Axe", "Flail"}, name) {
		return withName(Item{
			Damage:  8,
			Slots:   2,
			Hands:   1,
			Quality: 3,
		})
	}

	// Large weapon
	if Includes([]string{"Halberd", "War Hammer", "Long Sword", "Battle Axe"}, name) {
		return withName(Item{
			Damage:  10,
			Slots:   3,
			Hands:   2,
			Quality: 3,
		})
	}

	// Misc. items
	return Item{
		Name:       name,
		ArmorBonus: 0,
		Slots:      1,
		Quality:    1,
		Hands:      1,
		Damage:     0,
	}
}

func Includes[T comparable](arr []T, want T) bool {
	for _, v := range arr {
		if v == want {
			return true
		}
	}

	return false
}

func Sample[T any](arr []T) T {
	index := rand.Intn(len(arr))
	return arr[index]
}

func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

func Min(arr []int) int {
	cur := arr[0]

	for _, v := range arr {
		if cur > v {
			cur = v
		}
	}

	return cur
}
