package app

import org.apache.kafka.clients.producer.Callback
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerConfig
import org.apache.kafka.clients.producer.ProducerRecord
import org.apache.kafka.common.serialization.StringSerializer
import java.util.*


fun main() {
   val bootstrapServers = "127.0.0.1:9092"

   // create Producer properties
   val properties = Properties()
   properties.setProperty(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers)
   properties.setProperty(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer::class.java.name)
   properties.setProperty(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer::class.java.name)
   // create the producer
   val producer = KafkaProducer<String, String>(properties)
   // create a producer record
   val record = ProducerRecord<String, String>("first_topic", "hello world")
   // send data - async
   producer.send(record)
   // flush data
   producer.flush()
   // flush and close
   producer.close()
}