// Code generated by 'freedom new-crud'
package object

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cart struct {
	changes map[string]interface{}
	Id      int       `gorm:"primary_key" column:"id"`
	UserId  int       `gorm:"column:user_id"`  // 用户ID
	GoodsId int       `gorm:"column:goods_id"` // 商品id
	Num     int       `gorm:"column:num"`      // 数量
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}

func (obj *Cart) TableName() string {
	return "cart"
}

// TakeChanges .
func (obj *Cart) TakeChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// SetUserId .
func (obj *Cart) SetUserId(userId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.UserId = userId
	obj.changes["user_id"] = userId
}

// SetGoodsId .
func (obj *Cart) SetGoodsId(goodsId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.GoodsId = goodsId
	obj.changes["goods_id"] = goodsId
}

// SetNum .
func (obj *Cart) SetNum(num int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Num = num
	obj.changes["num"] = num
}

// SetCreated .
func (obj *Cart) SetCreated(created time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Created = created
	obj.changes["created"] = created
}

// SetUpdated .
func (obj *Cart) SetUpdated(updated time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Updated = updated
	obj.changes["updated"] = updated
}

// AddUserId .
func (obj *Cart) AddUserId(userId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.UserId += userId
	obj.changes["user_id"] = gorm.Expr("user_id + ?", userId)
}

// AddGoodsId .
func (obj *Cart) AddGoodsId(goodsId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.GoodsId += goodsId
	obj.changes["goods_id"] = gorm.Expr("goods_id + ?", goodsId)
}

// AddNum .
func (obj *Cart) AddNum(num int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Num += num
	obj.changes["num"] = gorm.Expr("num + ?", num)
}
