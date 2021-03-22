package base;

import (
	"github.com/gin-gonic/gin"
	"image-cfb-server/kernel/component"
	"net/http"
)

/**

 */
type BaseController struct {

}

func response(context *gin.Context,responseCode int, responseData map[string]interface{}, headers map[string]string)  {
	apiResponseStructure := &component.ApiResponseStructure{
		Code: responseCode,
		Data:	responseData,
	}

	component.ApiResponseComponent.Response(context, apiResponseStructure, http.StatusOK, headers)
}
