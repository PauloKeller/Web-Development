const express = require('express')
const bodyParser= require('body-parser')
const app = express()
const MongoClient = require('mongodb').MongoClient
var database
app.set('view engine', 'ejs')
app.use(express.static(__dirname + '/styles'));

app.use(bodyParser.urlencoded({extended: true}))

MongoClient.connect('mongodb://localhost:27017/star-wars-quotes', (err, client) => {
    if (err) return console.log(err)       
    
    database = client.db('star-wars-quotes');
    
    app.listen(3000, function() {
        console.log('listening on 3000')
    })
})

app.get('/', (req, res) => {
    database.collection('quotes').find().toArray(function(err, result) {
        if (err) console.log(err)

        res.render('index.ejs', {quotes: result})
    })
})

app.post('/quotes', (req, res) => {
    database.collection('quotes').save(req.body, (err, result) => {
        if (err) return console.log(err)

        console.log('saved to database')
        res.redirect('/')
    })
})



