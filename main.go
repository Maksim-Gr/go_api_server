package main

import (
	"fmt"
	"go_api_server/coffee"
)

func main() {
	fmt.Println("Listing all available coffees")
	coffees, err := coffee.GetCoffees()
	if err != nil {
		fmt.Println("Error getting coffee list", err)
		return
	}
	for _, element := range coffees.List {
		result := fmt.Sprintf("%s for %v", element.Name, element.Price)
		fmt.Println(result)
	}

	fmt.Println("Is Mocha available", coffee.IsCoffeeAvailable("Mocha"))
}
