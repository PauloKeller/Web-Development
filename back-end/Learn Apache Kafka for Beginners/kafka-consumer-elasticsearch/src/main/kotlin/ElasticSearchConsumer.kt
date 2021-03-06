import com.google.gson.JsonParser
import org.apache.http.HttpHost
import org.apache.http.auth.AuthScope
import org.apache.http.auth.UsernamePasswordCredentials
import org.apache.http.client.CredentialsProvider
import org.elasticsearch.client.RestClientBuilder
import org.apache.http.impl.client.BasicCredentialsProvider
import org.apache.kafka.clients.consumer.ConsumerConfig
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.apache.kafka.clients.consumer.KafkaConsumer
import org.apache.kafka.common.serialization.StringDeserializer
import org.elasticsearch.action.bulk.BulkRequest
import org.elasticsearch.action.index.IndexRequest
import org.elasticsearch.client.RequestOptions
import org.elasticsearch.client.RestClient
import org.elasticsearch.client.RestHighLevelClient
import org.elasticsearch.common.xcontent.XContentType
import org.slf4j.LoggerFactory
import java.lang.Exception
import java.time.Duration
import java.util.*
import kotlin.math.log

class ElasticSearchConsumer {
    fun createClient(): RestHighLevelClient {
        // credentials provider help supply username and password
        val hostname = "kafka-course-8642358721.us-west-2.bonsaisearch.net" // localhost or bonsai url
        val username = "ywl9uvf3bh" // needed only for bonsai
        val password = "xw1zqux8yz" // needed only for bonsai

        // credentials provider help supply username and password
        val credentialsProvider: CredentialsProvider = BasicCredentialsProvider()
        credentialsProvider.setCredentials(
            AuthScope.ANY,
            UsernamePasswordCredentials(username, password)
        )

        val builder: RestClientBuilder = RestClient.builder(
            HttpHost(hostname, 443, "https")
        ).setHttpClientConfigCallback { httpClientBuilder ->
            httpClientBuilder.setDefaultCredentialsProvider(credentialsProvider)
        }

        return RestHighLevelClient(builder)
    }

    fun createConsumer(topic: String): KafkaConsumer<String, String> {
        val properties = Properties()
        val bootstrapServe = "127.0.0.1:9092"
        val groupId = "kafka-demo-elasticsearch"

        properties.setProperty(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServe)
        properties.setProperty(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
        properties.setProperty(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer::class.java.name)
        properties.setProperty(ConsumerConfig.GROUP_ID_CONFIG, groupId)
        properties.setProperty(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest")
        properties.setProperty(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, "false")
        properties.setProperty(ConsumerConfig.MAX_POLL_RECORDS_CONFIG, "100")

        val consumer = KafkaConsumer<String, String>(properties)
        consumer.subscribe(listOf(topic))

        return consumer
    }

    fun extractIdFromTweet(tweetJson: String): String {
        val jsonPaser = JsonParser()
        return jsonPaser.parse(tweetJson).asJsonObject.get("id_str").asString
    }
}

fun main() {
    val logger = LoggerFactory.getLogger(ElasticSearchConsumer::class.java.name)

    val elasticSearchConsumer = ElasticSearchConsumer()
    val client = elasticSearchConsumer.createClient()
    val consumer = elasticSearchConsumer.createConsumer("twitter_tweets")

    while (true) {
        val records: ConsumerRecords<String, String> = consumer.poll(Duration.ofMillis(100)) //new in kafka 2.0.0
        val recordCount = records.count()
        logger.info("Received $recordCount records")

        val bulkRequest = BulkRequest()

        for (record in records) {
            // 2 strategies
            // kafka generic ID
            // val kafkaGenericId = record.topic() + "_" + record.partition() + "_" + record.offset()
            try {
                // twitter feed specific id
                // use the resource id
                val id = elasticSearchConsumer.extractIdFromTweet(record.value())

                // where we insert data into ElasticSearch
                val indexRequest = IndexRequest(
                    "twitter",
                    "tweets",
                    id // this is to make our consumer idempotent
                ).source(record.value(), XContentType.JSON)

                bulkRequest.add(indexRequest) // we add to our bulk request (take no time)
            } catch (e: Exception) {
                logger.warn("skipping bad data: " + record.value())
            }

        }

        if (recordCount > 0) {
            val bulkResponse = client.bulk(bulkRequest, RequestOptions.DEFAULT)

            logger.info("Committing offsets...")
            consumer.commitSync()
            logger.info("Offsets have been committed")
            Thread.sleep(1000)
        }
    }

    // close the client
    // client.close()
}