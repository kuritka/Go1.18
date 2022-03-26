package generics

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


func TestCount(t *testing.T){
	// arrange
	x := []int{1,2,6,4,5}
	y := strings.Split("live is short, buy the shoes"," ")
	// act
	intSize := GetSizePrimitive[int](x)
	strSize := GetSizePrimitive[string](y)
	// assert
	assert.Equal(t, 5, intSize)
	assert.Equal(t, 6, strSize)
}