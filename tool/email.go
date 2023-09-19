package tool

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func LoginByEmail(userEmail, acode string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", Conf.Server.Name, Conf.Email.Username)
	e.To = []string{userEmail}
	e.Subject = "验证码"
	e.HTML = []byte(fmt.Sprintf(loginEmailHtml, "登录", acode))
	if err := e.SendWithTLS(Conf.Email.SMTP.String(), smtp.PlainAuth("", Conf.Email.Username, Conf.Email.AuthCode, Conf.Email.SMTP.Host), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         Conf.Email.SMTP.Host,
	}); err != nil {
		return err
	}
	return nil
}

func ResetPasswordEmail(userEmail, acode string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", Conf.Server.Name, Conf.Email.Username)
	e.To = []string{userEmail}
	e.Subject = "验证码"
	e.HTML = []byte(fmt.Sprintf(loginEmailHtml, "重置密码", acode))
	if err := e.SendWithTLS(Conf.Email.SMTP.String(), smtp.PlainAuth("", Conf.Email.Username, Conf.Email.AuthCode, Conf.Email.SMTP.Host), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         Conf.Email.SMTP.Host,
	}); err != nil {
		return err
	}
	return nil
}

var loginEmailHtml = `
<head>
<base target="_blank" />
<style type="text/css">::-webkit-scrollbar{ display: none; }</style>
<style id="cloudAttachStyle" type="text/css">#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>
<style id="blockquoteStyle" type="text/css">blockquote{display:none;}</style>
<style type="text/css">
	body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}
	td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}
	pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%}
	th,td{font-family:arial,verdana,sans-serif;line-height:1.666}
	img{ border:0}
	header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}
	blockquote{margin-right:0px}
</style>
</head>
<body tabindex="0" role="listitem">
<table width="700" border="0" align="center" cellspacing="0" style="width:700px;">
<tbody>
<tr>
	<td>
		<div style="width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;">
			<table border="0" cellpadding="0" cellspacing="0" width="700" height="39" style="font:12px Tahoma, Arial, 宋体;">
				<tbody><tr><td width="210"></td></tr></tbody>
			</table>
		</div>
		<div style="width:680px;padding:0 10px;margin:0 auto;">
			<div style="line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;">
				<strong style="display:block;margin-bottom:15px;">尊敬的用户：<span style="color:#f60;font-size: 16px;"></span>您好！</strong>
				<strong style="display:block;margin-bottom:15px;">
					您正在进行<span style="color: red">%s</span>操作，请在验证码输入框中输入：<span style="color:#f60;font-size: 24px">%s</span>，以完成操作。
				</strong>
			</div>
		</div>
		<div style="width:700px;margin:0 auto;">
			<div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
				<p>此为系统邮件，请勿回复<br>
					请保管好您的邮箱，避免账号被他人盗用
				</p>
				<p>netdisk</p>
			</div>
		</div>
	</td>
</tr>
</tbody>
</table>
</body>
`