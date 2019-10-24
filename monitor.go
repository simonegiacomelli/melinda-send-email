package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"net/smtp"
	"os"
	"strconv"
)

func main() {

	ininame := "config.ini"
	handle, err := os.OpenFile(ininame, os.O_RDONLY|os.O_CREATE, 0666)
	_ = handle.Close()

	iniconf, err := ini.Load(ininame)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	smtpconf := iniconf.Section("smtp")
	host := smtpconf.Key("host").MustString("smtp.gmail.com")
	user := smtpconf.Key("user").MustString("")
	pass := smtpconf.Key("pass").MustString("")

	port := smtpconf.Key("port").MustInt(587)

	fmt.Println("Hostname:", host)
	fmt.Println("Port:", port)
	to := "simone.giacomelli@gmail.com"
	msg := "From: " + user + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"This is a really unimaginative message, I know."

	host_fqdn := host + ":" + strconv.Itoa(port)

	auth := smtp.PlainAuth("", user, pass, host)

	err = smtp.SendMail(host_fqdn, auth, user, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
