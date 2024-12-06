package day5

import (
	"strconv"
	"strings"
)

type Printer struct {
	Rules []string
}

func NewPrinter(rules []string) *Printer {
	return &Printer{
		Rules: rules,
	}
}

func (p *Printer) PrintQueue(input [][]int, adjust bool) (sumMiddlePagesNumber int) {
	var notValidRows = map[int][]int{}
	var mappedRules = [][]string{}
	for _, rule := range p.Rules {
		mappedRules = append(mappedRules, strings.Split(rule, "|"))
	}
	// not valid
	p.validateInput(mappedRules, input, notValidRows)
	if adjust {
		for iRow, currentRow := range input {
			if _, ok := notValidRows[iRow]; ok {
				// fmt.Println(currentRow)
				lenIdx := len(currentRow)
				middle := (lenIdx / 2)
				middleValue := input[iRow][middle]
				sumMiddlePagesNumber = sumMiddlePagesNumber + middleValue
			}
		}
		return
	}
	for iRow, currentRow := range input {
		if _, ok := notValidRows[iRow]; ok {
			continue
		}
		lenIdx := len(currentRow)
		middle := (lenIdx / 2)
		middleValue := input[iRow][middle]
		sumMiddlePagesNumber = sumMiddlePagesNumber + middleValue
	}
	return
}

func (p *Printer) validateInput(mappedRules [][]string, input [][]int, notValidRows map[int][]int) {
	for iRule := 0; iRule < len(mappedRules); iRule++ {
		for iRow := 0; iRow < len(input); iRow++ {
			var n1iCol, n2iCol = -1, -1
			for iCol := 0; iCol < len(input[iRow]); iCol++ {
				currentValue := strconv.Itoa(input[iRow][iCol])
				if currentValue == mappedRules[iRule][0] {
					n1iCol = iCol
				}
				if currentValue == mappedRules[iRule][1] {
					n2iCol = iCol
				}
				if n1iCol != -1 &&
					n2iCol != -1 &&
					n1iCol > n2iCol {
					notValidRows[iRow] = addNotValid(input[iRow], n1iCol, n2iCol)
					iRule = 0
					iRow--
					break
				}
			}
		}
	}
}

func addNotValid(input []int, n1iCol int, n2iCol int) []int {
	notValid := input
	temp := notValid[n1iCol]
	notValid[n1iCol] = notValid[n2iCol]
	notValid[n2iCol] = temp
	return notValid
}
