package main

import (
	"fmt"
)

func main() {
	
	creditSum := 1000
	percent := 9
	limit := 2500
	month := 0
	
	for creditSum <= limit {
		
		fmt.Println("В месяце номер ", month)
		sumPercent := creditSum * percent / 100

		fmt.Println("Сумма процента составляет ", sumPercent)
		creditSum = creditSum + sumPercent

		month ++
	}
	
	fmt. Println("Общая сумма кредита:",creditSum, "\nМесячная процентная ставка: 9%\nlimit: 2500\nПревысило за: ", month, "месяцев.")
}