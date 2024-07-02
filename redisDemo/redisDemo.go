package redisdemo

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mike504110403/goutils/log"
)

// RedisConnect : 連線 redis
func RedisConnect() (redis.Conn, error) {
	// 連線 redis
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return nil, err
	}
	// defer c.Close()
	return c, nil
}

// RedisGet : 取得 redis key 的 value
func RedisSet(key, value string) error {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return err
	}
	defer c.Close()

	_, err = c.Do("Set", key, value)
	if err != nil {
		log.Error(fmt.Sprintf("set key failed: %s", err.Error()))
		return err
	}
	return nil
}

// RedisGet : 取得 redis key 的 value
func RedisGet(key string) (string, error) {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return "", err
	}
	defer c.Close()

	value, err := c.Do("Get", key)
	if err != nil {
		log.Error(fmt.Sprintf("get key failed: %s", err.Error()))
		return "", err
	}
	if value == nil {
		log.Error(fmt.Sprintf("key not found: %s", key))
		return "", fmt.Errorf("key not found: %s", key)
	}
	return string(value.([]byte)), nil
}

// RedisMGet : 取得多個 redis key 的 value
func RedisMSet(keyValues map[string]string) error {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return err
	}
	defer c.Close()

	args := []interface{}{}
	for k, v := range keyValues {
		args = append(args, k, v)
	}
	_, err = c.Do("MSet", args...)
	if err != nil {
		log.Error(fmt.Sprintf("mset key failed: %s", err.Error()))
		return err
	}
	return nil
}

// RedisMGet : 取得多個 redis key 的 value
func RedisMGet(keys []string) (map[string]string, error) {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return nil, err
	}
	defer c.Close()

	args := []interface{}{}
	for _, k := range keys {
		args = append(args, k)
	}
	values, err := c.Do("MGet", args...)
	if err != nil {
		log.Error(fmt.Sprintf("mget key failed: %s", err.Error()))
		return nil, err
	}
	if values == nil {
		log.Error(fmt.Sprintf("key not found: %v", keys))
		return nil, fmt.Errorf("key not found: %v", keys)
	}
	valueMap := map[string]string{}
	for i, v := range values.([]interface{}) {
		valueMap[keys[i]] = string(v.([]byte))
	}
	return valueMap, nil
}

// RedisHashSet : 設定 redis hash key 的 field 的 value
func RedisHashSet(key, field, value string) error {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return err
	}
	defer c.Close()

	_, err = c.Do("HSet", key, field, value)
	if err != nil {
		log.Error(fmt.Sprintf("hset key failed: %s", err.Error()))
		return err
	}
	return nil
}

// RedisHashGet : 取得 redis hash key 的 field 的 value
func RedisHashGet(key, field string) (string, error) {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return "", err
	}
	defer c.Close()

	value, err := c.Do("HGet", key, field)
	if err != nil {
		log.Error(fmt.Sprintf("hget key failed: %s", err.Error()))
		return "", err
	}
	if value == nil {
		log.Error(fmt.Sprintf("field not found: %s", field))
		return "", fmt.Errorf("field not found: %s", field)
	}
	return string(value.([]byte)), nil
}

// RedisExpire : 設定 redis key 的過期時間
func RedisExpire(key string, seconds int) error {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return err
	}
	defer c.Close()

	_, err = c.Do("Expire", key, seconds)
	if err != nil {
		log.Error(fmt.Sprintf("expire key failed: %s", err.Error()))
		return err
	}
	return nil
}

// RedisLPush : 將 value 推入 redis list
func RedisLPush(key, value string) error {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return err
	}
	defer c.Close()

	_, err = c.Do("lpush", key, value)
	if err != nil {
		log.Error(fmt.Sprintf("lpush key failed: %s", err.Error()))
		return err
	}
	return nil
}

// RedisQueue : 取得 redis list 的 value
func RedisQueue(key string, resc chan string) {
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return
	}
	defer c.Close()

	for {
		r, err := redis.String(c.Do("lpop", key))
		if err != nil {
			if err == redis.ErrNil {
				log.Info(fmt.Sprintf("lpop key nil: %s", err.Error()))
				time.Sleep(1 * time.Second)
				continue
			}
			log.Error(fmt.Sprintf("lpop key failed: %s", err.Error()))
			return
		}
		resc <- r
		log.Info(fmt.Sprintf("lpop key success: %s", r))
	}
}

// QueueToChannelDemo : 將 redis list 的 value 推入 channel
func QueueToChannelDemo() {
	resc := make(chan string)
	go RedisQueue("testQueue", resc)
	// 單消費者模式
	// for {
	// 	select {
	// 	case r := <-resc:
	// 		log.Info(fmt.Sprintf("get queue value: %s", r))
	// 	}
	// }
	// 多消費者模式 -> 可啟用多個 goroutine跑這段
	for r := range resc {
		log.Info(fmt.Sprintf("get queue value: %s", r))
	}
}

func RedisPoolDemo() {
	// 建立 redis pool
	pool := &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	// 取得連線
	c := pool.Get()
	defer c.Close()

	// 設定 key
	_, err := c.Do("Set", "poolKey", "poolValue")
	if err != nil {
		log.Error(fmt.Sprintf("set key failed: %s", err.Error()))
		return
	}

	// 取得 key
	value, err := redis.String(c.Do("Get", "poolKey"))
	if err != nil {
		log.Error(fmt.Sprintf("get key failed: %s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("get key success: %s", value))
}

// RedisChannelDemo : redis channel demo
func RedisChannelDemo() {
	// 連線 redis
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return
	}
	defer c.Close()

	c.Send("SET", "username1", "value1")
	c.Send("SET", "username2", "value2")
	// 一次性發送命令
	c.Flush()

	// 接收命令結果 先進先出
	v, err := c.Receive()
	log.Info(fmt.Sprintf("receive: %v, err: %s\n", v, err.Error()))

	v, err = c.Receive()
	log.Info(fmt.Sprintf("receive: %v, err: %s\n", v, err.Error()))

	v, err = c.Receive() // 這裡會阻塞
	log.Info(fmt.Sprintf("receive: %v, err: %s\n", v, err.Error()))

	// 結果
	// 2021/07/07 16:00:00 receive: OK, err: <nil>
	// 2021/07/07 16:00:00 receive: OK, err: <nil>
}

// RedisTransectionDemo : redis 事務 demo
func RedisTransectionDemo() {
	// 連線 redis
	c, err := RedisConnect()
	if err != nil {
		log.Error(fmt.Sprintf("connect to redis failed: %s", err.Error()))
		return
	}
	defer c.Close()

	// 開始事務
	c.Send("MULTI")
	c.Send("SET", "username1", "value1")
	c.Send("SET", "username2", "value2")
	// 執行事務
	v, err := c.Do("EXEC")
	log.Info(fmt.Sprintf("exec: %v, err: %s\n", v, err.Error()))
}
