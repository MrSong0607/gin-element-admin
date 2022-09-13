package dao

import (
	"admin/dto"
	"admin/models"
	"admin/utils"
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_db     *gorm.DB
	_dbInit = sync.Once{}

	_cache     *redis.Pool
	_cacheInit = sync.Once{}
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type DbEntity interface {
	BeginTransaction(trans *gorm.DB)
}

func Db() *gorm.DB {
	_dbInit.Do(func() {
		var err error
		_db, err = gorm.Open(mysql.Open(utils.GetConfig().DbConnectionString), &gorm.Config{PrepareStmt: true, SkipDefaultTransaction: true})
		if err != nil {
			fmt.Println(err)
		}
	})

	if utils.GetConfig().Debug {
		return _db.Debug()
	}

	return _db
}

func Cache() *redis.Pool {
	_cacheInit.Do(func() {
		cfg := utils.GetConfig()

		_cache = &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", cfg.RedisConnectionString, redis.DialDatabase(0))
			},
		}
	})

	return _cache
}

func Migrate() {
	Db().AutoMigrate(&models.Account{}, &models.User{})
}

func BeginTransaction(entity ...DbEntity) *gorm.DB {
	trans := Db().Begin()
	for _, item := range entity {
		item.BeginTransaction(trans)
	}
	return trans
}

func Paginate(q dto.QueryBase) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, size := q.Page, q.Size

		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 20
		}

		if page < 1 {
			page = 1
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
