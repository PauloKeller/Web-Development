const express = require('express');
const graphql = require('graphql').graphql;
const request = require('request-promise');
let mongoose = require('mongoose')
const rootSchema = require('../schema/Schema')

function start(done, appPort) {
  const app = express();
  const PORT = appPort || 9000;

  mongoose.connect('mongodb://root:root1234@ds141631.mlab.com:41631/pretender', { useNewUrlParser:true })
  mongoose.connection.once('open', () => {
    console.log("connected")
  })

  app.get('/graphql', (req, res) => {
    const graphqlQuery = req.query.graphqlQuery;
    if (!graphqlQuery) {
      return res.status(500).send('You must provide a query');
    }

    return graphql(rootSchema, graphqlQuery)
      .then(response => response.data)
      .then((data) => res.json(data))
      .catch((err) => console.error(err));
  });

  return app.listen(PORT, () => {
    console.log('Server started at port [%s]', PORT);
    done();
  });
};

function stop(app, done) {
  app.close();
  mongoose.connection.close()
  done()
};

function graphqlQuery(app, query) {
  return request({
    baseUrl : `http://localhost:${app.address().port}`,
    uri : '/graphql',
    qs : {
      graphqlQuery : query
    },
    resolveWithFullResponse: true,
    json: true
  })
};

module.exports = {
  start,
  stop,
  graphqlQuery
};