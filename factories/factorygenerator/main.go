package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/factories/factorygenerator/domain"
)

func main() {
	developerFactory := domain.NewEmployeeFactory("Developer", 60000)
	managerFactory := domain.NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Cristiano")
	fmt.Println(developer)

	manager := managerFactory("Lionel")
	fmt.Println(manager)

	bossFactory := domain.NewEmployeeFactory2("CEO", 100000)
	// Can modify its property
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("John")
	fmt.Println(boss)
}
