package config

import (
	"bytes"
	"context"
	"example.com/http_demo/utils/zlog"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"io"
	"time"
)

type RemoteConfig struct {
	viper.RemoteProvider

	Username string
	Password string
}

func (c *RemoteConfig) Get(rp viper.RemoteProvider) (io.Reader, error) {
	c.RemoteProvider = rp

	return c.get()
}

func (c *RemoteConfig) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	c.RemoteProvider = rp

	return c.get()
}

func (c *RemoteConfig) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	c.RemoteProvider = rp
	rr := make(chan *viper.RemoteResponse)
	stop := make(chan bool)

	go func() {
		client, err := c.newClient()
		if err != nil {
			zlog.Error("Etcd Client Error", zap.Error(err))
			return
		}
		defer client.Close()

		for {
			ch := client.Watch(context.Background(), c.RemoteProvider.Path())
			select {
			case <-stop:
				return
			case res := <-ch:
				for _, event := range res.Events {
					rr <- &viper.RemoteResponse{
						Value: event.Kv.Value,
					}
				}
			}
		}
	}()

	return rr, stop
}

func (c *RemoteConfig) newClient() (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{c.Endpoint()},
		Username:             c.Username,
		Password:             c.Password,
		AutoSyncInterval:     time.Hour,
		DialTimeout:          time.Second * 5,
		DialKeepAliveTime:    time.Second * 15,
		DialKeepAliveTimeout: time.Second * 5,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *RemoteConfig) get() (io.Reader, error) {
	client, err := c.newClient()

	if err != nil {
		return nil, err
	}

	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, c.Path())
	cancel()

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(resp.Kvs[0].Value), nil
}
