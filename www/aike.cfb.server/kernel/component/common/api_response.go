package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {

}

/**
	响应结构体
 */
type ApiResponseStructure struct {
	Code int	`json:"code"`
	Message string	`json:"message"`
	Module string	`json:"module"`
	Data map[string]interface{}	`json:"data"`
	Now int64 `json:"now"`
}

var ApiResponseComponent ApiResponse = ApiResponse{}

/**
	通用响应
 */
func (response ApiResponse) Response(context *gin.Context, apiResponseStructure *ApiResponseStructure, status int, headers map[string]string) {
	// 批量设置header
	for headerKey,headerValue := range headers {
		context.Writer.Header().Set(headerKey, headerValue)
	}
	// 状态处理
	context.Status(status)

	context.JSON(http.StatusOK, apiResponseStructure)
}
