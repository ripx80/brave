package glue

/*InSlice checks if that value is in slice*/
func InSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
