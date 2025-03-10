/*
 * @Author: SpenserCai
 * @Date: 2023-09-29 15:37:14
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-10-19 14:30:01
 * @Description: file content
 */
package api

import (
	"log"
	"os"

	"github.com/phyzical/sd-webui-discord/api/business"
	"github.com/phyzical/sd-webui-discord/api/gen/restapi"
	"github.com/phyzical/sd-webui-discord/api/gen/restapi/operations"
	"github.com/phyzical/sd-webui-discord/api/middleware"
	"github.com/phyzical/sd-webui-discord/global"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
)

func BeforeRun() {
	// 清空原有的命令行参数
	os.Args = os.Args[:1]
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	global.ApiService = operations.NewAPIServiceAPI(swaggerSpec)
	global.ApiService.BearerAuth = middleware.ValidateJwt
	business.BusinessBase{}.SetLoginHandler()
	business.BusinessBase{}.SetAuthHandler()
	business.BusinessBase{}.SetUserInfoHandler()
	business.BusinessBase{}.SetUserHistoryHandler()
	business.BusinessBase{}.SetDiscordServerHandler()
	business.BusinessBase{}.SetOpenDiscordServerHandler()
	business.BusinessBase{}.SetClusterHandler()
	business.BusinessBase{}.SetCommunityHistoryHandler()
	business.BusinessBase{}.SetUserListHandler()
	business.BusinessBase{}.SetSetUserPrivateHandler()
	business.BusinessBase{}.SetSetUserEnableHandler()

}

func StartApiService() {
	BeforeRun()
	server := restapi.NewServer(global.ApiService)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "AiMediaService"
	parser.LongDescription = "AiMediaService API"
	server.ConfigureFlags()
	for _, optsGroup := range global.ApiService.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.Host = global.Config.WebSite.Api.Host
	server.Port = global.Config.WebSite.Api.Port
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
