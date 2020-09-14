package dao

import (
	"fmt"
	"testing"
)

func TestInnerGameInfo(t *testing.T) {
	result := InnerGameInfo("AA", "1")
	for _,value := range result{
		fmt.Println(value)
	}

}
