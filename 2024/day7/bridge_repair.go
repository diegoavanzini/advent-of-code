package day7

func getTotalCalibrationResult(calibrations []CalibrationResult) int {
	var result int
	for _, calibration := range calibrations {
		if _, ok := NewNode(calibration.Calibrations, 0, calibration.Result); ok {
			result += calibration.Result
		}
	}
	return result
}

func NewNode(i []int, prevValue int, result int) (*CalibrationNode, bool) {
	if len(i) == 0 {
		return nil, prevValue == result
	}
	value := i[0]
	addNode, res := NewNode(i[1:], value+prevValue, result)
	if res {
		return addNode, true
	}
	mulNode, res := NewNode(i[1:], value*prevValue, result)
	if res {
		return mulNode, true
	}
	return &CalibrationNode{
		Value:        value,
		AddNode:      addNode,
		MultiplyNode: mulNode,
	}, value == result
}

type CalibrationResult struct {
	Result       int
	Calibrations []int
}

type CalibrationNode struct {
	Value        int
	AddNode      *CalibrationNode
	MultiplyNode *CalibrationNode
}
