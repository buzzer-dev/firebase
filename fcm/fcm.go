package fcm

import (
	"context"

	"gorm.io/gorm"
)

type MessageType uint

const (
	User      MessageType = iota // 指定用户
	All                          // 所有用户
	LoginUser                    // 已登入用户
	Apple                        // Apple 用户
	Android                      // 安卓用户
)

func SaveAndPush(ctx context.Context, db *gorm.DB, msgType MessageType, userID uint, title, body, image string) (err error) {
	push := PushNotification{
		Type:   msgType,
		UserID: userID,
		Title:  title,
		Body:   body,
		Image:  image,
	}
	err = db.Create(&push).Error
	if err != nil {
		return
	}
	return push.Push(ctx, db, msgType)
}

func SaveAndPushWithDeeplink(ctx context.Context, db *gorm.DB, msgType MessageType, userID uint, title, body, image, deeplink string) (err error) {
	push := PushNotification{
		Type:     msgType,
		UserID:   userID,
		Title:    title,
		Body:     body,
		Image:    image,
		Deeplink: deeplink,
	}
	err = db.Create(&push).Error
	if err != nil {
		return
	}
	return push.Push(ctx, db, msgType)
}
