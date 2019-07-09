const Router = require('express').Router();
const Passport = require('passport');

// auth login
Router.get('/login', (req, res) => {
    res.render('Login', { user: req.user });
});

// auth logout
Router.get('/logout', (req, res) => {
    // handle with passport
    req.logout();
    res.redirect('/');
})

// auth with google
Router.get('/google', Passport.authenticate('google', {
    scope: ['profile']
}));

// callback route for google to redirect to
Router.get('/google/redirect', Passport.authenticate('google'), (req, res) => {
    res.redirect('/profile/');
})

module.exports = Router;