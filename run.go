package main

import (
	"flag"
	"strings"
	"net/url"
	"math/rand"
	"time"
	"os"
)

type resource struct {
	url string
	target string
	start int
	end int
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

func ruleResource() []resource  {
	var res []resource
	r1 := resource{
		url: "http://localhost:8888/",
		target:"",
		start:0,
		end:0,
	}
	r2 := resource{
		url: "http://localhost:8888/list/{$id}.html",
		target:"{$id}",
		start:1,
		end:21,
	}
	r3 := resource{
		url: "http://localhost:8888/movie/{$id}.html",
		target:"{$id}",
		start:1,
		end:12000,
	}
	res = append(append(append(res, r1), r2), r3)
	return res
}

func buildUrl(res []resource) []string {
	var list []string
	for _, resItem := range res {
		if len(resItem.target)==0 {
			list = append(list, resItem.url)
		}else {
			for i := resItem.start; i <= resItem.end; i++ {
				urlStr := strings.Replace(resItem.url,resItem.target,string(i),-1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}

func makeLog(current,refer,ua string) string {
	u := url.Values{}
	u.Set("time","1")
	u.Set("url",current)
	u.Set("refer",refer)
	u.Set("ua",ua)
	paramStr := u.Encode()
	logTemplate := "127.0.0.1 - - [24/Jul/2018:23:22:36 +0800] \"OPTIONS /dig?{$paramStr} HTTP/1.1\" 200 43 \"-\" \"{$ua}\""
	log := strings.Replace(logTemplate,"{$paramStr}",paramStr,-1)
	log = strings.Replace(log,"{$ua}",ua,-1)

	return log
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min)+min
}

func main() {
	total := flag.Int("total",100,"说明:日志总数")
	filePath := flag.String("filePath","./dig.log","说明:url 路径")
	// 需要 parse 来使生效
	flag.Parse()

	// 构造真实的 url 集合
	res := ruleResource()
	list := buildUrl(res)

	// 按照要求,生成$total 行日志
	var logStr string
	for i := 0; i <= *total; i++ {
		currentUrl := list[randInt(0,len(list)-1)]
		referUrl := list[randInt(0,len(list)-1)]
		ua := uaList[randInt(0,len(uaList)-1)]
		logStr = logStr + makeLog(currentUrl,referUrl,ua)+"\n"
	}
	// 覆盖写
	//ioutil.WriteFile(*filePath,[]byte(logStr),0644)
	fd, err := os.OpenFile(*filePath,os.O_RDWR|os.O_APPEND|os.O_CREATE,0644)
	if err != nil {
		println(err.Error())
	}
	fd.Write([]byte(logStr))
	fd.Close()
	println("done")
}
