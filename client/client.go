package client

import (
	"go-ent-demo/client/driver"
	"go-ent-demo/config"
	"go-ent-demo/entcore"

	"context"
	"fmt"
	"go-ent-demo/hooks"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"

	"go-ent-demo/util/log"
	"go-ent-demo/util/strutil"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var (
	dbOnce sync.Once
)

const (
	dbTypeMySQL      = "MySQL"
	dbTypePostgreSQL = "PostgreSQL"
)

var coreDBClient *entcore.Client

func CoreDBClient() *entcore.Client {
	dbOnce.Do(func() {
		cfg := config.GetConfig()
		coreDBClient = initDB(cfg.Database)
	})

	return coreDBClient
}

func CoreDBWithTx(c context.Context, client *entcore.Client, fn func(tx *entcore.Tx) error) error {
	tx, err := client.Tx(c)

	if err != nil {
		log.Errorf("start transaction failed: %v", err)
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Errorf("rolling back transaction: %v", err)
			err = errors.Wrapf(err, "rolling back transaction: %v", rollbackErr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Errorf("committing transaction: %v", err)
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

func CloseClient() error {
	if err := CoreDBClient().Close(); err != nil {
		log.Errorf("db close error:%v", err)
		return err
	}

	return nil
}

func initDB(cfg *config.DatabaseConfig) *entcore.Client {
	var client *entcore.Client

	switch cfg.Type {
	case dbTypeMySQL:
		client = initMySQL(cfg)
	case dbTypePostgreSQL:
		client = initPostgreSQL(cfg)
	default:
		//	log.Errorf("unsupported database type: %s", cfg.Type)
		panic("unsupported database type: " + cfg.Type)
	}

	// 设置Hooks
	client.Use(hooks.CreateUpdateHook())
	// client.Use(hooks.LoggingHook())

	// 设置拦截器 Interceptors
	// client.Intercept(interceptors.DataScopeQuery()) TODO

	if config.GetConfig().App.Env == "prod" {
		log.Info("run in prod mode, SQL isn't print")
	} else {
		client = client.Debug()
	}

	return client
}

func initMySQL(cfg *config.DatabaseConfig) *entcore.Client {
	// 对密码进行 URL 编码，处理密码中含有特殊字符如 `@` 的情况
	password := strutil.URLEncoding(string(cfg.Password))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		cfg.Username,
		password,
		cfg.Address,
		cfg.Port,
		cfg.DBName,
		cfg.ConnConfig)

	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime))
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime))

	if err := db.Ping(); err != nil {
		log.Errorf("failed to connect db: %v", err)
		panic(err)
	}

	client := entcore.NewClient(entcore.Driver(driver.NewEntDriverWrapper(drv)))
	//client := entcore.NewClient(entcore.Driver(drv), entcore.Log(logging))
	//client := entcore.NewClient(entcore.Driver(drv))
	return client
}

func initPostgreSQL(cfg *config.DatabaseConfig) *entcore.Client {
	// 对密码进行 URL 编码，处理密码中含有特殊字符
	password := strutil.URLEncoding(string(cfg.Password))

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		cfg.Username,
		password,
		cfg.Address,
		cfg.Port,
		cfg.DBName,
		cfg.ConnConfig,
	)

	drv, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Minute)

	client := entcore.NewClient(entcore.Driver(driver.NewEntDriverWrapper(drv)))

	return client
}

func initSQLite(cfg *config.DatabaseConfig) *entcore.Client {
	// dsn := "file::memory:?cache=shared&_foreign_keys=on" // 使用内存
	dsn := fmt.Sprintf("file:%s?_foreign_keys=on", "")

	drv, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Minute)

	client := entcore.NewClient(entcore.Driver(driver.NewEntDriverWrapper(drv)))

	return client
}
