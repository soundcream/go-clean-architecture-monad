package entity

import (
	datetimeutil "n4a3/clean-architecture/app/base/util/datetime"
	"time"
)

type BaseEntity struct {
	Id          int       `gorm:"primaryKey" column:"id" json:"id"`
	CreatedBy   string    `column:"created_by" len:"50" json:"createdBy" example:"system"`
	UpdatedBy   *string   `column:"updated_by" len:"50" json:"updatedBy" example:"admin_a"`
	CreatedDate time.Time `column:"created_date" json:"createdDate" example:"2020-01-01"`
	UpdatedDate time.Time `column:"updated_date" json:"updatedDate" example:"2020-01-01"`
}

func (e *BaseEntity) GetId() int {
	return e.Id
}

func (e *BaseEntity) SetInserter(user string) {
	e.CreatedDate = time.Now()
	e.CreatedBy = user
}

func (e *BaseEntity) SetUpdater(user string) {
	e.UpdatedDate = time.Now()
	e.UpdatedBy = &user
}

func (e *BaseEntity) Base() *BaseEntity {
	return e
}

func NewBase() *BaseEntity {
	return &BaseEntity{}
}

func NewBaseWithId(id int) *BaseEntity {
	return &BaseEntity{
		Id: id,
	}
}

type SoftDeleteEntity struct {
	IsDelete       bool       `column:"is_delete" json:"isDelete"`
	DeletedDate    *time.Time `column:"deleted_date" json:"deletedDate" example:"2020-01-01"`
	DeletedBy      *string    `column:"deleted_by" len:"50" json:"deletedBy" example:"admin_a"`
	DeleteReason   *string    `column:"deleted_reason" len:"100" json:"deletedReason" example:"Because Duplicated"`
	RestoredDate   *time.Time `column:"restore_date" json:"restoreDate" example:"2020-01-01"`
	RestoredBy     *string    `column:"restore_by" len:"50" json:"restoreBy" example:"admin_a"`
	RestoredReason *string    `column:"restored_reason" len:"100" json:"restoredReason" example:"Because wrong deleted"`
}

func (e *SoftDeleteEntity) SetDelete(user string, reason string) {
	e.IsDelete = true
	e.DeletedDate = datetimeutil.NowPtr()
	e.DeletedBy = &user
	e.DeleteReason = &reason
}

func (e *SoftDeleteEntity) SetRestore(user string, reason string) {
	e.IsDelete = false
	e.RestoredDate = datetimeutil.NowPtr()
	e.RestoredBy = &user
	e.RestoredReason = &reason
}

type Activation struct {
	IsActive        bool       `column:"is_active" json:"isActive" example:"false"`
	ActivatedDate   *time.Time `column:"activated_date" json:"activatedDate" example:"2020-01-01"`
	ActivatedBy     *string    `column:"activated_by" len:"50" json:"activatedBy" example:"admin_a"`
	DeactivatedDate *time.Time `column:"deactivated_date" json:"deactivatedDate" example:"2020-01-01"`
	DeactivatedBy   *string    `column:"deactivated_by" len:"50" json:"deactivatedBy" example:"admin_a"`
}

func (e *Activation) SetActivate(user string, reason string) {
	e.IsActive = true
	e.ActivatedDate = datetimeutil.NowPtr()
	e.ActivatedBy = &user
}

func (e *Activation) SetDeActivate(user string, reason string) {
	e.IsActive = false
	e.DeactivatedDate = datetimeutil.NowPtr()
	e.DeactivatedBy = &user
}
