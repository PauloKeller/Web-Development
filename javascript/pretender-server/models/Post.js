let mongoose = require('mongoose')
let Schema = mongoose.Schema

let PostShema = new Schema({
    title: String,
    body: String,
    isPretended: Boolean,
    authorID: String,
    created_at: Date
})

module.exports = mongoose.model('Post', PostShema)