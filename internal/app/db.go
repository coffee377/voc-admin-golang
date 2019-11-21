package app

import (
	"fmt"
	"github.com/coffee377/voc-admin/internal/app/config"
	"github.com/coffee377/voc-admin/internal/app/errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lexkong/log"
	"strings"
)

// 支持数据库类型
const (
	SQLITE = "sqlite3"
	MYSQL  = "mysql"
)

// 获取 orm 实列
func openDB() (*gorm.DB, error) {
	db := config.Global.DB
	var dialect, source string
	dialect = strings.ToLower(db.Adapter)
	switch dialect {
	case SQLITE:
		if db.Dir == "" {
			source = fmt.Sprintf("%s.db", db.Name)
		} else {
			source = fmt.Sprintf("%s/%s.db", strings.TrimSuffix(db.Dir, "/"), db.Name)
		}
		break
	case MYSQL:
		var params = ""
		for k, v := range db.Params {
			if k == "parsetime" {
				k = "parseTime"
			}
			params += fmt.Sprintf("%s=%v&", k, v)

		}
		params = strings.TrimSuffix(params, "&")
		if params != "" {
			params = "?" + params
		}
		source = fmt.Sprintf("%s:%s@%s(%s:%s)/%s%s", db.Username, db.Password, db.Protocol, db.Host, db.Port, db.Name, params)
		break
	default:
		return nil, errors.NewException("", "不受支持的数据库类型", nil)
	}

	return gorm.Open(dialect, source)
}

func InitDB() error {

	db, err := openDB()
	if err != nil {
		//panic("连接数据库失败")
		log.Errorf(err, "连接数据库失败")
	}
	defer db.Close()

	return err
}
