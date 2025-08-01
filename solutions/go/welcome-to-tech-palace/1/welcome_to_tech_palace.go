package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToTitle(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsStr := ""
    for i := 0; i<numStarsPerLine; i++ {
        starsStr = starsStr + "*"
    }
	return starsStr + "\n" + welcomeMsg + "\n" + starsStr
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "* \n")
}
