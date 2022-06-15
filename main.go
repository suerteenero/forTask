package main

import "fmt"

// Общая информация

var (
	appleCost     float64 = 5.99
	pearCost      int     = 7
	moneyInPocket int     = 23
)

func main() {
	//Представление того, что имеем
	fmt.Printf("Apple cost %v \n", appleCost)
	fmt.Printf("Pear cost %v \n", pearCost)
	fmt.Printf("In my pocket I have %v dollars \n", moneyInPocket)

	// Необходимые суммы и количество
	totalCost := (appleCost * float64(9)) + (float64(pearCost) * 8)
	numberOfPear := moneyInPocket / pearCost
	numberOfApple := int(float64(moneyInPocket) / appleCost)
	sumOfTwo := (float64(pearCost) + appleCost) * 2

	// Вывод возможностей приобретения
	fmt.Printf("1. For 9 apples and 8 pears I need %v UAH\n", totalCost)
	fmt.Printf("2. How many pears I can afford? - %v \n", numberOfPear)
	fmt.Printf("3. How many apples I can afford? - %v \n", numberOfApple)
	fmt.Printf("4. If I going buy 2 apples and 2 pears I need %v UAH, but I have only %d UAH\n", sumOfTwo, moneyInPocket)

}
