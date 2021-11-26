let mongoose = require('mongoose')
let Schema = mongoose.Schema

let VoteShema = new Schema({
    isPretended: Boolean,
    userID: String,
    postID: String,
    created_at: Date,
})

module.exports = mongoose.model('Vote', VoteShema)