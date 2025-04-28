package str

import (
	"strconv"
)

type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

type Signed interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type Float interface {
	~float32 | ~float64
}

func ParseUint[U Unsigned](s string) (U, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return U(u), err
	}
	return U(u), nil
}

func ParseInt[I Signed](s string) (I, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return I(i), err
	}
	return I(i), err
}

func ParseFloat[F Float](s string) (F, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return F(f), err
	}
	return F(f), nil
}

func FormatUint[U Unsigned](u U) string {
	return strconv.FormatUint(uint64(u), 10)
}

func FormatInt[I Signed](i I) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatFloat[F Float](f F) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}
