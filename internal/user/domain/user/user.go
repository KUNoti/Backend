package user

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	LOGINLINK       = "https://myapi.ku.th/auth/login"
	GETSCHEDULELINK = "https://myapi.ku.th/std-profile/getGroupCourse"
)

type Test struct {
	ID int
}

type UserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func LoginUser(c *gin.Context) {
	var reqBody map[string]string
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	encodedBody := map[string]string{
		"username": encodeString(reqBody["username"]),
		"password": encodeString(reqBody["password"]),
	}

	// headers := map[string]string{
	// 	"app-key": "apikey",
	// }

	encodedJSON, err := json.Marshal(encodedBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	resp, err := http.Post(LOGINLINK, "application/json", bytes.NewBuffer(encodedJSON))

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer resp.Body.Close()

	// Handle response logic here

	c.JSON(resp.StatusCode, gin.H{"data": "response data"}) // replace with actual response data
}

func encodeString(data string) string {
	block, _ := pem.Decode([]byte("kuPublicKey"))
	if block == nil {
		panic("Failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("Failed to parse public key")
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		panic("Failed to convert public key to RSA public key")
	}

	encryptedData, err := rsa.EncryptPKCS1v15(nil, rsaPubKey, []byte(data))
	if err != nil {
		panic("Failed to encrypt data")
	}

	return base64.StdEncoding.EncodeToString(encryptedData)
}

func getSchedule(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"test": "test"})
}
