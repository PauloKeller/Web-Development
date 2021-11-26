let mongoose = require('mongoose')
let Schema = mongoose.Schema

let UserShema = new Schema({
    name: String,
    email: String,
    points: Number,
    created_at: Date,
})

module.exports = mongoose.model('User', UserShema)