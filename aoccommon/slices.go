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
