# Firebase FCM 使用方式
## 設定
可使用檔案設定及環境設置，如果有環境設置，優先使用環境設置   
使用檔案:
* 需要使用 config.yaml 當作設定檔
* 在設定檔裡面設置 firebase 的憑證 json 位置
```yaml
FIREBASE:
  CREDENTIAL_FILE: xxxx/credential.json
```

使用環境設置:
* 必須要有檔案 config.yaml，內容必須包含
```yaml
FIREBASE:
  CREDENTIAL_FILE:
```
* 環境設置 key: FIREBASE_CREDENTIAL_FILE

## 使用方式
```go
import "github.com/buzzer-dev/firebase/fcm"
```
```go
var db *db *gorm.DB
var userID = user.ID
var title = "這是主旨"          // 可為空
var body = "這是內容"
var image = "http://影像URL"   // 可為空
err := fcm.SaveAndPush(context.TODO(), fcm.User, userID, db, title, body, image)
if err != nil {
  slog.Error("發送推播失敗", "error", err)
}
```
