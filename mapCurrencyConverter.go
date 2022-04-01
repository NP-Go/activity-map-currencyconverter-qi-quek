package main

import (
	"fmt"
	"strings"
)

//Declare Stuct here

type exchRateEuro struct { //my struct type is exchRateEuro, not struct itself
	currencyName string
	exchRate     float32
}

//Declare Map here

func main() {
	USD := exchRateEuro{
		currencyName: "US Dollar",
		exchRate:     1.318,
	}
	JPY := exchRateEuro{
		currencyName: "Japanese Yen",
		exchRate:     121.05,
	}
	GBP := exchRateEuro{
		currencyName: "Pound Sterlign",
		exchRate:     0.90630,
	}
	CNY := exchRateEuro{
		currencyName: "Chinese yuan renminbi",
		exchRate:     7.9944,
	}
	SGD := exchRateEuro{
		currencyName: "Singapore dollar",
		exchRate:     1.5743,
	}
	MYR := exchRateEuro{
		currencyName: "Malayisan ringgit",
		exchRate:     4.8390,
	}
	// mapStruct := make(map[string]exchRateEuro)

	mapStruct := map[string]exchRateEuro{ // struct is not a type, it is a keyword
		//to define my type
		"USD": USD,
		"JPY": JPY,
		"GBP": GBP,
		"CNY": CNY,
		"SGD": SGD,
		"MYR": MYR}

	for x, y := range mapStruct { //self note : for range in xx - to give index , for key, value in map
		// fmt.Println(x)
		// fmt.Println(y)
		fmt.Println(x, " : ", y.currencyName)
		fmt.Println(x, " : ", y.exchRate)
	}
	// fmt.Println(mapStruct['CNY'])

	var inputConv, outputConv string
	var amountCon float32

	fmt.Println("Input currency to convert:")
	fmt.Scanln(&inputConv)
	fmt.Println("Input amount to be converted:")
	fmt.Scan(&amountCon)
	// currMul := mapStruct[wantedConv].exchRate
	fmt.Println("Input currency wanted:")
	fmt.Scan(&outputConv)
	convertedVal := amountCon / mapStruct[strings.ToUpper(inputConv)].exchRate * mapStruct[strings.ToUpper(outputConv)].exchRate
	fmt.Printf("The converted value %v of %v to %v is %v", amountCon, inputConv, outputConv, convertedVal)

}
