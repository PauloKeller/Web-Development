const integrationServer = require("./IntegrationServer")
const chai = require('chai');

const expect = chai.expect;

describe('Server integration', () => {
  let app;

  beforeEach((done) => {
    app = integrationServer.start(done);
  });

  afterEach((done) => {
    integrationServer.stop(app, done);
    console.log("after")
  });

  it('Should resolve User', () => {
    let query = `{
      getUser(id:"5b594c937a7fac7782b85517") {
        id
        name
        email
        points
      }
    }`;
    let expected = {
      "getUser": {
        "id": "5b594c937a7fac7782b85517",
        "name": "tiao",
        "email": "tiao@tiao.com",
        "points": 0.3333333333333333
      }
    };

    return integrationServer
      .graphqlQuery(app, query)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body).to.have.deep.equals(expected)
    })
  })

  it('Should resolve Users', () => {
    const query = `{
      getUsers {
        id
        name
        points
        email
      }
    }`;

    const expected = {
      "id": "5b594c937a7fac7782b85517",
      "name": "tiao",
      "points": 0.3333333333333333,
      "email": "tiao@tiao.com"
    };

    return integrationServer
      .graphqlQuery(app, query)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.getUsers[0]).to.deep.equal(expected)
        expect(response.body.getUsers).to.have.lengthOf.at.least(3)
      })
  })

  it('Should resolve Post', () => {
    let query = `{
      getPost(id:"5b5939b7e8da5c551f4498e8") {
        id
        title
        body
        isPretended
      }
    }`;
    let expected = {
      "getPost": {
        "id": "5b5939b7e8da5c551f4498e8",
        "title": "Post test 1",
        "body": "Post body",
        "isPretended": true
      }
    };

    return integrationServer
      .graphqlQuery(app, query)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body).to.have.deep.equals(expected)
    })
  })

  it('Should resolve Posts', () => {
    const query = `{
      getPosts {
        id
        title
        body
        isPretended
      }
    }`;

    const expected = {
      "id": "5b5939b7e8da5c551f4498e8",
      "title": "Post test 1",
      "body": "Post body",
      "isPretended": true
    };

    return integrationServer
      .graphqlQuery(app, query)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.getPosts[0]).to.deep.equal(expected)
        expect(response.body.getPosts).to.have.lengthOf.at.least(3)
      })
  })

  it('Should resolve Vote', () => {
    let query = `{
      getVote(id:"5b5bd9c95b04675c28222484") {
        id
        isPretended
      }
    }`;
    let expected = {
      "getVote": {
        "id": "5b5bd9c95b04675c28222484",
        "isPretended": true
      }
    };

    return integrationServer
      .graphqlQuery(app, query)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body).to.have.deep.equals(expected)
    })
  })

  it('Should resolve Votes', () => {
    const query = `{
      getVotes {
        id
        isPretended
      }
    }`;

    const expected = {
      "id": "5b5bd9c95b04675c28222484",
      "isPretended": true
    };

    return integrationServer
      .graphqlQuery(app, query)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.getVotes[0]).to.deep.equal(expected)
        expect(response.body.getVotes).to.have.lengthOf.at.least(3)
      })
  })

  it('Should resolve Posts from other users', () => {
    let query = `{
      getPostsFromOtherUser(userID:"5b594c937a7fac7782b85517") {
        id
        title
        body
        isPretended
      }
    }`;
    let expected = {
        "id": "5b5939b7e8da5c551f4498e8",
        "title": "Post test 1",
        "body": "Post body",
        "isPretended": true
    };

    return integrationServer
      .graphqlQuery(app, query)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.getPostsFromOtherUser[0]).to.deep.equal(expected)
        expect(response.body.getPostsFromOtherUser).to.have.lengthOf.at.least(3)
    })
  })
})