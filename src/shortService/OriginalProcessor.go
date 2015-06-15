/*************************************************************************
	> File Name: OriginalProcessor.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/

package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shortlib"
	"fmt"
)

type OriginalProcessor struct {
	*shortlib.BaseProcessor
	Count_Channl chan CountChannl
}

const POST string = "POST"
const TOKEN string = "token"
const ORIGINAL_URL string = "original"
const SHORT_URL string = "short"

/*
*
*
*
{
	"original" : "http://XX.XX.com/XXTTYD",
	"token" : "DFEdafaeaqh43da"
}
*
*
*/
func (this *OriginalProcessor) ProcessRequest(method string, params map[string]string, body []byte, w http.ResponseWriter) error {
	if method != POST {
		return errors.New("Create short url must be POST the information")
	}
	var bodyInfo map[string]interface{}
	err := json.Unmarshal(body, &bodyInfo)
	if err != nil {
		return err
	}

	token, has_token := bodyInfo[TOKEN].(string)
	original_url, has_original_url := bodyInfo[ORIGINAL_URL].(string)

	if !has_token || !has_original_url {
		return errors.New("Post info errors")
	}

	if !shortlib.IsAllowToken(token) {
		return errors.New("Token is not allow")
	}

	if !shortlib.IsNormalUrl(original_url) {
		return errors.New("Url is not normal")
	}

	short_url, err := this.createUrl(original_url)
	if err != nil {
		return err
	}

	response, err := this.createResponseJson(short_url)
	if err != nil {
		return err
	}

	//add head information
	header := w.Header()
	header.Add("Content-Type", "application/json")
	header.Add("charset", "UTF-8")
	io.WriteString(w, response)

	return nil
}

//
//生成short url
//
//
func (this *OriginalProcessor) createUrl(original_url string) (string, error) {

	start := shortlib.TimeNow()
	short,err :=this.Lru.GetShortURL(original_url)
	if err == nil{
		fmt.Printf("[INFO] Match the short url : %v ===> %v\n",original_url,short)
		return short,nil
	}
/*
	count, err := this.RedisCli.NewShortUrlCount()
	if err != nil {
		return "", err
	}
*/	
	count_c := make(chan int64)
	ch:=CountChannl{0,count_c}
	this.Count_Channl <- ch
	count := <- count_c
	
	shortlib.DuringTime(start, "NewShortUrlCount")
	start = shortlib.TimeNow()
	short_url, err := shortlib.TransNumToString(count)
	if err != nil {
		return "", err
	}
	//将对应关系添加到LRU缓存中
	this.Lru.SetURL(original_url,short_url)
	shortlib.DuringTime(start, "TransNumToString")
	return short_url, nil

}

func (this *OriginalProcessor) createResponseJson(short_url string) (string, error) {

	json_res := make(map[string]interface{})
	json_res[SHORT_URL] = this.HostName + short_url

	res, err := json.Marshal(json_res)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
