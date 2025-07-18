package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter slice length (it must be an integer value) =")

	val, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading string , try again")
		return
	}
	val = strings.TrimSpace(val)

	value, err := strconv.Atoi(val)
	if err != nil || value <= 0 {
		fmt.Println("Error:Bad value try again\n", err)
		return
	}

	flSlice := make([]float64, value)
	for i := range flSlice {
		flSlice[i] = rand.Float64()
	}

	min, max, avg := MinMaxAvg(flSlice)
	fmt.Printf("Slice is := %v\n", flSlice)
	fmt.Printf("Min= %f , Max= %f , Avg= %f", min, max, avg)
}

func MinMaxAvg(slice []float64) (min, max, avg float64) {
	if len(slice) == 0 {
		fmt.Printf("Empty slice")
		return
	}

	min, max = slice[0], slice[0]
	sum := 0.0

	for _, v := range slice {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}

	avg = sum / float64(len(slice))
	return min, max, avg
}
