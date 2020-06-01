package bazi

import (
	"fmt"
	"testing"
)

func TestGetBazi(t *testing.T) {
	bz := GetBazi(2020,5,31,23,30,0,0)

	fmt.Println(bz.SiZhu.YearZhu.GanZhi.ToString())
	fmt.Println(bz.SiZhu.MonthZhu.GanZhi.ToString())
	fmt.Println(bz.SiZhu.DayZhu.GanZhi.ToString())
	fmt.Println(bz.SiZhu.HourZhu.GanZhi.ToString())
}