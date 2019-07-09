const Express = require('express');
const AuthRoutes = require('./routes/AuthRoutes');
const ProfileRoutes = require('./routes/ProfileRoutes');
const PassPortSetup = require('./config/PassportSetup');
const Mongoose = require('mongoose');
const Keys = require('./config/Keys');
const CookieSession = require('cookie-session');
const Passport = require('passport')

const App = Express();

// set up view engine
App.set('view engine', 'ejs');

App.use(CookieSession({
    maxAge: 24 * 60 * 60 * 1000,
    keys: [Keys.session.cookieKey]
}));

// initialize passport
App.use(Passport.initialize());
App.use(Passport.session());

// connect to mongodb
Mongoose.connect(Keys.mongodb.dbURI, { useNewUrlParser:true }, () => {
    console.log('connected to mongodb');
});

// set up routes
App.use('/auth', AuthRoutes);
App.use('/profile', ProfileRoutes);

// create home route
App.get('/', (req, res) => {
    res.render('Home', { user: req.user });
});

App.listen(3000, () => {
    console.log('app now listening for requests on port 3000');
});