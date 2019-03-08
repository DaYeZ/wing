// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package captcha

import (
	"net/smtp"
	"strings"
	"wing/logger"
)

// Email sender, including smtp authtication and user info
type EmailSender struct {
	smtp.Auth
	identity             string
	user, password, host string
	contentType          string
}

// Start sets TLS to true
func (e EmailSender) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return e.Auth.Start(&s)
}

// SetContentType sets mail content type
func (e *EmailSender) SetContentType(contentType string) *EmailSender {
	e.contentType = contentType
	logger.D("Set content type to:", contentType)
	return e
}

// Send sends a mail to given contacts
func (a *EmailSender) Send(mailto, subject, body string) error {
	from := a.identity
	if from == "" {
		from = a.user
	}
	contacts := strings.Split(mailto, ";")
	content := []byte("" +
		"To: " + mailto + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + a.contentType + "\r\n\r\n" + body)
	return smtp.SendMail(a.host, a, a.user, contacts, content)
}

// NewEmailSender create a email sender for given user
func NewEmailSender(ideneity, user, password, host string) *EmailSender {
	// default content type is html, you may set plain as
	// 'text/plain; charset=UTF-8'
	contentType := "text/html; charset=UTF-8"
	eu := &EmailSender{
		smtp.PlainAuth(ideneity, user, password, host),
		ideneity, user, password, host, contentType,
	}
	logger.D("New a NewEmailSender for", user, "with host:", host)
	return eu
}
