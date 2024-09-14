package configMgmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccessTokenResponse struct {
	Access_token string `json:"access_token"`
}
type AIResponse struct {
	ID               string `json:"id"`
	Object           string `json:"object"`
	Created          int    `json:"created"`
	Result           string `json:"result"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
}

func chatWithAI(question string, roleDescription string) (string, error) {
	accessToken := "24.5f49c62da0c44cf057adae93da0b43dd.2592000.1728439234.282335-115511015"
	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/ernie-speed-128k?access_token=%s", accessToken)

	payload := map[string][]map[string]interface{}{
		"messages": {
			{
				"role":        "user",
				"content":     question,
				"system":      roleDescription,
				"temperature": 0.5,
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	// 添加调试输出
	fmt.Printf("Request Payload:\n%s\n", string(jsonPayload))
	fmt.Printf("Response Body:\n%s\n", string(body))
	var response AIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	return response.Result, nil
}

//func main() {
//	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/ernie-speed-128k?access_token=%s", "24.5f49c62da0c44cf057adae93da0b43dd.2592000.1728439234.282335-115511015")
//
//	payload := []byte(`{
//		"messages": [
//			{
//				"role": "user",
//				"content": "介绍一下北京"
//			}
//		]
//	}`)
//
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
//	if err != nil {
//		fmt.Println("Error creating request:", err)
//		return
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Error making request:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Error reading response body:", err)
//		return
//	}
//	fmt.Println("Hello, World!")
//	fmt.Println(string(body))
//}
