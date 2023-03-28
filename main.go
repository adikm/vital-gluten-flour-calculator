package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
)

type input struct {
	flourProteinContent,
	glutenProteinContent,
	targetProteinContent,
	targetFlourWeight float64
}

func main() {
	var i input
	flag.Float64Var(&i.flourProteinContent, "flourProteinContent", 0.0, "your flour protein content, per 100g. For example: 10.3")
	flag.Float64Var(&i.glutenProteinContent, "glutenProteinContent", 0.0, "your gluten protein content, per 100g. For example: 78.5")
	flag.Float64Var(&i.targetProteinContent, "targetProteinContent", 0.0, "wanted flour protein content, per 100g. For example: 13.5")
	flag.Float64Var(&i.targetFlourWeight, "targetFlourWeight", 0.0, "target flour weight, for which needed amount of vital wheat gluten will be calculated, for example 600")
	flag.Parse()

	msgs := verifyInput(i.flourProteinContent, i.glutenProteinContent, i.targetProteinContent, i.targetFlourWeight)
	if len(msgs) > 0 {
		panic(strings.Join(msgs, " "))
	}

	flour, gluten := countFlourGlutenRatio(i)

	fmt.Printf("In order to reach the desired protein content of %.0f%% in total weight of mixture %.0fg\n", i.targetProteinContent, i.targetFlourWeight)
	fmt.Printf("you need to mix %.0fg of flour and %.0fg of vital wheat gluten\n", flour, gluten)
}

// countFlourGlutenRatio calculates needed flour and vital wheat gluten in order to achieve desired protein content in the mix
// it uses the following formula. All values are in grams.:
// ((targetProteinContent - glutenProteinContent) / (flourProteinContent - glutenProteinContent)) * targetFlourWeight
// for example: ((13.5 - 78) / (10.0-78)) * 600
// the result is 569g of flour and 31g of vital wheat gluten
func countFlourGlutenRatio(i input) (float64, float64) {
	flourPercentage := (i.targetProteinContent - i.glutenProteinContent) / (i.flourProteinContent - i.glutenProteinContent)
	flourWeight := i.targetFlourWeight * flourPercentage
	glutenWeight := i.targetFlourWeight - flourWeight
	return math.Round(flourWeight), math.Round(glutenWeight)
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
	} else if targetProteinContent < 0.0 {
		msgs = append(msgs, "'targetProteinContent' can't be less than 0.")
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

	if len(msgs) > 0 {
		msgs = append(msgs, "Check --help for details")
	}

	return msgs

}
