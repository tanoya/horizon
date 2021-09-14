package response

import (
	"net/http"

	"g.hz.netease.com/horizon/lib/log"
	"g.hz.netease.com/horizon/server/middleware/requestid"
	"github.com/gin-gonic/gin"
)

type DataWithTotal struct {
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

type Response struct {
	ErrorCode    string      `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	RequestID    string      `json:"requestId,omitempty"`
}

func NewResponse() *Response {
	return &Response{}
}

func NewResponseWithData(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, NewResponse())
}

func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponseWithData(data))
}

func Abort(c *gin.Context, httpCode int, errorCode, errorMessage string) {
	logger := log.GetLogger(c)
	rid, err := requestid.FromContext(c)
	if err != nil {
		logger.Errorf("error to get requestID from context, err: %v", err)
	}

	c.JSON(httpCode, &Response{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		RequestID:    rid,
	})
	c.Abort()
}

func AbortWithRequestError(c *gin.Context, errorCode, errorMessage string) {
	Abort(c, http.StatusBadRequest, errorCode, errorMessage)
}

func AbortWithInternalError(c *gin.Context, errorCode, message string) {
	Abort(c, http.StatusInternalServerError, errorCode, message)
}

func AbortWithNotFoundError(c *gin.Context, errorCode, message string) {
	Abort(c, http.StatusNotFound, errorCode, message)
}
