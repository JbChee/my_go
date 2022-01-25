package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"regexp"
	"runtime"
	"strconv"

	//"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func listfunc() {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func slicefunc() {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func runabc(iterNum int, strNum int) {
	var wait, next, over chan struct{}
	var firstwait chan struct{}
	wait = make(chan struct{})
	firstwait = wait
	fmt.Printf("%#v \n", firstwait)
	fmt.Printf("%#v \n", wait)

	for i := 0; i < strNum; i++ {
		next = make(chan struct{})
		over = next
		go echoStr(i, wait, next)
		wait = next
	}

	for i := 0; i < iterNum; i++ {
		firstwait <- struct{}{}
		<-over
	}

	fmt.Printf("%#v\n", wait)
	fmt.Printf("%#v\n", firstwait)

	close(firstwait)

}

func echoStr(threadNum int, wait chan struct{}, next chan struct{}) {

	str := string('A' + threadNum)

	for _ = range wait {
		fmt.Printf("%d : %s \n", threadNum, str)
		next <- struct{}{}
	}

	close(next)
	fmt.Printf("close %d : %s", threadNum, str)

}

func testtime() {

	nowTime := time.Now().Unix()
	startTime := 1617998400
	endTime := 1619812799
	if nowTime < int64(startTime) || nowTime > int64(endTime) {
		fmt.Println("false")
	}

	timeStr := time.Unix(int64(startTime), 0).Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	timeStr1 := time.Unix(int64(endTime), 0).Format("2006-01-02 15:04:05")
	fmt.Println(timeStr1)
	openTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	fmt.Println(openTime)
	closeTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr1, time.Local)
	fmt.Println(closeTime)
	if nowTime <= openTime.Unix() {
		fmt.Println("false")
	}
	t := closeTime.Unix() - nowTime
	fmt.Printf("t = %v", t)
}

//func testcron() {
//	fmt.Println("start cron")
//	c := cron.New()
//	//tim := "0 */1 * * *"
//	//榜单结算切换定时任务
//	//id, err := c.AddFunc("0 * * * * ?", func() {
//	id, err := c.AddFunc("0 30 * * ?", func() {
//		pln()
//	})
//	fmt.Println(id, err)
//
//	c.Start()
//}
func pln() {
	fmt.Println("start cron")
}

func testtime2() {
	t1 := time.Now().Add(-time.Minute * 1)
	fmt.Println(t1.Hour())
}

//test slice
func testSlice() {
	res := make([][]int, 0)
	path := []int{3, 2, 1}
	//path1 := []int{3,2,1}

	res = append(res, append([]int(nil), path...))

	//res = append(res, path1)
	//res = append(path1, res...)  //错误 path1必须是2维的
	fmt.Println("res = ", res)

}

func re() {
	content := "{{user_name}} won gift {{gift_name}}!"
	r, _ := regexp.Compile(`\{\{(.*?)\}\}`)
	ret := r.FindAllString(content, -1)
	tem := content
	for i, item := range ret {
		r1, _ := regexp.Compile(item)
		str := fmt.Sprintf("{{%d}}", i+1)
		rep := r1.ReplaceAllString(tem, str)
		tem = rep
		fmt.Println(tem)

	}

	//fmt.Println(ret)

	fmt.Println(tem)

}

func testessql(uid int64) (int, error) {
	type bigCoinSource struct {
		Uid  int64 `json:"uid"`
		Gold int   `json:"gold"`
	}

	type BigDataCoinDetail struct {
		Source bigCoinSource `json:"_source"`
	}
	type BigDataConsumeIncome struct {
		Total int64               `json:"total"`
		Hits  []BigDataCoinDetail `json:"hits"`
	}
	type BigDataConsumeDetail struct {
		Timeout   bool                 `json:"-"`
		DayIncome BigDataConsumeIncome `json:"hits"`
	}

	t := time.Now()
	stat := BigDataConsumeDetail{}
	//当天查询上一天更新的记录
	dayType1 := t.AddDate(0, 0, -1).Format("200601")   //202104
	dayType2 := t.AddDate(0, 0, -1).Format("20060102") //20210412
	ur := "http://es.dsj.inkept.cn/_sql?%s"
	//paresSql := fmt.Sprintf("select * from app_uid  where appname = 'vs' and index_type = 'anchor_data' and _type='%s' and uid=%d and substring(dt,0,7)='%s'", dayType, data.Uid, data.Date)
	paresSql := fmt.Sprintf("SELECT * FROM vs_coin_log_%s/%s where uid=%d", dayType1, dayType2, uid)
	fmt.Printf("paresSql = %v", paresSql)
	var u url.Values
	u = url.Values{}
	u.Set("sql", paresSql)
	ur = fmt.Sprintf(ur, u.Encode())
	//ur = fmt.Sprintf(ur,paresSql)

	fmt.Printf("urlString %s", ur)
	request, err := http.NewRequest("GET", ur, nil)
	//request, err := http.NewRequest("GET", paresSql, nil)
	//resp, err := http.Get(paresSql)
	if err != nil {
		fmt.Println("http.NewRequest error:%v", err)
		return 0, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 5000 * time.Millisecond}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do error:%v", err)
		return 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("readall failed, err=%s", err)
		return 0, err
	}
	fmt.Printf("body:%v", string(body))
	defer resp.Body.Close()
	err = json.Unmarshal(body, &stat)
	if err != nil {
		fmt.Println("Unmarshal error:%v", err)
		return 0, err
	}
	if len(stat.DayIncome.Hits) > 0 {
		return stat.DayIncome.Hits[0].Source.Gold, nil
	}
	return 0, nil

}

