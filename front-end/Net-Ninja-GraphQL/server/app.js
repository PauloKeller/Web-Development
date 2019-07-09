const express = require('express')
const graphqlHTTP = require('express-graphql')
const schema = require('./schema/schema')
const mongoose = require('mongoose')
const cors = require('cors')

const app = express()


// allow cross-origin requests
app.use(cors())

// connect to mlab database
// make sure to replace my db string & creds with your own
mongoose.connect('mongodb://paulo:1234@ds117010.mlab.com:17010/gql-ninja')
mongoose.connection.once('open', () => {
    console.log('connected to database')
})

app.use('/graphql', graphqlHTTP({
    schema,
    graphiql: true
}))

app.listen(4000, () => {
    console.log('now listening for requests on port 4000')
})