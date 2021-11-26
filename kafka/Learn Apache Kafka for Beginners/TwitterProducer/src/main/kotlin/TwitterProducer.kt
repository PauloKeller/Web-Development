import com.google.common.collect.Lists
import com.twitter.hbc.ClientBuilder
import com.twitter.hbc.core.Client
import com.twitter.hbc.core.Constants
import com.twitter.hbc.core.Hosts
import com.twitter.hbc.core.HttpHosts
import com.twitter.hbc.core.endpoint.StatusesFilterEndpoint
import com.twitter.hbc.core.processor.StringDelimitedProcessor
import com.twitter.hbc.httpclient.auth.Authentication
import com.twitter.hbc.httpclient.auth.OAuth1
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.ProducerConfig
import org.apache.kafka.clients.producer.ProducerRecord
import org.apache.kafka.common.serialization.StringSerializer
import java.util.*
import java.util.concurrent.BlockingQueue
import java.util.concurrent.LinkedBlockingQueue
import java.util.concurrent.TimeUnit


class TwitterProducer {
    private var consumerKey: String = "JoOi39PrDAiYvRUEBIFm4d58Q"
    private var consumerSecret: String = "q01AnjTjJmhquHUCD2xFDDqgEA1Akg6v6v7Uyfsgj5E2iGubTD"
    private var token: String = "968495699947122688-kPmynC3o6JsfkUuclQA8eFg7snRH1IM"
    private var secret: String = "l9xeyPUM4xRWUK5ww6lbCVVwRvBsyYVgvWro0O4Hxuau4"

    val terms: List<String> = Lists.newArrayList("bitcoin", "usa", "politics", "sporte", "soccer")

    fun run() {
        print("Setup")
        /** Set up your blocking queues: Be sure to size these properly based on expected TPS of your stream */
        val msgQueue: BlockingQueue<String> = LinkedBlockingQueue(100000)
        val client = createTwitterClient(msgQueue)
        client.connect()

        val producer = createKafkaProducer()

        Runtime.getRuntime().addShutdownHook(object: Thread()  {
            override fun run() {
                client.stop()
                producer.close()
            }
        })

        while (!client.isDone) {
            var msg: String? = null
            try {
                msg = msgQueue.poll(5, TimeUnit.SECONDS)
            } catch (e: InterruptedException) {
                e.printStackTrace()
                client.stop()
            }
            if (msg != null) {
                print(msg)
                producer.send(ProducerRecord("twitter_tweets", null, msg))
            }
        }
        print("End of application")
    }

    private fun createTwitterClient(msgQueue: BlockingQueue<String>): Client {
        /** Declare the host you want to connect to, the endpoint, and authentication (basic auth or oauth) */
        /** Declare the host you want to connect to, the endpoint, and authentication (basic auth or oauth)  */
        val hosebirdHosts: Hosts = HttpHosts(Constants.STREAM_HOST)
        val hosebirdEndpoint = StatusesFilterEndpoint()


        hosebirdEndpoint.trackTerms(terms)

        val hosebirdAuth: Authentication = OAuth1(consumerKey, consumerSecret, token, secret)

        val builder = ClientBuilder()
                .name("Hosebird-Client-01") // optional: mainly for the logs
                .hosts(hosebirdHosts)
                .authentication(hosebirdAuth)
                .endpoint(hosebirdEndpoint)
                .processor(StringDelimitedProcessor(msgQueue))

        return builder.build()
    }

    private fun createKafkaProducer(): KafkaProducer<String, String> {
        val bootstrapServers = "127.0.0.1:9092"
        val properties = Properties()
        properties.setProperty(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers)
        properties.setProperty(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer::class.java.name)
        properties.setProperty(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer::class.java.name)
        // create safe producer
        properties.setProperty(ProducerConfig.ENABLE_IDEMPOTENCE_CONFIG, "true")
        properties.setProperty(ProducerConfig.ACKS_CONFIG, "all")
        properties.setProperty(ProducerConfig.RETRIES_CONFIG, Integer.MAX_VALUE.toString())
        properties.setProperty(ProducerConfig.MAX_IN_FLIGHT_REQUESTS_PER_CONNECTION, "5")

        // high throughput producer (at the expense of a bit of latency and CPU usage)
        properties.setProperty(ProducerConfig.COMPRESSION_TYPE_CONFIG, "snappy")
        properties.setProperty(ProducerConfig.LINGER_MS_CONFIG, "20")
        properties.setProperty(ProducerConfig.BATCH_SIZE_CONFIG, (32 * 1024).toString())

        // create the producer
        return KafkaProducer(properties)
    }

}

fun main() {
    TwitterProducer().run()
}