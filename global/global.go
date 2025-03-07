/*
 * @Author: SpenserCai
 * @Date: 2023-08-16 11:05:26
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-09-29 19:08:51
 * @Description: file content
 */
package global

import (
	"github.com/bwmarrin/discordgo"
	"github.com/phyzical/sd-webui-discord/api/gen/restapi/operations"
	"github.com/phyzical/sd-webui-discord/cluster"
	"github.com/phyzical/sd-webui-discord/config"
	"github.com/phyzical/sd-webui-discord/user"
)

var (
	Config         *config.Config
	ClusterManager *cluster.ClusterService
	LongDBotChoice map[string][]*discordgo.ApplicationCommandOptionChoice
	UserCenterSvc  *user.UserCenterService
	ApiService     *operations.APIServiceAPI
)
