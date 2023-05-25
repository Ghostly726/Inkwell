package inkwell

func RemoveElement(arr []string, index int) []string {
	if index < 0 || index >= len(arr) {
		return arr
	}

	return append(arr[:index], arr[index+1:]...)
}
