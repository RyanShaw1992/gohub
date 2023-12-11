package sms

import (
	"gohub/pkg/config"
	"sync"
)

// Message是短信的结构体
type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

// SMS短信发送操作类
type SMS struct {
	Driver Driver
}

// once单例模式
var once sync.Once

// internalSMS内部使用的SMS对象
var internalSMS *SMS

// NewSMS单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}

func (sms *SMS)Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
