/*************************************************************************
	> File Name: Router.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com 
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/
 
 package shorlib
 
 import (
	"fmt"
	"net/http"
)

type Router struct {
	Configure  *Configure
}

const {
	SHORT_URL = 0
	ORIGINAL_URL = 1
	UNKOWN_URL = 2
}

//路由设置
//数据分发
func (this *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	action,err := this.ParseUrl(r.RequestURI)
	if err != nil {
		fmt.Printf("[ERROR]parse url fail : %v \n",err)
	}
	switch action{
		//请求的是短连接，需要返回跳转的原始连接
		case SHORT_URL:
		
		//请求的是长连接，申请一个短连接
		case ORIGINAL_URL:
		default:
			fmt.Printf("[ERROR]Unknow url...:%v \n",r.RequestURI)
	}
	
	return 
}


func (this *Router) ParseURL(url string) (action int, err error) {

	if isShortUrl(url){
		return SHORT_URL
	}else{
		return ORIGINAL_URL
	}

}


func (this *Router) isShortUrl(url string) bool{
	
	short_url_pattern := "XXXXX"
	url_reg_exp, err := regexp.Compile(short_url_pattern)
	if err != nil {
		return false
	}
	short_match := urlRegexp.FindStringSubmatch(url)
	if short_match == nil {
		return true
	}
	
	return false
	
}