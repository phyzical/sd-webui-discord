/*
 * @Author: SpenserCai
 * @Date: 2023-10-19 14:18:07
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-10-19 14:21:03
 * @Description: file content
 */
package business

import (
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/phyzical/sd-webui-discord/api/gen/models"
	ServiceOperations "github.com/phyzical/sd-webui-discord/api/gen/restapi/operations/admin"
	"github.com/phyzical/sd-webui-discord/global"
	DbotUser "github.com/phyzical/sd-webui-discord/user"
)

func (b BusinessBase) SetSetUserEnableHandler() {
	global.ApiService.AdminSetUserEnableHandler = ServiceOperations.SetUserEnableHandlerFunc(func(params ServiceOperations.SetUserEnableParams, principal interface{}) middleware.Responder {
		if !strings.Contains(principal.(DbotUser.UserInfo).Roles, "admin") {
			return ServiceOperations.NewSetUserEnableOK().WithPayload(&models.BaseResponse{
				Code:    -1,
				Message: "permission denied",
			})
		}
		err := global.UserCenterSvc.SetUserEnable(params.Body.UserID, params.Body.IsEnable)
		if err != nil {
			return ServiceOperations.NewSetUserEnableOK().WithPayload(&models.BaseResponse{
				Code:    -1,
				Message: err.Error(),
			})
		}
		return ServiceOperations.NewSetUserEnableOK().WithPayload(&models.BaseResponse{
			Code:    0,
			Message: "success",
		})
	})
}
