package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache interface {
	Get(ctx context.Context, id int64) (domain.User, error)
	Set(ctx context.Context, u domain.User) error
	Del(ctx context.Context, id int64) error
}

type UserRedisCache struct {
	client redis.Cmdable
	// 过期时间
	expiration time.Duration
}

func NewUserRedisCache(client redis.Cmdable) UserCache {
	return &UserRedisCache{client: client, expiration: time.Hour * 24 * 7}
}
func (cache *UserRedisCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := cache.key(id)
	var u domain.User
	data, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, err
	}
	// 要反序列化
	err = json.Unmarshal([]byte(data), &u)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (cache *UserRedisCache) Set(ctx context.Context, u domain.User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return cache.client.Set(ctx, cache.key(u.Id), data, cache.expiration).Err()

}
func (cache *UserRedisCache) Del(ctx context.Context, id int64) error {
	return cache.client.Del(ctx, cache.key(id)).Err()
}

// 整个系统的  key 都是tech_blog:$module:XXX
func (cache *UserRedisCache) key(id int64) string {
	return fmt.Sprintf("tech_blog:user:%d", id)
}
