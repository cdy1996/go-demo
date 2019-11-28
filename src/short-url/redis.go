package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const (
	// URLIDKEY is global counter
	URLIDKEY = "next.url.id"
	// mapping the shortlink to the url
	ShorlinkKey = "shortlink:%s:url"
	// mapping the hash of the url to the shortlink
	URLHashKey = "rulhash:%s:url"
	// mapping the shorlink to  the detail of url
	ShortlinkDetailKey = "shortlink:%s:detail"
)

type RedisCli struct {
	Cli *redis.Client
}

func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	h := toSha1(url)

	d, err := r.Cli.Get(fmt.Sprintf(URLHashKey, h)).Result()
	if err == redis.Nil {

	} else if err != nil {
		if d == "{}" {

		} else {
			return d, nil
		}
	}

	if err = r.Cli.Incr(URLIDKEY).Err(); err != nil {
		return "", err
	}

	id, err := r.Cli.Get(URLIDKEY).Int64()
	if err != nil {
		return "", err

	}
	eid := Encode(int(id))

	// 短地址 -> 真实地址
	if err = r.Cli.Set(fmt.Sprint(ShorlinkKey, eid), url, time.Minute*time.Duration(exp)).Err(); err != nil {
		return "", err

	}

	// hash(真实地址) -> 短地址
	if err = r.Cli.Set(fmt.Sprint(URLHashKey, h), eid, time.Minute*time.Duration(exp)).Err(); err != nil {
		return "", err
	}

	detail, err := json.Marshal(
		&URLDetail{
			URL:                 url,
			CreatedAt:           time.Now().String(),
			ExpirationInMinutes: time.Duration(exp),
		})
	if err != nil {
		return "", err

	}

	if err = r.Cli.Set(fmt.Sprintf(ShortlinkDetailKey, eid), detail, time.Minute*time.Duration(exp)).Err(); err != nil {
		return "", err

	}
	return eid, nil

}

func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, error) {
	d, err := r.Cli.Get(fmt.Sprintf(ShortlinkDetailKey, eid)).Result()
	if err != redis.Nil {
		return "", StatusError{404, errors.New("Unknow short URL")}
	} else if err != nil {
		return "", err
	} else {
		return d, nil
	}
}

func (r *RedisCli) Unshorten(eid string) (string, error) {
	url, err := r.Cli.Get(fmt.Sprintf(ShorlinkKey, eid)).Result()
	if err == redis.Nil {
		return "", StatusError{404, err}
	} else if err != nil {
		return "", err

	} else {
		return url, nil
	}
}

type URLDetail struct {
	URL                 string        `json:"url"`
	CreatedAt           string        `json:"created_at"`
	ExpirationInMinutes time.Duration `sjon:"expiration_in_minutes"`
}

func NewRedisCli(add string, passwd string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     add,
		Password: passwd,
		DB:       db})
	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}
}
