package service

import (
	"fmt"
	"math/rand"
)

// lunh algorithm
func LunhAlgorithm(number string) bool {
	number = deleteWhiteSpacesInCreditCard(number)
	var  s1 = 0;
	var  s2 = 0;


	for i := len(number) - 1; i >= 0; i -= 2 {
		var d = int(number[i] - '0')
		s1 += d
	}
	for i := len(number) - 2; i >= 0; i -= 2 {
		var d = int(number[i] - '0')
		if d >= 5 {
			d = d * 2 - 9
		} else {
			d *= 2
		}
		s2 += d
	}
	return (s1 + s2) % 10 == 0

}

func deleteWhiteSpacesInCreditCard(creditCard string)string {
	var result string
	for _, char := range creditCard {
		if char != ' ' {
			result += string(char)
		}
	}
	return result
}

func AddSpacesToCreditCard(creditCard string) string {
	var result string

	for i, digit := range creditCard {
	
		if i > 0 && i%4 == 0 {
			result += " "
		}
		result += string(digit)
	}

	return result
}

func CreateFakeCreditCard() (string, error) {
	
	var creditCard string
	for i := 0; i < 16; i++ {
		creditCard += fmt.Sprintf("%d", rand.Intn(10))
	}

	if len(creditCard) != 16 {
		return "", fmt.Errorf("tamanho inválido do cartão de crédito")
	}


	creditCardWithSpaces := AddSpacesToCreditCard(creditCard)

	return creditCardWithSpaces, nil
}