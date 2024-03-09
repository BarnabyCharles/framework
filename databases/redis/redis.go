package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ghodss/yaml"
	"github.com/go-redis/redis/v8"

	"github.com/BarnabyCharles/framework/config"
)

var Rdb *redis.Client

func WithCliRedis(ctx context.Context, serverName string, cli func(*redis.Client) (string, error)) (string, error) {
	content, err := config.GetNacosConfig(serverName, "DEFAULT_GROUP")
	if err != nil {
		return "", err
	}
	var rediscfg struct {
		Redis config.RedisConfig `json:"Redis" yaml:"Redis" mapstruture:"Redis"`
	}

	err = yaml.Unmarshal([]byte(content), &rediscfg)
	if err != nil {
		return "", errors.New("转换为结构体格式失败redis" + err.Error())
	}
	cfg := rediscfg.Redis

	Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:   0,
	})

	defer Rdb.Close()
	cli(Rdb)

	return "", nil
}
func GetRedisKey2(ctx context.Context, serverName, key string) (string, error) {
	content, err := config.GetNacosConfig(serverName, "DEFAULT_GROUP")
	if err != nil {
		return "", err
	}
	var rediscfg struct {
		Redis config.RedisConfig `json:"Redis" yaml:"Redis" mapstruture:"Redis"`
	}

	err = yaml.Unmarshal([]byte(content), &rediscfg)
	if err != nil {
		return "", errors.New("转换为结构体格式失败redis" + err.Error())
	}
	cfg := rediscfg.Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:   0,
	})
	result, err := Rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
func GetRedisKey(ctx context.Context, serverName, key string) (string, error) {
	var data string
	var err error
	log.Println("key=======================", key)
	data, err = WithCliRedis(ctx, serverName, func(client *redis.Client) (string, error) {
		code := client.Get(ctx, key).Val()
		if code == "" {
			return "", errors.New("获取key值失败!")
		}
		log.Println("====================data", code)
		return code, err
	})
	if err != nil {
		return "", errors.New("获取key失败！" + err.Error())
	}
	log.Println("data==============================================", data)
	return data, nil
}

func SetKey(ctx context.Context, serverName, key string, val int, duration time.Duration) error {
	var err error
	_, err = WithCliRedis(ctx, serverName, func(client *redis.Client) (string, error) {
		err = client.Set(ctx, key, val, duration).Err()
		return "", err
	})
	if err != nil {
		return errors.New("设置key失败！" + err.Error())
	}
	return nil
}

func ExpireKey(ctx context.Context, serverName, key string) bool {
	_, err := WithCliRedis(ctx, serverName, func(client *redis.Client) (string, error) {
		exists := client.Exists(ctx, key).String()
		if exists == "" {
			return "", errors.New("当前key不存在")
		}
		return "", nil
	})
	if err != nil {
		return false
	}
	return true
}

func IndexAdd(ctx context.Context, serverName, key string, duration time.Duration) error {
	_, err := WithCliRedis(ctx, serverName, func(client *redis.Client) (string, error) {
		_, err := client.Incr(ctx, key).Result()
		if err != nil {
			return "", errors.New("设置key失败！" + err.Error())
		}
		_, err = client.Expire(ctx, key, duration).Result()
		if err != nil {
			return "", errors.New("设置过期时间失败！" + err.Error())
		}
		return "", nil
	})
	if err != nil {
		return err
	}
	return nil
}
