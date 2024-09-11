package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		fallthrough
	case "jack":
		fallthrough
	case "queen":
		fallthrough
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	mine := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case mine == 22:
		return "P"
	case mine == 21 && dealer < 10:
		return "W"
	case mine == 21 && dealer >= 10:
		return "S"
	case mine >= 17 && mine <= 20:
		return "S"
	case mine >= 12 && mine <= 26 && dealer < 7:
		return "S"
	case mine >= 12 && mine <= 26 && dealer >= 7:
		return "H"
	default:
		return "H"
	}
}
