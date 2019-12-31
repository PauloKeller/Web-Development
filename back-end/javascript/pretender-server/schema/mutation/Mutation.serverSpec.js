let integrationServer = require("../../IntegrationServer")
let chai = require('chai')
let expect = chai.expect

describe('Server integration', () => {
    let app;
  
    beforeEach((done) => {
      app = integrationServer.start(done);
    });
  
    afterEach((done) => {
      integrationServer.stop(app, done);
      console.log("after")
    });

    it('Should add a User', () => {
    let mutation = `mutation {
      createUser(name:"T-User", email:"tuser@tuser.com") {
        name
        email
        points
      }
    }`;
    let expected = {
      "name": "T-User",
      "email": "tuser@tuser.com",
      "points": 0,
    };

    return integrationServer
      .graphqlMutation(app, mutation)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.createUser).to.deep.equal(expected)
    })
  })

  it('Should add a Post', () => {
    let mutation = `mutation {
      createPost(title:"T-Post", body:"T-Body", isPretended:true, authorID: "5b57dedd0c9ac11eb2a3a625") {
        title
        body
        isPretended
      }
    }`;
    let expected = {
      "title": "T-Post",
      "body": "T-Body",
      "isPretended": true,
    };

    return integrationServer
      .graphqlMutation(app, mutation)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.createPost).to.deep.equal(expected)
    })
  })

  it('Should add a Vote', () => {
    let mutation = `mutation {
      createVote(isPretended:true, userID: "5b594c937a7fac7782b85517", postID: "5b5bc1f0dfd65b46d23de55c") {
        isPretended
      }
    }`;
    let expected = {
      "isPretended": true,
    };

    return integrationServer
      .graphqlMutation(app, mutation)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.createVote).to.deep.equal(expected)
    })
  })

  it('Should login into the app', () => {
    let mutation = `mutation {
      login(email:"tiao@tiao.com") {
        id
        name
        email
        points
      }
    }`;
    let expected = {
      "id": "5b594c937a7fac7782b85517",
      "name": "tiao",
      "points": 0.3333333333333333,
      "email": 'tiao@tiao.com',
    };

    return integrationServer
      .graphqlMutation(app, mutation)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.login).to.deep.equal(expected)
    })
  })

  it('Should update Users points when a right hint are done', () => {
    let mutation = `mutation {
      updateUsersPointsWhenRightHint(postID: "5b5d2544996b1330a02b65f7") {
        id
        title
        body
        isPretended
      }
    }`;

    let expected = {
      "id": "5b5d2544996b1330a02b65f7",
      "title": "T-Post",
      "body": "T-Body",
      "isPretended": true,
    };

    return integrationServer
      .graphqlMutation(app, mutation)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.updateUsersPointsWhenRightHint).to.deep.equal(expected)
    })
  }).timeout(8000)

  it('Should update Users points when a wrong hint are done', () => {
    let mutation = `mutation {
      updateUsersPointsWhenWrongHint(postID: "5b5d2544996b1330a02b65f7") {
        id
        title
        body
        isPretended
      }
    }`;

    let expected = {
      "id": "5b5d2544996b1330a02b65f7",
      "title": "T-Post",
      "body": "T-Body",
      "isPretended": true,
    };

    return integrationServer
      .graphqlMutation(app, mutation)
      .then((response) => {
        expect(response.statusCode).to.equal(200)
        expect(response.body.updateUsersPointsWhenWrongHint).to.deep.equal(expected)
    })
  }).timeout(8000)

})