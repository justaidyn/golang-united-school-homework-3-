package homework

func average(input [15]float32) (result float32) {
	var total float32
	for _, number := range input {
		total = total + number
	}
	result = total / 15
	return
}
