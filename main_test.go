package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_countFlourGlutenRatio(t *testing.T) {
	tests := []struct {
		input          input
		expectedFlour  float64
		expectedGluten float64
	}{
		{
			input{10, 78, 13.5, 600},
			569,
			31,
		},
		{
			input{10, 75, 12.5, 800},
			769,
			31,
		},
		{
			input{9, 80, 13.5, 500},
			468,
			32,
		},
		{
			input{10, 75, 14, 700},
			657,
			43,
		},
		{
			input{8.5, 66, 13.8, 568},
			516,
			52,
		},
	}
	for i, tt := range tests {
		t.Run("testcase "+strconv.Itoa(i+1), func(t *testing.T) {
			flour, gluten := countFlourGlutenRatio(tt.input)
			if flour != tt.expectedFlour {
				t.Errorf("countFlourGlutenRatio() got flour = %v, expectedFlour %v", flour, tt.expectedFlour)
			}
			if gluten != tt.expectedGluten {
				t.Errorf("countFlourGlutenRatio() got gluten = %v, expectedGluten %v", gluten, tt.expectedGluten)
			}
		})
	}
}

func Test_verifyInput(t *testing.T) {
	tests := []struct {
		name string
		args input
		want []string
	}{
		{name: "all fields empty",
			args: input{},
			want: []string{"'flourProteinContent' flag not specified.", "'targetProteinContent' flag not specified.",
				"'glutenProteinContent' flag not specified.", "'targetFlourWeight' flag not specified.",
				"Check --help for details"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyInput(tt.args.flourProteinContent, tt.args.glutenProteinContent, tt.args.targetProteinContent, tt.args.targetFlourWeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verifyInput() = %v, expectedMsgs %v", got, tt.want)
			}
		})
	}
}
