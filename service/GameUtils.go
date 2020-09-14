package service

import (
	"Game/common/vo"
	"Game/dao"
	"Game/model"
	"Game/utils"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const RED = "01、02、07、08、12、13、18、19、23、24、29、30、34、35、40、45、46"
const BLUE = "03、04、09、10、14、15、20、25、26、31、36、37、41、42、47、48"
const GREEN = "05、06、11、16、17、21、22、27、28、32、33、38、39、43、44、49"

const TIAN_XIAO = "牛、兔、龙、马、猴、猪、"
const DI_XIAO = "鼠、虎、蛇、羊、鸡、狗"
const QIAN_XIAO = "鼠、牛、虎、兔、龙、蛇"
const HOU_XIAO = "马、羊、猴、鸡、狗、猪"
const JIA_XIAO = "牛、马、羊、鸡、狗、猪"
const YE_XIAO = "鼠、虎、龙、蛇、兔、猴"

const SSC_DA = "5,6,7,8,9"
const SSC_XIAO = "0,1,2,3,4"
const SSC_DAN = "1,3,5,7,9"
const SSC_SHUANG = "0,2,4,6,8"
const SSC_ZHI = "1,2,3,5,7"
const SSC_HE = "0,4,6,8,9"
const BAOZI = "000,111,222,333,444,555,666,777,888,999"
const SHUNZI = "012,021,102,120,201,210,123,132,213,231,312,321,234,243,324,342,423,432,345,354,435,453,534,543,456,465,546,564,645,654" +
	"567,576,657,675,756,765,678,687,768,786,867,876,789,798,879,897,978,987,890,809,980,908,098,089,901,910,190,109,019,091"

var PC28_DA_DAN = []string{"15", "17", "19", "21", "23", "25", "27"}
var PC28_XIAO_DAN = []string{"1", "3", "5", "7", "9", "11", "13"}
var PC28_DA_SHUANG = []string{"14", "16", "18", "20", "22", "24", "26"}
var PC28_XIAO_SHUANG = []string{"0", "2", "4", "6", "8", "10", "12"}
var PC28_RED = []string{"3", "6", "9", "12", "15", "18", "21", "24"}
var PC28_BLUE = []string{"2", "5", "8", "11", "17", "20", "23", "26"}
var PC28_GREEN = []string{"1", "4", "7", "10", "16", "19", "22", "25"}
var PC28_HE = []string{"0", "13", "14", "27"}

const GOLD = "06,07,20,21,28,29,36,37"
const WOOD = "02,03,10,11,18,19,32,33,40,41,48,49"
const WATER = "08,09,16,17,24,25,38,39,46,47"
const FIRE = "04,05,12,13,26,27,34,35,42,43"
const EARTH = "01,14,15,22,23,30,31,44,45"

const format = "2006-01-02"

var mp = make(map[string]string)

var markSixDataService = MarkSixDataService{}
var zodiacOddsRelationService = ZodiacOddsRelationService{}
var gameBettingService = GameBettingService{}

func init() {
	mapData := markSixDataService.SelectAllYear()
	for _, value := range mapData {
		str := value.Year + "-" + value.Month + "-" + value.Day
		mp[str] = value.Data
	}
}

type IGameUtils interface {
	GenerateRandomNumber(start int, end int, count int) []int
	CalculationWiningResults(nums []int) []string
	OperatorWingData(gameCode string, modelCode string, outNumbers []int) []string
	IsEffectiveDateStr(nowTime string, sTime string, eTime string) bool
	RandomNumber(modelCode string) []string
	YearOfNumber(index int) map[string][]string
}

type GameUtils struct{}

//生成count个[start,end)结束的不重复的随机数
func (g *GameUtils) RandomNumber(modelCode string) []string {
	//存放结果的slice
	var pk10 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var lotterySix = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	var change5 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var ft = []int{1, 2, 3, 4}
	nums := make([]string, 0)
	if "PK10" == modelCode {
		var r = rand.New(rand.NewSource(time.Now().UnixNano()))
		//随机数生成器，加入时间戳保证每次生成的随机数不一样
		for len(nums) < len(pk10) {
			//生成随机数
			num := r.Intn(len(pk10)) + 1
			//查重
			exist := false
			for _, v := range nums {
				var str string
				if len(nums) > 0 {
					if len(v) > 1 && len(v) < 3 {
						str = v[1:2]
					}
					if len(v) == 3 {
						str = v[1:3]
					}
					if len(v) == 1 {
						str = v
					}
				}
				if str == strconv.Itoa(num) {
					exist = true
					break
				}
			}
			if !exist {
				if len(nums) == 0 {
					nums = append(nums, strconv.Itoa(num))
				} else {
					nums = append(nums, ","+strconv.Itoa(num))
				}
			}
		}
	}
	if "Lottery" == modelCode {
		var lotterySlice = lotterySix
		var r = rand.New(rand.NewSource(time.Now().UnixNano()))
		//随机数生成器，加入时间戳保证每次生成的随机数不一样
		for i := len(lotterySix); i > 42; i-- {
			var indexNum int
			//生成随机数
			num := r.Intn(i)
			indexNum = lotterySlice[num]
			var strNum string
			if indexNum < 10 {
				strNum = "0" + strconv.Itoa(indexNum)
			} else {
				strNum = strconv.Itoa(indexNum)
			}
			if len(nums) == 0 {
				nums = append(nums, strNum)
			} else {
				nums = append(nums, ","+strNum)
			}
			lotterySlice = append(lotterySlice[:num], lotterySlice[num+1:]...)
		}
	}
	if "SSC" == modelCode {
		for {
			var r = rand.New(rand.NewSource(time.Now().UnixNano()))
			num := r.Intn(9) + 1
			if len(nums) == 0 {
				nums = append(nums, strconv.Itoa(num))
			} else {
				nums = append(nums, ","+strconv.Itoa(num))
			}
			if len(nums) >= 5 {
				break
			}
		}
	}
	if "11X5" == modelCode {
		var change5Slice = change5
		var r = rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 11; i > 6; i-- {
			var indexNum int
			//生成随机数
			num := r.Intn(i)
			indexNum = change5Slice[num]
			var strNum string
			if indexNum < 10 {
				strNum = "0" + strconv.Itoa(indexNum)
			} else {
				strNum = strconv.Itoa(indexNum)
			}
			if len(nums) == 0 {
				nums = append(nums, strNum)
			} else {
				nums = append(nums, ","+strNum)
			}
			change5Slice = append(change5Slice[:num], change5Slice[num+1:]...)
		}
	}
	if "K3" == modelCode {
		for {
			var r = rand.New(rand.NewSource(time.Now().UnixNano()))
			num := r.Intn(6)+1
			if len(nums) == 0 {
				nums = append(nums, strconv.Itoa(num))
			} else {
				nums = append(nums, ","+strconv.Itoa(num))
			}
			if len(nums) >= 3 {
				break
			}
		}
	}
	if "PC28" == modelCode {
		for {
			var r = rand.New(rand.NewSource(time.Now().UnixNano()))
			num := r.Intn(10)
			if len(nums) == 0 {
				nums = append(nums, strconv.Itoa(num))
			} else {
				nums = append(nums, ","+strconv.Itoa(num))
			}
			if len(nums) >= 3 {
				break
			}
		}
	}
	if "FT" == modelCode {
		var ftSlice = ft
		var r = rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 4; i > 0; i-- {
			//生成随机数
			num := r.Intn(i)
			nums = append(nums, strconv.Itoa(num))
			if len(nums) >= 1 {
				break
			}
			ftSlice = append(ftSlice[:num], ftSlice[num+1:]...)
		}
	}
	return nums
}

//生成1-10的开奖结果
func (g *GameUtils) CalculationWiningResults(modelCode string, nums []string) []string {
	lists := make([]string, 0)
	if "PK10" == modelCode {
		var one, two, three, four, five, six, seven, eight, nine, ten, sum int
		one, _ = strconv.Atoi(nums[0])
		two, _ = strconv.Atoi(nums[1])
		three, _ = strconv.Atoi(nums[2])
		four, _ = strconv.Atoi(nums[3])
		five, _ = strconv.Atoi(nums[4])
		six, _ = strconv.Atoi(nums[5])
		seven, _ = strconv.Atoi(nums[6])
		eight, _ = strconv.Atoi(nums[7])
		nine, _ = strconv.Atoi(nums[8])
		ten, _ = strconv.Atoi(nums[9])
		sum = one + ten
		lists = append(lists, strconv.Itoa(sum))
		if sum > 11 {
			lists = append(lists, "大")
		} else if sum <= 11 {
			lists = append(lists, "小")
		}
		if sum%2 == 0 {
			lists = append(lists, "双")
		} else {
			lists = append(lists, "单")
		}
		//1～5龙虎
		if one > ten {
			lists = append(lists, "龙")
		} else {
			lists = append(lists, "虎")
		}
		if two > nine {
			lists = append(lists, "龙")
		} else {
			lists = append(lists, "虎")
		}
		if three > eight {
			lists = append(lists, "龙")
		} else {
			lists = append(lists, "虎")
		}
		if four > seven {
			lists = append(lists, "龙")
		} else {
			lists = append(lists, "虎")
		}
		if five > six {
			lists = append(lists, "龙")
		} else {
			lists = append(lists, "虎")
		}
	}
	if "Lottery" == modelCode {
		stringListMap := g.YearOfNumber(time.Now().Year())
		for _, numsValue := range StringSliceCut(nums) {
			for index, v := range stringListMap {
				for _, value := range StringSliceCut(v) {
					if strings.Contains(numsValue, value) {
						if strings.Contains(RED, value) {
							str := value + "-" + "red" + "-" + index
							lists = append(lists, str)
						} else if strings.Contains(BLUE, value) {
							str := value + "-" + "blue" + "-" + index
							lists = append(lists, str)
						} else {
							str := value + "-" + "green" + "-" + index
							lists = append(lists, str)
						}
					}
				}
			}
		}
	}
	if "SSC" == modelCode {
		var one, two, three, four, five,  sum int
		one, _ = strconv.Atoi(nums[0])
		two, _ = strconv.Atoi(nums[1])
		three, _ = strconv.Atoi(nums[2])
		four, _ = strconv.Atoi(nums[3])
		five, _ = strconv.Atoi(nums[4])
		sum  = one + two + three + four + five
		firstThree := strconv.Itoa(one) + strconv.Itoa(two) + strconv.Itoa(three)
		midThree := strconv.Itoa(two) + strconv.Itoa(three) + strconv.Itoa(four)
		lastThree := strconv.Itoa(three) + strconv.Itoa(four) + strconv.Itoa(five)
		lists = append(lists, strconv.Itoa(sum))
		if sum >= 23 && sum <= 45 {
			lists = append(lists, "大")
		} else if sum >= 0 && sum <= 22 {
			lists = append(lists, "小")
		}
		if sum % 2 == 0 {
			lists = append(lists, "双")
		} else if sum % 2 != 0{
			lists = append(lists, "单")
		}
		if one > five {
			lists = append(lists, "龙")
		} else if one < five{
			lists = append(lists, "虎")
		} else if one == five {
			lists = append(lists, "和")
		}
		if strings.Contains(BAOZI,firstThree){
			lists = append(lists, "豹子")
		} else if strings.Contains(SHUNZI,firstThree){
			lists = append(lists, "顺子")
		} else if (((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI,firstThree)){
			lists = append(lists, "对子")
		} else if ((math.Abs(float64(one - two)) == 1 || math.Abs(float64(two - three)) == 1 || math.Abs(float64(three - one)) == 1 ||
					math.Abs(float64(one - two)) == 9 || math.Abs(float64(two - three)) == 9 ||
					math.Abs(float64(three - one)) == 9) && !strings.Contains(SHUNZI,firstThree)  &&
		 !(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI,firstThree))){
			lists = append(lists, "半顺")
		} else {
			lists = append(lists, "杂六")
		}
		if strings.Contains(BAOZI,midThree){
			lists = append(lists, "豹子")
		} else if strings.Contains(SHUNZI,midThree){
			lists = append(lists, "顺子")
		} else if (((two == three || three == four || four == two) && (two != three || three != four)) && !strings.Contains(BAOZI,midThree)){
			lists = append(lists, "对子")
		} else if ((math.Abs(float64(one - two)) == 1 || math.Abs(float64(two - three)) == 1 ||
					math.Abs(float64(three - one)) == 1 ||
					math.Abs(float64(one - two)) == 9 || math.Abs(float64(two - three)) == 9 ||
					math.Abs(float64(three - one)) == 9) && !strings.Contains(SHUNZI,midThree)  &&
			!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI,midThree))){
			lists = append(lists, "半顺")
		} else {
			lists = append(lists, "杂六")
		}
		if strings.Contains(BAOZI,lastThree){
			lists = append(lists, "豹子")
		} else if strings.Contains(SHUNZI,lastThree){
			lists = append(lists, "顺子")
		} else if ((three == four || four == five || five == three) && (three != four || four != five) && !strings.Contains(BAOZI,lastThree)){
			lists = append(lists, "对子")
		} else if ((math.Abs(float64(one - two)) == 1 || math.Abs(float64(two - three)) == 1 ||
					math.Abs(float64(three - one)) == 1 ||
					math.Abs(float64(one - two)) == 9 || math.Abs(float64(two - three)) == 9 ||
					math.Abs(float64(three - one)) == 9) && !strings.Contains(SHUNZI,lastThree)  &&
			!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI,lastThree))){
			lists = append(lists, "半顺")
		} else {
			lists = append(lists, "杂六")
		}
	}
	if "11X5" == modelCode {
		var one, two, three, four, five,  sum int
		one, _ = strconv.Atoi(nums[0])
		two, _ = strconv.Atoi(nums[1])
		three, _ = strconv.Atoi(nums[2])
		four, _ = strconv.Atoi(nums[3])
		five, _ = strconv.Atoi(nums[4])
		sum  = one + two + three + four + five
		digits := sum % 10
		if sum > 30 {
			lists = append(lists, "和大")
		} else if sum < 30 {
			lists = append(lists, "和小")
		}
		if sum % 2 == 0 {
			lists = append(lists, "和双")
		} else if sum % 2 !=0 {
			lists = append(lists, "和单")
		}
		if one > five {
			lists = append(lists, "龙")
		} else if one < five {
			lists = append(lists, "虎")
		}
		if digits >= 5{
			lists = append(lists, "尾大")
		} else if digits <= 4 {
			lists = append(lists, "尾小")
		}

	}
	if "K3" == modelCode {
		nums := StringSliceCut(nums)
		var one, two, three, sum int
		one, _ = strconv.Atoi(nums[0])
		two, _ = strconv.Atoi(nums[1])
		three, _ = strconv.Atoi(nums[2])
		oneStr := strconv.Itoa(one)
		twoStr := strconv.Itoa(two)
		threeStr := strconv.Itoa(three)
		flag := strings.Contains(oneStr, twoStr) && strings.Contains(twoStr, threeStr) && strings.Contains(threeStr, oneStr)
		sum = one + two + three
		if flag {
			lists = append(lists, "围骰")
		} else {
			lists = append(lists, strconv.Itoa(sum))
			if 11 <= sum && sum <= 17 {
				lists = append(lists, "大")
			} else if 4 <= sum && sum <= 10 {
				lists = append(lists, "小")
			}
			if sum % 2 == 0 {
				lists = append(lists, "双")
			} else if sum % 2 != 0 {
				lists = append(lists, "单")
			}
		}
	}
	if "PC28" == modelCode {
		nums := StringSliceCut(nums)
		var one, two, three, sum int
		one, _ = strconv.Atoi(nums[0])
		two, _ = strconv.Atoi(nums[1])
		three, _ = strconv.Atoi(nums[2])
		sum  = one + two + three
		lists = append(lists, nums[0])
		lists = append(lists, nums[1])
		lists = append(lists, nums[2])
		lists = append(lists, strconv.Itoa(sum))
	}
	if "FT" == modelCode {
		lists = append(lists, nums[0])
	}
	return StringJointSlice(lists)
}

