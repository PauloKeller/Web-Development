const Mongoose = require('mongoose');
const Schema = Mongoose.Schema;

const UserSchema = new Schema({
    username: String,
    googleId: String,
    thumbnail: String,
})

const User = Mongoose.model('user', UserSchema);

module.exports = User;