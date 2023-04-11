package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

var (
	kafkaAddr = []string{"192.168.0.164:9092"}
	topic     = "test-assign"
)

func main() {
	flag.Parse()

	if len(flag.Args()) >= 2 {
		startProducer(kafkaAddr, topic)
		return
	}

	// create and start consumer
	consumer, err := NewConsumer(kafkaAddr, nil)
	if err != nil {
		panic(err)
	}

	for {
		log.Println("consumer is running...")
		// will block
		err := Consume(context.Background(), consumer, topic, func(partition int32, partitionConsumer sarama.PartitionConsumer, message *sarama.ConsumerMessage) {
			log.Println("consumer new mq, paritition=", partition, ",topic:", message.Topic, ",offset:", message.Offset)
		})

		if err != nil {
			log.Println(err)
		} else {
			log.Println("consume exit")
		}
		time.Sleep(time.Second * 3)
	}

}

func startProducer(addrs []string, topic string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		p, offset, err := producer.SendMessage(&sarama.ProducerMessage{
			Key:   sarama.StringEncoder(strconv.Itoa(i)),
			Value: sarama.StringEncoder("hello" + strconv.Itoa(i)),
			Topic: topic,
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("produce success, partition:", p, ",offset:", offset)
		}
		time.Sleep(time.Second)
	}
	log.Println("exit producer.")
}
