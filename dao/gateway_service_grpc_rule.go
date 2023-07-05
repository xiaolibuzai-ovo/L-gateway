package dao

import (
	"context"
	"gorm.io/gorm"
)

type GrpcRule struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	ServiceID      int64  `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	Port           int    `json:"port" gorm:"column:port" description:"端口	"`
	HeaderTransfor string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (t *GrpcRule) TableName() string {
	return "gateway_service_grpc_rule"
}

func (t *GrpcRule) Find(ctx context.Context, db *gorm.DB, search *GrpcRule) (*GrpcRule, error) {
	model := &GrpcRule{}
	err := db.WithContext(ctx).Where(search).Find(model).Error
	return model, err
}

func (t *GrpcRule) Save(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *GrpcRule) ListByServiceID(ctx context.Context, db *gorm.DB, serviceID int64) ([]GrpcRule, int64, error) {
	var list []GrpcRule
	var count int64
	query := db.WithContext(ctx)
	query = query.Table(t.TableName()).Select("*")
	query = query.Where("service_id=?", serviceID)
	err := query.Order("id desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, 0, err
	}
	return list, count, nil
}
