package database

import (
	"backend/handler/security"
	"backend/handler/security/models"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type access struct {
	dbGorm *gorm.DB
	redis  *redis.Client
}

func NewAccessRepo(dbGorm *gorm.DB, redis *redis.Client) security.IAccessRepo {
	return access{dbGorm: dbGorm, redis: redis}
}

func (a access) InsertAccess(req models.Access) (res models.Access, err error) {
	err = a.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (a access) SelectAccessByFilter(req util.FilterQuery) (res models.Access, err error) {

	query, arguments := req.BuildQuery()

	result, _ := a.redis.Get(context.Background(), fmt.Sprintf("select-access:%s", query)).Bytes()

	if result == nil {

		if err = a.dbGorm.Where(query, arguments...).First(&res).Error; err != nil {
			return
		}

		result, _ = json.Marshal(res)

		a.redis.Set(context.Background(), fmt.Sprintf("select-access:%s", query), result, 5*time.Minute)

		return
	}

	if err = json.Unmarshal(result, &res); err != nil {
		return
	}

	return

}

func (a access) UpdateAccess(req models.Access) (res models.Access, err error) {
	err = a.dbGorm.Model(&models.Access{}).Where(`id = ?`, req.ID).Updates(req).First(&res).Error
	return
}
