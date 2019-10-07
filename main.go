// Random Numbers generators. Generates int, uint, float64 & more.
// func main() hash a number with crypto/sha256 package using fmt.Sprintf to convert numbers into strings.

package main

import (
	crypto "crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	math "math/rand"
	"sync"
)

// global value that performs all random number operations
var globalSource = math.New(&source{})

// math/rand Source using entropy from crypto/rand
type source struct {
	data [8]byte // 64 bits
	mtx  sync.Mutex
}

func main() {

	for i := 0; i <= 10; i++ {
		// Initialise r1 variable to a random float64 number
		r1 := rand.Float64()
		fmt.Printf("This is the random float64 number: %v\n", r1)
		// Convert r1 to string with fmt.Sprintln to pass it to sha256.New() function and hash it
		sr1 := fmt.Sprintln(r1)
		fmt.Printf("This is a converted float to string: %v\n", sr1)
		// Initialise variable h to sha256.New function and hash previously
		// converted float64 to string into a SHA256 hash
		if fmt.Sprintln(r1) != sr1 {
			fmt.Println("Ooops@! Something's wrong here!")
		} else {
			fmt.Println("OK, conversion success!")

		}
		h := sha256.New()
		h.Write([]byte(fmt.Sprint(r1)))
		fmt.Printf("This is the SHA256 hash of the converted float64 to string: %x\n\n", h.Sum(nil))

	}
}

// Float64 returns a pseudo-random number in [0.0,1.0)
func Float64() float64 { return globalSource.Float64() }

// needed by math/rand.Source, but we don't require seeding
func (src *source) Seed(_ int64) {}

// needed by math/rand.Source
func (src *source) Int63() int64 {
	return int64(src.Uint64() >> 1) // need 63 bit number
}

// Uint64 returns a randomly selected 64-bit unsigned integer.
func (src *source) Uint64() uint64 {
	src.mtx.Lock()
	defer src.mtx.Unlock()

	data := src.data[:]
	n, err := crypto.Read(data)
	if err != nil {
		panic("crypto.Read failed: " + err.Error())
	}
	if n != 8 {
		panic("read too few random bytes")
	}
	return binary.BigEndian.Uint64(data)
}

// Intn returns, as an int, a non-negative pseudo-random number in
// [0,n). It panics if n <= 0.
func Intn(n int) int {
	return globalSource.Intn(n)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an
// int64.
func Int63() int64 {
	return globalSource.Int63()
}

// Int63n returns, as an int64, a non-negative pseudo-random number in
// [0,n). It panics if n <= 0.
func Int63n(n int64) int64 {
	return globalSource.Int63n(n)
}

// Uint64 returns a randomly selected 64-bit unsigned integer.
func Uint64() uint64 {
	return globalSource.Uint64()
}
