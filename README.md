AliyunAfsAuthenticateClient 阿里云人机验证golang验证接口实现. [阿里云人机验证接口文档](https://help.aliyun.com/document_detail/66340.html?spm=a2c4g.11186623.6.575.4d6d196cvtY5al)
----

### 安装
```shell
go get github.com/jie123108/AliyunAfsAuthenticateClient
```

### 使用示例

```golang
package main

import (
	"fmt"
	afsAuthClient "github.com/jie123108/AliyunAfsAuthenticateClient"
)

func main() {
	var accessKeyId = "你的阿里云AccesskeyId"
	var accessSecret = "你的阿里云accessSecret"
	var captchaAppKey = "人机验证的AppKey"

	var scene, nc_token, sessionid, sig string

	scene = "nc_message_h5"
	nc_token = "captcha token"
	sessionid = "captcha sessionid"
	sig = "captcha sig"

	remoteIp := "127.0.0.1"
	afsClient, err := afsAuthClient.NewAfsAuthenticateClient(accessKeyId, accessSecret, "", captchaAppKey)
	resp, err := afsClient.AfsCheck(sessionid, nc_token, sig, scene, remoteIp)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("resp: %+v\n", resp)
}

正常情况下, 输出如下: 
{Code:100 Msg:pass_1 RequestId:xxxxxxxx Detail:{"sigSource":0} RiskLevel:}
```