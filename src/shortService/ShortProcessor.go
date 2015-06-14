/*************************************************************************
	> File Name: ShortProcessor.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com 
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/


package main
import (
	"shortlib"
	"net/http"

)

type ShortProcessor struct {
	BaseProcessor *shortlib.BaseProcessor
}


func (this *ShortProcessor)ProcessRequest(method string,params map[string]string,body []byte,w http.ResponseWriter) error{
	
	return nil
}