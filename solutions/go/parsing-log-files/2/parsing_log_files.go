package parsinglogfiles
import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	reg := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]\s.+`)
    return reg.FindString(text) != ""
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<[-=~*]*>`)
    return reg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	reg := regexp.MustCompile(`".*pass[wW]ord.*"`)
    passwordCount := 0
    for _, line := range lines {
        passwordCount += len(reg.FindAllString(line, -1))
    }
    return passwordCount
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line\d*`)
    return reg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	userReg := regexp.MustCompile(`User\s+(\w+)\s+`)
    for i, line := range lines {
        matches := userReg.FindStringSubmatch(line)
        if len(matches) == 0 {
            continue
        }

        fmt.Println(matches, matches[1])
        lines[i] = "[USR] " + matches[1] + " " + line
    }
    return lines
}
