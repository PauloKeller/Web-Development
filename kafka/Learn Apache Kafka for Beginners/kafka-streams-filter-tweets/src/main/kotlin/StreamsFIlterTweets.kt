import com.google.gson.JsonParser
import org.apache.kafka.common.serialization.Serdes
import org.apache.kafka.streams.KafkaStreams
import org.apache.kafka.streams.StreamsBuilder
import org.apache.kafka.streams.StreamsConfig
import java.lang.Exception
import java.util.*

fun extractUserFollowersInTweet(tweetJson: String): Int {
    val jsonPaser = JsonParser()
    return try {
        jsonPaser.parse(tweetJson)
                .asJsonObject
                .get("user")
                .asJsonObject
                .get("followers_count")
                .asInt
    } catch (e:Exception) {
        0
    }

}

fun main() {
    // create properties
    val properties = Properties()
    properties.setProperty(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, "127.0.0.1:9092")
    properties.setProperty(StreamsConfig.APPLICATION_ID_CONFIG, "demo-kafka-streams")
    properties.setProperty(StreamsConfig.DEFAULT_KEY_SERDE_CLASS_CONFIG, Serdes.StringSerde::class.java.name)
    properties.setProperty(StreamsConfig.DEFAULT_VALUE_SERDE_CLASS_CONFIG, Serdes.StringSerde::class.java.name)

    // create the topology
    val streamsBuilder = StreamsBuilder()

    // input topic
    val inputTopic = streamsBuilder.stream<String, String>("twitter_tweets")
    val filteredStream = inputTopic.filter { key, jsonTweet ->
        // filter for tweets which has a user of over 10000 followers
        extractUserFollowersInTweet(jsonTweet) > 1000
    }
    filteredStream.to("important_tweets")

    // build the topology
    val kafkaStreams = KafkaStreams(streamsBuilder.build(), properties)

    // start our streams application
    kafkaStreams.start()
}