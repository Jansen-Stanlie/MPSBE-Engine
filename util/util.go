package util

import (
	"fmt"
	"github.com/Jansen-Stanlie/MPSBE-Engine/pkg/common/db"
	"github.com/Jansen-Stanlie/MPSBE-Engine/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func LoadConfig(r *gin.Engine) *models.Config {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("current dir" + dir)

	viper.SetConfigFile(os.Getenv("GOPATH") + "/pkg/mod/github.com/!jansen-!stanlie/!m!p!s!b!e-!engine@v0.0.0-20230619082206-7b41d38bfe13/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	dbUrl := viper.Get("DB_URL").(string)

	h := db.Init(dbUrl)

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
