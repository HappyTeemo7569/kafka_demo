package alikafka

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 获取路径
func GetFullPath(file string) string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(workPath, "conf")
	fullPath := filepath.Join(configPath, file)
	return fullPath
}

// 生产者 消费者其实是一样的 主要是连接加上鉴权
func newSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	config.Net.SASL.Enable = true
	config.Net.SASL.User = "ali 用户名"
	config.Net.SASL.Password = "ali 密码"
	config.Net.SASL.Handshake = true

	//读配置文件
	certBytes, err := ioutil.ReadFile(GetFullPath("下载的鉴权文件"))
	if err != nil {
		fmt.Println("kafka client read cert file failed ", err.Error())
		return nil, err
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		return nil, errors.New("kafka client failed to parse root certificate")
	}
	config.Net.TLS.Config = &tls.Config{
		RootCAs:            clientCertPool,
		InsecureSkipVerify: true,
	}
	config.Net.TLS.Enable = true

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{"连接地址"}, config)
	//defer producer.Close()

	if err != nil {
		return nil, err
	}

	return producer, nil
}
