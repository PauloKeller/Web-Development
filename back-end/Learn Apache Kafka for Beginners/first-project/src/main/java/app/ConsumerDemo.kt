package app

import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.apache.kafka.clients.consumer.KafkaConsumer
import org.apache.kafka.common.serialization.StringDeserializer
import java.time.Duration
import java.util.*

fun main() {
    val properties = Properties()
    val bootstrapServe = "127.0.0.1:9092"
    val groupId = "my-fourth-application"
    val topic = "first_topic"
    // consumer configs
    properties.setProperty(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServe)
    properties.setProperty(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
    properties.setProperty(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
    properties.setProperty(ConsumerConfig.GROUP_ID_CONFIG, groupId)
    properties.setProperty(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest")

    //create consumer
    val consumer = KafkaConsumer<String, String>(properties)
    // subscribe consumer to our topic(s)
    consumer.subscribe(listOf(topic))
    // poll for new data
    while (true) {
        val records: ConsumerRecords<String, String> = consumer.poll(Duration.ofMillis(100)) //new in kafka 2.0.0

        for (record in records) {
            println("Key: ${record.key()}, Value: ${record.value()}, Partition: ${record.partition()}, Offset: ${record.offset()}")
        }
    }
}
