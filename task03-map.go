package homework

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var count int

	for i := len(input); i > 0; i-- {
		result = append(result, input[count])
		count++

	}
	return
}
