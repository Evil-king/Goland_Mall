package utils

import (
	"strconv"
	"strings"
	"time"
)
//获取期号
func GetPeriodNum(flag string, periodNum string, gameCode string) string {
	if "true" == flag {
		str := time.Now().String()
		str2 := strings.Split(str," ")[0]
		str1 := strings.Replace(str2,"-","",-1)
		str3 :=str1[2:len(str1)]
		periodNum = gameCode + str3
		initNum := "0001"
		return periodNum + initNum
	} else {
		num,_ := strconv.Atoi(periodNum[8:len(periodNum)])
		num++
		if len(strconv.Itoa(num)) == 1 {
			return gameCode + getPeriodNumFormat() + "000" + strconv.Itoa(num)
		} else if len(strconv.Itoa(num)) == 2 {
			return gameCode + getPeriodNumFormat() + "00" + strconv.Itoa(num)
		} else if len(strconv.Itoa(num)) == 3 {
			return gameCode + getPeriodNumFormat() + "0" + strconv.Itoa(num)
		} else {
			return gameCode + getPeriodNumFormat()  + strconv.Itoa(num)
		}
	}
}
//期号加1
func AddPeriodNum(periodNum string,gameCode string) string  {
	num,_ := strconv.Atoi(periodNum[8:len(periodNum)])
	num++
	if len(strconv.Itoa(num)) == 1 {
		return gameCode + getPeriodNumFormat() + "000" + strconv.Itoa(num)
	} else if len(strconv.Itoa(num)) == 2 {
		return gameCode + getPeriodNumFormat() + "00" + strconv.Itoa(num)
	} else if len(strconv.Itoa(num)) == 3 {
		return gameCode + getPeriodNumFormat() + "0" + strconv.Itoa(num)
	} else {
		return gameCode + getPeriodNumFormat()  + strconv.Itoa(num)
	}
}

func LastIssuePeriodNum(periodNum string,gameCode string) string  {
	num,_ := strconv.Atoi(periodNum[8:len(periodNum)])
	num--
	if len(strconv.Itoa(num)) == 1 {
		return gameCode + getPeriodNumFormat() + "000" + strconv.Itoa(num)
	} else if len(strconv.Itoa(num)) == 2 {
		return gameCode + getPeriodNumFormat() + "00" + strconv.Itoa(num)
	} else if len(strconv.Itoa(num)) == 3 {
		return gameCode + getPeriodNumFormat() + "0" + strconv.Itoa(num)
	} else {
		return gameCode + getPeriodNumFormat()  + strconv.Itoa(num)
	}
}

func getPeriodNumFormat() string {
	var timeLayoutStr = "2006-01-02"
	t:=time.Now()
	str := t.Format(timeLayoutStr)
	return strings.Replace(str,"-","",-1)[2:len(strings.Replace(str,"-","",-1))]
}

//切片转string
func SliceToString(data interface{}) string {
	var str string
	if v,ok:=data.([]int);ok{
		str += "["
		for k, v := range v {
			if k == 0 {
				str = str + strconv.Itoa(v)
			} else {
				str = str + " " + strconv.Itoa(v)
			}
		}
		str += "]"
	} else if v,ok:=data.([]string);ok{
		str += "["
		for k, v := range v {
			if k == 0 {
				str = str + v
			} else {
				str = str + " " + v
			}
		}
		str += "]"
	} else {
		panic("类型转化异常")
	}
	return str
}

//去除重复
func RemoveDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}