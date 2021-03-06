package database

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lucaspizzo/api-payment/infrastructure/config"
)

type Repository struct {
	db   *gorm.DB
	once sync.Once
}

func (r *Repository) Start() {
	_ = LoadGormPostGres(
		config.DB_USER,
		config.DB_PASS,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		false)
}

func (r *Repository) Stop() {
	defer r.db.Close()
}

func (r *Repository) GetInstance() *gorm.DB {
	r.once.Do(func() {
		var err error
		r.db, err = GetGormDb()
		if err != nil {
			panic(err.Error())
		}
		r.db.SingularTable(true)
		r.db.LogMode(true)
	})
	return r.db
}
