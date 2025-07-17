package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BaseConverter struct {
	Celsius float64
}

func (t BaseConverter) ToKelvin() float64 {
	return t.Celsius + 273.15
}

func (t BaseConverter) ToFahrenheit() float64 {
	return t.Celsius*1.8 + 32
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please, enter temperature in °C: ")
		tempInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("There is error in string reading, try again")
			return
		}
		celsius, err := strconv.ParseFloat(strings.TrimSpace(tempInput), 64)
		if err != nil {
			fmt.Println("Error: enter correct number(it must be number of type float)")
			continue
		}

		converter := BaseConverter{
			Celsius: celsius}

		for {
			fmt.Print("Convert in : K-kelvine, F-farenheit or Q-exit [K/F/Q]: ")
			choose, err := reader.ReadString('\n')
			if err != nil {
				fmt.Print("There is error in string reading, try again")
				return
			}
			choose = strings.TrimSpace(strings.ToUpper(choose))

			switch choose {
			case "K":
				fmt.Printf("Result: %.3fK\n", converter.ToKelvin())
				return
			case "F":
				fmt.Printf("Result: %.3f°F\n", converter.ToFahrenheit())
				return
			case "Q":
				fmt.Println("Exit")
				return
			default:
				fmt.Println("Error: choose K, F or Q")
			}
		}
	}
}
