package configMgmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ZentaoClient struct {
	BaseURL string
	Token   string
}

// 创建一个新的ZenTao客户端
func NewZentaoClient(baseUrl string) *ZentaoClient {
	return &ZentaoClient{
		BaseURL: baseUrl,
	}
}

// 获取访问令牌
func (zc *ZentaoClient) GetToken(username, password string) (string, error) {
	url := fmt.Sprintf("%s/user-login", zc.BaseURL)
	payload := map[string]string{"account": username, "password": password}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	token, ok := result["token"].(string)
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}

// 创建一个新的bug
func (zc *ZentaoClient) CreateBug(projectID int, title, content string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/bug-create", zc.BaseURL)
	payload := map[string]interface{}{
		"project": projectID,
		"title":   title,
		"content": content,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+zc.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//func main() {
//	client := NewZentaoClient("http://your-zentao-url.com/api")
//
//	// 获取访问令牌
//	token, err := client.GetToken("your-username", "your-password")
//	if err != nil {
//		fmt.Println("Error getting token:", err)
//		return
//	}
//	client.Token = token
//
//	// 创建一个新的bug
//	bug, err := client.CreateBug(1, "这是一个测试标题", "这是bug的详细描述")
//	if err != nil {
//		fmt.Println("Error creating bug:", err)
//		return
//	}
//	fmt.Println("New bug created:", bug)
//}
