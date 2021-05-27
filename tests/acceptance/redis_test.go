package acceptance

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/newrelic/infra-integrations-sdk/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNRIRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	if err := setUpRedis(); err != nil {
		panic(err)
	}
	RunSpecs(t, "NRI Redis Suite")

}

func setUpRedis() error {
	maxTries := 20
	envVars := make([]string, 0)
	ports := []string{"6379:6379"}
	if stdout, stderr, err := dockerComposeRunMode(envVars, ports, "redis", true); err != nil {
		log.Info(stdout)
		log.Warn(stderr)
		log.Fatal(err)
	}
	client := redisClient("localhost", "6379")
	for ; maxTries > 0; maxTries-- {
		log.Info("try to establish de connection with the redis...")
		pong, err := client.Ping().Result()
		if err != nil {
			log.Warn(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Println(pong, err)
		// Output: PONG <nil>
		return nil
	}
	return errors.New("redis connection cannot be established")
}

func redisClient(hostname, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", hostname, port),
	})
	return client
}
