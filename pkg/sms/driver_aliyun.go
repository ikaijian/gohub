// Aliyun Driver
package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	jsoniter "github.com/json-iterator/go"
	"gohub/app/models/sms"
	"gohub/pkg/logger"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Aliyun 实现 sms.Driver interface
type Aliyun struct{}

// Send 实现 sms.Driver interface 的 Send 方法
func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	// 打印配置信息用于调试
	logger.DebugJSON("短信[阿里云]", "配置信息", config)

	client, err := CreateClient(tea.String(config["access_key_id"]), tea.String(config["access_key_secret"]))
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}
	logger.DebugJSON("短信[阿里云]", "配置信息", config)

	param, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "短信模板参数解析错误", err.Error())
		return false
	}
	// 打印请求参数用于调试
	logger.DebugJSON("短信[阿里云]", "请求参数", map[string]string{
		"PhoneNumbers":  phone,
		"SignName":      config["sign_name"],
		"TemplateCode":  message.Template,
		"TemplateParam": string(param),
	})
	// 发送参数
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(config["sign_name"]),
		TemplateCode:  tea.String(message.Template),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(string(param)),
	}
	// 添加更多的运行时选项
	runtime := &util.RuntimeOptions{
		ReadTimeout:    tea.Int(5000),
		ConnectTimeout: tea.Int(5000),
	}

	response, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if response.Body == nil || *response.Body.Code != "OK" {
		if response.Body != nil {
			logger.ErrorString("短信[阿里云]", "服务商返回错误", *response.Body.Message)
		} else {
			logger.ErrorString("短信[阿里云]", "服务商返回错误", "response body is nil")
		}
		return false
	}

	if err != nil {
		var errs = &tea.SDKError{}
		if _t, ok := err.(*tea.SDKError); ok {
			errs = _t
		} else {
			errs.Message = tea.String(err.Error())
		}

		var r dysmsapi20170525.SendSmsResponseBody
		err = json.Unmarshal([]byte(*errs.Data), &r)

		logger.LogIf(err)
		smsModel := sms.Sms{
			Phone:         phone,
			SignName:      config["sign_name"],
			TemplateCode:  message.Template,
			TemplateParam: string(param),
			RequestId:     tea.StringValue(r.RequestId),
			BizId:         "",
			Code:          tea.StringValue(r.Code),
			Message:       tea.StringValue(r.Message),
		}
		smsModel.SaveSmsLog()
		return false
	}
	smsModel := sms.Sms{
		Phone:         phone,
		SignName:      config["sign_name"],
		TemplateCode:  message.Template,
		TemplateParam: string(param),
		RequestId:     tea.StringValue(response.Body.RequestId),
		BizId:         tea.StringValue(response.Body.BizId),
		Code:          tea.StringValue(response.Body.Code),
		Message:       tea.StringValue(response.Body.Message),
	}
	smsModel.SaveSmsLog()

	logger.DebugString("短信[阿里云]", "发送成功", "")
	return true
}

// CreateClient
/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	clientConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		// 访问的域名
		Endpoint: tea.String("dysmsapi.aliyuncs.com"),
	}

	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(clientConfig)

	return _result, _err
}
