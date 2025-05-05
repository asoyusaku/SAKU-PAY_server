package response

import (
	"SAKU-PAY/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func LineVerify(token string) (model.Response, error) {
	var response model.Response

	req, err := http.NewRequest("GET", "https://api.line.me/v2.1/profile", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return model.Response{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return model.Response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return model.Response{}, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return model.Response{}, err
	}
	return response, nil
}
