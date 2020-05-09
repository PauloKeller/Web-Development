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
import java.util.concurrent.BlockingQueue
import java.util.concurrent.LinkedBlockingQueue
import java.util.concurrent.TimeUnit


class TwitterProducer {
    private var consumerKey: String = "JoOi39PrDAiYvRUEBIFm4d58Q"
    private var consumerSecret: String = "q01AnjTjJmhquHUCD2xFDDqgEA1Akg6v6v7Uyfsgj5E2iGubTD"
    private var token: String = "968495699947122688-kPmynC3o6JsfkUuclQA8eFg7snRH1IM"
    private var secret: String = "l9xeyPUM4xRWUK5ww6lbCVVwRvBsyYVgvWro0O4Hxuau4"

    fun run() {
        print("Setup")
        /** Set up your blocking queues: Be sure to size these properly based on expected TPS of your stream */
        val msgQueue: BlockingQueue<String> = LinkedBlockingQueue(100000)
        val client = createTwitterClient(msgQueue)
        client.connect()

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
            }
        }
        print("End of application")
    }

    private fun createTwitterClient(msgQueue: BlockingQueue<String>): Client {
        /** Declare the host you want to connect to, the endpoint, and authentication (basic auth or oauth) */
        /** Declare the host you want to connect to, the endpoint, and authentication (basic auth or oauth)  */
        val hosebirdHosts: Hosts = HttpHosts(Constants.STREAM_HOST)
        val hosebirdEndpoint = StatusesFilterEndpoint()

        val terms: List<String> = Lists.newArrayList("bitcoin", "api")
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

}

fun main() {
    TwitterProducer().run()
}