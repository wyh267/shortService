/*************************************************************************
	> File Name: OriginalProcessor.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com 
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/


package main


import (
	"shortlib"
	"net/http"
	"errors"
	"io"
	"encoding/json"
	"fmt"
)

type OriginalProcessor struct {
	BaseProcessor *shortlib.BaseProcessor
}

const POST string = "POST"
const TOKEN string = "token"
const ORIGINAL_URL string = "original"

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
func (this *OriginalProcessor) ProcessRequest(method string,params map[string]string,body []byte,w http.ResponseWriter) error{
	fmt.Printf("[INFO]in OriginalProcessor ProcessRequest")	
	if method != POST{
		return errors.New("Create short url must be POST the information")
	}

	var bodyInfo map[string]interface{}
	err := json.Unmarshal(body,bodyInfo)
	if err != nil{
		return err
	}
	
	token,has_token:=bodyInfo[TOKEN].(string)
	original_url,has_original_url:=bodyInfo[ORIGINAL_URL].(string)

	if !has_token || !has_original_url {
		return errors.New("Post info errors")
	}


	if !shortlib.IsAllowToken(token){
		return errors.New("Token is not allow")
	}

	if !shortlib.IsNormalUrl(original_url){
		return errors.New("Url is not normal")
	}

	short_url,err := this.createUrl(original_url)
	if err!=nil{
		return err
	}

	response,err := this.createResponseJson(short_url)
	if err != nil {
		return err
	}

	io.WriteString(w,response)


	return nil
}


func (this *OriginalProcessor) createUrl(original_url string) (string,error){

	return "string",nil

}


func (this *OriginalProcessor) createResponseJson(short_url string)(string,error){

	return "string",nil
}
