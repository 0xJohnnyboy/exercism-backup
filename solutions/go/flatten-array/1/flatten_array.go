package flatten
import "fmt"
import "regexp"
import "strconv"

func Flatten(nested interface{}) []interface{} {
    str := fmt.Sprintf("%v", nested)
    reg := regexp.MustCompile(`(-?\d+)`)
    var out []interface{}
    empty := []interface{}{}
    matches := reg.FindAllString(str, -1)
    for _, m := range matches {
        num, err := strconv.Atoi(m)
        if err != nil {
            return empty
        }
        out = append(out, num)
    }

    if len(out) == 0 {
        return empty
    }

    return out
}