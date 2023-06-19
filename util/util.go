package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/Jansen-Stanlie/MPSBE-Engine/pkg/common/db"
	"github.com/Jansen-Stanlie/MPSBE-Engine/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

func LoadConfig() *models.Config {
	//https://github.com/Jansen-Stanlie/MPSBE-Engine/blob/main/envs/.env
	viper.SetConfigFile(os.Getenv("GOPATH") + "/pkg/mod/github.com/!jansen-!stanlie/!m!p!s!b!e-!engine@v0.0.0-20230619082206-7b41d38bfe13/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	dbUrl := viper.Get("DB_URL").(string)
	fmt.Println("port : " + port)
	fmt.Println("dbUrl : " + dbUrl)

	h := db.Init(dbUrl)

	r := gin.Default()

	s := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	config := &models.Config{
		Database: h,
		Server:   s,
	}

	return config
}

func decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
