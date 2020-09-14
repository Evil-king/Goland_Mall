package service

import (
	"fmt"
	"testing"
)
var gameUtils = &GameUtils{}
func TestGameUtils_GenerateRandomNumber(t *testing.T) {
	var gameUtils = &GameUtils{}
	result := gameUtils.RandomNumber("FT")
	//result := gameUtils.RandomNumber("PK10")
	//result := gameUtils.RandomNumber("Lottery")
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//num := r.Intn(6)+1
	//fmt.Println(num)
	fmt.Println(result)
}

func TestGameUtils_CalculationWiningResults(t *testing.T) {
	outNumber := gameUtils.RandomNumber("Lottery")
	fmt.Println("开出号码为:",outNumber)
	result :=gameUtils.CalculationWiningResults("Lottery",outNumber)
	fmt.Println(result)
}

func TestYearOfNumber(t *testing.T) {
	result :=gameUtils.YearOfNumber(2020)
	fmt.Println(result)
}

func TestJudgmentCurrentYearIsBronYear(t *testing.T) {
	result:= gameUtils.JudgmentCurrentYearIsBronYear()
	fmt.Println("result=",result)
}

func TestGameUtils_JudgeStrWhichChineseZodiac(t *testing.T) {
	result:= gameUtils.JudgeStrWhichChineseZodiac("")
	fmt.Println("result=",result)
}

func TestGameUtils_GetChineseZodiacValue(t *testing.T) {
	result:= gameUtils.GetChineseZodiacValue("马")
	fmt.Println("result=",result)
}

func TestGameUtils_GetChineseZodiacAllValue(t *testing.T) {
	result:= gameUtils.GetChineseZodiacAllValue()
	fmt.Println("result=",result)
}

func TestGameUtils_SliceToString(t *testing.T) {
	outNumber := gameUtils.RandomNumber("Lottery")
	fmt.Println("outNumber=",outNumber)
	result:=gameUtils.SliceToString(outNumber)
	fmt.Println("result=",result)
}

func TestGameUtils_OperatorWingData(t *testing.T) {
	outNumber := gameUtils.RandomNumber("SSC")
	result:=gameUtils.OperatorWingData("CI","SSC",outNumber)
	fmt.Println("result=",result)
}

type chineseZodiacValue struct {
	Name string
	Data string
}

func Test(t *testing.T) {

	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	//	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	//删除第i个元素
	//a = append(a[:40], a[40+1:]...)
	//for i:=0;i<49;i++{
	//	fmt.Println(a[12])
	//}
	//var str string = "01"

	//for _,value:=range mp{
	//	dynamic := make(map[string]string)
	//	json.Unmarshal([]byte(value), &dynamic)
	//	fmt.Println(dynamic)
	//	for objKey,objValue :=range dynamic{
	//		obj:=chineseZodiacValue{
	//			Name: objKey,
	//			Data: objValue,
	//		}
	//		fmt.Println(obj.Data)
	//	}
	//}

	//str := `{"鼠":"01,13,25,37,49","鸡":"04,16,28,40","猴":"05,17,29,41","兔":"10,22,34,46","狗":"03, 15, 27, 39","蛇":"08, 20, 32, 44","龙":"09, 21, 33, 45","猪":"02, 14, 26, 38","羊":"06, 18, 30, 42","牛":"12, 24, 36, 48","马":"07, 19, 31, 43","虎":"11, 23, 35, 47"}`
	//
	//dynamic := make(map[string]interface{})
	//json.Unmarshal([]byte(str), &dynamic)
	//
	//
}
