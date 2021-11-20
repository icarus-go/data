package result

import "pmo-test4.yz-intelligence.com/kit/data/result/constant"

type ApiResult struct {
	Code   constant.ResponseType  `json:"code" swaggertype:"integer"` //  Code 错误码 , {"0":"成功"}
	Msg    string                 `json:"msg"`                        //  Msg 消息
	Data   interface{}            `json:"data"`                       //  Data 数据
	DataKV map[string]interface{} `json:"-"`                          //  DataKV KV返回
}
