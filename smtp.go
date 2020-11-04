package main

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

func main()  {
	host := "smtp.126.com"
	port := "25"
	from := "xxx@126.com"
	password := "xxxx"
	to := []string{
		"xxxx@qq.com",
	}

	//构筑消息体
	header := make(map[string]string)
	header["From"] = from
	header["To"] = "xxxx@qq.com"
	header["Subject"] = "当前时段统计报表"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	body := "报表内容一切正常"

	message := ""
	for k,v := range header {
		message += fmt.Sprintf("%s: %s\r\n",k,v)
	}

	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	//预先认证
	auth := smtp.PlainAuth("",from,password,host)
	//发送邮件
	smtp.SendMail(host+":"+port,auth,from,to,[]byte(message))
}