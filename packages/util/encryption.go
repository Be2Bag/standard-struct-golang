package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HashData(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))    // ใส่ข้อมูล
	hashedData := hash.Sum(nil) // รับผลลัพธ์เป็น byte slice

	hashedString := hex.EncodeToString(hashedData)
	return hashedString
}

func EncryptAES256GCM(data string, key string, nonce string) (string, error) {
	hashedKey := sha256.Sum256([]byte(key))

	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, []byte(nonce), []byte(data), nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES256GCM(encryptedText string, key string, nonce string) string {
	hashedKey := sha256.Sum256([]byte(key))

	// ตรวจสอบขนาดของ Key
	if len(key) != 32 {
		log.Println("key length must be 32 bytes")
		return encryptedText
	}
	// ตรวจสอบขนาดของ Nonce
	if len(nonce) != 12 {
		log.Println("nonce length must be 12 bytes")
		return encryptedText
	}

	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		log.Printf("fail to create cipher text %v", err)
		return encryptedText
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("fail: %v", err)
		return encryptedText
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		log.Printf("fail to decode: %v", err)
		return encryptedText
	}

	plaintext, err := aesGCM.Open(nil, []byte(nonce), ciphertext, nil)
	if err != nil {
		log.Printf("fail: %v", err)
		return encryptedText
	}

	return string(plaintext)
}
