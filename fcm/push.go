package fcm

import (
	"context"
	"log/slog"
	"time"

	"firebase.google.com/go/v4/messaging"
	"gorm.io/gorm"
)

type PushNotification struct {
	gorm.Model
	Type     MessageType
	UserID   uint
	Title    string
	Body     string
	Image    string
	Deeplink string
	FcmID    *string
	PushAt   *time.Time
}

func (p *PushNotification) Push(ctx context.Context, db *gorm.DB, msgType MessageType) (err error) {
	// Add timeout to prevent blocking on OAuth token refresh
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}
	notification := &messaging.Notification{
		Title:    p.Title,
		Body:     p.Body,
		ImageURL: p.Image,
	}

	// Build data payload with deeplink
	data := make(map[string]string)
	if p.Deeplink != "" {
		data["deeplink"] = p.Deeplink
	}

	message := &messaging.Message{
		Notification: notification,
		Data:         data,
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Sound: "default",
				},
			},
		},
	}
	switch msgType {

	case All:
		message.Topic = "allUser"

	case LoginUser:
		message.Topic = "loginUser"

	case Apple:
		message.Topic = "apple"

	case Android:
		message.Topic = "android"

	default:
		if p.UserID != 0 {
			// get User fcm Token
			var fcmToken = FcmToken{UserID: p.UserID}
			message.Token, err = fcmToken.GetUserToken(db)
			if err != nil {
				return
			}
		} else {
			message.Topic = "allUser"
		}
	}

	slog.Info("push message", "userID", p.UserID, "token", message.Token, "title", message.Notification.Title, "body", message.Notification.Body, "image", message.Notification.ImageURL, "deeplink", p.Deeplink)
	response, err := client.Send(ctx, message)
	if err != nil {
		return
	}

	// 儲存
	p.FcmID = &response
	now := time.Now()
	p.PushAt = &now
	err = db.Updates(p).Error
	return
}
