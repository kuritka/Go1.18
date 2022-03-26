package generics

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)



func TestCount(t *testing.T){
	// arrange
	x := []int{1,2,6,4,5}
	x32 := []X32{1,2,6,4,5}
	y := strings.Split("live is short, buy the shoes"," ")
	// act
	intSize := GetSizePrimitive[int](x)
	x32Size := GetSizePrimitive[X32](x32)	// X's underlying type is ~int
	strSize := GetSizePrimitive[string](y)
	// assert
	assert.Equal(t, 5, x32Size)
	assert.Equal(t, 5, intSize)
	assert.Equal(t, 6, strSize)
}


func TestIs64Bit(t *testing.T){
	var a float64 = 5
	var b float32 = 5
	var c int64 = 5
	assert.True(t, is64Bit(a))
	assert.False(t, is64Bit(b))
	assert.True(t, is64Bit(c))
}

func TestGetMax(t *testing.T) {
	// arrange
	x := []int{1,2,6,4,5}
	y := strings.Split("live is short, buy the shoes"," ")
	// act
	intSize := GetMax[int](x)
	stringSize := GetMax[string](y)
	// assert
	assert.Equal(t, 6, intSize)
	assert.Equal(t, "short,", stringSize)
}