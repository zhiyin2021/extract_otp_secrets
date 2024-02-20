google authenticator 导出离线密钥解码导入其他 F2A app

```bash
git clone https://github.com/zhiyin2021/extract_otp_secrets
cd extract_otp_secrets
go run . -data=otpauth-migration://offline?data=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

![demo](demo.jpg){width=300}
