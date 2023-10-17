package redis

import (
	"crypto/tls"

	"github.com/go-redis/redis/v8"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v8"

	"coba/pkg/database"
	"coba/pkg/dotenv"
)

// InitClient initialize redis client
func InitClient() database.RedisTemplate {
	var (
		opts    *redis.Options
		redisDB = dotenv.REDISDB()
	)

	if dotenv.ISREDISTLS() {
		opts = &redis.Options{
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Addr:     dotenv.REDISURL(),
			Password: dotenv.REDISPASS(),
			DB:       redisDB,
		}
	} else {
		opts = &redis.Options{
			Addr:     dotenv.REDISURL(),
			Password: dotenv.REDISPASS(),
			DB:       redisDB,
		}
	}

	client := redis.NewClient(opts)
	if dotenv.IsAppEnvProduction() {
		client.AddHook(nrredis.NewHook(opts))
	}

	return client
}

// InitClientFr initialize redis client forgerock
func InitClientFr() database.RedisTemplate {
	var (
		opts    *redis.Options
		redisDB = dotenv.GetInt("REDIS_DB_FR", 0)
	)

	if dotenv.ISREDISTLS() {
		opts = &redis.Options{
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Addr:     dotenv.GetString("REDIS_URL_FR", ""),
			Password: dotenv.GetString("REDIS_PASS_FR", ""),
			DB:       redisDB,
		}
	} else {
		opts = &redis.Options{
			Addr:     dotenv.GetString("REDIS_URL_FR", ""),
			Password: dotenv.GetString("REDIS_PASS_FR", ""),
			DB:       redisDB,
		}
	}

	client := redis.NewClient(opts)
	if dotenv.IsAppEnvProduction() {
		client.AddHook(nrredis.NewHook(opts))
	}

	return client
}

// InitCluster iniialize cluster client
func InitCluster() *redis.ClusterClient {
	opts := &redis.ClusterOptions{
		Addrs:    []string{dotenv.REDISURL()},
		Password: dotenv.REDISPASS(),
	}
	client := redis.NewClusterClient(opts)
	if dotenv.IsAppEnvProduction() {
		client.AddHook(nrredis.NewHook(nil))
	}
	return client
}
