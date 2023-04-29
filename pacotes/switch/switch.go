package main

import "fmt"

func retornDia(dia int) string {
	switch dia {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "nenhum valor corresponde a um dia da semana"
	}
}

func main() {

	var result = retornDia(3);
	fmt.Println(result);
}
