func main() {
emoji := "perro"

	if emoji == "perro" {
		fmt.Println("es un perro")
	} else if emoji := "carita" {
		fmt.Println("no es un perro")
	} else {
		fmt.Println("este es el else")
	}
	
	emoji := "gato"

	if emoji == "gato" {
		fmt.Println("es un gato")
	} else if emoji := "carita" {
		fmt.Println("no es un perro")
	} else {
		fmt.Println("este es el else")
	}
	
	// SWITCH
	
	perro := 34
	switch perro {
	case perro == 34 || perro == "perro":     //se pueden anidar variables en el switch
		fmt.Println("si es")
	case 23:
		fmt.Println("no es")
	default:
		fmt.Println("default")
	}
for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	// for continuo similar a ciclo while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ 
	}
}
