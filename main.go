package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, value := range reader {
		fmt.Printf("Veuillez saisir une réponse pour cette valeur %v ", value[0])
		var reponse string
		fmt.Scanf("%v \n", &reponse)
		fmt.Println(reponse)
		if reponse != value[1] {
			fmt.Printf("Votre réponse %v est différente de la réponse %v \n ", reponse, value[1])
		} else {
			fmt.Printf(" La question %v est bien égale à la réponse %v \n ", value[0], reponse)
		}
	}
}
