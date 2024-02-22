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

func GenerateOTPcode() int {
	// Set a seed for randomness based on current time
	rand.Seed(time.Now().UnixNano())

	// Generate 4 digits random number
	return rand.Intn(900000) + 100000
}

// func Encrypt(input string) string  {
// 	plainText := []byte(input)
// 	// hasher.Write([]byte(numValue))
// 	hashByte := sha256.Sum256(plainText)

// 	return hex.EncodeToString(hashByte[:])
// }

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