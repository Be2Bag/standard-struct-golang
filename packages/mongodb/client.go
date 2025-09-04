package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// กำหนดรูปแบบ config ที่จะใช้ต่อกับ database
type Config struct {
	Host       string
	Port       string
	Username   string
	Password   string
	DbName     string
	Connection *string
}

// กำหนดรูปแบบ Instance ของตัว mongo
type Client struct {
	cfg Config
	db  *mongo.Client
}

// Set Instance ของตัว mongo ด้วย Config
func NewWithConfig(cfg Config) Client {
	return Client{
		cfg: cfg,
	}
}

// Set Instance ของตัว mongo ด้วย connection string
func NewWithConnectionString(conn string) Client {
	return Client{
		cfg: Config{
			Host:       "",
			Port:       "",
			Username:   "",
			Password:   "",
			DbName:     "",
			Connection: &conn,
		},
	}
}
