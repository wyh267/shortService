/*************************************************************************
	> File Name: Utils.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com
	> Created Time: 日  6/14 18:05:47 2015
 ************************************************************************/

package shortlib

import (
	"container/list"
	"fmt"
	"time"
)

type CountChannl struct {
	Ok           int64
	CountOutChan chan int64
}

type CreateCountFunc func() (int64, error)

func IsAllowToken(token string) bool {
	return true

}

func IsNormalUrl(url string) bool {
	return true
}

func TransNumToString(num int64) (string, error) {

	startTime := TimeNow()
	var base int64
	base = 62
	baseHex := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	output_list := list.New()
	for num/base != 0 {
		output_list.PushFront(num % base)
		num = num / base
	}
	output_list.PushFront(num % base)
	str := ""
	for iter := output_list.Front(); iter != nil; iter = iter.Next() {
		str = str + string(baseHex[int(iter.Value.(int64))])
	}
	DuringTime(startTime, "TransNumToString")
	return str, nil
}

func TransStringToNum(str string) (int64, error) {

	return 0, nil
}

func TimeNow() time.Time {
	return time.Now()
}

func DuringTime(start time.Time, taskname string) {

	endTime := time.Now()
	fmt.Printf("[INFO] [ %v ] COST Time %v \n", taskname, endTime.Sub(start))

}

func IsShortUrl(short_url string) error {
	return nil
}

func CreateCounter(count_type string, count_chan chan CountChannl, rediscli *RedisAdaptor) CreateCountFunc {

	if count_type == "inner" {
		return func() (int64, error) {
			count_c := make(chan int64)
			ch := CountChannl{0, count_c}
			count_chan <- ch
			count := <-count_c
			return count, nil
		}
	} else {
		return func() (int64, error) {
			count, err := rediscli.NewShortUrlCount()
			if err != nil {
				return 0, err
			}
			return count, nil
		}
	}

}
