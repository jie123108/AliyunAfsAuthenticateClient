package AliyunAfsAuthenticateClient

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AfsAuthenticateClient struct {
	client        *sdk.Client
	captchaAppKey string
}

func NewAfsAuthenticateClient(accessKeyId, accessSecret, regionId string, captchaAppKey string) (afsClient *AfsAuthenticateClient, err error) {
	if regionId == "" {
		regionId = "cn-hangzhou"
	}
	client, err := sdk.NewClientWithAccessKey(regionId, accessKeyId, accessSecret)
	if err != nil {
		return
	}
	afsClient = &AfsAuthenticateClient{}
	afsClient.client = client
	afsClient.captchaAppKey = captchaAppKey

	return
}

type AfsCheckResp struct {
	Code      int
	Msg       string
	RequestId string
	Detail    string
	RiskLevel string
}

func (this *AfsAuthenticateClient) AfsCheck(captchaSessionId, captchaToken, captchaSig, captchaScene, remoteIp string) (resp *AfsCheckResp, err error) {
	request := &requests.RpcRequest{}
	request.InitWithApiInfo("afs", "2018-01-12", "AuthenticateSig", "", "")
	request.Method = "GET"
	request.Domain = "afs.aliyuncs.com"
	request.QueryParams["SessionId"] = captchaSessionId
	request.QueryParams["Token"] = captchaToken
	request.QueryParams["Sig"] = captchaSig
	request.QueryParams["Scene"] = captchaScene
	request.QueryParams["Platform"] = "3"
	request.QueryParams["AppKey"] = this.captchaAppKey
	request.QueryParams["RemoteIp"] = remoteIp

	resp = &AfsCheckResp{Code: 900, Msg: "UNKONW_ERROR"}
	response := &responses.BaseResponse{}

	err = this.client.DoAction(request, response)
	if err != nil {
		fmt.Printf("DoAction failed! err: %v", err)
		return
	}

	err = json.Unmarshal(response.GetHttpContentBytes(), resp)
	return
}
