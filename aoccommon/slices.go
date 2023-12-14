package aoccommon

func Map[Ti, To any](slice []Ti, f func(Ti) To) []To {
	output := make([]To, len(slice))
	for idx := range slice {
		output[idx] = f(slice[idx])
	}
	return output
}

func Reduce[Ti, To any](slice []Ti, f func(Ti, To) To, init To) To {
	for idx := range slice {
		init = f(slice[idx], init)
	}
	return init
}

func Transpose[T any](slices2D [][]T) [][]T {
	transposed := make([][]T, 0)
	for range slices2D[0] {
		transposed = append(transposed, make([]T, len(slices2D)))
	}
	for i := range slices2D {
		for j, value := range slices2D[i] {
			transposed[j][i] = value
		}
	}
	return transposed
}

func RotateRight[T any](slice2D [][]T) [][]T {
	if len(slice2D) == 0 {
		return make([][]T, 0)
	}
	output := make([][]T, len(slice2D[0]))

	for i := range output {
		output[i] = make([]T, len(slice2D))
	}

	for i := range slice2D {
		for j := range slice2D[i] {
			output[j][len(output[0])-1-i] = slice2D[i][j]
		}
	}

	return output
}
