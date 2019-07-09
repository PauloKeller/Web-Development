const Passport = require('passport');
const GoogleStrategy = require('passport-google-oauth20');
const Keys = require('./Keys');
const User = require('../models/User');

Passport.serializeUser((user, done) => {
    done(null, user.id);
});

Passport.deserializeUser((id, done) => {
    User.findById(id).then((user) => {
        done(null, user);
    }); 
});

Passport.use(
    new GoogleStrategy({
        // options for the google strat
        callbackURL: '/auth/google/redirect',
        clientID: Keys.google.clientID,
        clientSecret: Keys.google.clientSecret
    }, (accessToken, refreshToken, profile, done) => {
        // check if user already exists in our db
        User.findOne({ googleId: profile.id }).then((currentUser) => {
            if (currentUser) {
                // already have the user
                console.log('user is:', currentUser);
                done(null, currentUser);
            } else {
                // create user in db
                new User({
                    username: profile.displayName,
                    googleId: profile.id,
                    thumbnail: profile._json.image.url
                }).save().then((newUser) => {
                    console.log('new user created:' + newUser);
                    done(null, newUser);
                });
            }
        })         
    })
);