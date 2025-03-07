/*
 * @Author: SpenserCai
 * @Date: 2023-10-09 12:48:04
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-10-09 12:50:53
 * @Description: file content
 */
package business

import (
	"github.com/go-openapi/runtime/middleware"
	ServiceOperations "github.com/phyzical/sd-webui-discord/api/gen/restapi/operations/system"
	"github.com/phyzical/sd-webui-discord/global"
)

func (b BusinessBase) SetOpenDiscordServerHandler() {
	global.ApiService.SystemOpenDiscordServerHandler = ServiceOperations.OpenDiscordServerHandlerFunc(func(params ServiceOperations.OpenDiscordServerParams) middleware.Responder {
		return ServiceOperations.NewOpenDiscordServerFound().WithLocation(global.Config.Discord.DiscordServerUrl)
	})
}
