package service

import (
	"context"
	"fmt"
	"monitor-video/common"
	"monitor-video/setting"
	"net/http"
	"strings"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/rtc"
	"github.com/qiniu/go-sdk/v7/storage"
)

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
	H int `json:"h"`
	W int `json:"w"`
}

type input struct {
	UserId   string   `json:"userId"`
	Tag      string   `json:"tag,omitempty"`
	Position position `json:"position"`
}

type config struct {
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	Fps         int    `json:"fps"`
	Kbps        int    `json:"kbps"`
	StretchMode string `json:"stretchMode"`
}

type output struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type recordData struct {
	Type    string   `json:"type"`
	Inputs  []input  `json:"inputs"`
	Config  config   `json:"config"`
	Outputs []output `json:"outputs"`
}

type recordRes struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

var (
	accessKey  string
	secretKey  string
	appId      string
	bucketName string
	httpClient *http.Client
	mac        *auth.Credentials
	rtcHost    = "rtc.qiniuapi.com"
)

func init() {
	accessKey = setting.LoadStringParam("qiniu", "ACCESS_KEY")
	secretKey = setting.LoadStringParam("qiniu", "SECRET_KEY")
	appId = setting.LoadStringParam("qiniu", "APP_ID")
	bucketName = setting.LoadStringParam("qiniu", "BUCKET_NAME")
	httpClient = &http.Client{Timeout: 10 * time.Second}
	mac = auth.New(accessKey, secretKey)
}

// https://developer.qiniu.com/rtc/12205/rtc-recorded-directly
func InitStorage(expireDay int) error {
	url := common.BuildURL(rtcHost, "/v3/apps/"+appId)
	data := map[string]interface{}{
		"storage": map[string]interface{}{
			"bucket":    bucketName,
			"expireDay": expireDay,
		},
	}
	info := common.PostReq(httpClient, mac, url, &data, nil)
	return info.Err
}

// https://developer.qiniu.com/rtc/8813/roomToken
func GetRoomToken(roomName, userId, permission string) (token string, err error) {
	manager := rtc.NewManager(mac)
	token, err = manager.GetRoomToken(rtc.RoomAccess{
		AppID:      appId,
		RoomName:   roomName,
		UserID:     userId,
		ExpireAt:   time.Now().Unix() + 8*3600,
		Permission: permission,
	})
	return
}

// https://developer.qiniu.com/rtc/8815/api-room
func CloseRoom(roomName string) error {
	url := common.BuildURL(rtcHost, "/v3/apps/"+appId+"/rooms/"+roomName)
	info := common.DelReq(httpClient, mac, url, nil)
	return info.Err
}

// https://developer.qiniu.com/rtc/12205/rtc-recorded-directly
func RecordTask(roomName, userId string) (recordRes, error) {
	url := common.BuildURL(rtcHost, "/v4/apps/"+appId+"/rooms/"+roomName+"/jobs")
	data := recordData{
		Type: "basic",
		Inputs: []input{
			{UserId: userId, Position: position{X: 0, Y: 0, Z: 0, H: 720, W: 1080}},
			{UserId: userId, Tag: "screen", Position: position{X: 1080, Y: 0, Z: 0, H: 720, W: 1080}},
		},
		Config:  config{Height: 720, Width: 2160, Fps: 25, StretchMode: "aspectFit"},
		Outputs: []output{{Type: "file", URL: roomName + "_" + userId}},
	}
	ret := recordRes{}
	info := common.PostReq(httpClient, mac, url, &data, &ret)
	return ret, info.Err
}

// https://developer.qiniu.com/rtc/12205/rtc-recorded-directly
func SaveVideo(roomName, userId string) error {
	url := common.BuildURL(rtcHost, "/v4/apps/"+appId+"/saveas")
	data := map[string]interface{}{
		"streamName": roomName + "_" + userId,
		"fname":      roomName + "_" + userId + "_video",
		"expireDays": 1,
	}
	info := common.PostReq(httpClient, mac, url, &data, nil)
	return info.Err
}

// https://developer.qiniu.com/dora/1248/audio-and-video-transcoding-avthumb#11
func VideoToMp4(videoName string) (*storage.PfopRet, error) {
	cfg := storage.Config{UseHTTPS: true}
	operationManager := storage.NewOperationManager(mac, &cfg)

	key := strings.Replace(videoName, ".m3u8", ".mp4", 1)
	fopAvthumb := fmt.Sprintf("avthumb/mp4/s/2160x720/vb/500k|saveas/%s", storage.EncodedEntry(bucketName, key))

	fopBatch := []string{fopAvthumb}
	fops := strings.Join(fopBatch, ";")

	force := int64(1)
	notifyURL := "https://api-dev.kp-para.cn/paracraft-compete/v0/competeMonitor/callback"
	pipeline := ""
	persistentType := int64(0)

	request := storage.PfopRequest{
		BucketName: bucketName,
		ObjectName: videoName,
		Fops:       fops,
		NotifyUrl:  notifyURL,
		Force:      force,
		Type:       persistentType,
		Pipeline:   pipeline,
	}

	persistentId, err := operationManager.PfopV2(context.Background(), &request)

	return persistentId, err
}

// https://developer.qiniu.io/kodo/1238/go#8
func CheckMp4Status(persistentId string) (storage.PrefopRet, error) {
	cfg := storage.Config{UseHTTPS: true}
	operationManager := storage.NewOperationManager(mac, &cfg)
	ret, err := operationManager.Prefop(persistentId)
	return ret, err
}
