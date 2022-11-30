package entity

import (
	"backend/config"
	"backend/handler/user"
	"backend/handler/user/entity/database"
	"backend/logs"
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	goutil "github.com/muhammadrivaldy/go-util"
)

type UserEntity struct {
	UserRepo     user.IUserRepo
	UserTypeRepo user.IUserTypeRepo
}

func NewUserEntity(conf config.Configuration) (UserEntity, error) {

	clientMysql, err := goutil.NewMySQL(
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Schema.User.Address,
		conf.Database.Schema.User.Database,
		strings.Split(conf.Database.Parameters, ","))
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return UserEntity{}, err
	}

	driver, err := mysql.WithInstance(clientMysql, &mysql.Config{})
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return UserEntity{}, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		conf.Database.Schema.User.MigrationPath,
		conf.Database.Schema.User.Database,
		driver)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return UserEntity{}, err
	}

	// do a migration up
	m.Up()

	dbGorm, err := goutil.NewGorm(clientMysql, "mysql", goutil.LoggerSilent)
	if err != nil {
		logs.Logging.Error(context.Background(), err)
		return UserEntity{}, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Address,
		DB:   0,
	})

	return UserEntity{
		UserRepo:     database.NewUserRepo(dbGorm, redisClient),
		UserTypeRepo: database.NewUserTypeRepo(dbGorm),
	}, nil
}
