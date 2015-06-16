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

func (this *RedisAdaptor) SetUrl(short_url, original_url string) error {

	key := fmt.Sprintf("short:%v", short_url)
	_, err := this.conn.Do("SET", key, original_url)
	if err != nil {
		return err
	}
	return nil
}

func (this *RedisAdaptor) GetUrl(short_url string) (string, error) {

	key := fmt.Sprintf("short:%v", short_url)
	original_url, err := redis.String(this.conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return original_url, nil
}
