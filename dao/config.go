package dao

import (
	"fmt"
	"log"
	"netdisk/tool"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBMgr db总会话
var DBMgr *gorm.DB

// 初始化
func init() {
	db, err := gorm.Open(tool.Conf.DataSource.SourceName, tool.Conf.DataSource.String())
	if err != nil {
		log.Fatal(fmt.Errorf("DB failure: %w", err))
		return
	}
	original := db.DB()
	original.SetMaxOpenConns(15)
	DBMgr = db

}