package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	a := "Welcome to the Tech Palace, "
    return a + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	a := strings.Repeat("*", numStarsPerLine)
    b := a + "\n" + welcomeMsg + "\n" + a
    return  b
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	a := strings.ReplaceAll(oldMsg, "*", " ")
    return strings.TrimSpace(a)
}
