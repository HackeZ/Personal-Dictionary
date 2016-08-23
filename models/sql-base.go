package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	utils "Personal-Dictionary/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var o orm.Ormer

func Connect() {

	var dns string

	// init DB Config.
	dbType := beego.AppConfig.String("DB_TYPE")
	dbHost := beego.AppConfig.String("DB_HOST")
	dbPort := beego.AppConfig.String("DB_PORT")
	dbUser := beego.AppConfig.String("DB_USER")
	dbPass := beego.AppConfig.String("DB_PASS")
	dbName := beego.AppConfig.String("DB_NAME")
	dbPath := beego.AppConfig.String("DB_PATH")
	dbSslmode := beego.AppConfig.String("DB_SSLMODE")

	switch dbType {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
		break
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", dbName, dbHost, dbUser, dbPass, dbPort, dbSslmode)
	case "sqlite3":
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if dbPath == "" {
			dbPath = "./"
		}
		dns = fmt.Sprintf("%s%s.db", dbPath, dbName)
		break
	default:
		beego.Critical("Database driver is not allowed:", dbType)
	}
	orm.RegisterDataBase("default", dbType, dns)
}

// Syncdb 初次启动初始化数据库
func Syncdb() {
	createDB()
	Connect()
	o = orm.NewOrm()
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	insertUser()
	fmt.Println("database init is complete.\nPlease restart the application")
}

func createDB() {

	var dns string
	var sqlstring string

	// init DB Config.
	dbType := beego.AppConfig.String("DB_TYPE")
	dbHost := beego.AppConfig.String("DB_HOST")
	dbPort := beego.AppConfig.String("DB_PORT")
	dbUser := beego.AppConfig.String("DB_USER")
	dbPass := beego.AppConfig.String("DB_PASS")
	dbName := beego.AppConfig.String("DB_NAME")
	dbPath := beego.AppConfig.String("DB_PATH")
	dbSslmode := beego.AppConfig.String("DB_SSLMODE")

	switch dbType {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", dbUser, dbPass, dbHost, dbPort)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", dbName)
		break
	case "postgres":
		dns = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", dbHost, dbUser, dbPass, dbPort, dbSslmode)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", dbName)
	case "sqlite3":
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if dbPath == "" {
			dbPath = "./"
		}
		dns = fmt.Sprintf("%s%s.db", dbPath, dbName)
		os.Remove(dns)
		sqlstring = "create table init (n varchar(32));drop table init;"
		break
	default:
		beego.Critical("Database driver is not allowed:", dbType)
	}

	db, err := sql.Open(dbType, dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", dbName, " created")
	}
	defer db.Close()
}

func insertUser() {
	fmt.Println("=== insert User start... ===")
	u := new(User)
	u.Username = "HackerZ"
	salt := utils.GetNewSalt(8)
	u.Salt = salt
	u.Password = utils.PassEncode("admin", salt)
	u.Email = "767110505@qq.com"

	o := orm.NewOrm()
	o.Insert(u)
	fmt.Println("=== insert User done... ===")
}
