package redis

import (
	"flash_sale/pkg/setting"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get search a key
func Get(key string) (map[string]interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	flashSale, err := GetFlashSaleData(conn, key)
	if err != nil {
		return nil, err
	}

	return flashSale, nil
}

// HGETALL search keys
func HGetAll(pattern string) ([]interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	var cursor int64 = 0
	var keys []string
	scanValues, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", pattern, "COUNT", 10))
	if err != nil {
		return nil, err
	}
	scanValues, err = redis.Scan(scanValues, &cursor, &keys)
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, len(keys))
	for i, key := range keys {
		flashSale := make(map[string]interface{})
		flashSale, err = GetFlashSaleData(conn, key)
		if err != nil {
			return nil, err
		}
		result[i] = flashSale
	}
	return result, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// HMSet set key
func HMSet(key string, data map[string]interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	args := make([]interface{}, 0, len(data)*2+1)
	args = append(args, key)

	for field, value := range data {
		args = append(args, field, value)
	}

	_, err := conn.Do("HMSET", args...)
	if err != nil {
		return err
	}
	return nil
}

// GetFlashSaleData set key
func GetFlashSaleData(conn redis.Conn, key string) (map[string]interface{}, error) {
	values, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	flashSale := make(map[string]interface{})
	for j := 0; j < len(values); j += 2 {
		field := string(values[j].([]byte))
		valueStr := string(values[j+1].([]byte))

		if field == "id" || field == "product_id" || field == "stock" || field == "discount_percent" {
			valueInt, err := strconv.Atoi(valueStr)
			if err != nil {
				flashSale[field] = valueStr
			} else {
				flashSale[field] = valueInt
			}
		} else {
			flashSale[field] = valueStr
		}
	}
	return flashSale, nil
}
