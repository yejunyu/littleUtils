package main

import (
	"flag"
	"strings"
	"net/url"
	"math/rand"
	"time"
	"io/ioutil"
	"fmt"
	"strconv"
)

type resource struct {
	url    string
	target string
	start  int
	end    int
}

var uaList = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.96 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.2; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.96 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.96 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.96 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.96 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:53.0) Gecko/20100101 Firefox/53.0",
	"Mozilla/5.0 (Windows NT 6.2; Win64; x64; rv:53.0) Gecko/20100101 Firefox/53.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:53.0) Gecko/20100101 Firefox/53.0",
	"Mozilla/5.0 (X11; Linux x86_64) Gecko/20100101 Firefox/53.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) Gecko/20100101 Firefox/53.0",
	"Mozilla/5.0 (Windows NT 6.1; rv:39.0) Gecko/20100101 Firefox/39.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393",
}

var ipList = []string{
	"120.92.118.60",
	"120.92.43.21",
	"120.92.116.150",
	"120.92.45.215",
	"120.92.118.123",
	"120.92.88.23",
	"218.60.8.99",
	"112.115.57.20",
	"121.40.225.209",
	"121.42.167.160",
}

var monList = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

func genRanTime() string {
	var year = strconv.Itoa(randInt(2017, 2018))
	var mon = monList[randInt(0, len(monList)-1)]
	var day = strconv.Itoa(randInt(1, 31))
	if len(day) == 1 {
		day = "0" + day
	}
	var hour = strconv.Itoa(randInt(0, 23))
	if len(hour) == 1 {
		hour = "0" + string(hour)
	}
	var min = strconv.Itoa(randInt(0, 59))
	if len(min) == 1 {
		min = "0" + string(min)
	}
	var sec = strconv.Itoa(randInt(0, 59))
	if len(sec) == 1 {
		sec = "0" + string(sec)
	}
	return fmt.Sprintf("[%s/%s/%s:%s:%s:%s +0800]", day, mon, year, hour, min, sec)
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{
		url:    "http://localhost:8888/",
		target: "",
		start:  0,
		end:    0,
	}
	r2 := resource{
		url:    "http://localhost:8888/videos/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    12000,
	}
	r3 := resource{
		url:    "http://localhost:8888/acticle/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    12000,
	}
	res = append(append(append(res, r1), r2), r3)
	return res
}

func buildUrl(res []resource) []string {
	var list []string
	for _, resItem := range res {
		if len(resItem.target) == 0 {
			list = append(list, resItem.url)
		} else {
			for i := resItem.start; i <= resItem.end; i++ {
				urlStr := strings.Replace(resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}

func makeLog(current, refer, ua, ip string) string {
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramStr := u.Encode()
	logTemplate := "{$ip} - - {$time} \"OPTIONS /dig?{$paramStr} HTTP/1.1\" 200 {$flow} \"-\" \"{$ua}\""
	log := strings.Replace(logTemplate, "{$paramStr}", paramStr, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)
	log = strings.Replace(log, "{$ip}", ip, -1)
	log = strings.Replace(log,"{$time}",genRanTime(),-1)
	log = strings.Replace(log,"{$flow}",strconv.Itoa(randInt(10,10000)),-1)

	return log
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min + 1
}

func main() {
	total := flag.Int("total", 100, "说明:日志总数")
	filePath := flag.String("filePath", "./dig.log", "说明:url 路径")
	// 需要 parse 来使生效
	flag.Parse()

	// 构造真实的 url 集合
	res := ruleResource()
	list := buildUrl(res)

	// 按照要求,生成$total 行日志
	var logStr string
	for i := 0; i <= *total; i++ {
		currentUrl := list[randInt(0, len(list)-1)]
		referUrl := list[randInt(0, len(list)-1)]
		ua := uaList[randInt(0, len(uaList)-1)]
		ip := ipList[randInt(0, len(ipList)-1)]
		logStr = logStr + makeLog(currentUrl, referUrl, ua, ip) + "\n"
	}
	// 覆盖写
	ioutil.WriteFile(*filePath, []byte(logStr), 0644)
	//fd, err := os.OpenFile(*filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	//if err != nil {
	//	println(err.Error())
	//}
	//fd.Write([]byte(logStr))
	//fd.Close()
	println("done")
}
