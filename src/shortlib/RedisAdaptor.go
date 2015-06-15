/*************************************************************************
  > File Name: RedisAdaptor.go
  > Author: Wu Yinghao
  > Mail: wyh817@gmail.com
  > Created Time: 二  6/ 9 15:29:05 2015
 ************************************************************************/
package shortlib

import (
	//"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisAdaptor struct {
	conn   redis.Conn
	config *Configure
}

const SHORT_URL_COUNT_KEY string = "short_url_count"

func NewRedisAdaptor(config *Configure) (*RedisAdaptor, error) {
	redis_cli := &RedisAdaptor{}
	redis_cli.config = config

	host, _ := config.GetRedisHost()
	port, _ := config.GetRedisPort()

	connStr := fmt.Sprintf("%v:%v", host, port)
	fmt.Printf(connStr)
	conn, err := redis.Dial("tcp", connStr)
	if err != nil {
		return nil, err
	}

	redis_cli.conn = conn

	return redis_cli, nil
}

func (this *RedisAdaptor) Release() {
	this.conn.Close()
}

func (this *RedisAdaptor) InitCountService() error {

	_, err := this.conn.Do("SET", SHORT_URL_COUNT_KEY, 0)
	if err != nil {
		return err
	}
	return nil

}

func (this *RedisAdaptor) NewShortUrlCount() (int64, error) {

	count, err := redis.Int64(this.conn.Do("INCR", SHORT_URL_COUNT_KEY))
	if err != nil {
		return 0, err
	}

	return count, nil

}

/*
func (this *RedisAdaptor) GetValue(cid int64, source string) (string, error) {
	key := fmt.Sprintf("%v:%v", cid, source)
	value, err := redis.String(this.conn.Do("GET", key))
	if err != nil {
		return "ERR", err
	}

	return value, nil
}

func (this *RedisAdaptor) Append(cid int64, source, v string) (int64, error) {


	key := fmt.Sprintf("%v:%v", cid, source)
	value := fmt.Sprintf("%v|", v)
	count, err := this.conn.Do("APPEND", key, value)
	if err != nil {
		return 0, err
	}

	res, ok := count.(int64)
	if !ok {
		return 0, errors.New("ERR")
	}

	return res, nil

}

func (this *RedisAdaptor) SetValue(cid int64, source, v string) (int64, error) {

	key := fmt.Sprintf("%v:%v", cid, source)
	count, err := this.conn.Do("SET", key, v)
	if err != nil {
		return 0, err
	}

	res, ok := count.(int64)
	if !ok {
		return 0, errors.New("COUNT ERR")
	}

	return res, nil

}



*/
