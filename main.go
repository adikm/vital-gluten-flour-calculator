package main

import (
	"flag"
	"fmt"
)

func main() {
	var flourProteinContent float64
	var glutenProteinContent float64
	var targetProteinContent float64
	var targetFlourWeight float64
	flag.Float64Var(&flourProteinContent, "flourProteinContent", 0.0, "your flour protein content, per 100g. For example: 10.3")
	flag.Float64Var(&glutenProteinContent, "glutenProteinContent", 0.0, "your gluten protein content, per 100g. For example: 78.5")
	flag.Float64Var(&targetProteinContent, "targetProteinContent", 0.0, "wanted flour protein content, per 100g. For example: 13.5")
	flag.Float64Var(&targetFlourWeight, "targetFlourWeight", 0, "target flour weight, for which needed amount of vital wheat gluten will be calculated, for example 600")

	flag.Parse()
	if flourProteinContent == 0.0 {
		panic("'flourProteinContent' flag not specified. Check --help for details")
	}
	if targetProteinContent == 0.0 {
		panic("'targetProteinContent' flag not specified. Check --help for details")
	}
	if glutenProteinContent == 0.0 {
		panic("'glutenProteinContent' flag not specified. Check --help for details")
	}
	if targetFlourWeight == 0.0 {
		panic("'targetFlourWeight' flag not specified. Check --help for details")
	}
	//([Target protein content usersFlourPercentage] - [All-purpose flour protein content usersFlourPercentage]) x [Total flour weight] = [Amount of vital wheat gluten to substitute for all-purpose flour]
	usersFlourPercentage := (targetProteinContent - glutenProteinContent) / (flourProteinContent - glutenProteinContent)
	usersFlourGrams := targetFlourWeight * usersFlourPercentage

	fmt.Printf("You need to use %.0fg of your flour\n", usersFlourGrams)
	fmt.Printf("and %.0fg of vital gluten\n", targetFlourWeight-usersFlourGrams)
}
