package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	// if list is nil return false
	if list == nil {
		return false
	}
	// if list is empty return false
	if len(list) == 0 {
		return false
	}

	// iterate elements of list
	for _, x := range list {
		if x == num {
			return true
		}
	}
	// if element is target return true
	return false
}
