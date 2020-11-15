package service

import (
	"fmt"

	redis "github.com/go-redis/redis/v7"

	"gitthub.com/marcellop71/mosaic/abe/log"
)

type StorageRedis struct{}

var rc *redis.Client

// init
func (abeStorage StorageRedis) Init(config Config) {
	conf := config.Storage.Redis[config.Active.Label]
	rc = redis.NewClient(&redis.Options{
		Addr: conf.Addr, Password: conf.Password, DB: 0})
	if rc == nil {
		log.Panic("error init new Redis client")
	}
	_, err := rc.Ping().Result()
	if err != nil {
		log.Panic("error init new Redis client")
	}
	rc.FlushAll()
}

// set
func (abeStorage StorageRedis) Set(key string, value string) {
	err := rc.Set(key, value, 0).Err()
	if err != nil {
		log.Error("%s", err)
	}
}

// get
func (abeStorage StorageRedis) Get(key string) string {
	value, err := rc.Get(key).Result()
	if err != nil {
		log.Error("%s", err)
	}
	return value
}

// scan
func (abeStorage StorageRedis) Scan(key string) []string {
	var err error
	var keys []string
	var keys_batch []string
	var cursor uint64
	batch_size := int64(100)
	cursor = 0
	key_ := fmt.Sprintf("%s*", key)
	for {
		keys_batch, cursor, err = rc.Scan(cursor, key_, batch_size).Result()
		if (err != nil) {
			log.Error("%s", err)
		}
		if (len(keys_batch) > 0) {
			keys = append(keys, keys_batch...)
		}
		if (cursor == 0) {
			break
		}
	}
	return keys
}
