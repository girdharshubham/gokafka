# Golang Kafka Consumer
This template utilizes Confluent's GoLang client for Apache Kafka.  

# Getting Started
 * **Install librdkafka**: 
This template for Go depends on the availability of librdkafka(v1.3.0 or later), if you wish to build librdkafka from
source, follow the following steps to build it from the version provided in this template:  
    ```
    cd librdkafka
    ./configure --prefix /usr
    make
    sudo make install
    ```
 * **Install Confluent Client**:
 You can utilize the `go get` command to download the Confluent Client for Apache Kafka  
    ```
    go get -u gopkg.in/confluentinc/confluent-kafka-go.v1/kafka
    ```
 * **Building the Code**
    To build the source code, use the following command:
    ```
    go build main.go
    ```
 * **Running the Code**
    ```
    ./main
    ```
**Configurations can be found and updated in the [application.conf](/application.json) file**  
