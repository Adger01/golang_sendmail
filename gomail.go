package main

import "gopkg.in/gomail.v2"

//使用gomail进行发邮件
func main() {
	//smtp服务器
	host := "smtp.126.com"
	//smtp端口
	post := 465
	//发件人
	from := "xxxx@126.com"
	//发件人密码
	password := "xxxxxx"
	//接收人
	to := "xxxx@qq.com"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "测试邮件标题")
	m.SetBody("text/html", "测试邮件内容")

	d := gomail.NewDialer(host, post, from, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
