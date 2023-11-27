package cache

import (
	"context"
	"encoding/json"
	"job-portal-api/internal/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RDBLayer struct {
	rdb *redis.Client
}
type Caching interface {
	AddToTheCache(ctx context.Context, jid uint, jobData models.Job) error
	GetTheCacheData(ctx context.Context, jid uint) (string, error)
}

func NewRDBLayer(rdb *redis.Client) Caching {
	return &RDBLayer{
		rdb: rdb,
	}
}
func (c *RDBLayer) AddToTheCache(ctx context.Context, jid uint, jobData models.Job) error {
	jobID := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		return err
	}
	err = c.rdb.Set(ctx, jobID, val, 10*time.Minute).Err()
	return err
}

func (c *RDBLayer) GetTheCacheData(ctx context.Context, jid uint) (string, error) {
	jobID := strconv.FormatUint(uint64(jid), 10)
	str, err := c.rdb.Get(ctx, jobID).Result()
	return str, err
}
