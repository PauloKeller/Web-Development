package app

import jdk.nashorn.internal.runtime.regexp.joni.exception.InternalException
import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.apache.kafka.clients.consumer.KafkaConsumer
import org.apache.kafka.common.errors.WakeupException
import org.apache.kafka.common.serialization.StringDeserializer
import java.time.Duration
import java.util.*
import java.util.concurrent.CountDownLatch

fun main() {
    val bootstrapServe = "127.0.0.1:9092"
    val groupId = "my-sixth-application"
    val topic = "first_topic"
    // latch for dealing with multiple threads
    val latch = CountDownLatch(1)
    // create the runnable
    val consumerRunnable = ConsumerRunnable(bootstrapServe, groupId, topic, latch)
    // start the thread
    val thread = Thread(consumerRunnable)
    thread.start()
    // add a shutdown hook
    Runtime.getRuntime().addShutdownHook(Thread {
        println("Caught shutdown hook")
        consumerRunnable.shutdown()
        try {
            latch.await()
        } catch (e: InterruptedException) {
            e.printStackTrace()
        }
        println("Application has exited")
    })

    try {
        latch.await()
    } catch (e: InternalException) {
        e.printStackTrace()
    } finally {
        println("App is closing")
    }
}

class ConsumerRunnable: Runnable {
    private var consumer: KafkaConsumer<String, String>? = null
    private var properties: Properties = Properties()
    private var latch: CountDownLatch = CountDownLatch(1)

    constructor(bootstrapServe: String, groupId: String, topic: String, latch: CountDownLatch) {
        // consumer configs
        properties.setProperty(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServe)
        properties.setProperty(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
        properties.setProperty(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
        properties.setProperty(ConsumerConfig.GROUP_ID_CONFIG, groupId)
        properties.setProperty(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest")
        //create consumer
        consumer = KafkaConsumer<String, String>(properties)
        // subscribe consumer to our topic(s)
        consumer!!.subscribe(listOf(topic))
        this.latch = latch
    }

    override fun run() {
        // poll for new data
        try {
            while (true) {
                val records: ConsumerRecords<String, String> = consumer!!.poll(Duration.ofMillis(100)) //new in kafka 2.0.0

                for (record in records) {
                    println("Key: ${record.key()}, Value: ${record.value()}, Partition: ${record.partition()}, Offset: ${record.offset()}")
                }
            }
        } catch (e: WakeupException) {
            println("Received shutdown signal!")
        } finally {
            consumer!!.close()
            // tell our main code we're done with consumer
            latch.countDown()
        }

    }

    fun shutdown() {
        // the wakeup method is a special method to interrupt consumer.poll
        // it will throw the exception WakeUpException
        consumer!!.wakeup()
    }

}
