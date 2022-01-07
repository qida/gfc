package logs

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

const (
	Rb助手 = iota
	Rb调试
	Rb错误
	Rb重要
	Rb监控
	Rb日常
	Rb工作
	Rb打卡
	Rb伙伴
	Rb服务
	Rb正贤
	Rb积分
	Rb宣易
)

func Send2Ding(index int8, content string) (err error) {
	req := httplib.Post("http://api.sunqida.cn/v1/logs/ding")
	req.Param("token", "test")
	req.Param("robot", fmt.Sprintf("%d", index))
	req.Param("msg", content)
	req.DoRequest()
	return
}

func Send2Dingf(index int8, format string, content ...interface{}) (err error) {

	req := httplib.Post("http://api.sunqida.cn/v1/logs/ding")
	req.Param("token", "test")
	req.Param("robot", fmt.Sprintf("%d", index))
	req.Param("msg", fmt.Sprintf(format, content...))
	req.DoRequest()
	return
}
