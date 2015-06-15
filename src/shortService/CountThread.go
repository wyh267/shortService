/*************************************************************************
	> File Name: CountThread.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com 
	> Created Time: 一  6/15 16:19:18 2015
 ************************************************************************/

package main

import (
	"fmt"
)

type CountChannl struct{
	Ok int64
	CountOutChan chan int64
}


func CountThread(count_chan_in chan CountChannl){

	var count int64
	count = 1000
	fmt.Printf("Running CountThread")
	for{
		fmt.Printf("Get CountChannl")
		select {
		case ok:=<-count_chan_in:
			count=count+1
			ok.CountOutChan <- count

		}

	}
}
