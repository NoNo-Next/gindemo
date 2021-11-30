package entity

type ResBody struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (b *ResBody) Fail(data interface{}) ResBody {
	b.Code = 404
	b.Msg = "服务器出错"
	b.Data = data
	return *b
}

func (b *ResBody) Success( msg string ,data interface{}) ResBody {
	b.Code = 200
	b.Msg = "访问成功"
	if msg != "" {
		b.Msg = msg
	}
	b.Data = data
	return *b
}