package main

import "fmt"

// Константы конвертации
const usdToEur = 0.86 // 1 USD = 0.86 EUR
const usdToRub = 80.0 // 1 USD = 80 RUB

// Константа для конвертации EUR в RUB на основе первых двух
const eurToRub = usdToRub / usdToEur // 1 EUR = usdToRub / usdToRub RUB

func main() {
	fmt.Println("💰 Калькулятор конвертации EUR в RUB")
	fmt.Println("=====================================")

	var eurAmount float64
	fmt.Print("Введите сумму в EUR: ")
	_, err := fmt.Scan(&eurAmount)

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Конвертируем используя константу eurToRub
	rubAmount := eurAmount * eurToRub

	fmt.Printf("\nРезультат конвертации:\n")
	fmt.Printf("EUR: %.2f\n", eurAmount)
	fmt.Printf("RUB: %.2f\n", rubAmount)
	fmt.Printf("Курс: 1 EUR = %.2f RUB\n", eurToRub)
	fmt.Printf("Расчет курса: usdToRub / usdToEur = %.2f / %.2f = %.2f\n",
		usdToRub, usdToEur, eurToRub)
}
