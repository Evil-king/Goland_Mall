package service

import (
	"fmt"
	"testing"
)

func TestDrawOperator(t *testing.T) {
	DrawOperator("AA")
}

func TestCreatePeriodNum(t *testing.T) {
	CreatePeriodNum("AA")
}

func TestGenerateRandomNumber(t *testing.T) {
	str := GenerateRandomNumber(1, 11, 10)
	fmt.Println(str)
}

func TestCalculationWiningResults(t *testing.T) {
	str := GenerateRandomNumber(1, 11, 10)
	fmt.Println(str)
	result := CalculationWiningResults(str)
	fmt.Println(result)
}
