package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var flourProteinContent float64
	var glutenProteinContent float64
	var targetProteinContent float64
	var targetFlourWeight float64
	flag.Float64Var(&flourProteinContent, "flourProteinContent", 0.0, "your flour protein content, per 100g. For example: 10.3")
	flag.Float64Var(&glutenProteinContent, "glutenProteinContent", 0.0, "your gluten protein content, per 100g. For example: 78.5")
	flag.Float64Var(&targetProteinContent, "targetProteinContent", 0.0, "wanted flour protein content, per 100g. For example: 13.5")
	flag.Float64Var(&targetFlourWeight, "targetFlourWeight", 0.0, "target flour weight, for which needed amount of vital wheat gluten will be calculated, for example 600")
	flag.Parse()

	msgs := verifyInput(flourProteinContent, glutenProteinContent, targetProteinContent, targetFlourWeight)

	if len(msgs) > 0 {
		panic(strings.Join(msgs, " "))
	}
	//([Target protein content usersFlourPercentage] - [All-purpose flour protein content usersFlourPercentage]) x [Total flour weight] = [Amount of vital wheat gluten to substitute for all-purpose flour]
	usersFlourPercentage := (targetProteinContent - glutenProteinContent) / (flourProteinContent - glutenProteinContent)
	usersFlourGrams := targetFlourWeight * usersFlourPercentage

	fmt.Printf("You need to use %.0fg of your flour\n", usersFlourGrams)
	fmt.Printf("and %.0fg of vital gluten\n", targetFlourWeight-usersFlourGrams)
}

// verifyInput checks all cmd flags and returns errorMessages if validation failed
func verifyInput(flourProteinContent, glutenProteinContent, targetProteinContent, targetFlourWeight float64) []string {
	var msgs []string

	if flourProteinContent == 0.0 {
		msgs = append(msgs, "'flourProteinContent' flag not specified.")
	} else if flourProteinContent < 0.0 {
		msgs = append(msgs, "'flourProteinContent' can't be less than 0.")
	}

	if targetProteinContent == 0.0 {
		msgs = append(msgs, "'targetProteinContent' flag not specified.")
	} else if flourProteinContent != 0.0 && targetProteinContent < flourProteinContent {
		msgs = append(msgs, "'targetProteinContent' must be bigger than your flour protein content.")
	}

	if glutenProteinContent == 0.0 {
		msgs = append(msgs, "'glutenProteinContent' flag not specified.")
	} else if glutenProteinContent < 0.0 {
		msgs = append(msgs, "'glutenProteinContent' can't be less than 0.")
	}

	if targetFlourWeight == 0.0 {
		msgs = append(msgs, "'targetFlourWeight' flag not specified.")
	} else if targetFlourWeight < 0.0 {
		msgs = append(msgs, "'targetFlourWeight' can't be less than 0.")
	}

	msgs = append(msgs, "Check --help for details")

	return msgs

}
