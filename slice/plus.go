package slice

// PlusStrings adds two []string arrays together, returning a new copied []string.
func PlusStrings(s, plus []string) []string {

	a := make([]string, 0, len(s)+len(plus))
	return append(append(a, s...), plus...)

}
