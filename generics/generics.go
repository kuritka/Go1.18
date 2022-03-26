package generics

type X32 int

func GetSizePrimitive[V ~int | float64 | string](arr []V) int {
	cnt := len(arr)
	return cnt
}
//
//// GetMax gets largest/ longest element from the slice
func GetMax[T int | string](slice []T) T {
	var strMax int
	var ret T
	for _, value := range slice {
		switch t := (interface{})(value).(type) {
		case string:
			if len(t) > strMax {
				strMax = len(t)
				ret = value
			}
		default:
			if value > ret {
				ret = value
			}
		}
	}
	return ret
}


func returnZero[T any](s ...T) T {
	var zero T
	return zero
}

func is64Bit[T float32 | float64 | int32 | int64 ](v T) bool {
	switch (interface{})(v).(type) {
	case float32:
		return false
	case float64:
		return true
	case int64:
		return true
	}
	return false
}
