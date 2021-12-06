package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
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

func testcron() {
	fmt.Println("start cron")
	c := cron.New()
	//tim := "0 */1 * * *"
	//榜单结算切换定时任务
	//id, err := c.AddFunc("0 * * * * ?", func() {
	id, err := c.AddFunc("0 30 * * ?", func() {
		pln()
	})
	fmt.Println(id, err)

	c.Start()
}
func pln() {
	fmt.Println("start cron")
}


func testtime2(){
	t1 := time.Now().Add(-time.Minute*1)
	if t1.Unix() >1619431200{
		fmt.Println("true")
	}



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

func main() {
	//listfunc()
	//slicefunc()
	//runabc(2,3)

	//testtime()
	//go testcron()
	//time.Sleep(time.Second*60*60)

	//res, err := testessql(100256)
	//fmt.Println(res, err)

	testtime2()
}
