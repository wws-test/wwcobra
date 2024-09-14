package configMgmt

type AccessTokenResponse struct {
	Access_token string `json:"access_token"`
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
