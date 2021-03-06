package rpubsub

import (
	"github.com/go-redis/redis/v8"
	"github.com/jhump/protoreflect/desc"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber/cli"
)

type Redis struct {
	Options *cli.Options
	Client  *redis.Client
	MsgDesc *desc.MessageDescriptor
	log     *logrus.Entry
}

func NewClient(opts *cli.Options) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     opts.RedisPubSub.Address,
		Username: opts.RedisPubSub.Username,
		Password: opts.RedisPubSub.Password,
		DB:       opts.RedisPubSub.Database,
	}), nil
}
