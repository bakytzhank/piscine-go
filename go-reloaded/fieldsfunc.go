

func main() {
    // Return true if comma or colon char.
    f := func(c rune) bool {
        return c == ',' || c == ':'
    }
    // Separate into fields with func.
    fields := strings.FieldsFunc("cat,dog:bird", f)
    fmt.Println(fields)
}