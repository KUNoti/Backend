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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }

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

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"result": string(body),
	})
}

func encodeString(data string) string {
	// kuKey := "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAytOhlq/JPcTN0fX+VqObE5kwIaDnEtso2KGHdi9y7uTtQA6pO4fsPNJqtXOdrcfDgp/EQifPwVRZpjdbVrD6FgayrQQILAnARKzVmzwSMDdaP/hOB6i9ouKsIhN9hQUmUhbhaMkh7UXoxGW+gCSK8dq0+FJVnlt1dtJByiVAJRi2oKSdLRqNjk8yGzuZ6SrEFzAgYZwmQiywUF6V1ZaMUQDz8+nr9OOVU3c6Z2IQXCbOv6S7TAg0VhriFL18ZxUPS6759SuKC63VOOSf4EEHy1m0qBgpCzzlsB7D4ssF9x0ZVXLREFrqikP71Hg6tSGcu4YBKL+VwIDWWaXzz6szxeDXdYTA3l35P7I9uBUgMznIjTjNaAX4AXRsJcN9fpF7mVq4eK1CorBY+OOzOc+/yVBpKysdaV/yZ+ABEhX93B2kPLFSOPUKjSPK2rtqE6h2NSl5BFuGEoVBerKn+ymOnmE4/SDBSe5S6gIL5vwy5zNMsxWUaUF5XO9Ez+2v8+yPSvQydj3pw5Rlb07mAXcI18ZYGClO6g/aKL52KYnn1FZ/X3r8r/cibfDbuXC6FRfVXJmzikVUqZdTp0tOwPkh4V0R63l2RO9Luy7vG6rurANSFnUA9n842KkRtBagQeQC96dbC0ebhTj+NPmskklxr6/6Op/P7d+YY76WzvQMvnsCAwEAAQ==\n-----END PUBLIC KEY-----"
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
