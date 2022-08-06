package main

import "fmt"

type Item struct {
	Name       string
	ArmorBonus int
	Slots      int
	Quality    int
	Hands      int
	Damage     int
}

func (i Item) String() string {
	if i.ArmorBonus != 0 {
		return fmt.Sprintf("%s (%d slot(s), %d quality)",
			i.Name, i.Slots, i.Quality)
	}

	if i.Damage != 0 {
		return fmt.Sprintf("%s (d%d damage, %d slot(s) %d hand(s), %d quality)",
			i.Name, i.Damage, i.Slots, i.Hands, i.Quality)
	}

	return i.Name
}

type Inventory = []Item

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

func NewInventory() Inventory {
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

func ArmorDefense(items []Item) int {
	count := 0

	for _, item := range items {
		count += item.ArmorBonus
	}

	if count > 0 {
		return count + 10
	}

	return 11
}

func Slots(items []Item) int {
	count := 0
	for _, item := range items {
		count += item.Slots
	}
	return count
}

func rollArmor() []string {
	roll := Roll(20)
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
	roll := Roll(20)
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
