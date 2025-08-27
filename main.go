// Домашнее задание номер 1
package main

import "fmt"

const usdToEur = 0.86
const usdToRub = 80

func main() {
	var sumEur float64
	fmt.Println("Введите сумму в евро: ")
	fmt.Scan(&sumEur)
	sumUsd := sumEur / usdToEur
	sumRub := sumUsd * usdToRub

	fmt.Println("Итоговая сумма в рублях: ", sumRub)
}
