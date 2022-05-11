package homework

func Reverse(input []int64) (result []int64) {
	//Place your code here
	var reverseSlice []int64
	for i := len(input) - 1; i >= 0; i-- {
		reverseSlice = append(reverseSlice, input[i])
	}
	return
}
