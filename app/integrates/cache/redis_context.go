package cache

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/global"
	"strconv"
	"time"
)

type RedisContext interface {
	// SetWithKey cache data with key
	SetWithKey(key string, value interface{}) core.Either[core.Unit, core.ErrContext]
	// SetWithKeyAndExp cache data with key and set expiration
	SetWithKeyAndExp(key string, value interface{}, expiration time.Duration) core.Either[core.Unit, core.ErrContext]
	// GetWithKey get cache data with key
	GetWithKey(key string) core.Either[string, core.ErrContext]
	// Publish publish redis message pub-sub with channel
	Publish(channel string, message interface{}) core.Either[core.Unit, core.ErrContext]
	// Subscribe subscribe redis message pub-sub with channel
	Subscribe(channel string, fn func(*redis.Message))
}

type redisContext struct {
	context context.Context
	client  *redis.Client
	config  global.RedisConfig
}

func NewRedisContext(config global.RedisConfig) RedisContext {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.DB,
		Protocol: config.Protocol,
	})
	ctx := context.Background()
	return &redisContext{
		context: ctx,
		client:  client,
		config:  config,
	}
}

func (r *redisContext) SetWithKey(key string, value interface{}) core.Either[core.Unit, core.ErrContext] {
	err := r.client.Set(r.context, key, value, 0).Err()
	if err != nil {
		log.Error(err)
		return core.LeftEither[core.Unit, core.ErrContext](core.NewErrorCode(core.Invalid))
	}
	return core.NewRightEither[core.Unit, core.ErrContext](core.NewUnitPtr())
}

func (r *redisContext) SetWithKeyAndExp(key string, value interface{}, expiration time.Duration) core.Either[core.Unit, core.ErrContext] {
	err := r.client.Set(r.context, key, value, expiration).Err()
	if err != nil {
		log.Error(err)
		return core.LeftEither[core.Unit, core.ErrContext](core.NewErrorCode(core.Invalid))
	}
	return core.NewRightEither[core.Unit, core.ErrContext](core.NewUnitPtr())
}

func (r *redisContext) GetWithKey(key string) core.Either[string, core.ErrContext] {
	exists, err := r.client.Exists(r.context, key).Result()
	if err != nil {
		log.Error(err)
		return core.LeftEither[string, core.ErrContext](core.NewErrorCode(core.Invalid))
	}
	if exists > 0 {
		value, err := r.client.Get(r.context, key).Result()
		if err != nil {
			log.Error(err)
			return core.LeftEither[string, core.ErrContext](core.NewErrorCode(core.Invalid))
		}
		return core.RightEither[string, core.ErrContext](value)
	}
	return core.RightEither[string, core.ErrContext]("")
}

func (r *redisContext) Publish(channel string, message interface{}) core.Either[core.Unit, core.ErrContext] {
	if err := r.client.Publish(r.context, channel, message).Err(); err != nil {
		log.Error("Cannot Publish Redis Message", err)
		return core.LeftEither[core.Unit, core.ErrContext](core.NewErrorCode(core.Invalid))
	}
	return core.RightEither[core.Unit, core.ErrContext](core.NewUnit())
}

func (r *redisContext) Subscribe(channel string, fn func(*redis.Message)) {
	subscriber := r.client.Subscribe(r.context, channel)
	for {
		msg, err := subscriber.ReceiveMessage(r.context)
		if err != nil {
			log.Error("Cannot ReceiveMessage Redis Message", err)
			break
		} else {
			fn(msg)
		}
	}
}
