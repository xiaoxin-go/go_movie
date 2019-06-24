package libs

type HttpResponse struct{
	Code int `json:"code"`
	HttpData interface{} `json:"data"`
	Message string `json:"message"`
}

func (this *HttpResponse) Success(data interface{}, message string)*HttpResponse{
	this.Code = 200
	this.HttpData = data
	this.Message = message
	return this
}

func (this *HttpResponse) ServerError(message string)*HttpResponse{
	this.Code = 500
	this.Message = message
	return this
}

func (this *HttpResponse) ParamsError(message string)*HttpResponse{
	this.Code = 400
	this.Message = message
	return this
}

type MyError struct{
	Message string
}

func (this *MyError) Error()string{
	return this.Message
}