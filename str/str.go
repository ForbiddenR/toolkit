package str

import "strconv"

func ParseUint[U ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint](s string) (U, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return U(u), err
	}
	return U(u), nil
}

func ParseInt[I ~int8 | ~int16 | ~int32 | ~int64 | ~int](s string) (I, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return I(i), err
	}
	return I(i), err
}

func FormatUint[U ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint](u U) string {
	return strconv.FormatUint(uint64(u), 10)
}

func FormatInt[I ~int8 | ~int16 | ~int32 | ~int64 | ~int](i I) string {
	return strconv.FormatInt(int64(i), 10)
}
