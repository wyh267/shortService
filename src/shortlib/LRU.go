/*************************************************************************
	> File Name: LRU.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com
	> Created Time: 一  6/15 17:07:37 2015
 ************************************************************************/

package shortlib

import (
	"container/list"
	"errors"
	"fmt"
)

type UrlElement struct {
	Original string
	Short    string
}

type LRU struct {
	HashShortUrl  map[string]*list.Element
	HashOriginUrl map[string]*list.Element
	ListUrl       *list.List
	RedisCli      *RedisAdaptor
}

func NewLRU(redis_cli *RedisAdaptor) (*LRU, error) {

	lru := &LRU{make(map[string]*list.Element), make(map[string]*list.Element), list.New(), redis_cli}
	fmt.Printf("New LRU Monery Pool\n")
	return lru, nil
}

func (this *LRU) GetOriginalURL(short_url string) (string, error) {

	element, ok := this.HashOriginUrl[short_url]
	//没有找到key,从Redis获取
	if !ok {
		return "", errors.New("No URL")
	}

	//调整位置
	this.ListUrl.MoveToFront(element)
	ele, _ := element.Value.(UrlElement)
	return ele.Original,nil
}

func (this *LRU) GetShortURL(original_url string) (string, error) {

	element, ok := this.HashOriginUrl[original_url]
	//没有找到key，返回错误，重新生成url
	if !ok {
		return "", errors.New("No URL")
	}

	//调整位置
	this.ListUrl.MoveToFront(element)
	ele, _ := element.Value.(UrlElement)
	/*
		fmt.Printf("Short_Url : %v \n",short_url)

		for iter:=this.ListUrl.Front();iter!=nil;iter=iter.Next(){
			fmt.Printf("Element:%v ElementAddr:%v\n",iter.Value,iter)
		}
		fmt.Printf("\n\n\n")
		for key,value := range this.HashUrl{
			fmt.Printf("Key:%v ==== Value:%v\n",key,value)
		}
	*/
	return ele.Short, nil

}

func (this *LRU) SetURL(original_url, short_url string) error {

	ele := this.ListUrl.PushFront(UrlElement{original_url, short_url})
	this.HashShortUrl[short_url] = ele
	this.HashOriginUrl[original_url] = ele
	//	fmt.Printf("PushFront:::: ELEMENT : %v.\t HASH: %v\n",ele,this.HashUrl[original_url])
	return nil

}

func (this *LRU) checkList() error {

	return nil
}
