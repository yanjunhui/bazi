package bazi

import (
	"fmt"
	"log"
	"time"
)

// 八字
type TBazi struct {
	SolarDate   TDate       // 新历日期
	LunarDate   TDate       // 农历日期
	BaziDate    TDate       // 八字日期
	PreviousJie TDate       // 上一个节(气)
	NextJie     TDate       // 下一个节(气)
	SiZhu       TSiZhu      // 四柱
}

// 计算
func calc(bazi *TBazi, nSex int) {
	// 通过立春获取当年的年份
	bazi.BaziDate.Year = GetLiChun2(bazi.SolarDate)
	// 通过节气获取当前后的两个节
	bazi.PreviousJie, bazi.NextJie = GetJieQi(bazi.SolarDate)
	// 八字所在的节气是上一个的节气
	bazi.BaziDate.JieQi = bazi.PreviousJie.JieQi
	// 节气0 是立春 是1月
	bazi.BaziDate.Month = bazi.BaziDate.JieQi/2 + 1

	// 通过八字年来获取年柱
	bazi.SiZhu.YearZhu = GetZhuFromYear(bazi.BaziDate.Year)
	// 通过年干支和八字月
	bazi.SiZhu.MonthZhu = GetZhuFromMonth(bazi.BaziDate.Month, bazi.SiZhu.YearZhu.Gan.Value)

	// 通过公历 年月日计算日柱
	bazi.SiZhu.DayZhu = GetZhuFromDay(bazi.SolarDate.Year, bazi.SolarDate.Month, bazi.SolarDate.Day, bazi.SolarDate.Hour)

	//获取时柱
	bazi.SiZhu.HourZhu = GetZhuFromHour(bazi.SolarDate.Hour, bazi.SiZhu.DayZhu.Gan.Value)

	//23点后换日
	if bazi.SolarDate.Hour == 23{
		t, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%.2d-%.2d", bazi.SolarDate.Year, bazi.SolarDate.Month, bazi.SolarDate.Day))
		if err == nil{
			t = t.AddDate(0,0,1)
			bazi.SiZhu.DayZhu = GetZhuFromDay(bazi.SolarDate.Year, bazi.SolarDate.Month, t.Day(), 0)
		}
	}
}

// 从新历获取八字(年, 月, 日, 时, 分, 秒, 性别男1,女0)
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) TBazi {
	var bazi TBazi

	if !GetDateIsValid(nYear, nMonth, nDay) {
		log.Println("无效的日期", nYear, nMonth, nDay)
		return bazi
	}

	// 新历年
	bazi.SolarDate.Year = nYear
	bazi.SolarDate.Month = nMonth
	bazi.SolarDate.Day = nDay
	bazi.SolarDate.Hour = nHour
	bazi.SolarDate.Minute = nMinute
	bazi.SolarDate.Second = nSecond

	// 转农历
	var nTimeStamp = Get64TimeStamp(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	bazi.LunarDate = GetLunarDateFrom64TimeStamp(nTimeStamp)

	// 进行计算
	calc(&bazi, nSex)

	return bazi
}

// 从农历获取八字
func GetBaziFromLunar(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int, isLeap bool) TBazi {
	nYear, nMonth = ChangeLunarLeap(nYear, nMonth, isLeap)

	var bazi TBazi

	if !GetLunarDateIsValid(nYear, nMonth, nDay) {
		log.Println("无效的日期", nYear, nMonth, nDay)
		return bazi
	}

	// 农历年
	bazi.LunarDate.Year = nYear
	bazi.LunarDate.Month = nMonth
	bazi.LunarDate.Day = nDay
	bazi.LunarDate.Hour = nHour
	bazi.LunarDate.Minute = nMinute
	bazi.LunarDate.Second = nSecond

	// 转新历
	var nTimeStamp = GetLunar64TimeStamp(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	bazi.LunarDate = GetDateFrom64TimeStamp(nTimeStamp)

	// 进行计算
	calc(&bazi, nSex)

	return bazi

}

func (bazi *TBazi) String() string {
	return fmt.Sprintf("%s %s %s %s",
		bazi.SiZhu.YearZhu.GanZhi.ToString(),
		bazi.SiZhu.MonthZhu.GanZhi.ToString(),
		bazi.SiZhu.DayZhu.GanZhi.ToString(),
		bazi.SiZhu.HourZhu.GanZhi.ToString())
}
