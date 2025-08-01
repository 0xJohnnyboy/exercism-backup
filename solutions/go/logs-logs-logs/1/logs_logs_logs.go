package logs
import "fmt"
import "unicode/utf8"


// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
        switch fmt.Sprintf("%U", char) {
            case "U+2757": 
            	return "recommendation"
            case "U+1F50D": 
            	return "search"
            case "U+2600": 
            	return "weather"
        }
    }
    
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var newLogLine string
	for _, char := range log {
        if char == oldRune {
            newLogLine += string(newRune)
            continue
        }
        newLogLine += string(char)
    }
    
    return newLogLine
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
