package objects

import (
	"time"
)

// Goods .
type Goods struct {
	changes map[string]interface{}
	ID      int       `gorm:"primary_key" column:"id"`
	Name    string    `gorm:"column:name"`  // 商品名称
	Price   int       `gorm:"column:price"` // 价格
	Stock   int       `gorm:"column:stock"` // 库存
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}

func (obj *Goods) TableName() string {
	return "goods"
}

// Save .
func (obj *Goods) Save() map[string]interface{} {
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

// SetID .
func (obj *Goods) SetID(iD int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.ID = iD
	obj.changes["id"] = iD
}

// SetName .
func (obj *Goods) SetName(name string) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Name = name
	obj.changes["name"] = name
}

// SetPrice .
func (obj *Goods) SetPrice(price int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Price = price
	obj.changes["price"] = price
}

// SetStock .
func (obj *Goods) SetStock(stock int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Stock = stock
	obj.changes["stock"] = stock
}

// SetCreated .
func (obj *Goods) SetCreated(created time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Created = created
	obj.changes["created"] = created
}

// SetUpdated .
func (obj *Goods) SetUpdated(updated time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Updated = updated
	obj.changes["updated"] = updated
}
