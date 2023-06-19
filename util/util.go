package util

import (
	"github.com/Jansen-Stanlie/MPSBE-Engine/pkg/common/db"
	"github.com/Jansen-Stanlie/MPSBE-Engine/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func LoadConfig() *models.Config {
	viper.SetConfigFile("./envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

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
