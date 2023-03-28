# vital-gluten-flour-calculator [![Test](https://github.com/adikm/vital-gluten-flour-calculator/actions/workflows/test.yml/badge.svg)](https://github.com/adikm/vital-gluten-flour-calculator/actions/workflows/test.yml)

Vital Wheat Gluten Flour Calculator

For you all software engineers and sourdough bakers out there, a small utility written in Golang calculating the ratio
of vital wheat gluten you need to use in your mix in order to achieve desired protein level.
A recipe calls for 500 grams of 14% bread flour, but all you've got is 10% all-purpose and some pure gluten?
No worries, it'll tell you exactly how much of what you need to use in order to achieve perfect result!

### The formula

It makes use of the following formula, all variables in grams:
```text
((targetProteinContent - glutenProteinContent) / (flourProteinContent - glutenProteinContent)) * targetFlourWeight
```

For example:
`((13.5 - 78) / (10.0-78)) * 600`
gives you a mix ~569g of flour and ~31g of vital wheat gluten that you need to mix together, so you get a mixture of 13.5%
protein content

### Usage

If you have already Go installed on your machine, simply run the following command:

```shell
go run . -flourProteinContent 10 -glutenProteinContent 78 -targetProteinContent 13.5 -targetFlourWeight 600
```

You should get the result similar to this:

```text
In order to reach the desired protein content of 14% in total weight of mixture 600g
you need to mix 569g of flour and 31g of vital wheat gluten
```