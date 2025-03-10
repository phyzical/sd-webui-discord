/*
 * @Author: SpenserCai
 * @Date: 2023-09-29 21:26:43
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-10-19 14:09:38
 * @Description: file content
 */
package business

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/phyzical/sd-webui-discord/api/gen/models"
	ServiceOperations "github.com/phyzical/sd-webui-discord/api/gen/restapi/operations/user"
	"github.com/phyzical/sd-webui-discord/global"
	DbotUser "github.com/phyzical/sd-webui-discord/user"
)

func (b BusinessBase) SetUserInfoHandler() {
	global.ApiService.UserUserInfoHandler = ServiceOperations.UserInfoHandlerFunc(func(params ServiceOperations.UserInfoParams, principal interface{}) middleware.Responder {
		userInfo, err := global.UserCenterSvc.GetUserInfo(principal.(DbotUser.UserInfo).Id)
		if err != nil {
			return ServiceOperations.NewUserInfoOK().WithPayload(&models.UserInfo{
				Code:    -1,
				Message: err.Error(),
			})
		}
		count, err := global.UserCenterSvc.GetUserImageTotal(principal.(DbotUser.UserInfo).Id)
		if err != nil {
			return ServiceOperations.NewUserInfoOK().WithPayload(&models.UserInfo{
				Code:    -1,
				Message: err.Error(),
			})
		}
		return ServiceOperations.NewUserInfoOK().WithPayload(&models.UserInfo{
			Code:    0,
			Message: "success",
			Data: &models.UserInfoData{
				User: &models.UserItem{
					ID:       userInfo.Id,
					Username: userInfo.Name,
					Avatar: func() string {
						if userInfo.Avatar == "" {
							return "https://cdn.discordapp.com/embed/avatars/0.png"
						}
						return userInfo.Avatar
					}(),
					Enable:       userInfo.Enable,
					IsPrivate:    userInfo.IsPrivate,
					StableConfig: userInfo.StableConfig,
					Roles:        userInfo.Roles,
					Created:      userInfo.Created,
					ImageCount:   int32(count),
				},
			},
		})
	})
}
