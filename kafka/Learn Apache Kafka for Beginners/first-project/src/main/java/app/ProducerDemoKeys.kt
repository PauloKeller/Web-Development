package app

import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerConfig
import org.apache.kafka.clients.producer.ProducerRecord
import org.apache.kafka.common.serialization.StringSerializer
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import java.util.*

fun main() {
   val logger:Logger = LoggerFactory.getLogger("ProducerDemoWithCallBack")

   val bootstrapServers = "127.0.0.1:9092"

   // create Producer properties
   val properties = Properties()
   properties.setProperty(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers)
   properties.setProperty(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer::class.java.name)
   properties.setProperty(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer::class.java.name)
   // create the producer
   val producer = KafkaProducer<String, String>(properties)
   for (i in 0..10) {
      // create a producer record
      val topic = "first_topic"
      val value = "hello world $i"
      val key = "id_$i"
      val record = ProducerRecord<String, String>(topic, key, value)

      println("Key: $key")
      // send data - async
      producer.send(record) { metadata, exception ->
         // executes every time a record is successfully sent or an exception is thrown
         if (exception == null) {
            // the record was successfully sent
            println("Received new metadata. \n" +
                    "Topic: ${metadata.topic()}\n" +
                    "Partition: ${metadata.partition()}\n" +
                    "Offset: ${metadata.offset()}\n"+
                    "Timestamp: ${metadata.timestamp()}")
         } else {
            logger.error("Error while producing", exception)
         }
      }.get() // don't do this in production
   }

   // flush data
   producer.flush()
   // flush and close
   producer.close()
}