package data

import (
	"time"
	"url-shorten/app/shorten/internal/conf"

	mgorm "github.com/HarryBird/mo-kit/db/gorm"
	mlog "github.com/HarryBird/mo-kit/log/kratos/gorm"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewShortenRepo)

// Data .
type Data struct {
	db  *gorm.DB
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

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "mod", "repo.data"))

	d := &Data{
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}
