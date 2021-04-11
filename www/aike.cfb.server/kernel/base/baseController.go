package base;

import (
	"aike-cfb-server/kernel/component/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**

 */
type BaseController struct {

}

/**
	响应
 */
func response(context *gin.Context,responseCode int, responseData map[string]interface{}, headers map[string]string)  {
	apiResponseStructure := &common.ApiResponseStructure{
		Code: responseCode,
		Data:	responseData,
	}

	common.ApiResponseComponent.Response(context, apiResponseStructure, http.StatusOK, headers)
}
