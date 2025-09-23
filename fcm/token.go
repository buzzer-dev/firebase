package fcm

import (
	"time"

	"gorm.io/gorm"
)

type FcmToken struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	Token     string
	Platform  *string
	IP        *string
}

func (f FcmToken) GetUserToken(db *gorm.DB) (token string, err error) {
	err = db.Model(f).Unscoped().
		Model(FcmToken{}).
		Select("token").
		Where("user_id=?", f.UserID).
		Order("updated_at DESC").
		Limit(1).
		Scan(&token).Error

	return
}
