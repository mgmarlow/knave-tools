package main

import "fmt"

type Traits struct {
	Physique   string
	Face       string
	Skin       string
	Hair       string
	Clothing   string
	Virtue     string
	Vice       string
	Speech     string
	Background string
	Misfortune string
	Alignment  string
}

var physique = []string{
	"Athletic", "Brawny", "Corpulent", "Delicate", "Gaunt",
	"Hulking", "Lanky", "Ripped", "Rugged", "Scrawny",
	"Short", "Sinewy", "Slender", "Flabby", "Statuesque",
	"Stout", "Tiny", "Towering", "Willowy", "Wiry",
}

var face = []string{
	"Bloated", "Blunt", "Bony", "Chiseled", "Delicate",
	"Elongated", "Patrician", "Pinched", "Hawkish", "Broken",
	"Impish", "Narrow", "Ratlike", "Round", "Sunken",
	"Sharp", "Soft", "Square", "Wide", "Wolfish",
}

var skin = []string{
	"Battle Scar", "Birthmark", "Burn Scar", "Dark", "Makeup",
	"Oily", "Pale", "Perfect", "Pierced", "Pockmarked",
	"Reeking", "Tattooed", "Rosy", "Rough", "Sallow",
	"Sunburned", "Tanned", "War Paint", "Weathered", "Whip Scar",
}

var hair = []string{
	"Bald", "Braided", "Bristly", "Cropped", "Curly",
	"Disheveled", "Dreadlocks", "Filthy", "Frizzy", "Greased",
	"Limp", "Long", "Luxurious", "Mohawk", "Oily",
	"Ponytail", "Silky", "Topknot", "Wavy", "Wispy",
}

var clothing = []string{
	"Antique", "Bloody", "Ceremonial", "Decorated", "Eccentric",
	"Elegant", "Fashionable", "Filthy", "Flamboyant", "Stained",
	"Foreign", "Frayed", "Frumpy", "Livery", "Oversized",
	"Patched", "Perfumed", "Rancid", "Torn", "Undersized",
}

var virtue = []string{
	"Ambitious", "Cautious", "Courageous", "Courteous", "Curious",
	"Disciplined", "Focused", "Generous", "Gregarious", "Honest",
	"Honorable", "Humble", "Idealistic", "Just", "Loyal",
	"Merciful", "Righteous", "Serene", "Stoic", "Tolerant",
}

var vice = []string{
	"Aggressive", "Arrogant", "Bitter", "Cowardly", "Cruel",
	"Deceitful", "Flippant", "Gluttonous", "Greedy", "Irascible",
	"Lazy", "Nervous", "Prejudiced", "Reckless", "Rude",
	"Suspicious", "Vain", "Vengeful", "Wasteful", "Whiny",
}

var speech = []string{
	"Blunt", "Booming", "Breathy", "Cryptic", "Drawling",
	"Droning", "Flowery", "Formal", "Gravelly", "Hoarse",
	"Mumbling", "Precise", "Quaint", "Rambling", "Rapid-fire",
	"Dialect", "Slow", "Squeaky", "Stuttering", "Whispery",
}

var background = []string{
	"Alchemist", "Beggar", "Butcher", "Burglar", "Charlatan",
	"Cleric", "Cook", "Cultist", "Gambler", "Herbalist",
	"Magician", "Mariner", "Mercenary", "Merchant", "Outlaw",
	"Performer", "Pickpocket", "Smuggler", "Student", "Tracker",
}

var misfortunes = []string{
	"Abandoned", "Addicted", "Blackmailed", "Condemned", "Cursed",
	"Defrauded", "Demoted", "Discredited", "Disowned", "Exiled",
	"Framed", "Haunted", "Kidnapped", "Mutilated", "Poor",
	"Pursued", "Rejected", "Replaced", "Robbed", "Suspected",
}

var alignment = []string{
	"Law", "Law", "Law", "Law", "Law",
	"Neutrality", "Neutrality", "Neutrality", "Neutrality", "Neutrality",
	"Neutrality", "Neutrality", "Neutrality", "Neutrality", "Neutrality",
	"Chaos", "Chaos", "Chaos", "Chaos", "Chaos",
}

func NewTraits() Traits {
	return Traits{
		Physique:   Sample(physique),
		Face:       Sample(face),
		Skin:       Sample(skin),
		Hair:       Sample(hair),
		Clothing:   Sample(clothing),
		Virtue:     Sample(virtue),
		Vice:       Sample(vice),
		Speech:     Sample(speech),
		Background: Sample(background),
		Misfortune: Sample(misfortunes),
		Alignment:  Sample(alignment),
	}
}

func (t Traits) toString() string {
	fmtStr := "%s. Wears %s clothes, and has %s speech. Has a %s physique, " +
		"a %s face, %s skin, and %s hair. Is %s, but %s. Has been %s in " +
		"the past. Favors %s."

	return fmt.Sprintf(fmtStr,
		t.Background,
		t.Clothing,
		t.Speech,
		t.Physique,
		t.Face,
		t.Skin,
		t.Hair,
		t.Virtue,
		t.Vice,
		t.Misfortune,
		t.Alignment,
	)
}
