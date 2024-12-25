package dingtalk

import (
    "bytes"
    "dingding_golang/pkg/config"
    "dingding_golang/pkg/logger"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "sync"
    "time"

    "go.uber.org/zap"

)

type DingTalkClient struct {
    accessToken string
    expireTime time.Time
    mutex      sync.RWMutex
}

var client *DingTalkClient
var once sync.Once

func GetClient() *DingTalkClient {
    once.Do(func() {
        client = &DingTalkClient{}
    })
    return client
}

func (c *DingTalkClient) GetAccessToken() (string, error) {
    c.mutex.RLock()
    if c.accessToken != "" && time.Now().Before(c.expireTime) {
        token := c.accessToken
        c.mutex.RUnlock()
        return token, nil
    }
    c.mutex.RUnlock()

    return c.refreshAccessToken()
}

func (c *DingTalkClient) refreshAccessToken() (string, error) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s",
        config.Get().Dingtalk.AppKey,
        config.Get().Dingtalk.AppSecret)

    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("获取access token失败: %v", err)
    }
    defer resp.Body.Close()

    var result struct {
        ErrCode     int    `json:"errcode"`
        ErrMsg      string `json:"errmsg"`
        AccessToken string `json:"access_token"`
        ExpiresIn   int    `json:"expires_in"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", fmt.Errorf("解析access token响应失败: %v", err)
    }

    if result.ErrCode != 0 {
        return "", fmt.Errorf("获取access token失败: %s", result.ErrMsg)
    }

    c.accessToken = result.AccessToken
    c.expireTime = time.Now().Add(time.Duration(result.ExpiresIn-60) * time.Second)

    return c.accessToken, nil
}

func (c *DingTalkClient) SendMessage(userId string, message interface{}) error {
    token, err := c.GetAccessToken()
    if err != nil {
        return err
    }

    url := fmt.Sprintf("https://oapi.dingtalk.com/chat/send?access_token=%s", token)

    // 构建完整的消息体
    msgBody := map[string]interface{}{
        "touser":  userId,
        "msgtype": message.(map[string]interface{})["msgtype"],
    }
    // 合并其他消息字段
    for k, v := range message.(map[string]interface{}) {
        if k != "msgtype" {
            msgBody[k] = v
        }
    }

    jsonData, err := json.Marshal(msgBody)
    if err != nil {
        return fmt.Errorf("消息序列化失败: %v", err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("发送消息失败: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("读取响应失败: %v", err)
    }

    var result struct {
        ErrCode int    `json:"errcode"`
        ErrMsg  string `json:"errmsg"`
    }

    if err := json.Unmarshal(body, &result); err != nil {
        return fmt.Errorf("解析响应失败: %v", err)
    }

    if result.ErrCode != 0 {
        logger.Error("钉钉发送消息失败",
            zap.Int("errCode", result.ErrCode),
            zap.String("errMsg", result.ErrMsg))
        return fmt.Errorf("钉钉发送消息失败: %s", result.ErrMsg)
    }

    return nil
} 