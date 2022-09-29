
# 环境安装

## docker安装
```
yum update
yum install docker

# 启动
systemctl start docker
# 加入开机启动
systemctl enable docker

# 检查是否启动
docker version

# 测试一下
docker run hello-world
```

## docker-compose安装
```
# 第一步 下载二进制文件到/usr/local/bin/位置
curl -L https://github.com/docker/compose/releases/download/1.24.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

# 下载慢就用这个
curl -L https://get.daocloud.io/docker/compose/releases/download/v2.11.1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose


# 第二步 赋予可执行权限
chmod +x /usr/local/bin/docker-compose

# 查看版本号
docker-compose version
```

## kafka安装

### 单机版
```
version: "3"
services:
  zookeeper:
    image: 'bitnami/zookeeper:latest'
    ports:
      - '2181:2181'
    environment:
      # 匿名登录--必须开启
      - ALLOW_ANONYMOUS_LOGIN=yes
    #volumes:
      #- ./zookeeper:/bitnami/zookeeper
  # 该镜像具体配置参考 https://github.com/bitnami/bitnami-docker-kafka/blob/master/README.md
  kafka:
    image: 'bitnami/kafka:2.8.0'
    ports:
      - '9092:9092'
      - '9999:9999'
    environment:
      - KAFKA_BROKER_ID=1
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092
      # 客户端访问地址，更换成自己的主机IP （如果要外网访问就是服务器IP）
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://127.0.0.1:9092
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      # 允许使用PLAINTEXT协议(镜像中默认为关闭,需要手动开启)
      - ALLOW_PLAINTEXT_LISTENER=yes
      # 关闭自动创建 topic 功能
      - KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=false
      # 全局消息过期时间 6 小时(测试时可以设置短一点)
      - KAFKA_CFG_LOG_RETENTION_HOURS=6
      # 开启JMX监控
      - JMX_PORT=9999
    #volumes:
      #- ./kafka:/bitnami/kafka
    depends_on:
      - zookeeper
  # Web 管理界面 另外也可以用exporter+prometheus+grafana的方式来监控 https://github.com/danielqsj/kafka_exporter
  kafka_manager:
    image: 'hlebalbau/kafka-manager:latest'
    ports:
      - "9000:9000"
    environment:
      ZK_HOSTS: "zookeeper:2181"
      APPLICATION_SECRET: letmein
    depends_on:
      - zookeeper
      - kafka
```

### 集群版
```
version: '3.8'
services:
  zookeeper:
    image: 'bitnami/zookeeper:latest'
    ports:
      - '2181:2181'
    environment:
      # 匿名登录--必须开启
      - ALLOW_ANONYMOUS_LOGIN=yes
    #volumes:
      #- ./zookeeper:/bitnami/zookeeper
  kafka1:
    image: wurstmeister/kafka
    restart: always
    container_name: kafka1
    ports:
      - "9092:9092"
      - "9977:9977"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ADVERTISED_HOST_NAME: kafka1
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9092
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://127.0.0.1:9092
      KAFKA_ADVERTISED_PORT: 9092
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      JMX_HOSTNAME : 127.0.0.1
      JMX_PORT: 9977
      KAFKA_JMX_OPTS: "-Djava.rmi.server.hostname=127.0.0.1 -Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false  -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.rmi.port=9977"
    depends_on:
      - zookeeper

  kafka2:
    image: wurstmeister/kafka
    restart: always
    container_name: kafka2
    ports:
      - "9093:9093"
      - "9988:9988"
    environment:
      KAFKA_BROKER_ID: 2
      KAFKA_ADVERTISED_HOST_NAME: kafka2
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9093
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://127.0.0.1:9093
      KAFKA_ADVERTISED_PORT: 9093
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      JMX_HOSTNAME : 127.0.0.1
      JMX_PORT: 9988
      KAFKA_JMX_OPTS: "-Djava.rmi.server.hostname=127.0.0.1 -Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false  -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.rmi.port=9988"
    depends_on:
      - zookeeper

  kafka3:
    image: wurstmeister/kafka
    restart: always
    container_name: kafka3
    ports:
      - "9094:9094"
      - "9999:9999"
    environment:
      KAFKA_BROKER_ID: 3
      KAFKA_ADVERTISED_HOST_NAME: kafka3
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9094
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://127.0.0.1:9094
      KAFKA_ADVERTISED_PORT: 9094
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      JMX_HOSTNAME : 127.0.0.1
      JMX_PORT: 9999
      KAFKA_JMX_OPTS: "-Djava.rmi.server.hostname=127.0.0.1 -Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false  -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.rmi.port=9999"
    depends_on:
      - zookeeper
```

## 管理工具
管理工具：

[https://github.com/didi/KnowStreaming/blob/master/docs/install_guide/单机部署手册.md](https://github.com/didi/KnowStreaming/blob/master/docs/install_guide/%E5%8D%95%E6%9C%BA%E9%83%A8%E7%BD%B2%E6%89%8B%E5%86%8C.md)

初始账户 和 密码 admin