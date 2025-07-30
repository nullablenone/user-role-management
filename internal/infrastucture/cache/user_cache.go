package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"manajemen-user/internal/domain/user"
	"time"

	"github.com/redis/go-redis/v9"
)

type userCacheRepository struct {
	next  user.Repository
	redis *redis.Client
	ctx   context.Context
}

func NewUserCacheRepository(next user.Repository, redis *redis.Client) user.Repository {
	return &userCacheRepository{
		next:  next,
		redis: redis,
		ctx:   context.Background(),
	}
}

func (r *userCacheRepository) GetAllUsers() ([]user.User, error) {
	cacheKey := "users:all"

	val, err := r.redis.Get(r.ctx, cacheKey).Result()
	if err == nil {
		log.Println(" Hitting Cache: GetAllUsers ")
		var users []user.User
		json.Unmarshal([]byte(val), &users)
		return users, nil
	}

	users, err := r.next.GetAllUsers()
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(users)
	r.redis.Set(r.ctx, cacheKey, data, time.Hour*1)

	return users, nil
}

func (r *userCacheRepository) GetUsersByID(id string) (*user.User, error) {
	cacheKey := fmt.Sprintf("user:%s", id)

	val, err := r.redis.Get(r.ctx, cacheKey).Result()
	if err == nil {
		log.Println(" Hitting Cache: GetUsersByID ")
		var user user.User
		json.Unmarshal([]byte(val), &user)
		return &user, nil
	}

	user, err := r.next.GetUsersByID(id)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(user)
	r.redis.Set(r.ctx, cacheKey, data, time.Hour*1)

	return user, nil
}

func (r *userCacheRepository) FindByEmailWithRole(email string) (*user.User, error) {
	return r.next.FindByEmailWithRole(email)
}

func (r *userCacheRepository) CreateUsers(user *user.User) error {
	err := r.next.CreateUsers(user)
	if err != nil {
		return err
	}

	log.Println(" Invalidating Cache: users:all ")
	r.redis.Del(r.ctx, "users:all")

	return nil
}

func (r *userCacheRepository) SaveUsers(user *user.User) error {
	err := r.next.SaveUsers(user)
	if err != nil {
		return err
	}

	userKey := fmt.Sprintf("user:%d", user.ID)
	log.Printf(" Invalidating Cache: %s dan users:all \n", userKey)
	r.redis.Del(r.ctx, userKey)
	r.redis.Del(r.ctx, "users:all")

	return nil
}

func (r *userCacheRepository) DeleteUsers(user *user.User) error {
	err := r.next.DeleteUsers(user)
	if err != nil {
		return err
	}

	userKey := fmt.Sprintf("user:%d", user.ID)
	log.Printf(" Invalidating Cache: %s dan users:all \n", userKey)
	r.redis.Del(r.ctx, userKey)
	r.redis.Del(r.ctx, "users:all")

	return nil
}
