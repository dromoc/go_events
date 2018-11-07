go get github.com/gorilla/handlers
go get github.com/Shopify/sarama
go get -u github.com/mitchellh/mapstructure

sudo service mongod start

docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
docker run -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=127.0.0.1 --env ADVERTISED_PORT=9092 spotify/kafka

docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it kafkadocker_kafka_1 bash
docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -u root -it bear-clickhouse bash
