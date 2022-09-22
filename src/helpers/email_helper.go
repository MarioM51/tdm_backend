package helpers

import (
	"crypto/tls"
	"time"

	simpleMail "github.com/xhit/go-simple-mail/v2"
)

/*
const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzNZPksELBi3PSRnhy3qlTO8PyLk3INYk6g/cbFylbPDoQd6j
5nm/G8cLLiINxrm0NWKtAYXNCv2vxOfz74RtO086+tHMLt2BRxuePJvvcvzFzutR
D12gso0FG8Rcj34Ux+yFlwfpkmYgjq9YLPSdWy3+nFx9pk28RS2DrZdIpbbPCS6i
kjAf8AGGUKHfSJLc8bmIyzNZrl4N4d8ZX3ipArijTSgrmp7c7e4xf90YIljKZtOU
f0M9tyPpeCVISlFmoh+vlGJ+UiW9EL/dmzJc51/TR0cZDswCElXaeLndj1TZx4Cl
Jr02BNFS3HieXRTdOeECWrwcc6BeNJyNElQoFwIDAQABAoIBAFl3zf/Kg5cqURyb
ymzG4AZvcJR6maKlBjCZxuwptzOTMc77gNk9GgT29mIrC5teJ2Ed/XTpzTpcvfYi
XgCi9dbu71L4adeadknpvT652Wd/EqMjUx+EBJmYUL/lD4y5RDhijaL/GL0SEGxi
GABiw4w81bXUDCmkUHRiwd4Dcv61xkPWqSC/BvnQo7V4UMmoUbLtvNjGxHNEAQne
LBkZmezx1eFemUMPKyiGUSB5oSFx/eQxTulV4pPmfixInF1kCcst6G8F/+KLs1KZ
vJRIpu/IOAtYaN4kfB6m6FtytfUiWX7xTydccupCYgto3CNhRzCNAIyG4KCaq/Z3
KxfccYECgYEA/Q3koaqRwo42S3lpyn1JFNQjMmV0HNaAZAXrC9CD2EY2IsPHJ3PW
P3kx03UeIxlzrcIZM1snxj21hES2CXfqtFF+YC1hMEbnAI/0OQoeFCXlKba5xAy0
ccJyJJQ9aH3dqPbzUgjdeBJXdsHfZHItd5soSaOXtNVV1nBEkc1ruFcCgYEAzzi6
0FJNEHi8mBKqHTNy9kvULEnCFqoN+N1mdsp+fLyVXWx2EmawKY3xSZ4UMtmFhgv3
iEehskl1iAchuDw0dfSEewvDKwkFm7ANDVGP9UstO8EuameF5N59M/9zC562VUzo
0CVhg8Yo3f4mpkxogVrAVNLfEQaRCKD2jIioNkECgYBk3wDgIKnxr9acx00QVlin
YNiW4jIivK55MJK9JuUndPVnbjsY0uf4bUsbS3gz7ZVbEiARhKiaMUcF7o3RwGdi
cYm6tNwk7l5urvNfOVU8Gs76jcgHCjlzj1sIkb7YxDNzgt0DOl1t24HZ6PYviAPv
xX2NvRRgFRoeXKo4pHXoCQKBgDvgf5KchWoiCRTEJ+WiLTDf/mIBuhSEdN8ZUnc2
0c/HSj2hjoiIpZSMUFFeXSXIVt3B7XeygxWaRlzU+rhapSoESpencXCo/bbb6xmT
HM7bNynSC1Bxs23LoE/7G0obRUJmo8spUgEarEphGtfosjXWfEbGW/B2fkgJAtTe
1ARBAoGBALrVVLSgIgWLNqboHFuMNA/62Ydql/IWaPi7YKww35i7Sof6FMK64Ug/
c4Z94dKb9cU4ThYVaRsb22z/fzjhGWeLTDZC04npQCOZ+sckJAijwtobtqzQEu4A
5m7UHwUb28nW0nykq4ET6UovX1lHA/An9FLzGiItognPyRzc7GaD
-----END RSA PRIVATE KEY-----
*/

type EmailSender struct {
	server *simpleMail.SMTPServer
}

func NewEmailSender(consts Constants) *EmailSender {
	var server = simpleMail.NewSMTPClient()

	server.Host = consts.SMTP.Host
	server.Port = consts.SMTP.Port
	server.Username = consts.SMTP.Username
	server.Password = consts.SMTP.Password

	server.Encryption = simpleMail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second

	var emailSender = &EmailSender{
		server: server,
	}
	return emailSender
}

func (emailSender *EmailSender) SendEmail(addressee, subject, htmlBody string) error {
	emailSender.server.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// SMTP client
	smtpClient, err := emailSender.server.Connect()

	if err != nil {
		return err
	}

	email := simpleMail.NewMSG()
	email.SetFrom("Tienda de Mario <tiendaDeArtesaniasTlaxco@gmail.com>").
		AddTo(addressee).
		SetSubject(subject).
		//SetListUnsubscribe("<mailto:unsubscribe@example.com?subject=https://example.com/unsubscribe>")
		SetBody(simpleMail.TextHTML, htmlBody)
		//Attach(&mail.File{FilePath: "/path/to/image.png", Name:"Gopher.png", Inline: true})

	/* add dkim signature
	if privateKey != "" {
		options := dkim.NewSigOptions()
		options.PrivateKey = []byte(privateKey)
		options.Domain = "example.com"
		options.Selector = "default"
		options.SignatureExpireIn = 3600
		options.Headers = []string{"from", "date", "mime-version", "received", "received"}
		options.AddSignatureTimestamp = true
		options.Canonicalization = "relaxed/relaxed"

		email.SetDkim(options)
	}
	*/

	// always check error after send
	if email.Error != nil {
		return email.Error
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil

}
