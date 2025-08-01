package flatten

func Flatten(nested interface{}) []interface{} {
    out := []interface{}{}

    switch v := nested.(type) {
        case []interface{}: 
        	for _, el := range v {
                out = append(out, Flatten(el)...)
            }
        case interface{}:
            out = append(out, v)
    }

    return out
}