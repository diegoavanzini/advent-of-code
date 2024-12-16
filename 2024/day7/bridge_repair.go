package day7

import (
	"fmt"
	"strconv"
)

func getTotalCalibrationResult(calibrations []CalibrationResult, concatStrategy bool) int {
	var result int
	for _, calibration := range calibrations {
		if _, ok := NewNode(calibration.Calibrations, 0, 0, calibration.Result, concatStrategy); ok {
			result += calibration.Result
		}
	}
	return result
}

func NewNode(i []int, cumulativeValue int, prevValue int, result int, concatStrategy bool) (*CalibrationNode, bool) {
	if len(i) == 0 {
		return nil, cumulativeValue == result
	}
	currentValue := i[0]
	addNode, res := NewAddNode(i[1:], cumulativeValue, currentValue, result, concatStrategy)
	if res {
		return addNode, true
	}
	mulNode, res := NewMultiplyNode(i[1:], cumulativeValue, currentValue, result, concatStrategy)
	if res {
		return mulNode, true
	}
	node := &CalibrationNode{
		Value:        currentValue,
		AddNode:      addNode,
		MultiplyNode: mulNode,
	}
	if concatStrategy {
		concatNode, res := NewConcatNode(i[1:], cumulativeValue, currentValue, result)
		if res {
			return concatNode, true
		}
		node.ConcatNode = concatNode
	}
	return node, currentValue == result
}

func NewConcatNode(i []int, prevValue, currentValue, result int) (*CalibrationNode, bool) {
	v, err := strconv.Atoi(fmt.Sprintf("%d%d", prevValue, currentValue))
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return NewNode(i, v, currentValue, result, true)
}

func NewAddNode(i []int, cumulativeValue, currentValue, result int, concatStrategy bool) (*CalibrationNode, bool) {
	return NewNode(i, currentValue+cumulativeValue, currentValue, result, concatStrategy)
}

func NewMultiplyNode(i []int, cumulativeValue, currentValue, result int, concatStrategy bool) (*CalibrationNode, bool) {
	return NewNode(i, currentValue*cumulativeValue, currentValue, result, concatStrategy)
}

type CalibrationResult struct {
	Result       int
	Calibrations []int
}

type CalibrationNode struct {
	Value        int
	AddNode      *CalibrationNode
	MultiplyNode *CalibrationNode
	ConcatNode   *CalibrationNode
}
