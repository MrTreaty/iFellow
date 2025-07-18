package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Time struct {
	Hours   int
	Minutes int
}

func (t Time) HourConvert() float64 {
	hConvert := float64(t.Hours%12)*30 + float64(t.Minutes)*0.5
	return hConvert
}

func (t Time) MinuteConvert() float64 {
	minConvert := float64(t.Minutes) * 6
	return minConvert
}

func (t Time) AngleBetween() float64 {
	hourAngle := t.HourConvert()
	minuteAngle := t.MinuteConvert()

	fmt.Printf("Hour arrow angle is %.2f, minute arrow angle is %.2f", hourAngle, minuteAngle)
	diff := math.Abs(hourAngle - minuteAngle)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var hours int
	for {
		fmt.Print("Enter hours (0-12): ")
		hourInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in string reading, try again")
			return
		}
		hourInput = strings.TrimSpace(hourInput)
		h, err := strconv.Atoi(hourInput)
		if err != nil || h < 0 || h > 12 {
			fmt.Println("Error! Enter integer number between 0 and 12")
			continue
		}
		hours = h
		break
	}

	var minutes int
	for {
		fmt.Print("Enter minutes (0-59): ")
		minuteInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in string reading, try again")
			return
		}
		minuteInput = strings.TrimSpace(minuteInput)
		m, err := strconv.Atoi(minuteInput)
		if err != nil || m < 0 || m > 59 {
			fmt.Println("Error! Enter integer number between 0 and 59")
			continue
		}
		minutes = m
		break
	}

	time := Time{Hours: hours, Minutes: minutes}
	angle := time.AngleBetween()

	fmt.Printf("\nTime is: %02d:%02d\n", time.Hours, time.Minutes)
	fmt.Printf("Angle between arrows: %.1f°\n", angle)
}
