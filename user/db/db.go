/*
 * @Author: SpenserCai
 * @Date: 2023-08-30 21:21:40
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-10-14 14:55:00
 * @Description: file content
 */
package db

import (
	"errors"
	"log"
	"reflect"
	"strings"

	"github.com/phyzical/sd-webui-discord/config"
	"github.com/phyzical/sd-webui-discord/user/db/db_backend"
	"gorm.io/gorm"
)

type BotDb struct {
	Db *gorm.DB
}

func (botDb *BotDb) Close() error {
	db, err := botDb.Db.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	return err
}

func (botDb *BotDb) CreateOrUpdateDb() error {
	err := botDb.Db.AutoMigrate(&db_backend.UserInfo{})
	if err != nil {
		log.Println("AutoMigrate `user_info` table Error:" + err.Error() + ", try to check db columns")
	}
	err = botDb.Db.AutoMigrate(&db_backend.History{})
	if err != nil {
		log.Println("AutoMigrate `history` table Error:" + err.Error() + ", try to check db columns")
	}
	return nil
}

func NewBotDb(dbCfg *config.DbConfig) (*BotDb, error) {
	dbType := dbCfg.Type

	dbCreateName := "CreateDb" + strings.ToUpper(dbType[:1]) + dbType[1:] + "Connect"
	pkgValue := reflect.ValueOf(db_backend.DbBackend{})
	methodValue := pkgValue.MethodByName(dbCreateName)
	if !methodValue.IsValid() {
		return nil, errors.New("db type not support")
	}

	createDbFunc := methodValue.Interface().(func(string) (*gorm.DB, error))
	db, err := createDbFunc(dbCfg.DSN)
	if err != nil {
		return nil, err
	}

	botDb := &BotDb{
		Db: db,
	}
	err = botDb.CreateOrUpdateDb()
	if err != nil {
		return nil, err
	}
	return botDb, nil
}
