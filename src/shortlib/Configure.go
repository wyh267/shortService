package shortlib

import (
	"errors"
	"github.com/ewangplay/config"
	"strconv"
)

type Configure struct {
	ConfigureMap map[string]string
}

func NewConfigure(filename string) (*Configure, error) {
	config := &Configure{}

	config.ConfigureMap = make(map[string]string)
	err := config.ParseConfigure(filename)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (this *Configure) loopConfigure(sectionName string, cfg *config.Config) error {

	if cfg.HasSection(sectionName) {
		section, err := cfg.SectionOptions(sectionName)
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(sectionName, v)
				if err == nil {
					this.ConfigureMap[v] = options
				}
			}

			return nil
		}
		return errors.New("Parse Error")
	}

	return errors.New("No Section")
}

func (this *Configure) ParseConfigure(filename string) error {
	cfg, err := config.ReadDefault(filename)
	if err != nil {
		return err
	}

	this.loopConfigure("sever", cfg)
	this.loopConfigure("service", cfg)
	this.loopConfigure("log", cfg)
	this.loopConfigure("mysql", cfg)
	this.loopConfigure("redis", cfg)
	this.loopConfigure("weibosender", cfg)
	this.loopConfigure("weixinsender", cfg)

	return nil
}

//服务信息
func (this *Configure) GetPort() (int, error) {

	portstr, ok := this.ConfigureMap["port"]
	if ok == false {
		return 9090, errors.New("No Port set, use default")
	}

	port, err := strconv.Atoi(portstr)
	if err != nil {
		return 9090, err
	}

	return port, nil
}

//监控服务的端口信息
func (this *Configure) GetMPort() (int, error) {

	portstr, ok := this.ConfigureMap["mport"]
	if ok == false {
		return 9999, errors.New("No Port set, use default")
	}

	port, err := strconv.Atoi(portstr)
	if err != nil {
		return 9999, err
	}

	return port, nil
}

//log文件信息
func (this *Configure) GetLogfile() (string, error) {

	logfile, ok := this.ConfigureMap["logfile"]

	if ok == false {
		return "log.log", errors.New("No logfile,use defualt")
	}

	return logfile, nil

}

//服务控制信息
func (this *Configure) GetUrlPattern() (string, error) {

	UrlPattern, ok := this.ConfigureMap["urlpattern"]
	if ok == false || UrlPattern == "" {
		return "/v(\\d)/(contents|control|hello|email|sms|template)/", errors.New("No UrlPattern,use defualt")
	}

	return UrlPattern, nil

}

//数据库连接配置信息
func (this *Configure) GetMysqlUserName() (string, error) {

	mysqlusername, ok := this.ConfigureMap["mysqlusername"]

	if ok == false {
		return "root", errors.New("No mysqlusername,use defualt")
	}

	return mysqlusername, nil
}

func (this *Configure) GetMysqlPassword() (string, error) {

	mysqlpassword, ok := this.ConfigureMap["mysqlpassword"]

	if ok == false {
		return "12345", errors.New("No mysqlpassword,use defualt")
	}

	return mysqlpassword, nil
}

func (this *Configure) GetMysqlHost() (string, error) {

	mysqlhost, ok := this.ConfigureMap["mysqlhost"]

	if ok == false {
		return "127.0.0.1", errors.New("No mysqlhost,use defualt")
	}

	return mysqlhost, nil
}

func (this *Configure) GetMysqlPort() (string, error) {

	mysqlport, ok := this.ConfigureMap["mysqlport"]

	if ok == false {
		return "3306", errors.New("No mysqlport,use defualt")
	}

	return mysqlport, nil
}

func (this *Configure) GetMysqlDBname() (string, error) {

	mysqlDBname, ok := this.ConfigureMap["mysqlDBname"]

	if ok == false {
		return "test", errors.New("No mysqlDBname,use defualt")
	}

	return mysqlDBname, nil
}

func (this *Configure) GetMysqlCharset() (string, error) {

	mysqlcharset, ok := this.ConfigureMap["mysqlcharset"]

	if ok == false {
		return "utf8", errors.New("No mysqlcharset,use defualt")
	}

	return mysqlcharset, nil
}

func (this *Configure) GetMysqlMaxConns() (int, error) {

	mysqlmaxconnsstr, ok := this.ConfigureMap["mysqlmaxconns"]
	if ok == false {
		return 9090, errors.New("No mysqlmaxconns set, use default")
	}

	mysqlmaxconns, err := strconv.Atoi(mysqlmaxconnsstr)
	if err != nil {
		return 2000, err
	}

	return mysqlmaxconns, nil
}

func (this *Configure) GetMysqlMaxIdleConns() (int, error) {

	mysqlmaxidleconnsstr, ok := this.ConfigureMap["mysqlmaxidleconns"]
	if ok == false {
		return 9090, errors.New("No mysqlmaxidleconns set, use default")
	}

	mysqlmaxidleconns, err := strconv.Atoi(mysqlmaxidleconnsstr)
	if err != nil {
		return 1000, err
	}

	return mysqlmaxidleconns, nil
}

func (this *Configure) GetRedisHost() (string, error) {
	redishost, ok := this.ConfigureMap["redishost"]

	if ok == false {
		return "127.0.0.1", errors.New("No redishost,use defualt")
	}

	return redishost, nil
}

func (this *Configure) GetRedisPort() (string, error) {
	redisport, ok := this.ConfigureMap["redisport"]

	if ok == false {
		return "6379", errors.New("No redisport,use defualt")
	}

	return redisport, nil
}

func (this *Configure) GetWeiboSenderHost() (string, error) {
	weibosenderhost, ok := this.ConfigureMap["weibosenderhost"]

	if ok == false {
		return "127.0.0.1", errors.New("No weibosenderhost,use defualt")
	}

	return weibosenderhost, nil
}

func (this *Configure) GetWeiboSenderPort() (string, error) {
	weibosenderport, ok := this.ConfigureMap["weibosenderport"]

	if ok == false {
		return "6379", errors.New("No weibosenderport,use defualt")
	}

	return weibosenderport, nil
}

func (this *Configure) GetWeixinSenderHost() (string, error) {
	weixinsenderhost, ok := this.ConfigureMap["weixinsenderhost"]

	if ok == false {
		return "127.0.0.1", errors.New("No weixinsenderhost,use defualt")
	}

	return weixinsenderhost, nil
}

func (this *Configure) GetWeixinSenderPort() (string, error) {
	weixinsenderport, ok := this.ConfigureMap["weixinsenderport"]

	if ok == false {
		return "6379", errors.New("No weixinsenderport,use defualt")
	}

	return weixinsenderport, nil
}
