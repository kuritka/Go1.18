package generics


func GetSizePrimitive[V int | int64 | float64 | string](arr []V) int {
	cnt := len(arr)
	return cnt
}
