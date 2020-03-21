// Code generated by 'freedom new-crud'
package object

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	changes    map[string]interface{}
	Id         int       `gorm:"primary_key" column:"id"`
	OrderNo    string    `gorm:"column:order_no"`
	UserId     int       `gorm:"column:user_id"`     // 用户id
	TotalPrice int       `gorm:"column:total_price"` // 总价
	Status     string    `gorm:"column:status"`      // 支付,未支付，发货，完成
	Created    time.Time `gorm:"column:created"`
	Updated    time.Time `gorm:"column:updated"`
}

func (obj *Order) TableName() string {
	return "order"
}

// TakeChanges .
func (obj *Order) TakeChanges() map[string]interface{} {
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

// SetOrderNo .
func (obj *Order) SetOrderNo(orderNo string) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.OrderNo = orderNo
	obj.changes["order_no"] = orderNo
}

// SetUserId .
func (obj *Order) SetUserId(userId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.UserId = userId
	obj.changes["user_id"] = userId
}

// SetTotalPrice .
func (obj *Order) SetTotalPrice(totalPrice int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.TotalPrice = totalPrice
	obj.changes["total_price"] = totalPrice
}

// SetStatus .
func (obj *Order) SetStatus(status string) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Status = status
	obj.changes["status"] = status
}

// SetCreated .
func (obj *Order) SetCreated(created time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Created = created
	obj.changes["created"] = created
}

// SetUpdated .
func (obj *Order) SetUpdated(updated time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Updated = updated
	obj.changes["updated"] = updated
}

// AddUserId .
func (obj *Order) AddUserId(userId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.UserId += userId
	obj.changes["user_id"] = gorm.Expr("user_id + ?", userId)
}

// AddTotalPrice .
func (obj *Order) AddTotalPrice(totalPrice int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.TotalPrice += totalPrice
	obj.changes["total_price"] = gorm.Expr("total_price + ?", totalPrice)
}
