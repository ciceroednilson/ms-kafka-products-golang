# Project to learn Golang and Apache Kafka.

I created a Producer and a Consumer for Apache Kafka using the Golang.   

## 🚀 Starting.

The principal objective of this project is to show how to use Golang with Apache Kafka.

## 🛠 Stack.

<ol>
  <li>Golang</li>
  <li>VsCode</li>
  <li>Postman</li>
  <li>Mysql</li>
  <li>Mysql Workbench</li>
  <li>Docker</li>
  <li>Docker compose</li>
  <li>Apache Kafka</li>
</ol>

## ⚙️ Architecture General.

![docs_readme/archicture_new.png](docs_readme/diagrams_architecture_source.drawio.png)

## ⚙️ Architecture of Producer.

![docs_readme/producer.drawio.png](docs_readme/producer.drawio.png)

## ⚙️ Architecture of Consumer.

![docs_readme/consumer.drawio.png](docs_readme/consumer.drawio.png)


## ⚙️ Creating a Database on MySQL.

~~~~sql
CREATE DATABASE `db_products`;
~~~~

## ⚙️ Creating a Table on Database.

~~~~sql
CREATE TABLE IF NOT EXISTS tb_product(
   `id` 		   INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
   `description`   VARCHAR(100) NOT NULL,
   `price`   	   FLOAT NOT NULL,
   `total`   	   INT NOT NULL,
   `created`   	   timestamp NOT NULL
);
~~~~

## ⚙ ️Execute the Docker compose to create your infrastructure with Kafka.

Path of file: ms-kafka-products-golang/docker-compose/docker-compose.yml

If you don't have a Docker compose installed on your PC, you shall follow the steps below.

~~~~shell
curl -SL https://github.com/docker/compose/releases/download/v2.24.6/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose

sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

chmod +x /usr/local/bin/docker-compose
~~~~

Execute the command below to create of the infrastructure.

~~~~shell
sudo docker-compose up -d --build
~~~~

After you execute the command of Docker compose, you shall see the containers below.

![docs_readme/containers.png](docs_readme/containers.png)

## ⚙️ UI for Apache Kafka.

If everything is executing successfully, you can see this UI for Apache Kafka

[http://localhost:8080/](http://localhost:8080/)

![docs_readme/kafka_ui.png](docs_readme/kafka_ui.png)

## ⚙️ Create a new Topic.

this new Topic, we go to use to posts and consume some messages.

![docs_readme/topic.png](docs_readme/topic.png)

## ⚙️ Creating Swagger of files in the project of the Producer.

~~~~shell
swag init -g cmd/main.go handlers/producthandler/producthandler.go
~~~~

## ⚙️ Swagger of the Producer.

![docs_readme/swagger.png](docs_readme/swagger.png)

~~~~shell
swag init -g cmd/main.go handlers/producthandler/producthandler.go 
~~~~

## ⚙️ Download the dependencies.

Execute this command in the root directory of two applications(ms-products-producer-golang and ms-products-consumer-golang).
~~~~shell
go mod tidy
~~~~

## ⚙️ Execute the command below in two projects to start applications.

~~~~shell
 go run cmd/main.go
~~~~

## ⚙️ Test with Postman.

![docs_readme/postman.png](docs_readme/postman.png)

## ⚙️ Message on the Topic.

After you post a message, you shall see in UI of Kafka.

![docs_readme/message_kafka.png](docs_readme/message_kafka.png)

## ⚙️ See your consumer on Apache Kafka.

![docs_readme/consumer_kafka.png](docs_readme/consumer_kafka.png)

## ⚙️ See your information on the Database..

![docs_readme/mysql_data.png](docs_readme/mysql_data.png)

## 📌 Documentation swaggo

https://github.com/swaggo/http-swagger


## 📌 Version.

1.0

## ✒️ Author.

Cícero Ednilson - ciceroednilson@gmail.com
