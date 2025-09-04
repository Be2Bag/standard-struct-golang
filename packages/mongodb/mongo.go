package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Client) Connect() error {
	var connectionString string
	if c.cfg.Connection != nil && *c.cfg.Connection != "" {
		//ในกรณีมี connection string ให้ใช้ connection string
		connectionString = *c.cfg.Connection
	} else {
		//ในกรณีไม่มี connection string ให้ใช้ username host port DbName มาต่อเป็น connection string
		if c.cfg.Username == "" || c.cfg.Password == "" {
			//ในกรณีไม่มี user password
			connectionString = fmt.Sprintf("mongodb://%s:%s/%s?authSource=admin", c.cfg.Host, c.cfg.Port, c.cfg.DbName)
		} else {
			connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
				c.cfg.Username, c.cfg.Password, c.cfg.Host, c.cfg.Port, c.cfg.DbName)
		}
	}
	// connect mongo db
	client, err := mongo.Connect(
		context.TODO(), //ใช่เป็น TODO เนื่องจากไม่แน่ใจ context
		options.Client().ApplyURI(connectionString),       //connectionString
		options.Client().SetConnectTimeout(time.Second*5), //จำกัดเวลาการ connect
		options.Client().SetTimeout(time.Second*60),       // Every query should not over 1min
	)
	//ดักเมื่อการเชื่อมต่อเกิด error
	if err != nil {
		return err
	}
	//กำหนด mongo client ให้ Instance
	c.db = client
	// คืนค่าด้วยการ Ping หา server
	return c.Ping()
}

func (c *Client) Ping() error {
	// คืนค่าด้วยการ Ping หา server
	return c.db.Ping(context.TODO(), nil)
}

func (c *Client) GetClient() *mongo.Client {
	//คืน Client ของ mong ในกรณีที่จำเป็นต้องมีการใช้ mongo.Client นอก package
	return c.db
}

func (c *Client) Close() error {
	//disconnect database
	if c.db != nil {
		return c.db.Disconnect(context.TODO())
	}
	return nil
}
