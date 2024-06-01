package pkg

func cmt1(x string) bool {
	// A
	if len(x) > 0 {
		return false
	}
	// B
	return true
}

func cmt2(x string) bool {
	if len(x) > 0 { // A
		return false
	}
	return true // B
}

func cmt3(x string) bool {
	if len(x) > 0 {
		return false // A
	}
	return true // B
}

func cmt4(x string) bool {
	if len(x) > 0 {
		return false // A
	}
	return true
	// B
}

// Hard to test, as the diag line adds a comment.
//func cmt5(x string) bool {
//	if len(x) > 0 { // diag(`should use 'return len(x) == 0'`)
//		return false
//	}
//	return true // A
//}

func cmt6(x string) bool {
	if len(x) > 0 { //@diag(`should use 'return len(x) == 0'`)
		return false // A
	}
	return true
}
