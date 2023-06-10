package starter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("Meee")
	assert.Equal(t, "Hello Meee. Welcome!", greeting)

	another_greeting := SayHello("Fern")
	assert.Equal(t, "Hello Fern. Welcome!", another_greeting)
}

// func TestNormalSayHello (){
// 	assert.Equal(t,"Hello World!",NormalSayHello())
// }

func TestEvenOrOdd(t *testing.T) {
	t.Run("Check Non Negative Odd  Numbers", func(t *testing.T) {
		assert.Equal(t, "50 is an Odd  number", OddOrEven(3))
		assert.Equal(t, "42 is an Odd  number", OddOrEven(3))
		assert.Equal(t, "0 is an Odd  number", OddOrEven(3))
	})

	t.Run("Check Non Negative Odd  Numbers", func(t *testing.T) {
		assert.Equal(t, "-50 is an Odd  number", OddOrEven(3))
		assert.Equal(t, "-42 is an Odd  number", OddOrEven(3))
	})

	t.Run("Check Non Negative Odd  Numbers", func(t *testing.T) {
		assert.Equal(t, "3 is an Odd  number", OddOrEven(3))
		assert.Equal(t, "9 is an Odd  number", OddOrEven(3))
		assert.Equal(t, "55 is an Odd  number", OddOrEven(3))
	})

	t.Run("Check Non Negative Odd  Numbers", func(t *testing.T) {
		assert.Equal(t, "50 is an Odd  number", OddOrEven(3))
		assert.Equal(t, "42 is an Odd  number", OddOrEven(3))
		assert.Equal(t, "0 is an Odd  number", OddOrEven(3))
	})
	t.Run("Check Negative Odd Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", OddOrEven(-45))
		assert.Equal(t, "-1 is an odd number", OddOrEven(-1))
		assert.Equal(t, "-101 is an odd number", OddOrEven(-101))
	})
}
