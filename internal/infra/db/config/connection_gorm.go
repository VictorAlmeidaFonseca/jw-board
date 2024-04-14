package config

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Pool struct {
	Conn *gorm.DB
}

var (
	instance *Pool
	once     sync.Once
)

func (*Pool) NewPool() (*Pool, error) {
	var err error

	once.Do(func() {
		db, dbErr := gorm.Open(sqlite.Open("data/jw-board.db"), &gorm.Config{})
		if dbErr != nil {
			err = dbErr
			return
		}
		instance = &Pool{
			Conn: db,
		}
	})

	return instance, err
}

func (*Pool) DestructPool() {
	if instance != nil {
		sqlDB, _ := instance.Conn.DB()
		sqlDB.Close()
	}
}
