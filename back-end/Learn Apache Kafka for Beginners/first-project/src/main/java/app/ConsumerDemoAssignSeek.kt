package app

import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.apache.kafka.clients.consumer.KafkaConsumer
import org.apache.kafka.common.TopicPartition
import org.apache.kafka.common.serialization.StringDeserializer
import java.time.Duration
import java.util.*

fun main() {
    val properties = Properties()
    val bootstrapServe = "127.0.0.1:9092"
    val topic = "first_topic"
    // consumer configs
    properties.setProperty(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServe)
    properties.setProperty(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
    properties.setProperty(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
    properties.setProperty(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest")

    //create consumer
    val consumer = KafkaConsumer<String, String>(properties)

    // assign and seek mostly used to replay data or fetch a specific message

    // assign
    val partitionToReadFrom = TopicPartition(topic, 0)
    val offsetToReadFrom = 15L
    consumer.assign(listOf(partitionToReadFrom))

    // seek
    consumer.seek(partitionToReadFrom, offsetToReadFrom)

    val numberOfMessagesToRead = 5
    var keepOnReading = true
    var numberOfMessagesReadSoFar = 0

    // poll for new data
    while (keepOnReading) {
        val records: ConsumerRecords<String, String> = consumer.poll(Duration.ofMillis(100)) //new in kafka 2.0.0

        for (record in records) {
            numberOfMessagesReadSoFar += 1
            println("Key: ${record.key()}, Value: ${record.value()}, Partition: ${record.partition()}, Offset: ${record.offset()}")
            if (numberOfMessagesReadSoFar >=  numberOfMessagesToRead)  {
                keepOnReading = false
                break
            }
        }

        println("Exiting the app")
    }
}
