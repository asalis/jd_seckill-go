package service

import (
	"github.com/blinkbean/dingtalk"
	"github.com/unknwon/goconfig"
	"github.com/asalis/jd_seckill/log"
)

type Dingtalk struct {
	conf *goconfig.ConfigFile
}

func NewDingtalk(conf *goconfig.ConfigFile) *Dingtalk {
	return &Dingtalk{conf: conf}
}

func (this *Dingtalk) Send(title, msg string) error {
	cli := dingtalk.InitDingTalkWithSecret(
		this.conf.MustValue("dingtalk", "access_token", ""),
		this.conf.MustValue("dingtalk", "secret", ""),
	)
	markdown := []string{
		"### " + title,
		"---------",
		msg,
	}
	log.Println("正在发送通知...")
	err := cli.SendMarkDownMessageBySlice(title, markdown)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("钉钉机器人推送成功")
	}

	return nil
}
