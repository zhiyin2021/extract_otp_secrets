package main

import (
	"encoding/base32"
	"encoding/base64"
	"extract_otp/gg"
	"flag"
	"fmt"
	"strings"

	"github.com/skip2/go-qrcode"
	"google.golang.org/protobuf/proto"
)

func main() {
	var data string
	flag.StringVar(&data, "data", "", "input otpauth-migration")
	flag.Parse()

	if data == "" {
		flag.PrintDefaults()
		return
	}
	data = strings.Replace(data, "otpauth-migration://offline?data=", "", -1)
	data = strings.Replace(data, "%2B", "+", -1)
	body, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	var payload gg.MigrationPayload
	err = proto.Unmarshal(body, &payload)
	if err != nil {
		panic(err)
	}
	for _, otp := range payload.OtpParameters {
		secret := base32.StdEncoding.EncodeToString(otp.Secret)
		typ := "hotp"
		if otp.Type == 2 {
			typ = "totp"
		}
		code := fmt.Sprintf("otpauth://%s/%s:%s?secret=%s", typ, otp.Issuer, otp.Name, secret)

		// 生成彩色 ASCII 二维码
		qr, err := qrcode.New(code, qrcode.Low)
		if err != nil {
			panic(err)
		}

		// 将二维码转换为 ASCII 字符
		ascii := qr.ToSmallString(false)

		// 输出二维码
		fmt.Printf("\n%s\n%s\n", code, ascii)

	}
}
