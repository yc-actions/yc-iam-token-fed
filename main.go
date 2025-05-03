package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	ctx := context.Background()
	gCtx, err := githubactions.Context()
	if err != nil {
		githubactions.Fatalf("failed to get context: %v", err)
	}
	audience := `https://github.com/` + gCtx.RepositoryOwner

	idToken, err := githubactions.GetIDToken(ctx, audience)
	if err != nil {
		githubactions.Fatalf("failed to get ID token: %v", err)
	}

	iamToken, err := postIDToken(idToken, audience)
	if err != nil {
		githubactions.Fatalf("failed to get IAM token: %v", err)
	}
	githubactions.SetOutput("access_token", iamToken.AccessToken)
	githubactions.AddMask(iamToken.AccessToken)
	githubactions.Infof("IAM token fetched successfully")
}

type IAMTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func postIDToken(token string, audience string) (*IAMTokenResponse, error) {
	uri := "https://auth.yandex.cloud/oauth/token"

	// Construct the URL-encoded data
	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange")
	data.Set("requested_token_type", "urn:ietf:params:oauth:token-type:access_token")
	data.Set("audience", audience)
	data.Set("subject_token", token)
	data.Set("subject_token_type", "urn:ietf:params:oauth:token-type:id_token")

	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", body)
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	var iamTokenResponse IAMTokenResponse
	err = json.Unmarshal(body, &iamTokenResponse)

	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}
	return &iamTokenResponse, nil
}
