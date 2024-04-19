package config

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
)

type Pool struct {
	Conn *gorm.DB
}

var (
	instance *Pool
	once     sync.Once
)

func init() {
	pool := &Pool{}
	_, err := pool.NewPool()
	if err != nil {
		panic(err)
	}
}

func NewConnection() *Pool {
	return &Pool{}
}

func (*Pool) NewPool() (*Pool, error) {
	var err error

	once.Do(func() {
		db, dbErr := gorm.Open(sqlite.Open("data/jw-board.db"), &gorm.Config{})
		if dbErr != nil {
			fmt.Println("Connection Error: ", err)
			err = dbErr
		}

		db.AutoMigrate(&entity.PersonRole{}, &entity.Person{}, &entity.Role{}, &entity.Assignment{})

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
