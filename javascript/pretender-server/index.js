let express = require('express')
let graphqlHTTP = require('express-graphql')
let schema = require('./schema/Schema')
let cors = require('cors')
let favicon = require('serve-favicon')
let mongoose = require('mongoose')

let app = express()

mongoose.connect('mongodb://root:root1234@ds141631.mlab.com:41631/pretender', { useNewUrlParser:true })
mongoose.connection.once('open', () => {
    console.log('Connected to database')
})

app.use(express.static('public'))
app.use(favicon(__dirname + '/public/favicon.ico'));
app.use(cors())

app.use('/graphql', graphqlHTTP({
    schema,
    graphiql: true
}))

app.listen(process.env.PORT || 4000, () =>{
    console.log('Now listening for requests')
})