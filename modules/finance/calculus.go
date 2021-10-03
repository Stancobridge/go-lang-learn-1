package finance

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func CalculateTime(time string) (string, error) {

	if time == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), time)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func WorkWithSlice() []byte {
	// a := []byte{3, 4, 5, 6, 7, 8, 43}
	// t := make([]byte, len(a), (cap(a)+1)*2)
	// copy(t, a)

	// // for i := range a {
	// // 	t[i] = a[i]
	// // }

	// // c := a[:cap(a)]

	// // x := make([]string, 5, 20)

	// // return cap(a)
	// // return a[3]
	// return t

	p := []byte{0}
	// fmt.Println(len(p))
	p = AppendByte(p, []byte{7, 11, 13}...)
	return p
}

// Ruyil => Ruil

// func copySlice (from []int, to []byte) []byte{

// }

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) {
		newSlice := make([]byte, n)
		copy(newSlice, slice)
		// fmt.Println(newSlice)
		slice = newSlice
		// newSlice = nil // free memory
		// fmt.Println(newSlice)
	}

	slice = slice[0:n]
	copy(slice[m:n], data)
	// fmt.Println(cap(slice))
	return slice
}
