package do

import (
	"github.com/aide-family/magicbox/safety"
	"github.com/aide-family/sovereign/internal/biz/vobj"
	"gorm.io/gorm"
)

type Namespace struct {
	BaseModel

	Name     string                      `gorm:"column:name;type:varchar(100);not null;uniqueIndex"`
	Metadata *safety.Map[string, string] `gorm:"column:metadata;type:json;"`
	Status   vobj.GlobalStatus           `gorm:"column:status;type:integer;not null;default:0"`
}

func (Namespace) TableName() string {
	return "namespaces"
}

func (n *Namespace) BeforeCreate(tx *gorm.DB) (err error) {
	if n.BaseModel.BeforeCreate(tx) != nil {
		return
	}
	if !n.Status.Exist() || n.Status.IsUnknown() {
		n.Status = vobj.GlobalStatusEnabled
	}
	return
}
