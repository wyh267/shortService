/*************************************************************************
	> File Name: Utils.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com 
	> Created Time: 日  6/14 18:05:47 2015
 ************************************************************************/

package shortlib

import (
//	"fmt"
	"container/list"
)



func IsAllowToken(token string) bool {
	return true

}


func IsNormalUrl(url string) bool {
	return true
}



func TransNumToString(num int64)(string,error){
	
	var base int64
	base=62
	baseHex:="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	output_list := list.New()
	for num/base != 0 {
		output_list.PushFront(num%base)
		num=num/base
	}
	output_list.PushFront(num%base)
	str := ""
	for iter:=output_list.Front();iter!=nil;iter=iter.Next(){
		str = str + string(baseHex[int(iter.Value.(int64))])
	}
	return str,nil
}



func TransStringToNum(str string)(int64,error){

	return 0,nil
}



