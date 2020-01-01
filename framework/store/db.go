package store

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
}

var (
	groups map[string]group
	once   sync.Once
)

type group map[string]*gorm.DB

type dbConf struct {
	user   string
	pwd    string
	host   string
	port   int
	dbName string
}

func init() {
	var dbConfMap map[string]map[string]dbConf
	if err := viper.UnmarshalKey("db", &dbConfMap); err != nil {
		log.Println("get db config faield")
		panic(err)
	}
	once.Do(func() {
		for groupName, groupCfg := range dbConfMap {
			g := make(map[string]*gorm.DB)

			for instanceRW, instanceCfg := range groupCfg {
				dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
					instanceCfg.user,
					instanceCfg.pwd,
					instanceCfg.host,
					instanceCfg.port,
					instanceCfg.dbName)
				db, err := gorm.Open("mysql", dbDSN)
				if err != nil {
					panic(err)
				}

				db.SingularTable(true)
				db.LogMode(true)
				db.DB().SetMaxIdleConns(10)
				db.DB().SetMaxOpenConns(100)
				g[instanceRW] = db
			}
			groups[groupName] = g
		}

		if _, ok := groups["default"]; !ok {
			panic("Don't found default group")
		}
	})
}

func CloseDB() {
	defer db.Close()
}

// Grp 返回指定实例组
func Grp(name string) group {
	return groups[name]
}

// R 返回"default"实例组的只读实例
// 业务逻辑使用指定实例组的只读实例：db.R().Where().Find()
func R() *gorm.DB {
	return groups["default"]["r"]
}

// R 返回实例组的只读实例
// 业务逻辑使用指定实例组的只读实例：db.Grp("other").R().Where().Find()
func (g group) R() *gorm.DB {
	return g["r"]
}

// R 返回"default"实例组的写实例
// 业务逻辑使用指定实例组的只读实例：db.R().Where().Find()
func W() *gorm.DB {
	return groups["default"]["w"]
}

// W 返回实例组的写实例
// 业务逻辑使用指定实例组的写实例：db.W("default").Set("data", 1)
func (g group) W() *gorm.DB {
	return g["w"]
}
