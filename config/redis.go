package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	viper.SetConfigName("config") // ไม่ต้องใส่ .yaml
	viper.SetConfigType("yaml")   // ระบุประเภทไฟล์
	viper.AddConfigPath(".")      // path ปัจจุบัน (หรือใช้ os.Getenv สำหรับ path .env)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")

	log.Println("host", host)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port), // หรือจาก ENV
		Password: password,                         // ใส่ถ้ามี
		DB:       0,                                // db index
	})

	_, err = RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}
