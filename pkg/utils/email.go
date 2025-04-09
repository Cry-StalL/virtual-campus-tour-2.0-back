package utils

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var emailConfig *EmailConfig

// InitEmailConfig 初始化邮件配置
func InitEmailConfig(host string, port int, username, password, from string) {
	emailConfig = &EmailConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

// SendVerificationCode 发送验证码邮件
func SendVerificationCode(to, code string) error {
	if emailConfig == nil {
		return fmt.Errorf("邮件配置未初始化")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/plain", fmt.Sprintf("您的验证码是：%s，有效期为5分钟。", code))

	d := gomail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.Username, emailConfig.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("发送邮件失败：%v", err)
		return fmt.Errorf("邮件发送失败")
	}

	return nil
}
