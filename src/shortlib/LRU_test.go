/*************************************************************************
	> File Name: LRU_test.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com
	> Created Time: 一  6/15 19:18:06 2015
 ************************************************************************/

package shortlib

import (
	"testing"
)

func Test_SetURL(t *testing.T) {

	lru, _ := NewLRU(nil)
	err := lru.SetURL("key6", "value6")
	err = lru.SetURL("key5", "value5")
	err = lru.SetURL("key4", "value4")
	err = lru.SetURL("key3", "value3")
	err = lru.SetURL("key2", "value2")
	err = lru.SetURL("key1", "value1")
	lru.GetShortURL("key3")
	if err != nil {
		t.Error("Fail....", err)
	} else {
		t.Log("OK...\n")
	}

}
