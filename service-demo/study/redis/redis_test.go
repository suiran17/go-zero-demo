package redis

import (
	"log"
	"testing"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func TestRedis(t *testing.T) {
	r, err := redis.NewRedis(redis.RedisConf{
		Host:        "dev.in",
		Type:        "",
		Pass:        "",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})
	if err != nil {
		log.Fatal("redis conn fail: ", err.Error())
	}

	key := "key1"
	redis.NewRedisLock(r, key)
}

/*
redis 锁
*/
func TestRedisLock(t *testing.T) {
	{
		conf := redis.RedisConf{
			Host: "dev.in:16379",
			Type: "node",
			Pass: "",
			Tls:  false,
		}

		rds := redis.MustNewRedis(conf)

		lock := redis.NewRedisLock(rds, "test")

		// 设置过期时间
		lock.SetExpire(10)

		// 尝试获取锁
		acquire, err := lock.Acquire()

		switch {
		case err != nil:
			// deal err
		case acquire:
			// 获取到锁
			defer lock.Release() // 释放锁
			// 业务逻辑
			//
			// 验证是自己的锁
			// ....

		case !acquire:
			// 没有拿到锁 wait?
		}
	}
}