//组装中奖结果
func (g *GameUtils) OperatorWingData(gameCode string, modelCode string, outNumbers []string) []string {
	lists := make([]string, 0)
	var outNumbersStr = StringSliceCut(outNumbers)
	log.Println("outNumbersStr=",outNumbersStr)
	if "PK10" == modelCode {
		var one, two, three, four, five, six, seven, eight, nine, ten int
		one, _ = strconv.Atoi(outNumbersStr[0])
		two, _ = strconv.Atoi(outNumbersStr[1])
		three, _ = strconv.Atoi(outNumbersStr[2])
		four, _ = strconv.Atoi(outNumbersStr[3])
		five, _ = strconv.Atoi(outNumbersStr[4])
		six, _ = strconv.Atoi(outNumbersStr[5])
		seven, _ = strconv.Atoi(outNumbersStr[6])
		eight, _ = strconv.Atoi(outNumbersStr[7])
		nine, _ = strconv.Atoi(outNumbersStr[8])
		ten, _ = strconv.Atoi(outNumbersStr[9])
		gameInfoInnerAggregationList := dao.SelectGameInfoByGameCode(gameCode, modelCode)
		for i := 0; i < len(gameInfoInnerAggregationList); i++ {
			gameInfoInnerAggregation := gameInfoInnerAggregationList[i]
			//选号
			if "选号" == gameInfoInnerAggregation.GroupName {
				if "冠军" == gameInfoInnerAggregation.MethodName && strconv.Itoa(one) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "亚军" == gameInfoInnerAggregation.MethodName && strconv.Itoa(two) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第三名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(three) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第四名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(four) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第五名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(five) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第六名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(six) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第七名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(seven) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第八名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(eight) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第九名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(nine) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if "第十名" == gameInfoInnerAggregation.MethodName && strconv.Itoa(ten) == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
			}
			//双面 冠军
			if "双面" == gameInfoInnerAggregation.GroupName {
				//双面 冠军
				if "冠军" == gameInfoInnerAggregation.MethodName {
					if one%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if one%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if one > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if one < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if one > ten && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if one < ten && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				//双面 亚军
				if "亚军" == gameInfoInnerAggregation.MethodName {
					if two%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if two%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if two > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if two < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if two > nine && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if two < nine && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				//双面 第三名
				if "第三名" == gameInfoInnerAggregation.MethodName {
					if three%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if three%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if three > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if three < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if three > eight && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if three < eight && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				//双面 第四名
				if "第四名" == gameInfoInnerAggregation.MethodName {
					if four%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if four%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if four > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if four < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if four > seven && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if four < seven && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				//双面 第五名
				if "第五名" == gameInfoInnerAggregation.MethodName {
					if five%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if five%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if five > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if five < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if five > six && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if five < six && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				//双面 第6-10名
				if "第六名" == gameInfoInnerAggregation.MethodName {
					if six%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if six%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if six > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if six < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第七名" == gameInfoInnerAggregation.MethodName {
					if seven%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if seven%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if seven > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if seven < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第八名" == gameInfoInnerAggregation.MethodName {
					if eight%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if eight%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if eight > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if eight < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第九名" == gameInfoInnerAggregation.MethodName {
					if nine%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if nine%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if nine > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if nine < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第十名" == gameInfoInnerAggregation.MethodName {
					if ten%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if ten%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if ten > 5 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if ten < 5 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				//双面-冠亚和值 冠亚军和
				if "冠亚军和" == gameInfoInnerAggregation.MethodName {
					if (one+two)%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if (one+two)%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one+two) > 11 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					} else if (one+two) < 11 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			//冠亚和值-冠亚和-和值
			if "冠亚和值" == gameInfoInnerAggregation.GroupName {
				sum := one + two
				bettingNameInt,_ := strconv.Atoi(gameInfoInnerAggregation.BettingName)
				if "冠亚军和" == gameInfoInnerAggregation.MethodName && sum == bettingNameInt {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+gameInfoInnerAggregation.GroupName+"&"+gameInfoInnerAggregation.MethodName+"&"+strconv.Itoa(one+two))
				}
			}
		}
	}
	if "Lottery" == modelCode {
		gameInfoInnerAggregationList := dao.SelectGameInfoByGameCode(gameCode, modelCode)
		for i := 0; i < len(gameInfoInnerAggregationList); i++ {
			gameInfoInnerAggregation := gameInfoInnerAggregationList[i]
			if "总和" == gameInfoInnerAggregation.GroupName {
				if "双面" == gameInfoInnerAggregation.MethodName {
					one, _ := strconv.Atoi(outNumbers[0])
					two, _ := strconv.Atoi(outNumbers[1])
					three, _ := strconv.Atoi(outNumbers[2])
					four, _ := strconv.Atoi(outNumbers[3])
					five, _ := strconv.Atoi(outNumbers[4])
					six, _ := strconv.Atoi(outNumbers[5])
					totalSum := one + two + three + four + five + six
					if totalSum > 175 && "大" == gameInfoInnerAggregation.MethodName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					} else if totalSum == 175 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else if totalSum < 175 && "小" == gameInfoInnerAggregation.MethodName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					} else if totalSum%2 == 0 && "双" == gameInfoInnerAggregation.MethodName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					} else if totalSum%2 != 0 && "单" == gameInfoInnerAggregation.MethodName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "特码" == gameInfoInnerAggregation.GroupName {
				if "双面" == gameInfoInnerAggregation.MethodName {
					luckly, _ := strconv.Atoi(outNumbers[6])
					tenDigits := luckly / 10
					digits := luckly % 10
					sum := tenDigits + digits
					if luckly == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else {
						if luckly >= 25 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if luckly < 25 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 7 && "合大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 7 && "合小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 5 && "尾小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "合双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "合单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if luckly >= 25 {
							signNums := "25、27、29、31、33、35、37、39、41、43、45、47"
							doubleNums := "26、28、30、32、34、36、38、40、42、44、46、48"
							if strings.Contains(signNums, strconv.Itoa(luckly)) && "大单" == gameInfoInnerAggregation.BettingName {
								lists = append(lists, gameInfoInnerAggregation.Id+"&"+
									gameInfoInnerAggregation.GroupName+"&"+
									gameInfoInnerAggregation.MethodName+"&"+
									gameInfoInnerAggregation.BettingName)
							}
							if strings.Contains(doubleNums, strconv.Itoa(luckly)) && "大双" == gameInfoInnerAggregation.BettingName {
								lists = append(lists, gameInfoInnerAggregation.Id+"&"+
									gameInfoInnerAggregation.GroupName+"&"+
									gameInfoInnerAggregation.MethodName+"&"+
									gameInfoInnerAggregation.BettingName)
							}
						}
						if luckly <= 24 {
							doubleNums := "02、04、06、08、10、12、14、16、18、20、22、24"
							signNums := "01、03、05、07、09、11、13、15、17、19、21、23"
							if strings.Contains(signNums, strconv.Itoa(luckly)) && "小单" == gameInfoInnerAggregation.BettingName {
								lists = append(lists, gameInfoInnerAggregation.Id+"&"+
									gameInfoInnerAggregation.GroupName+"&"+
									gameInfoInnerAggregation.MethodName+"&"+
									gameInfoInnerAggregation.BettingName)
							}
							if strings.Contains(doubleNums, strconv.Itoa(luckly)) && "小双" == gameInfoInnerAggregation.BettingName {
								lists = append(lists, gameInfoInnerAggregation.Id+"&"+
									gameInfoInnerAggregation.GroupName+"&"+
									gameInfoInnerAggregation.MethodName+"&"+
									gameInfoInnerAggregation.BettingName)
							}
						}
						if strings.Contains(g.QueryYearByMarkSixData("TIAN_XIAO"), strconv.Itoa(luckly)) && "天肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains(g.QueryYearByMarkSixData("DI_XIAO"), strconv.Itoa(luckly)) && "地肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains(g.QueryYearByMarkSixData("QIAN_XIAO"), strconv.Itoa(luckly)) && "前肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains(g.QueryYearByMarkSixData("HOU_XIAO"), strconv.Itoa(luckly)) && "后肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains(g.QueryYearByMarkSixData("JIA_XIAO"), strconv.Itoa(luckly)) && "家肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains(g.QueryYearByMarkSixData("YE_XIAO"), strconv.Itoa(luckly)) && "野肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "选号" == gameInfoInnerAggregation.MethodName {
					for j := 1; j < 50; j++ {
						var result string
						if j < 10 {
							result = "0" + strconv.Itoa(j)
						} else {
							result = strconv.Itoa(j)
						}
						intResult, _ := strconv.Atoi(result)
						six, _ := strconv.Atoi(outNumbers[6])
						if intResult == six && result == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+result)
						}
					}
				}
				if "五行" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(GOLD, outNumbers[6]) && "金" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(WOOD, outNumbers[6]) && "木" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(WATER, outNumbers[6]) && "水" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(FIRE, outNumbers[6]) && "火" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(EARTH, outNumbers[6]) && "土" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "正码" == gameInfoInnerAggregation.GroupName {
				if "双面-正码-1" == gameInfoInnerAggregation.MethodName {
					one, _ := strconv.Atoi(outNumbers[0])
					tenDigits := one / 10
					digits := one % 10
					sum := tenDigits + digits
					if one == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else {
						if one >= 25 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if one < 25 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 7 && "合大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 7 && "合小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 5 && "尾小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "合双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "合单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "双面-正码-2" == gameInfoInnerAggregation.MethodName {
					two, _ := strconv.Atoi(outNumbers[1])
					tenDigits := two / 10
					digits := two % 10
					sum := tenDigits + digits
					if two == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else {
						if two >= 25 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if two < 25 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 7 && "合大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 7 && "合小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 5 && "尾小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "合双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "合单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "双面-正码-3" == gameInfoInnerAggregation.MethodName {
					three, _ := strconv.Atoi(outNumbers[2])
					tenDigits := three / 10
					digits := three % 10
					sum := tenDigits + digits
					if three == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else {
						if three >= 25 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if three < 25 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 7 && "合大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 7 && "合小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 5 && "尾小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "合双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "合单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "双面-正码-4" == gameInfoInnerAggregation.MethodName {
					four, _ := strconv.Atoi(outNumbers[3])
					tenDigits := four / 10
					digits := four % 10
					sum := tenDigits + digits
					if four == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else {
						if four >= 25 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if four < 25 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 7 && "合大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 7 && "合小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 5 && "尾小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "合双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "合单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "双面-正码-5" == gameInfoInnerAggregation.MethodName {
					five, _ := strconv.Atoi(outNumbers[4])
					tenDigits := five / 10
					digits := five % 10
					sum := tenDigits + digits
					if five == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else {
						if five >= 25 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if five < 25 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 7 && "合大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 7 && "合小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 5 && "尾小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "合双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "合单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "双面-正码-6" == gameInfoInnerAggregation.MethodName {
					six, _ := strconv.Atoi(outNumbers[5])
					tenDigits := six / 10
					digits := six % 10
					sum := tenDigits + digits
					if six == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					} else {
						if six >= 25 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if six < 25 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 7 && "合大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 7 && "合小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum < 5 && "尾小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "合双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "合单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "任选一" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < len(outNumbers); i++ {
						for j := 1; j < 50; j++ {
							var result string
							if j < 10 {
								result = "0" + strconv.Itoa(j)
							} else {
								result = strconv.Itoa(j)
							}
							if strings.Contains(outNumbers[i], result) && strings.Contains(outNumbers[i], gameInfoInnerAggregation.BettingName) {
								lists = append(lists, gameInfoInnerAggregation.Id+"&"+
									gameInfoInnerAggregation.GroupName+"&"+
									gameInfoInnerAggregation.MethodName+"&"+outNumbers[i])
							}
						}
					}
				}
				if "正一特" == gameInfoInnerAggregation.MethodName {
					for i := 1; i < 50; i++ {
						var result string
						if i < 10 {
							result = "0" + strconv.Itoa(i)
						} else {
							result = strconv.Itoa(i)
						}
						intResult, _ := strconv.Atoi(result)
						intOne, _ := strconv.Atoi(outNumbers[0])
						if intResult == intOne && strings.Contains(result, gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "正二特" == gameInfoInnerAggregation.MethodName {
					for i := 1; i < 50; i++ {
						var result string
						if i < 10 {
							result = "0" + strconv.Itoa(i)
						} else {
							result = strconv.Itoa(i)
						}
						intResult, _ := strconv.Atoi(result)
						intOne, _ := strconv.Atoi(outNumbers[1])
						if intResult == intOne && strings.Contains(result, gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "正三特" == gameInfoInnerAggregation.MethodName {
					for i := 1; i < 50; i++ {
						var result string
						if i < 10 {
							result = "0" + strconv.Itoa(i)
						} else {
							result = strconv.Itoa(i)
						}
						intResult, _ := strconv.Atoi(result)
						intOne, _ := strconv.Atoi(outNumbers[2])
						if intResult == intOne && strings.Contains(result, gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "正四特" == gameInfoInnerAggregation.MethodName {
					for i := 1; i < 50; i++ {
						var result string
						if i < 10 {
							result = "0" + strconv.Itoa(i)
						} else {
							result = strconv.Itoa(i)
						}
						intResult, _ := strconv.Atoi(result)
						intOne, _ := strconv.Atoi(outNumbers[3])
						if intResult == intOne && strings.Contains(result, gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "正五特" == gameInfoInnerAggregation.MethodName {
					for i := 1; i < 50; i++ {
						var result string
						if i < 10 {
							result = "0" + strconv.Itoa(i)
						} else {
							result = strconv.Itoa(i)
						}
						intResult, _ := strconv.Atoi(result)
						intOne, _ := strconv.Atoi(outNumbers[4])
						if intResult == intOne && strings.Contains(result, gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "正六特" == gameInfoInnerAggregation.MethodName {
					for i := 1; i < 50; i++ {
						var result string
						if i < 10 {
							result = "0" + strconv.Itoa(i)
						} else {
							result = strconv.Itoa(i)
						}
						intResult, _ := strconv.Atoi(result)
						intOne, _ := strconv.Atoi(outNumbers[5])
						if intResult == intOne && strings.Contains(result, gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
			}
			if "连码" == gameInfoInnerAggregation.GroupName {
				if "四全中" == gameInfoInnerAggregation.MethodName && "四全中" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
				if "三全中" == gameInfoInnerAggregation.MethodName && "三全中" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
				if "二全中" == gameInfoInnerAggregation.MethodName && "二全中" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
				if "三中二" == gameInfoInnerAggregation.MethodName && "中三" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
				if "三中二" == gameInfoInnerAggregation.MethodName && "中二" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbers))
				}
				if "二中特" == gameInfoInnerAggregation.MethodName && "中二" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
				if "二中特" == gameInfoInnerAggregation.MethodName && "中特" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
				if "特串" == gameInfoInnerAggregation.MethodName && "特串" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
			}
			if "生肖" == gameInfoInnerAggregation.GroupName {
				if "特肖" == gameInfoInnerAggregation.MethodName {
					six := outNumbersStr[6]
					if strings.Contains(g.GetChineseZodiacValue("鼠"), six) && "鼠" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("牛"), six) && "牛" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("虎"), six) && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("兔"), six) && "兔" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("龙"), six) && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("蛇"), six) && "蛇" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("马"), six) && "马" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("羊"), six) && "羊" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("猴"), six) && "猴" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("鸡"), six) && "鸡" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("狗"), six) && "狗" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(g.GetChineseZodiacValue("猪"), six) && "猪" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "合肖" == gameInfoInnerAggregation.MethodName {
					luckly, _ := strconv.Atoi(outNumbersStr[6])
					if luckly == 49 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"[\"和局\"]")
					} else {
						if "2合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "3合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "4合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "5合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "6合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "7合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "8合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "9合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "10合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
						if "11合肖" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName+"&"+"[\""+g.JudgeStrWhichChineseZodiac(outNumbersStr[6])+"\"]")
						}
					}
				}
				if "正肖" == gameInfoInnerAggregation.MethodName {
					var result []string
					for i := 0; i < len(outNumbersStr)-1; i++ {
						if strings.Contains(g.GetChineseZodiacValue("鼠"), outNumbersStr[i]) {
							result = append(result, "\"鼠\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("牛"), outNumbersStr[i]) {
							result = append(result, "\"牛\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("虎"), outNumbersStr[i]) {
							result = append(result, "\"虎\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("兔"), outNumbersStr[i]) {
							result = append(result, "\"兔\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("龙"), outNumbersStr[i]) {
							result = append(result, "\"龙\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("蛇"), outNumbersStr[i]) {
							result = append(result, "\"蛇\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("马"), outNumbersStr[i]) {
							result = append(result, "\"马\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("羊"), outNumbersStr[i]) {
							result = append(result, "\"羊\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猴"), outNumbersStr[i]) {
							result = append(result, "\"猴\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("鸡"), outNumbersStr[i]) {
							result = append(result, "\"鸡\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("狗"), outNumbersStr[i]) {
							result = append(result, "\"狗\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猪"), outNumbersStr[i]) {
							result = append(result, "\"猪\"")
						}
					}
					for j := 0; j < len(result); j++ {
						if strings.Contains(result[j], gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "一肖" == gameInfoInnerAggregation.MethodName {
					var result []string
					for i := 0; i < len(outNumbersStr); i++ {
						if strings.Contains(g.GetChineseZodiacValue("鼠"), outNumbersStr[i]) {
							result = append(result, "\"鼠\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("牛"), outNumbersStr[i]) {
							result = append(result, "\"牛\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("虎"), outNumbersStr[i]) {
							result = append(result, "\"虎\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("兔"), outNumbersStr[i]) {
							result = append(result, "\"兔\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("龙"), outNumbersStr[i]) {
							result = append(result, "\"龙\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("蛇"), outNumbersStr[i]) {
							result = append(result, "\"蛇\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("马"), outNumbersStr[i]) {
							result = append(result, "\"马\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("羊"), outNumbersStr[i]) {
							result = append(result, "\"羊\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猴"), outNumbersStr[i]) {
							result = append(result, "\"猴\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("鸡"), outNumbersStr[i]) {
							result = append(result, "\"鸡\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("狗"), outNumbersStr[i]) {
							result = append(result, "\"狗\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猪"), outNumbersStr[i]) {
							result = append(result, "\"猪\"")
						}
					}
					for j := 0; j < len(result); j++ {
						if strings.Contains(result[j], gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "2连肖" == gameInfoInnerAggregation.MethodName ||
					"3连肖" == gameInfoInnerAggregation.MethodName ||
					"4连肖" == gameInfoInnerAggregation.MethodName ||
					"5连肖" == gameInfoInnerAggregation.MethodName {
					var result []string
					for i := 0; i < len(outNumbersStr); i++ {
						if strings.Contains(g.GetChineseZodiacValue("鼠"), outNumbersStr[i]) {
							result = append(result, "\"鼠\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("牛"), outNumbersStr[i]) {
							result = append(result, "\"牛\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("虎"), outNumbersStr[i]) {
							result = append(result, "\"虎\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("兔"), outNumbersStr[i]) {
							result = append(result, "\"兔\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("龙"), outNumbersStr[i]) {
							result = append(result, "\"龙\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("蛇"), outNumbersStr[i]) {
							result = append(result, "\"蛇\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("马"), outNumbersStr[i]) {
							result = append(result, "\"马\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("羊"), outNumbersStr[i]) {
							result = append(result, "\"羊\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猴"), outNumbersStr[i]) {
							result = append(result, "\"猴\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("鸡"), outNumbersStr[i]) {
							result = append(result, "\"鸡\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("狗"), outNumbersStr[i]) {
							result = append(result, "\"狗\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猪"), outNumbersStr[i]) {
							result = append(result, "\"猪\"")
						}
					}
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(result))
				}
				if "总肖" == gameInfoInnerAggregation.MethodName {
					var result []string
					for i := 0; i < len(outNumbersStr); i++ {
						if strings.Contains(g.GetChineseZodiacValue("鼠"), outNumbersStr[i]) {
							result = append(result, "\"+鼠+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("牛"), outNumbersStr[i]) {
							result = append(result, "\"+牛+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("虎"), outNumbersStr[i]) {
							result = append(result, "\"+虎+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("兔"), outNumbersStr[i]) {
							result = append(result, "\"+兔+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("龙"), outNumbersStr[i]) {
							result = append(result, "\"+龙+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("蛇"), outNumbersStr[i]) {
							result = append(result, "\"+蛇+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("马"), outNumbersStr[i]) {
							result = append(result, "\"+马+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("羊"), outNumbersStr[i]) {
							result = append(result, "\"+羊+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猴"), outNumbersStr[i]) {
							result = append(result, "\"+猴+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("鸡"), outNumbersStr[i]) {
							result = append(result, "\"+鸡+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("狗"), outNumbersStr[i]) {
							result = append(result, "\"+狗+\"")
						}
						if strings.Contains(g.GetChineseZodiacValue("猪"), outNumbersStr[i]) {
							result = append(result, "\"+猪+\"")
						}
					}
					operatorStr := utils.RemoveDuplicateElement(result) //去除重复
					if len(operatorStr) == 2 || len(operatorStr) == 3 || len(operatorStr) == 4 && "234肖" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if len(operatorStr) == 5 && "5肖" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if len(operatorStr) == 6 && "6肖" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if len(operatorStr) == 7 && "7肖" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if len(operatorStr)%2 == 0 && "总肖双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if len(operatorStr)%2 != 0 && "总肖单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "波色" == gameInfoInnerAggregation.GroupName {
				if "特码-波色" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(RED, outNumbers[6]) && "红波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(BLUE, outNumbers[6]) && "蓝波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(GREEN, outNumbers[6]) && "绿波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "特码-半波" == gameInfoInnerAggregation.MethodName {
					//红
					if strings.Contains("01、07、13、19、23、29、35、45", outNumbers[6]) && "红单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("02、08、12、18、24、30、34、40、46", outNumbers[6]) && "红双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("29、30、34、35、40、45、46", outNumbers[6]) && "红大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("01、02、07、08、12、13、18、19、23、24", outNumbers[6]) && "红小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					//蓝
					if strings.Contains("03、09、15、25、31、37、41、47", outNumbers[6]) && "蓝单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("04、10、14、20、26、36、42、48", outNumbers[6]) && "蓝双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("25、26、31、36、37、41、42、47、48", outNumbers[6]) && "蓝大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("03、04、09、10、14、15、20", outNumbers[6]) && "蓝小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					//绿
					if strings.Contains("05、11、17、21、27、33、39、43", outNumbers[6]) && "绿单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("06、16、22、28、32、38、44", outNumbers[6]) && "绿双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("27、28、32、33、38、39、43、44", outNumbers[6]) && "绿大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("05、06、11、16、17、21、22", outNumbers[6]) && "绿小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}

				}
				if "特码-半半波" == gameInfoInnerAggregation.MethodName {
					//红
					if strings.Contains("29、35、45", outNumbers[6]) && "红大单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("30、34、40、46", outNumbers[6]) && "红大双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("01、07、13、19、23", outNumbers[6]) && "红小单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("02、08、12、18、24", outNumbers[6]) && "红小双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					//蓝
					if strings.Contains("25、31、37、41、47", outNumbers[6]) && "蓝大单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("26、36、42、48", outNumbers[6]) && "蓝大双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("03、09、15", outNumbers[6]) && "蓝小单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("04、10、14、20", outNumbers[6]) && "蓝小双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					//绿
					if strings.Contains("27、33、39、43", outNumbers[6]) && "绿大单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("28、32、38、44", outNumbers[6]) && "绿大双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("05、11、17、21", outNumbers[6]) && "绿小单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("06、16、22", outNumbers[6]) && "绿小双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "正码-1" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(RED, outNumbers[0]) && "红波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(BLUE, outNumbers[0]) && "蓝波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(GREEN, outNumbers[0]) && "绿波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "正码-2" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(RED, outNumbers[1]) && "红波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(BLUE, outNumbers[1]) && "蓝波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(GREEN, outNumbers[1]) && "绿波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "正码-3" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(RED, outNumbers[2]) && "红波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(BLUE, outNumbers[2]) && "蓝波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(GREEN, outNumbers[2]) && "绿波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "正码-4" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(RED, outNumbers[3]) && "红波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(BLUE, outNumbers[3]) && "蓝波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(GREEN, outNumbers[3]) && "绿波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "正码-5" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(RED, outNumbers[4]) && "红波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(BLUE, outNumbers[4]) && "蓝波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(GREEN, outNumbers[4]) && "绿波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "正码-6" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(RED, outNumbers[5]) && "红波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(BLUE, outNumbers[5]) && "蓝波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains(GREEN, outNumbers[5]) && "绿波" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "气色波" == gameInfoInnerAggregation.MethodName {
					var redCount float64
					var blueCount float64
					var greenCount float64
					var redLucklyCount int
					var blueLucklyCount int
					var greenLucklyCount int
					for i := 0; i < len(outNumbers); i++ {
						if strings.Contains(RED, outNumbers[i]) {
							redCount++
						}
						if strings.Contains(BLUE, outNumbers[i]) {
							blueCount++
						}
						if strings.Contains(GREEN, outNumbers[i]) {
							greenCount++
						}
					}
					if strings.Contains(RED, outNumbers[6]) {
						redCount = redCount + 1.5
						redLucklyCount++
					}
					if strings.Contains(BLUE, outNumbers[6]) {
						blueCount = blueCount + 1.5
						blueLucklyCount++
					}
					if strings.Contains(GREEN, outNumbers[6]) {
						greenCount = greenCount + 1.5
						greenLucklyCount++
					}
					/**
					 * 开出的6个正码各以1个色波计算，特码以1.5个色波计算。得分最高
					 */
					if redCount-blueCount >= 0 && redCount-greenCount >= 0 {
						if "红波" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
					if blueCount-redCount >= 0 && blueCount-greenCount >= 0 {
						if "蓝波" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
					if greenCount-redCount >= 0 && greenCount-blueCount >= 0 {
						if "绿波" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
					}
					/**
					 * 6个正码开出3蓝3绿，特码开出1.5红
					 * 6个正码开出3蓝3红，特码开出1.5绿
					 * 6个正码开出3绿3红，特码开出1.5蓝
					 */
					if (blueCount == 3 && greenCount == 3) && redLucklyCount == 1 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					}
					if (blueCount == 3 && redCount == 3) && greenLucklyCount == 1 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					}
					if (redCount == 3 && greenCount == 3) && blueLucklyCount == 1 {
						lists = append(lists, "0"+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+"和局")
					}
				}
			}
			if "头尾" == gameInfoInnerAggregation.GroupName {
				luckly, _ := strconv.Atoi(outNumbersStr[6])
				tenDigits := luckly / 10
				digits := luckly % 10
				if "特码-头数" == gameInfoInnerAggregation.MethodName {
					if strings.Contains("0头", strconv.Itoa(tenDigits)) && "0头" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("1头", strconv.Itoa(tenDigits)) && "1头" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("2头", strconv.Itoa(tenDigits)) && "2头" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("3头", strconv.Itoa(tenDigits)) && "3头" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("4头", strconv.Itoa(tenDigits)) && "4头" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "特码-尾数" == gameInfoInnerAggregation.MethodName {
					if strings.Contains("0尾", strconv.Itoa(digits)) && "0尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("1尾", strconv.Itoa(digits)) && "1尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("2尾", strconv.Itoa(digits)) && "2尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("3尾", strconv.Itoa(digits)) && "3尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("4尾", strconv.Itoa(digits)) && "4尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("5尾", strconv.Itoa(digits)) && "5尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("6尾", strconv.Itoa(digits)) && "6尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("7尾", strconv.Itoa(digits)) && "7尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("8尾", strconv.Itoa(digits)) && "8尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
					if strings.Contains("9尾", strconv.Itoa(digits)) && "9尾" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+"&"+
							gameInfoInnerAggregation.GroupName+"&"+
							gameInfoInnerAggregation.MethodName+"&"+
							gameInfoInnerAggregation.BettingName)
					}
				}
				if "正特-尾数" == gameInfoInnerAggregation.MethodName {
					var result []string
					for i := 0; i < len(outNumbersStr); i++ {
						oneDigitsInt, _ := strconv.Atoi(outNumbersStr[i])
						oneDigits := oneDigitsInt % 10
						if strings.Contains("0尾", strconv.Itoa(oneDigits)) {
							result = append(result, "0")
						}
						if strings.Contains("1尾", strconv.Itoa(oneDigits)) {
							result = append(result, "1")
						}
						if strings.Contains("2尾", strconv.Itoa(oneDigits)) {
							result = append(result, "2")
						}
						if strings.Contains("3尾", strconv.Itoa(oneDigits)) {
							result = append(result, "3")
						}
						if strings.Contains("4尾", strconv.Itoa(oneDigits)) {
							result = append(result, "4")
						}
						if strings.Contains("5尾", strconv.Itoa(oneDigits)) {
							result = append(result, "5")
						}
						if strings.Contains("6尾", strconv.Itoa(oneDigits)) {
							result = append(result, "6")
						}
						if strings.Contains("7尾", strconv.Itoa(oneDigits)) {
							result = append(result, "7")
						}
						if strings.Contains("8尾", strconv.Itoa(oneDigits)) {
							result = append(result, "8")
						}
						if strings.Contains("9尾", strconv.Itoa(oneDigits)) {
							result = append(result, "9")
						}
					}
					for j := 0; j < len(result); j++ {
						if strings.Contains("0", result[j]) && "0尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("1", result[j]) && "1尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("2", result[j]) && "2尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("3", result[j]) && "3尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("4", result[j]) && "4尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("5", result[j]) && "5尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("6", result[j]) && "6尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("7", result[j]) && "7尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("8", result[j]) && "8尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}
						if strings.Contains("9", result[j]) && "9尾" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+"&"+
								gameInfoInnerAggregation.GroupName+"&"+
								gameInfoInnerAggregation.MethodName+"&"+
								gameInfoInnerAggregation.BettingName)
						}

					}
				}
				if "2连尾" == gameInfoInnerAggregation.MethodName ||
					"3连尾" == gameInfoInnerAggregation.MethodName ||
					"4连尾" == gameInfoInnerAggregation.MethodName ||
					"5连尾" == gameInfoInnerAggregation.MethodName {
					var result []string
					for i := 0; i < len(outNumbersStr); i++ {
						oneDigitsInt, _ := strconv.Atoi(outNumbersStr[i])
						oneDigits := oneDigitsInt % 10
						if strings.Contains("0尾", strconv.Itoa(oneDigits)) {
							result = append(result, "0")
						}
						if strings.Contains("1尾", strconv.Itoa(oneDigits)) {
							result = append(result, "1")
						}
						if strings.Contains("2尾", strconv.Itoa(oneDigits)) {
							result = append(result, "2")
						}
						if strings.Contains("3尾", strconv.Itoa(oneDigits)) {
							result = append(result, "3")
						}
						if strings.Contains("4尾", strconv.Itoa(oneDigits)) {
							result = append(result, "4")
						}
						if strings.Contains("5尾", strconv.Itoa(oneDigits)) {
							result = append(result, "5")
						}
						if strings.Contains("6尾", strconv.Itoa(oneDigits)) {
							result = append(result, "6")
						}
						if strings.Contains("7尾", strconv.Itoa(oneDigits)) {
							result = append(result, "7")
						}
						if strings.Contains("8尾", strconv.Itoa(oneDigits)) {
							result = append(result, "8")
						}
						if strings.Contains("9尾", strconv.Itoa(oneDigits)) {
							result = append(result, "9")
						}
					}
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+"&"+g.SliceToString(result))
				}
			}
			if "自选不中" == gameInfoInnerAggregation.GroupName {
				if "5不中" == gameInfoInnerAggregation.MethodName ||
					"6不中" == gameInfoInnerAggregation.MethodName ||
					"7不中" == gameInfoInnerAggregation.MethodName ||
					"8不中" == gameInfoInnerAggregation.MethodName ||
					"9不中" == gameInfoInnerAggregation.MethodName ||
					"10不中" == gameInfoInnerAggregation.MethodName ||
					"11不中" == gameInfoInnerAggregation.MethodName {
					lists = append(lists, gameInfoInnerAggregation.Id+"&"+
						gameInfoInnerAggregation.GroupName+"&"+
						gameInfoInnerAggregation.MethodName+"&"+
						gameInfoInnerAggregation.BettingName+g.SliceToString(outNumbersStr))
				}
			}
		}
	}
	if "SSC" == modelCode {
		var one, two, three, four, five int
		one, _ = strconv.Atoi(outNumbersStr[0])
		two, _ = strconv.Atoi(outNumbersStr[1])
		three, _ = strconv.Atoi(outNumbersStr[2])
		four, _ = strconv.Atoi(outNumbersStr[3])
		five, _ = strconv.Atoi(outNumbersStr[4])
		sum := one + two + three + four + five
		firstThree := strconv.Itoa(one) + strconv.Itoa(two) + strconv.Itoa(three)
		midThree := strconv.Itoa(two) + strconv.Itoa(three) + strconv.Itoa(four)
		lastThree := strconv.Itoa(three) + strconv.Itoa(four) + strconv.Itoa(five)
		gameInfoInnerAggregationList := dao.SelectGameInfoByGameCode(gameCode, modelCode)
		for i := 0; i < len(gameInfoInnerAggregationList); i++ {
			gameInfoInnerAggregation := gameInfoInnerAggregationList[i]
			if "双面" == gameInfoInnerAggregation.GroupName {
				if "总和" == gameInfoInnerAggregation.MethodName {
					if sum >= 23 && "总和大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if sum < 23 && "总和小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if sum%2 != 0 && "总和单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if sum%2 == 0 && "总和双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第一球" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(SSC_DA, outNumbersStr[0]) && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_XIAO, outNumbersStr[0]) && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_DAN, outNumbersStr[0]) && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_SHUANG, outNumbersStr[0]) && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_ZHI, outNumbersStr[0]) && "质" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_HE, outNumbersStr[0]) && "合" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第二球" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(SSC_DA, outNumbersStr[1]) && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_XIAO, outNumbersStr[1]) && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_DAN, outNumbersStr[1]) && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_SHUANG, outNumbersStr[1]) && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_ZHI, outNumbersStr[1]) && "质" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_HE, outNumbersStr[1]) && "合" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第三球" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(SSC_DA, outNumbersStr[2]) && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_XIAO, outNumbersStr[2]) && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_DAN, outNumbersStr[2]) && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_SHUANG, outNumbersStr[2]) && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_ZHI, outNumbersStr[2]) && "质" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_HE, outNumbersStr[2]) && "合" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第四球" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(SSC_DA, outNumbersStr[3]) && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_XIAO, outNumbersStr[3]) && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_DAN, outNumbersStr[3]) && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_SHUANG, outNumbersStr[3]) && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_ZHI, outNumbersStr[3]) && "质" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_HE, outNumbersStr[3]) && "合" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第五球" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(SSC_DA, outNumbersStr[4]) && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_XIAO, outNumbersStr[4]) && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_DAN, outNumbersStr[4]) && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_SHUANG, outNumbersStr[4]) && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_ZHI, outNumbersStr[4]) && "质" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SSC_HE, outNumbersStr[4]) && "合" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "龙虎斗" == gameInfoInnerAggregation.GroupName {
				if "龙虎斗" == gameInfoInnerAggregation.MethodName {
					if one > five && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if one < five && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if one == five && "和" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "选号" == gameInfoInnerAggregation.GroupName {
				if "第一球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 10; i++ {
						if i == one && strings.Contains(strconv.Itoa(i), gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName+
								"&"+strconv.Itoa(i))
						}
					}
				}
				if "第二球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 10; i++ {
						if i == one && strings.Contains(strconv.Itoa(i), gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName+
								"&"+strconv.Itoa(i))
						}
					}
				}
				if "第三球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 10; i++ {
						if i == one && strings.Contains(strconv.Itoa(i), gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName+
								"&"+strconv.Itoa(i))
						}
					}
				}
				if "第四球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 10; i++ {
						if i == one && strings.Contains(strconv.Itoa(i), gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName+
								"&"+strconv.Itoa(i))
						}
					}
				}
				if "第五球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 10; i++ {
						if i == one && strings.Contains(strconv.Itoa(i), gameInfoInnerAggregation.BettingName) {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName+
								"&"+strconv.Itoa(i))
						}
					}
				}
				if "全五中一" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 10; i++ {
						for j := 0; j < len(outNumbersStr); j++ {
							intJ, _ := strconv.Atoi(outNumbersStr[j])
							bettingNameInt, _ := strconv.Atoi(gameInfoInnerAggregation.BettingName)
							if i == intJ && i == bettingNameInt {
								lists = append(lists, gameInfoInnerAggregation.Id+
									"&"+gameInfoInnerAggregation.GroupName+
									"&"+gameInfoInnerAggregation.MethodName+
									"&"+gameInfoInnerAggregation.BettingName+
									"&"+strconv.Itoa(i))
							}
						}
					}
				}
			}
			if "前中后三" == gameInfoInnerAggregation.GroupName {
				if "前三" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(BAOZI, firstThree) && "豹子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SHUNZI, firstThree) && "顺子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, firstThree)) && "对子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (math.Abs(float64(one-two)) == 1 ||
						math.Abs(float64(two-three)) == 1 ||
						math.Abs(float64(three-one)) == 1 || math.Abs(float64(one-two)) == 1 ||
						math.Abs(float64(one-two)) == 9 ||
						math.Abs(float64(two-three)) == 9 ||
						math.Abs(float64(one-two)) == 9) && !strings.Contains(SHUNZI, firstThree) &&
						!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, firstThree)) &&
						"半顺" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (!strings.Contains(BAOZI, firstThree) && !strings.Contains(SHUNZI, firstThree) &&
						!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, firstThree)) &&
						!((math.Abs(float64(one-two)) == 1 || math.Abs(float64(two-three)) == 1 || math.Abs(float64(three-one)) == 1 ||
							math.Abs(float64(one-two)) == 9 || math.Abs(float64(two-three)) == 9 || math.Abs(float64(three-one)) == 9) && !strings.Contains(SHUNZI, firstThree))) &&
						"杂六" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "中三" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(BAOZI, midThree) && "豹子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SHUNZI, midThree) && "顺子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, midThree)) && "对子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (math.Abs(float64(one-two)) == 1 ||
						math.Abs(float64(two-three)) == 1 ||
						math.Abs(float64(three-one)) == 1 || math.Abs(float64(one-two)) == 1 ||
						math.Abs(float64(one-two)) == 9 ||
						math.Abs(float64(two-three)) == 9 ||
						math.Abs(float64(one-two)) == 9) && !strings.Contains(SHUNZI, midThree) &&
						!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, midThree)) &&
						"半顺" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (!strings.Contains(BAOZI, midThree) && !strings.Contains(SHUNZI, midThree) &&
						!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, midThree)) &&
						!((math.Abs(float64(one-two)) == 1 || math.Abs(float64(two-three)) == 1 || math.Abs(float64(three-one)) == 1 ||
							math.Abs(float64(one-two)) == 9 || math.Abs(float64(two-three)) == 9 || math.Abs(float64(three-one)) == 9) && !strings.Contains(SHUNZI, midThree))) &&
						"杂六" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "后三" == gameInfoInnerAggregation.MethodName {
					if strings.Contains(BAOZI, lastThree) && "豹子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if strings.Contains(SHUNZI, lastThree) && "顺子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, lastThree)) && "对子" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (math.Abs(float64(one-two)) == 1 ||
						math.Abs(float64(two-three)) == 1 ||
						math.Abs(float64(three-one)) == 1 || math.Abs(float64(one-two)) == 1 ||
						math.Abs(float64(one-two)) == 9 ||
						math.Abs(float64(two-three)) == 9 ||
						math.Abs(float64(one-two)) == 9) && !strings.Contains(SHUNZI, lastThree) &&
						!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, lastThree)) &&
						"半顺" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if (!strings.Contains(BAOZI, lastThree) && !strings.Contains(SHUNZI, lastThree) &&
						!(((one == two || one == three || two == three) && (one != three || three != two)) && !strings.Contains(BAOZI, lastThree)) &&
						!((math.Abs(float64(one-two)) == 1 || math.Abs(float64(two-three)) == 1 || math.Abs(float64(three-one)) == 1 ||
							math.Abs(float64(one-two)) == 9 || math.Abs(float64(two-three)) == 9 || math.Abs(float64(three-one)) == 9) && !strings.Contains(SHUNZI, lastThree))) &&
						"杂六" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
		}
	}
	if "11X5" == modelCode {
		var one, two, three, four, five int
		one, _ = strconv.Atoi(outNumbersStr[0])
		two, _ = strconv.Atoi(outNumbersStr[1])
		three, _ = strconv.Atoi(outNumbersStr[2])
		four, _ = strconv.Atoi(outNumbersStr[3])
		five, _ = strconv.Atoi(outNumbersStr[4])
		sum := one + two + three + four + five
		digits := sum % 10
		gameInfoInnerAggregationList := dao.SelectGameInfoByGameCode(gameCode, modelCode)
		for i := 0; i < len(gameInfoInnerAggregationList); i++ {
			gameInfoInnerAggregation := gameInfoInnerAggregationList[i]
			if "双面" == gameInfoInnerAggregation.GroupName {
				if "总和" == gameInfoInnerAggregation.MethodName {
					if sum > 30 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if sum < 30 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if sum == 30 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"大&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"小&和")
					}
					if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if digits >= 5 && "尾大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if digits <= 4 && "尾小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one > five && "龙" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					} else if one < five && "虎" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
				if "第一球" == gameInfoInnerAggregation.MethodName {
					if one == 11 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第一球&大&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第一球&小&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第一球&单&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第一球&双&和")
					} else {
						if one >= 06 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if one <= 05 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if one%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "第二球" == gameInfoInnerAggregation.MethodName {
					if two == 11 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第二球&大&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第二球&小&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第二球&单&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第二球&双&和")
					} else {
						if two >= 06 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if two <= 05 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if two%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if two%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "第三球" == gameInfoInnerAggregation.MethodName {
					if three == 11 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第三球&大&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第三球&小&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第三球&单&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第三球&双&和")
					} else {
						if three >= 06 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if three <= 05 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if three%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if three%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "第四球" == gameInfoInnerAggregation.MethodName {
					if four == 11 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第四球&大&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第四球&小&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第四球&单&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第四球&双&和")
					} else {
						if four >= 06 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if four <= 05 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if four%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if four%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
					}
				}
				if "第五球" == gameInfoInnerAggregation.MethodName {
					if five == 11 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第五球&大&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第五球&小&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第五球&单&和")
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"第五球&双&和")
					} else {
						if five >= 06 && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if five <= 05 && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if five%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if five%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
					}
				}
			}
			if "选号" == gameInfoInnerAggregation.GroupName {
				if "第一球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 12; i++ {
						var str string
						if i < 10 {
							str = "0" + strconv.Itoa(i)
						} else {
							str = strconv.Itoa(i)
						}
						intStr, _ := strconv.Atoi(str)
						intBettingName, _ := strconv.Atoi(gameInfoInnerAggregation.BettingName)
						if one == intStr && one == intBettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+str)
						}
					}
				}
				if "第二球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 12; i++ {
						var str string
						if i < 10 {
							str = "0" + strconv.Itoa(i)
						} else {
							str = strconv.Itoa(i)
						}
						intStr, _ := strconv.Atoi(str)
						intBettingName, _ := strconv.Atoi(gameInfoInnerAggregation.BettingName)
						if two == intStr && two == intBettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+str)
						}
					}
				}
				if "第三球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 12; i++ {
						var str string
						if i < 10 {
							str = "0" + strconv.Itoa(i)
						} else {
							str = strconv.Itoa(i)
						}
						intStr, _ := strconv.Atoi(str)
						intBettingName, _ := strconv.Atoi(gameInfoInnerAggregation.BettingName)
						if three == intStr && three == intBettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+str)
						}
					}
				}
				if "第四球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 12; i++ {
						var str string
						if i < 10 {
							str = "0" + strconv.Itoa(i)
						} else {
							str = strconv.Itoa(i)
						}
						intStr, _ := strconv.Atoi(str)
						intBettingName, _ := strconv.Atoi(gameInfoInnerAggregation.BettingName)
						if four == intStr && four == intBettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+str)
						}
					}
				}
				if "第五球" == gameInfoInnerAggregation.MethodName {
					for i := 0; i < 12; i++ {
						var str string
						if i < 10 {
							str = "0" + strconv.Itoa(i)
						} else {
							str = strconv.Itoa(i)
						}
						intStr, _ := strconv.Atoi(str)
						intBettingName, _ := strconv.Atoi(gameInfoInnerAggregation.BettingName)
						if five == intStr && five == intBettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+str)
						}
					}
				}
			}
			if "任选" == gameInfoInnerAggregation.GroupName {
				if "一中一" == gameInfoInnerAggregation.MethodName {
					if "一中一" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "二中二" == gameInfoInnerAggregation.MethodName {
					if "二中二" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "三中三" == gameInfoInnerAggregation.MethodName {
					if "二中二" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "四中四" == gameInfoInnerAggregation.MethodName {
					if "四中四" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "五中五" == gameInfoInnerAggregation.MethodName {
					if "五中五" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "六中五" == gameInfoInnerAggregation.MethodName {
					if "六中五" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "七中五" == gameInfoInnerAggregation.MethodName {
					if "七中五" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "八中五" == gameInfoInnerAggregation.MethodName {
					if "八中五" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+g.SliceToString(outNumbersStr))
					}
				}
				if "前二组选" == gameInfoInnerAggregation.MethodName {
					if "前二组选" == gameInfoInnerAggregation.BettingName {
						var oneStr string
						var twoStr string
						if one < 10 {
							oneStr = "0" + strconv.Itoa(one)
						} else {
							oneStr = strconv.Itoa(one)
						}
						if two < 10 {
							twoStr = "0" + strconv.Itoa(two)
						} else {
							twoStr = strconv.Itoa(two)
						}
						result := "[" + oneStr + twoStr + "]"
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+result)
					}
				}
				if "前三组选" == gameInfoInnerAggregation.MethodName {
					if "前三组选" == gameInfoInnerAggregation.BettingName {
						var oneStr string
						var twoStr string
						var threeStr string
						if one < 10 {
							oneStr = "0" + strconv.Itoa(one)
						} else {
							oneStr = strconv.Itoa(one)
						}
						if two < 10 {
							twoStr = "0" + strconv.Itoa(two)
						} else {
							twoStr = strconv.Itoa(two)
						}
						if three < 10 {
							threeStr = "0" + strconv.Itoa(three)
						} else {
							threeStr = strconv.Itoa(three)
						}
						result := "[" + oneStr + twoStr + threeStr + "]"
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+result)
					}
				}
			}
			if "直选" == gameInfoInnerAggregation.GroupName {
				if "前二直选" == gameInfoInnerAggregation.MethodName {
					if "前二直选" == gameInfoInnerAggregation.BettingName {
						var oneStr string
						var twoStr string
						if one < 10 {
							oneStr = "0" + strconv.Itoa(one)
						} else {
							oneStr = strconv.Itoa(one)
						}
						if two < 10 {
							twoStr = "0" + strconv.Itoa(two)
						} else {
							twoStr = strconv.Itoa(two)
						}
						result := "[" + oneStr + twoStr + "]"
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+result)
					}
				}
				if "前三直选" == gameInfoInnerAggregation.MethodName {
					if "前三直选" == gameInfoInnerAggregation.BettingName {
						var oneStr string
						var twoStr string
						var threeStr string
						if one < 10 {
							oneStr = "0" + strconv.Itoa(one)
						} else {
							oneStr = strconv.Itoa(one)
						}
						if two < 10 {
							twoStr = "0" + strconv.Itoa(two)
						} else {
							twoStr = strconv.Itoa(two)
						}
						if three < 10 {
							threeStr = "0" + strconv.Itoa(three)
						} else {
							threeStr = strconv.Itoa(three)
						}
						result := "[" + oneStr + twoStr + threeStr + "]"
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName+
							"&"+result)
					}
				}
			}
		}
	}
	if "K3" == modelCode {
		var one, two, three int
		one, _ = strconv.Atoi(outNumbersStr[0])
		two, _ = strconv.Atoi(outNumbersStr[1])
		three, _ = strconv.Atoi(outNumbersStr[2])
		oneStr := strconv.Itoa(one)
		twoStr := strconv.Itoa(two)
		threeStr := strconv.Itoa(three)
		flag := strings.Contains(oneStr, twoStr) && strings.Contains(twoStr, threeStr) && strings.Contains(threeStr, oneStr)
		sum := one + two + three
		sumStr := oneStr + twoStr + threeStr
		gameInfoInnerAggregationList := dao.SelectGameInfoByGameCode(gameCode, modelCode)
		for i := 0; i < len(gameInfoInnerAggregationList); i++ {
			gameInfoInnerAggregation := gameInfoInnerAggregationList[i]
			if "三军" == gameInfoInnerAggregation.GroupName {
				if "三军" == gameInfoInnerAggregation.MethodName {
					if (one == 1 || two == 1 || three == 1) && "1" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 2 || two == 2 || three == 2) && "2" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 3 || two == 3 || three == 3) && "3" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 4 || two == 4 || three == 4) && "4" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 5 || two == 5 || three == 5) && "5" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "围骰" == gameInfoInnerAggregation.GroupName {
				if sum == 3 && "111" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if sum == 6 && "222" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if sum == 9 && "333" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if sum == 12 && "444" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if sum == 15 && "555" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if sum == 18 && "666" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
				if (sum == 18 || sum == 15 || sum == 12 || sum == 9 || sum == 6 || sum == 3) && flag && "全骰" == gameInfoInnerAggregation.BettingName {
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+gameInfoInnerAggregation.BettingName)
				}
			}
			if "点数和值" == gameInfoInnerAggregation.GroupName {
				if "点数和值" == gameInfoInnerAggregation.MethodName {
					if !flag {
						if (11 < sum && sum < 17) && "大" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if (4 < sum && sum < 7) && "小" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						} else if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 4 && "4" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 5 && "5" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 6 && "6" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 7 && "7" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 8 && "8" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 9 && "9" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 10 && "10" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 11 && "11" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 12 && "12" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 13 && "13" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 14 && "14" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 15 && "15" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 16 && "16" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if sum == 17 && "17" == gameInfoInnerAggregation.BettingName {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
					}
				}
			}
			if "长牌" == gameInfoInnerAggregation.GroupName {
				if "长牌" == gameInfoInnerAggregation.MethodName {
					if StrIsContains(sumStr, "1,2") && StrIsContains(gameInfoInnerAggregation.BettingName, "1,2") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "1,3") && StrIsContains(gameInfoInnerAggregation.BettingName, "1,3") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "1,4") && StrIsContains(gameInfoInnerAggregation.BettingName, "1,4") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "1,5") && StrIsContains(gameInfoInnerAggregation.BettingName, "1,5") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "1,6") && StrIsContains(gameInfoInnerAggregation.BettingName, "1,6") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "2,3") && StrIsContains(gameInfoInnerAggregation.BettingName, "2,3") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "2,4") && StrIsContains(gameInfoInnerAggregation.BettingName, "2,4") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "2,5") && StrIsContains(gameInfoInnerAggregation.BettingName, "2,5") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "2,6") && StrIsContains(gameInfoInnerAggregation.BettingName, "2,6") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "3,4") && StrIsContains(gameInfoInnerAggregation.BettingName, "3,4") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "3,5") && StrIsContains(gameInfoInnerAggregation.BettingName, "3,5") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "3,6") && StrIsContains(gameInfoInnerAggregation.BettingName, "3,6") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "4,5") && StrIsContains(gameInfoInnerAggregation.BettingName, "4,5") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "4,6") && StrIsContains(gameInfoInnerAggregation.BettingName, "4,6") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "5,6") && StrIsContains(gameInfoInnerAggregation.BettingName, "5,6") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "短牌" == gameInfoInnerAggregation.GroupName {
				if "短牌" == gameInfoInnerAggregation.MethodName {
					if StrIsContains(sumStr, "11") && StrIsContains(gameInfoInnerAggregation.BettingName, "1,1") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "22") && StrIsContains(gameInfoInnerAggregation.BettingName, "2,2") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "33") && StrIsContains(gameInfoInnerAggregation.BettingName, "3,3") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "44") && StrIsContains(gameInfoInnerAggregation.BettingName, "4,4") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "55") && StrIsContains(gameInfoInnerAggregation.BettingName, "5,5") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if StrIsContains(sumStr, "66") && StrIsContains(gameInfoInnerAggregation.BettingName, "6,6") {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
		}
	}
	if "PC28" == modelCode {
		var one, two, three int
		one, _ = strconv.Atoi(outNumbersStr[0])
		two, _ = strconv.Atoi(outNumbersStr[1])
		three, _ = strconv.Atoi(outNumbersStr[2])
		flag := one == two && one == three && two == three
		sum := one + two + three
		result := strconv.Itoa(one) + "," + strconv.Itoa(two) + "," + strconv.Itoa(three)
		gameInfoInnerAggregationList := dao.SelectGameInfoByGameCode(gameCode, modelCode)
		for i := 0; i < len(gameInfoInnerAggregationList); i++ {
			gameInfoInnerAggregation := gameInfoInnerAggregationList[i]
			if "双面" == gameInfoInnerAggregation.GroupName {
				if "双面" == gameInfoInnerAggregation.MethodName {
					if sum >= 14 && "大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if sum <= 13 && "小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if sum%2 != 0 && "单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if sum%2 == 0 && "双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if SliceCompareString(strconv.Itoa(sum), PC28_DA_DAN) && "大单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if SliceCompareString(strconv.Itoa(sum), PC28_DA_SHUANG) && "大双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if SliceCompareString(strconv.Itoa(sum), PC28_XIAO_DAN) && "小单" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if SliceCompareString(strconv.Itoa(sum), PC28_XIAO_SHUANG) && "小双" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if sum >= 22 && "极大" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if sum <= 5 && "极小" == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "和值" == gameInfoInnerAggregation.GroupName {
				for i := 0; i < 28; i++ {
					if sum == i && strconv.Itoa(i) == gameInfoInnerAggregation.BettingName {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+strconv.Itoa(sum))
					}
				}
			}
			if "其他" == gameInfoInnerAggregation.GroupName {
				if "色波、豹子" == gameInfoInnerAggregation.MethodName{
					if SliceCompareString(strconv.Itoa(sum),PC28_RED) && "红波" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if SliceCompareString(strconv.Itoa(sum),PC28_HE) && "红波" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, "0"+"&"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "色波、豹子" == gameInfoInnerAggregation.MethodName{
					if SliceCompareString(strconv.Itoa(sum),PC28_BLUE) && "蓝波" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if SliceCompareString(strconv.Itoa(sum),PC28_HE) && "蓝波" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, "0"+"&"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "色波、豹子" == gameInfoInnerAggregation.MethodName{
					if SliceCompareString(strconv.Itoa(sum),PC28_GREEN) && "绿波" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if SliceCompareString(strconv.Itoa(sum),PC28_GREEN) && "绿波" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, "0"+"&"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "色波、豹子" == gameInfoInnerAggregation.MethodName{
					if flag && "豹子" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "特码包三" == gameInfoInnerAggregation.GroupName {
				if "特码包三" == gameInfoInnerAggregation.MethodName{
					lists = append(lists, gameInfoInnerAggregation.Id+
						"&"+gameInfoInnerAggregation.GroupName+
						"&"+gameInfoInnerAggregation.MethodName+
						"&"+"["+result+"]")
				}
			}
		}
	}
	if "FT" == modelCode {
		one, _ := strconv.Atoi(outNumbersStr[0])
		gameInfoInnerAggregationList := dao.SelectGameInfoByGameCode(gameCode, modelCode)
		for i := 0; i < len(gameInfoInnerAggregationList); i++ {
			gameInfoInnerAggregation := gameInfoInnerAggregationList[i]
			if "双面" == gameInfoInnerAggregation.GroupName{
				if "双面" == gameInfoInnerAggregation.MethodName{
					if (one == 3 || one == 4) && "大" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 1 || one == 2) && "小" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 1 || one == 3) && "单" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 2 || one == 4) && "双" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "番" == gameInfoInnerAggregation.GroupName{
				if "番" == gameInfoInnerAggregation.MethodName{
					if (one == 1) && "1番" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 2) && "2番" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 3) && "3番" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 4) && "4番" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "念" == gameInfoInnerAggregation.GroupName{
				if "念" == gameInfoInnerAggregation.MethodName{
					if "1念2" == gameInfoInnerAggregation.BettingName{
						if one == 1 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 2 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "1念3" == gameInfoInnerAggregation.BettingName{
						if one == 1 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 3 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "1念4" == gameInfoInnerAggregation.BettingName{
						if one == 1 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 4 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "2念1" == gameInfoInnerAggregation.BettingName{
						if one == 2 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 1 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "2念3" == gameInfoInnerAggregation.BettingName{
						if one == 2 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 3 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "2念4" == gameInfoInnerAggregation.BettingName{
						if one == 2 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 4 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "3念1" == gameInfoInnerAggregation.BettingName{
						if one == 3 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 1 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "3念2" == gameInfoInnerAggregation.BettingName{
						if one == 3 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 2 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "3念4" == gameInfoInnerAggregation.BettingName{
						if one == 3 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 4 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "4念1" == gameInfoInnerAggregation.BettingName{
						if one == 4 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 1 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "4念2" == gameInfoInnerAggregation.BettingName{
						if one == 4 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 2 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "4念3" == gameInfoInnerAggregation.BettingName{
						if one == 4 {
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 3 {
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
				}
			}
			if "角" == gameInfoInnerAggregation.GroupName{
				if "角" == gameInfoInnerAggregation.MethodName{
					if (one == 1 || one == 2) && "12角" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 2 || one == 3) && "23角" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 3 || one == 4) && "24角" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 1 || one == 4) && "14角" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "通" == gameInfoInnerAggregation.GroupName{
				if "23一通" == gameInfoInnerAggregation.MethodName{
					if one == 2 || one == 3 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 4 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "24一通" == gameInfoInnerAggregation.MethodName{
					if one == 2 || one == 4 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 3 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "34一通" == gameInfoInnerAggregation.MethodName{
					if one == 4 || one == 3 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 2 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "13二通" == gameInfoInnerAggregation.MethodName{
					if one == 1 || one == 3 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 4 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "14二通" == gameInfoInnerAggregation.MethodName{
					if one == 1 || one == 4 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 2 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "34二通" == gameInfoInnerAggregation.MethodName{
					if one == 3 || one == 4 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 1 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "12三通" == gameInfoInnerAggregation.MethodName{
					if one == 1 || one == 2 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 4 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "14三通" == gameInfoInnerAggregation.MethodName{
					if one == 1 || one == 4 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 2 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "24三通" == gameInfoInnerAggregation.MethodName{
					if one == 2 || one == 4 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 1 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "12四通" == gameInfoInnerAggregation.MethodName{
					if one == 1 || one == 2 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 3 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "13四通" == gameInfoInnerAggregation.MethodName{
					if one == 1 || one == 3 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 2 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
				if "23四通" == gameInfoInnerAggregation.MethodName{
					if one == 2 || one == 3 {
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if one == 1 {
						lists = append(lists, "0"+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+"和")
					}
				}
			}
			if "三门" == gameInfoInnerAggregation.GroupName{
				if "三门" == gameInfoInnerAggregation.MethodName{
					if (one == 1 || one == 2 || one == 3) && "123中" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 1 || one == 2 || one == 4) && "124中" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 1 || one == 3 || one == 4) && "134中" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
					if (one == 2 || one == 3 || one == 4) && "234中" == gameInfoInnerAggregation.BettingName{
						lists = append(lists, gameInfoInnerAggregation.Id+
							"&"+gameInfoInnerAggregation.GroupName+
							"&"+gameInfoInnerAggregation.MethodName+
							"&"+gameInfoInnerAggregation.BettingName)
					}
				}
			}
			if "正" == gameInfoInnerAggregation.GroupName{
				if "正" == gameInfoInnerAggregation.MethodName{
					if "1正" == gameInfoInnerAggregation.BettingName{
						if one == 1{
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 4 || one == 2{
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "3正" == gameInfoInnerAggregation.BettingName{
						if one == 2{
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 1 || one == 3{
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "3正" == gameInfoInnerAggregation.BettingName{
						if one == 3{
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 2 || one == 4{
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
					if "4正" == gameInfoInnerAggregation.BettingName{
						if one == 4{
							lists = append(lists, gameInfoInnerAggregation.Id+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+gameInfoInnerAggregation.BettingName)
						}
						if one == 1 || one == 3{
							lists = append(lists, "0"+
								"&"+gameInfoInnerAggregation.GroupName+
								"&"+gameInfoInnerAggregation.MethodName+
								"&"+"和")
						}
					}
				}
			}
		}
	}
	return lists
}

//判断当前时间是否在游戏计划时间之内
func (g *GameUtils) IsEffectiveDateStr(nowTime string, sTime string, eTime string) bool {
	//var timeLayoutStr = "15:04"
	//st, _ := time.ParseInLocation(timeLayoutStr, sTime, time.Local)   //string转time
	//et, _ := time.ParseInLocation(timeLayoutStr, eTime, time.Local)   //string转time
	//nt, _ := time.ParseInLocation(timeLayoutStr, nowTime, time.Local) //string转time
	//log.Println("判断时间是否在计划时间内", st.Before(nt) && et.After(nt))
	log.Println("判断时间是否在计划时间内", strings.Compare(nowTime, sTime) >= 0 && strings.Compare(nowTime, eTime) <= 0)
	return strings.Compare(nowTime, sTime) >= 0 && strings.Compare(nowTime, eTime) <= 0
}

func (g *GameUtils) YearOfNumber(index int) map[string][]string {
	shengxiao := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	year := int(math.Abs(float64(time.Now().Year()-index))) % 12
	return StitchingData(shengxiao, year)
}

func StitchingData(years []string, index int) map[string][]string {
	var mp = make(map[string][]string)
	var str string
	var list []string
	for i, j := 1, index; i < 50; i++ {
		list = mp[years[j]]
		if i < 10 {
			str = "0" + strconv.Itoa(i)
		} else {
			str = strconv.Itoa(i)
		}
		if mp[years[j]] == nil {
			list = make([]string, 0)
		}
		if len(list) == 0 {
			list = append(list, str)
		} else {
			list = append(list, ","+str)
		}
		mp[years[j]] = list
		j = j - 1
		if j < 0 {
			j = 11
		}
	}
	return mp
}

//[0 ,0 ,4]--->[0 0 4]
func StringSliceCut(str []string) []string {
	var resultStr []string
	var strArray []string
	var newStr string
	for index, v := range str {
		if index != 0 {
			strArray = strings.Split(v, ",")
			newStr = strArray[1]
		} else {
			newStr = v
		}
		resultStr = append(resultStr, newStr)
	}
	return resultStr
}

//[0 0 4]--->[0 ,0 ,4]
func StringJointSlice(str []string) []string {
	var resultStr []string
	for index, value := range str {
		if index != 0 {
			resultStr = append(resultStr, ","+value)
		} else {
			resultStr = append(resultStr, value)
		}
	}
	return resultStr
}

//将切片转成string类型
func (g *GameUtils) SliceToString(str []string) string {
	var strResult string
	for index, value := range str {
		if index != 0 {
			strResult += "," + value
		} else {
			strResult += value
		}
	}
	return "[" + strResult + "]"
}

func (g *GameUtils) JudgmentCurrentYearIsBronYear() string {
	now := time.Now().Format("2006-01-02")
	markSixData := markSixDataService.SelectByYear(strconv.Itoa(time.Now().Year()))
	nextMarkSixData := markSixDataService.SelectByYear(strconv.Itoa(time.Now().Year() + 1))
	currentYear := markSixData.Year + "-" + markSixData.Month + "-" + markSixData.Day
	lastYear := nextMarkSixData.Year + "-" + nextMarkSixData.Month + "-" + nextMarkSixData.Day
	if time.Now().Format("2006-01-02") == currentYear {
		return markSixData.ChineseZodiac
	}
	lastYearTime, _ := time.Parse(format, lastYear)
	currentYearTime, _ := time.Parse(format, currentYear)
	nowTime, _ := time.Parse(format, now)
	if nowTime.Before(lastYearTime) && nowTime.After(currentYearTime) {
		return nextMarkSixData.ChineseZodiac
	}
	return ""
}

func (g *GameUtils) JudgeStrWhichChineseZodiac(str string) string {
	nowStr := time.Now().Format("2006-01-02")
	nowTime, _ := time.Parse(format, nowStr)
	year := time.Now().Year()
	for index, value := range mp {
		index1Time, _ := time.Parse(format, index)
		yearStr := strings.Split(index, "-")[0]
		if yearStr == strconv.Itoa(year) {
			if index == nowStr {
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(value), &valueMap)
				for index, value := range valueMap {
					if strings.Contains(value, str) {
						return index
					}
				}
			}
			if nowTime.Before(index1Time) {
				markSixData := markSixDataService.SelectByYear(strconv.Itoa(year - 1))
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(markSixData.Data), &valueMap)
				for index, value := range valueMap {
					if strings.Contains(value, str) {
						return index
					}
				}
			}
			if nowTime.After(index1Time) {
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(value), &valueMap)
				for index, value := range valueMap {
					if strings.Contains(value, str) {
						return index
					}
				}
			}
		}
	}
	return ""
}

func (g *GameUtils) GetChineseZodiacValue(chineseZodiac string) string {
	nowStr := time.Now().Format("2006-01-02")
	nowTime, _ := time.Parse(format, nowStr)
	year := time.Now().Year()
	for index, value := range mp {
		index1Time, _ := time.Parse(format, index)
		yearStr := strings.Split(index, "-")[0]
		if yearStr == strconv.Itoa(year) {
			if index == nowStr {
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(value), &valueMap)
				for index, value := range valueMap {
					if chineseZodiac == index {
						return value
					}
				}
			}
			if nowTime.Before(index1Time) {
				markSixData := markSixDataService.SelectByYear(strconv.Itoa(year - 1))
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(markSixData.Data), &valueMap)
				for index, value := range valueMap {
					if chineseZodiac == index {
						return value
					}
				}
			}
			if nowTime.After(index1Time) {
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(value), &valueMap)
				for index, value := range valueMap {
					if chineseZodiac == index {
						return value
					}
				}
			}
		}
	}
	return ""
}

func (g *GameUtils) GetChineseZodiacAllValue() []vo.MarkSixDataVo {
	var markSixDataList []*model.MarkSixData
	var markSixDataVoList []vo.MarkSixDataVo
	year := time.Now().Year()
	var nowStr = time.Now().Format("2006-01-02")
	nowTime, _ := time.Parse(format, nowStr)
	markSixDataList = markSixDataService.SelectAllYear()
	for _, value := range markSixDataList {
		dbData := value.Year + "-" + value.Month + "-" + value.Day
		dbDataTime, _ := time.Parse(format, dbData)
		if strconv.Itoa(year) == value.Year {
			if nowStr == dbData {
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(value.Data), &valueMap)
				for name, data := range valueMap {
					markSixDataVo := vo.MarkSixDataVo{
						Name: name,
						Data: data,
					}
					markSixDataVoList = append(markSixDataVoList, markSixDataVo)
				}
				return markSixDataVoList
			}
			beforeMarkSixData := markSixDataService.SelectByYear(strconv.Itoa(year + 1))
			beforeDbData := beforeMarkSixData.Year + "-" + beforeMarkSixData.Month + "-" + beforeMarkSixData.Day
			beforeDbDataTime, _ := time.Parse(format, beforeDbData)
			if nowTime.After(dbDataTime) && nowTime.Before(beforeDbDataTime) {
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(value.Data), &valueMap)
				for name, data := range valueMap {
					markSixDataVo := vo.MarkSixDataVo{
						Name: name,
						Data: data,
					}
					markSixDataVoList = append(markSixDataVoList, markSixDataVo)
				}
				return markSixDataVoList
			}

			afterMarkSixData := markSixDataService.SelectByYear(strconv.Itoa(year - 1))
			if !afterMarkSixData.IsEmpty() {
				afterDbData := afterMarkSixData.Year + "-" + afterMarkSixData.Month + "-" + afterMarkSixData.Day
				afterDbDataTime, _ := time.Parse(format, afterDbData)
				if nowTime.Before(dbDataTime) && nowTime.After(afterDbDataTime) {
					//将value转换为map类型
					valueMap := make(map[string]string)
					json.Unmarshal([]byte(value.Data), &valueMap)
					for name, data := range valueMap {
						markSixDataVo := vo.MarkSixDataVo{
							Name: name,
							Data: data,
						}
						markSixDataVoList = append(markSixDataVoList, markSixDataVo)
					}
					return markSixDataVoList
				}
			}
		}
	}
	return markSixDataVoList
}

func (g *GameUtils) updateOdds() {
	//说明是本命年
	chineseZodiac := g.JudgmentCurrentYearIsBronYear()
	//更新本命年odds
	zodiacVoList := zodiacOddsRelationService.GetZodiacList()
	//遍历集合
	for _, value := range zodiacVoList {
		if "特肖" == value.ItemsName {
			zodiacOddsRelation := zodiacOddsRelationService.SelectByItemsName(value.ItemsName)
			id, _ := strconv.Atoi(value.Id)
			if chineseZodiac == value.BettingName {
				gameBettingService.UpdateGameBettingOddsByIdInZodiac(
					int64(id),
					zodiacOddsRelation.NatalOdds,
					value.BettingName)
			} else {
				gameBettingService.UpdateGameBettingOddsById(int64(id), zodiacOddsRelation.NoNatalOdds)
			}
		}
		if "正肖" == value.ItemsName {
			zodiacOddsRelation := zodiacOddsRelationService.SelectByItemsName(value.ItemsName)
			id, _ := strconv.Atoi(value.Id)
			if chineseZodiac == value.BettingName {
				gameBettingService.UpdateGameBettingOddsByIdInZodiac(
					int64(id),
					zodiacOddsRelation.NatalOdds,
					value.BettingName)
			} else {
				gameBettingService.UpdateGameBettingOddsById(int64(id), zodiacOddsRelation.NoNatalOdds)
			}
		}
		if "一肖" == value.ItemsName {
			zodiacOddsRelation := zodiacOddsRelationService.SelectByItemsName(value.ItemsName)
			id, _ := strconv.Atoi(value.Id)
			if chineseZodiac == value.BettingName {
				gameBettingService.UpdateGameBettingOddsByIdInZodiac(
					int64(id),
					zodiacOddsRelation.NatalOdds,
					value.BettingName)
			} else {
				gameBettingService.UpdateGameBettingOddsById(int64(id), zodiacOddsRelation.NoNatalOdds)
			}
		}
		if "连肖" == value.ItemsName {
			zodiacOddsRelation := zodiacOddsRelationService.SelectByItemsName(value.ItemsName)
			id, _ := strconv.Atoi(value.Id)
			if strings.Contains(value.BettingName, "-") {
				gameBettingService.UpdateGameBettingOddsByIdInZodiac(
					int64(id),
					zodiacOddsRelation.NatalOdds,
					value.BettingName)
			}
			if strings.Contains(value.BettingName, "/") {
				gameBettingService.UpdateGameBettingOddsById(int64(id), zodiacOddsRelation.NoNatalOdds)
			}
		}
	}
}

func (g *GameUtils) QueryYearByMarkSixData(name string) string {
	nowStr := time.Now().Format("2006-01-02")
	nowTime, _ := time.Parse(format, nowStr)
	year := time.Now().Year()
	markSixData := markSixDataService.SelectByYear(strconv.Itoa(year))
	dbData := markSixData.Year + "-" + markSixData.Month + "-" + markSixData.Day
	dbDataTime, _ := time.Parse(format, dbData)
	intYear,_ := strconv.Atoi(markSixData.Year)
	if year == intYear {

		if nowStr == dbData {
			valueMap := make(map[string]string)
			json.Unmarshal([]byte(markSixData.Data), &valueMap)
			return g.SliceToString(getData(name,valueMap))
		}

		beforeMarkSixData := markSixDataService.SelectByYear(strconv.Itoa(year + 1))
		beforeDbData := beforeMarkSixData.Year + "-" + beforeMarkSixData.Month + "-" + beforeMarkSixData.Day
		beforeDbDataTime, _ := time.Parse(format, beforeDbData)
		if nowTime.After(dbDataTime) && nowTime.Before(beforeDbDataTime) {
			valueMap := make(map[string]string)
			json.Unmarshal([]byte(markSixData.Data), &valueMap)
			return g.SliceToString(getData(name,valueMap))
		}

		afterMarkSixData := markSixDataService.SelectByYear(strconv.Itoa(year - 1))
		if !afterMarkSixData.IsEmpty() {
			afterDbData := afterMarkSixData.Year + "-" + afterMarkSixData.Month + "-" + afterMarkSixData.Day
			afterDbDataTime, _ := time.Parse(format, afterDbData)
			if nowTime.Before(dbDataTime) && nowTime.After(afterDbDataTime) {
				//将value转换为map类型
				valueMap := make(map[string]string)
				json.Unmarshal([]byte(afterMarkSixData.Data), &valueMap)
				return g.SliceToString(getData(name,valueMap))
			}
		}
	}
	return ""
}

func getData(name string, mapType map[string]string ) []string {
	var result []string
	if "TIAN_XIAO" == name {
		for index,value := range mapType{
			if strings.Contains(TIAN_XIAO,index){
				result = append(result, value)
			}
		}
	} else if "DI_XIAO" == name {
		for index,value := range mapType{
			if strings.Contains(DI_XIAO,index){
				result = append(result, value)
			}
		}
	} else if "QIAN_XIAO" == name {
		for index,value := range mapType{
			if strings.Contains(QIAN_XIAO,index){
				result = append(result, value)
			}
		}
	} else if "HOU_XIAO" == name {
		for index,value := range mapType{
			if strings.Contains(TIAN_XIAO,index){
				result = append(result, value)
			}
		}
	} else if "JIA_XIAO" == name {
		for index,value := range mapType{
			if strings.Contains(JIA_XIAO,index){
				result = append(result, value)
			}
		}
	} else if "YE_XIAO" == name {
		for index,value := range mapType{
			if strings.Contains(YE_XIAO,index){
				result = append(result, value)
			}
		}
	}
	return result
}

//判断str是否包含nums="1,2"中
func StrIsContains(str string, nums string) bool {
	strArray := strings.Split(nums, ",")
	if strings.Contains(str, strArray[0]) && strings.Contains(str, strArray[1]) {
		return true
	}
	return false
}
//判断str是否包含nums切片["1,2"]中
func SliceCompareString(str1 string, nums []string) bool {
	for _, value := range nums {
		if strings.Contains(str1, value) {
			return true
		}
	}
	return false
}
