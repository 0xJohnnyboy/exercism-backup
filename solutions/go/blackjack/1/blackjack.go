package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    value := 0
    switch card {
        case "ace": value = 11
        case "king": value = 10
        case "queen": value = 10
        case "jack": value = 10
        case "ten": value = 10
        case "nine": value = 9
        case "eight": value = 8
        case "seven": value = 7
        case "six": value = 6
        case "five": value = 5
        case "four": value = 4
        case "three": value = 3
        case "two": value = 2
        default: value = 0
    }
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    value := "H"
    handValue := ParseCard(card1) + ParseCard(card2)
	switch {
        case card1 == "ace" && card2 == "ace": value = "P"
        case handValue == 21 && ParseCard(dealerCard) < 10: value = "W"
        case handValue == 21 && ParseCard(dealerCard) >= 10: value = "S"
        case handValue >= 17 && handValue <= 20: value = "S"
        case handValue >= 12 && handValue <= 16 && ParseCard(dealerCard) < 7: value = "S"
        case handValue >= 12 && handValue <= 16 && ParseCard(dealerCard) >= 7: value = "H"    
        default: value = "H"
    }

    return value
}
