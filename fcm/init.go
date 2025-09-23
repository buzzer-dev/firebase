package fcm

import (
	"github.com/buzzer-dev/firebase/fcm/config"
	"context"
	"log/slog"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	ctx := context.TODO()
	opt := option.WithCredentialsFile(config.GetConfig(ctx).Firebase.CredentialFile)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		slog.ErrorContext(ctx, "fail to init fcm", "error", err)
		os.Exit(1)
	}
}
