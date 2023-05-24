package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec int, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res
}

/*
you can convert a number from one base to another by:
dec - 10
base - 2
10 % 2 = 0
res = "0"

dec = dec / base
dec = 5
base = 2
5 % 2 = 1
res = "10"

dec = dec / base
dec = 2
base = 2
2 % 2 = 0
res = "010"

dec = dec / base
dec = 1
base = 2
1 % 2 = 1
res = "1010"

dec = dec / base
dec = 0
base = 2
0 % 2 = 0
res = "1010"

*/
