const Router = require('express').Router();

const authCheck = (req, res, next) => {
    if (!req.user) {
        // user not logged in 
        res.redirect('/auth/login');
    } else {
        // logged in 
        next();
    }
};

Router.get('/', authCheck, (req, res) => {
    res.render('Profile', { user: req.user });
});

module.exports = Router;