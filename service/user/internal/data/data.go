package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	slog "log"
	"os"
	"time"
	"user/internal/conf"
	"user/internal/data/ent"

	"github.com/go-redis/redis/extra/redisotel"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewEntDb, NewRedis, NewUserRepo)

// Data .
type Data struct {
	db    *gorm.DB
	entDb *ent.Client
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, entDb *ent.Client, redis *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, redis: redis, entDb: entDb}, cleanup, nil
}

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {

	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			//IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, // 表名是否加 s
		},
	})

	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		panic("failed to connect database")
	}

	// 初始化表
	migrateTables(db)

	return db
}

// NewRedis 初始化redis
func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}

// NewEntDb NewENtDb 初始化db
func NewEntDb(c *conf.Data) *ent.Client {

	client, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func migrateTables(db *gorm.DB) {
	list, err := db.Migrator().GetTables()
	if err != nil {
		panic(err)
	}
	existTable := make(map[string]string)
	for _, s := range list {
		existTable[s] = s
	}

	tableMap := make(map[string]schema.Tabler)
	tableMap["user"] = &User{}
	for tableName, tabler := range tableMap {

		if _, ok := existTable[tableName]; ok {
			continue
		}

		if err = db.Migrator().AutoMigrate(tabler); err != nil {
			log.Fatal(err)
		}
	}
}
