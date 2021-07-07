package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// ローカル環境
const LOG_FOLDER string = `C:\Users\wenchao.che\Desktop\api_server_log`

// AWSテスト環境(111)
//const LOG_FOLDER string = `C:\Users\BPOServersAdmin\Desktop\api_server_log`

// 本番、検証環境
//const LOG_FOLDER string = `C:\Users\Administrator\Desktop\api_server_log`

// 操作者ID, 内容
const LOG_FORMAT_NORMAL string = "Info: %s %s"

// 操作者ID, 内容, エラーコード, エラー詳細
const LOG_FORMAT_ERROR string = "Error: %s %s(エラーコード: %d  エラー詳細: %s)"

func SendMail(destMail string) bool {

	result := true

	from := "chi8@chi8.store"
	password := "System123"

	// smtp server configuration.
	smtpHost := "smtp.exmail.qq.com"
	smtpPort := "465"

	// verify code (4 digits)
	verifyCode := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	e := email.NewEmail()
	e.From = "吃吧 <" + from + ">"
	e.To = []string{destMail}
	e.Subject = "邮箱验证"
	e.Text = []byte("验证码：" + verifyCode)

	err := e.SendWithTLS(smtpHost+":"+smtpPort, auth, &tls.Config{ServerName: smtpHost})
	if err != nil {
		fmt.Println(err.Error())
		result = false
	}

	return result

}

// Convert a string encoding from UTF-8 to ShiftJIS
func ToShiftJIS(str string) (string, error) {
	return TransformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewEncoder())
}

func TransformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

// ログ出力
func WriteLog(content string) {

	var logPath string = fmt.Sprintf(`%s\log_%s.txt`, LOG_FOLDER, time.Now().Format("2006-01-02"))

	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.Ldate|log.Ltime)
	logContent := fmt.Sprintf("%s", content)
	logger.Println(logContent)
}
