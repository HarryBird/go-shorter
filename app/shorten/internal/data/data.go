package data

import (
	"time"
	"url-shorten/app/shorten/internal/conf"

	mredis "github.com/HarryBird/mo-kit/cache/redis/goredis"
	mgorm "github.com/HarryBird/mo-kit/db/gorm"
	mlog "github.com/HarryBird/mo-kit/log/kratos/gorm"
	"github.com/go-kratos/kratos/v2/log"
	redis "github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewShortenRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	log *log.Helper
}

// NewDB Init DB Client
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	dblog := log.NewHelper(log.With(logger, "mod", "repo.db"))
	db, err := mgorm.NewDB(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: mlog.New(logger, mlog.Config{
			SlowThreshold: 1000 * time.Millisecond,
			LogLevel:      mlog.Info,
		}),
	})

	if err != nil {
		dblog.Fatalf("repo: failed opening connection to mysql: %v", err)
	}

	if conf.Database.Pool != nil {
		if err := mgorm.WithPool(db, conf.Database.Pool); err != nil {
			dblog.Fatalf("repo: failed set connection pool: %v", err)
		}
	}

	// if err := db.AutoMigrate(&Order{}); err != nil {
	//     log.Fatal(err)
	// }
	return db
}

func NewRedis(conf *conf.Data, logger log.Logger) *redis.Client {
	rlog := log.NewHelper(log.With(logger, "mod", "repo.redis"))
	rdb, err := mredis.NewRedis(conf.Redis, logger)
	if err != nil {
		rlog.Fatalf("repo: failed opening connection to redis %v", err)
	}

	rlog.Infof("repo: redis connection options: %+v", rdb.Options())

	return rdb
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "mod", "repo.data"))

	d := &Data{
		db:  db,
		rdb: rdb,
		log: log,
	}
	return d, func() {

	}, nil
}