func caculateNum(money int64) {

	addScore := (money / 8) * 100
	addLotterNums := (money / 8) * 1

	moneyChange := float64(money/8) * 0.9999

	fmt.Println(addScore, addLotterNums, moneyChange)
	//fmt.Println(addScore, addLotterNums)

}

func randLiveid() string {
	rand.Seed(time.Now().Unix())
	liveids := []string{"1621834055025730", "xx"}
	fmt.Println(liveids[rand.Intn(len(liveids))])
	return liveids[rand.Intn(len(liveids))]
}

func makeMap() {
	m := make(map[int32]string, 10)
	m[0] = "EDDYCJY1"
	m[1] = "EDDYCJY2"
	m[2] = "EDDYCJY3"
	m[3] = "EDDYCJY4"
	m[4] = "EDDYCJY5"

	for k, v := range m {
		fmt.Printf("k: %v, v: %v", k, v)
	}

	fmt.Println(m)
}

func testTicker() {
	// 使用time.Ticker
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case tickerRunMsg := <-ticker.C:
				fmt.Println("Tick at", tickerRunMsg)
				TimeStop()

			}

		}
	}()
	time.Sleep(time.Second * 6)
}

func TimeStop(){
	time.Sleep(time.Second*6)
	fmt.Println("in Tick at")
}



func test_type(x interface{}) {

	//var a string

	switch v := x.(type) {
	case string:
		fmt.Println("xxx string = ", v)
	}

	var a interface{} = "gggg"
	v, ok := a.(string)
	if ok {
		tem := fmt.Sprintf("xx = %v, %v", v, ok)
		fmt.Println(tem)
	}

}

func ParseGameIdToRound(gameId string) int64 {

	index := 0
	for i, str := range []byte(gameId[8:]) {
		if str != 48 {
			index = i
			break
		}
	}
	round := cast.ToInt64(gameId[8+index:])
	fmt.Println("a=", round)

	ret, _ := strconv.ParseInt(gameId[8:], 10, 64)
	fmt.Println("ret = ", ret)
	return round
}

func test_runtime_Gosched() {

	done := false

	go func() {
		done = true
	}()

	for !done {
		runtime.Gosched()//用于让出CPU时间片
		println("not done !") // 并不内联执行
	}

	println("done !")

}


func GetVipBuyRecTable(ts time.Time) string {
	TABLE_VIP_BUY_REC := "vins_vip_open_record"
	year := ts.Year()
	month := int(ts.Month())
	_dateStr := fmt.Sprintf("%d%02d", year, month)
	fmt.Printf("GetVipBuyRecTable: %v", fmt.Sprintf("%s_%s", TABLE_VIP_BUY_REC, _dateStr))
	return fmt.Sprintf("%s_%s", TABLE_VIP_BUY_REC, _dateStr)
	//return tableNameStr
}

//获取某一天的0点时间
func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

//获取本周周一的日期
func GetSundayOfWeek(t time.Time) (dayStr string) {
	dayObj := GetZeroTime(t)
	if t.Weekday() == time.Monday {
		//修改hour、min、sec = 0后格式化
		dayStr = dayObj.Format("2006_01_02")
	} else {
		offset := int(time.Monday - t.Weekday())
		if offset > 0 {
			offset = -6
		}
		dayStr = dayObj.AddDate(0, 0, offset).Format("2006_01_02")
	}
	fmt.Println(dayStr)
	return
}

func main() {
	//listfunc()
	//slicefunc()
	//runabc(2,3)

	//testtime()
	//go testcron()
	//time.Sleep(time.Second*60*60)

	//res, err := testessql(100256)
	//fmt.Println(res, err)

	//testtime2()

	//re()
	//testSlice()

	//caculateNum(8)
	//屏蔽
	//str := "AZIZI3.5.50_Android"
	//str1 := "AZIZI3.5.30_Iphone"
	//cvVersion := str[5:11]
	//cvVersion1 := str1[5:11]
	//fmt.Println(cvVersion == "3.5.50")
	//fmt.Println(cvVersion1)
	//
	//
	//rand.Seed(time.Now().Unix())
	//liveids := []string{"1621834055025730","xx"}
	//fmt.Println(liveids[rand.Intn(len(liveids))])

	//a := 2288
	//b := 8
	//
	//tem := float64(a) / float64(b)
	//fmt.Println(math.Ceil(tem))
	//a := randLiveid()
	//fmt.Println(a)

	//makeMap()

	//go testTicker()
	//time.Sleep(time.Second * 10)

	//var a interface{} = "gggg"
	//test_type(a)

	//str := "3.5.8"
	//str1 := "3.4.12"
	////a,_ := strconv.ParseFloat(str1, 64)
	////b,_ := strconv.ParseFloat(str, 64)
	////fmt.Println(a,b)
	//if  str> str1{
	//	fmt.Println("xxxxxxxxxxxx")
	//}else{
	//	fmt.Println("uyyyyyyyyy")
	//}

	//str := "1,2"
	//str1 := "1"
	//
	//ret := strings.Split(str,",")
	//ret1 := strings.Split(str1,",")
	//fmt.Println(ret)
	//fmt.Println(ret1)
	////1489582166978
	////1629773714367
	//fmt.Println(time.Unix(0,1629773714367*int64(time.Millisecond)).Format("2006-01-02"))

	//_ = ParseGameIdToRound("20210826000323")

	//test_runtime_Gosched()

	//GetVipBuyRecTable(time.Now())

	//测试时间
	GetSundayOfWeek(time.Now().Add(time.Hour * 24 * -1 ))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24 ))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*2))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*3))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*4))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*5))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*6))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*7))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*8))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*9))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*10))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*11))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*12))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*13))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*14))
	GetSundayOfWeek(time.Now().Add(time.Hour * 24*15))
}

