package iteration


func Repeat(char string) string {
	var re string
	for i :=0; i < 5; i++ {
		re = re + char
	}
	return re
}
