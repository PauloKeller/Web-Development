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
   // create a producer record
   for (i in 0..10) {
      val record = ProducerRecord<String, String>("first_topic", "hello world $i")
      // send data - async
      producer.send(record) { metadata, exception ->
         // executes every time a record is successfully sent or an exception is thrown
         if (exception == null) {
            // the record was successfully sent
            print("Received new metadata. \n" +
                    "Topic: ${metadata.topic()}\n" +
                    "Partition: ${metadata.partition()}\n" +
                    "Offset: ${metadata.offset()}\n"+
                    "Timestamp: ${metadata.timestamp()}")
         } else {
            logger.error("Error while producing", exception)
         }
      }
   }

   // flush data
   producer.flush()
   // flush and close
   producer.close()
}