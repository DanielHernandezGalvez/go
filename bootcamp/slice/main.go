package main

import "fmt"

func main() {
	 len(): # de elementos en el slice
	cap(): # de elementos del array donde apunta el slice, a partir del Γ­ndice de donde se creΓ³ el slice
		food := [5]string{"π­", "π", "π", "π", "π"}
	   	fruits := food[1:3] // "π", "π"
	   	fruits = append(fruits, "π", "π", "π")

	   	fmt.Println("food", food)
	   	fmt.Println("fruits", fruits)
	   	fmt.Println("tamaΓ±o", len(fruits))
	   	fmt.Println("capacidad", cap(fruits)) 

	fruits := []string{"π", "π"}
	 fruits := make([]string, 0, 3)

	fruits = append(fruits, "π", "π", "π", "π")
	fmt.Println("fruits", fruits)
	fmt.Println("tamaΓ±o", len(fruits))
	fmt.Println("capacidad", cap(fruits)) 

	fruits := []string{}
	fmt.Println("fruits", fruits)
	fmt.Println("tamaΓ±o", len(fruits))
	fmt.Println("capacidad", cap(fruits))
	fmt.Println(fruits == nil)
}
