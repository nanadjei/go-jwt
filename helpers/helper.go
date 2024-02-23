package helpers

import (
	// "crypto/sha256"
	// "encoding/hex"
	"math/rand"
	"time"
	"os"
	"crypto/aes"
	"crypto/cipher"
 	"encoding/base64"
)

// Generate a random 6 digit code
func GenerateOTPcode() int {
	// Set a seed for randomness based on current time
	rand.Seed(time.Now().UnixNano())
	// Generate 5 digits random number
	return rand.Intn(900000) + 100000
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
} 

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("APP_ENCRYPTION_KEY")))
	if err != nil {
	return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("APP_ENCRYPTION_KEY")))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}