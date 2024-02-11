package user

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"accesstoken"`
}

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }

// var globalAccessToken string

func LoginUser(c *gin.Context) {
	var requestBody LoginUserRequest

	loginLink := viper.GetString("LOGINLINK")
	appKey := viper.GetString("API_KEY")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody = LoginUserRequest{
		Username: encodeString(requestBody.Username),
		Password: encodeString(requestBody.Password),
	}

	fmt.Println("Encode Username: ", requestBody.Username)
	fmt.Println("Encode Password: ", requestBody.Password)

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parse request body: " + err.Error()})
		return
	}

	req, err := http.NewRequest(http.MethodPost, loginLink, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Error to create request: " + err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("app-key", appKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Error to send request: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Error to read response body: " + err.Error()})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal([]byte(body), &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON response"})
		return
	}

	// var accessTokenResponse AccessTokenResponse
	// if err := json.Unmarshal([]byte(body), &accessTokenResponse); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing AccessTokenResponse"})
	// 	return
	// }

	// globalAccessToken = accessTokenResponse.AccessToken

	c.JSON(http.StatusOK, jsonResponse)
}

const AuthorizationHeaderKey = "authorization"

func GetSchedule(c *gin.Context) {
	getScheduleLink := viper.GetString("GETSCHEDULELINK")
	appKey := viper.GetString("API_KEY")

	stdId := c.Query("stdId")

	accessToken := c.Request.Header.Get(AuthorizationHeaderKey)
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing access token"})
		return
	}

	params := url.Values{}
	params.Add("stdId", stdId)
	params.Add("academicYear", "2566")
	params.Add("semester", "2")

	u, _ := url.ParseRequestURI(getScheduleLink)
	u.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"errors": "Error to create request " + err.Error()})
		return
	}

	fmt.Println("u.String()", u.String())

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-access-token", strings.Split(accessToken, " ")[1])
	req.Header.Set("app-key", appKey)

	fmt.Println("request Header", req.Header)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Error to send request: " + err.Error()})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Error to read response body: " + err.Error()})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal([]byte(body), &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func encodeString(data string) string {
	kuKey := viper.GetString("KUPUBLICKEY")
	kuPublicKey := strings.Replace(kuKey, `\n`, "\n", -1)

	block, _ := pem.Decode([]byte(kuPublicKey))
	if block == nil {
		log.Fatal("failed to decode public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalf("failed to decode public key: %v", err)
	}

	rsaKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("got unexpected key type: %T", rsaKey)
	}

	encryptedData, err := rsa.EncryptOAEP(
		sha1.New(),
		rand.Reader,
		rsaKey,
		[]byte(data),
		nil,
	)
	if err != nil {
		log.Fatalf("failed to encrypted data: %v", err)
	}

	return base64.StdEncoding.EncodeToString(encryptedData)
}
