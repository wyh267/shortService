/*************************************************************************
	> File Name: ShortProcessor.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/

package main

import (
	"fmt"
	"net/http"
	"shortlib"
)

type ShortProcessor struct {
	*shortlib.BaseProcessor
}

func (this *ShortProcessor) ProcessRequest(method,request_url string, params map[string]string, body []byte, w http.ResponseWriter,r *http.Request) error {

	err:=shortlib.IsShortUrl(request_url)
	if err != nil {
		return err
	}

	original_url,err := this.GetOriginalURL(request_url)
	if err != nil {
		return err
	}

	fmt.Printf("REQUEST_URL: %v --- ORIGINAL_URL : %v \n",request_url,original_url)
	http.Redirect(w,r,original_url,http.StatusMovedPermanently)	
	return nil
}



func (this *ShortProcessor) GetOriginalURL(request_url string)(string,error){

	original_url,err := this.Lru.GetOriginalURL(request_url)
	//没有从LRU获取到地址
	if err != nil {
		return "",err
	}

	return original_url,nil



}
