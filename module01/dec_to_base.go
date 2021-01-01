package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	val := DecToBaseR(dec, base, "")
	//fmt.Println(val)
	return val
}

func DecToBaseR(dec, base int, val string) string {
	//fmt.Println("Dec: ", fmt.Sprint(dec), " Base: ", fmt.Sprint(base), " Val: ", val)
	if dec == 0 {
		return val
	}
	return DecToBaseR(dec/base, base, GetVal(dec, base)+val)
}

func GetVal(dec, base int) string {
	val := dec % base
	valS := ""
	switch val {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		valS = fmt.Sprint(val)
	case 10:
		valS = "A"
	case 11:
		valS = "B"
	case 12:
		valS = "C"
	case 13:
		valS = "D"
	case 14:
		valS = "E"
	case 15:
		valS = "F"
	}
	return valS
}
