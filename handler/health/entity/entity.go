package database

import (
	"backend/config"
	health "backend/handler/health"
	db "backend/handler/health/entity/database"
	"backend/logs"
	"context"
	"strings"

	goutil "github.com/muhammadrivaldy/go-util"
)

type HealthEntity struct {
	HealthRepo health.IHealthRepo
}

func NewHealthEntity(conf config.Configuration) (HealthEntity, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Schema.Security.Address,
		conf.Database.Schema.Security.Database,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return HealthEntity{}, err
	}

	dbGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return HealthEntity{}, err
	}

	return HealthEntity{
		HealthRepo: db.NewHealthRepo(dbGorm),
	}, nil
}
