let { 
    GraphQLList,
    GraphQLID,
    GraphQLNonNull,
    GraphQLFloat,
    GraphQLString,
    GraphQLBoolean
  } = require('graphql')

let  chai = require('chai')
let { UserType, VoteType, PostType  } = require('./Types')

let expect = chai.expect

describe('UserType', () => {
    it('Should have a non null id field of type ID', () => {
        expect(UserType.getFields()).to.have.property('id')
        expect(UserType.getFields().id.type).to.deep.equals(new GraphQLNonNull(GraphQLID))
    })

    it('Should have a non null name field of type String', () => {
        expect(UserType.getFields()).to.have.property('name')
        expect(UserType.getFields().name.type).to.deep.equals(new GraphQLNonNull(GraphQLString))
    })
    
    it('Should have a non null email field of type String', () => {
        expect(UserType.getFields()).to.have.property('email')
        expect(UserType.getFields().email.type).to.deep.equals(new GraphQLNonNull(GraphQLString))
    })
    
    it('Should have a non null points field of type Float', () => {
        expect(UserType.getFields()).to.have.property('points')
        expect(UserType.getFields().points.type).to.deep.equals(new GraphQLNonNull(GraphQLFloat))
    })

    it('Should have a posts field of type PostType List', () => {
        expect(UserType.getFields()).to.have.property('posts')
        expect(UserType.getFields().posts.type).to.deep.equals(new GraphQLList(PostType))
    })
})

describe('PostType', () => {
    it('Should have a non null id field of type ID', () => {
        expect(PostType.getFields()).to.have.property('id')
        expect(PostType.getFields().id.type).to.deep.equals(new GraphQLNonNull(GraphQLID))
    })

    it('Should have a non null title field of type String', () => {
        expect(PostType.getFields()).to.have.property('title')
        expect(PostType.getFields().title.type).to.deep.equals(new GraphQLNonNull(GraphQLString))
    })
    
    it('Should have a non null body field of type String', () => {
        expect(PostType.getFields()).to.have.property('body')
        expect(PostType.getFields().body.type).to.deep.equals(new GraphQLNonNull(GraphQLString))
    })
    
    it('Should have a author field of type UserType', () => {
        expect(PostType.getFields()).to.have.property('author')
        expect(PostType.getFields().author.type).to.deep.equals(UserType)
    })

    it('Should have a non null isPretended field of type Boolean', () => {
        expect(PostType.getFields()).to.have.property('isPretended')
        expect(PostType.getFields().isPretended.type).to.deep.equals(new GraphQLNonNull(GraphQLBoolean))
    })

    it('Should have a votes field of type VoteType List', () => {
        expect(PostType.getFields()).to.have.property('votes')
        expect(PostType.getFields().votes.type).to.deep.equals(new GraphQLList(VoteType))
    })
})

describe('VoteType', () => {
    it('Should have a non null id field of type ID', () => {
        expect(VoteType.getFields()).to.have.property('id')
        expect(VoteType.getFields().id.type).to.deep.equals(new GraphQLNonNull(GraphQLID))
    })

    it('Should have a user field of type UserType', () => {
        expect(VoteType.getFields()).to.have.property('user')
        expect(VoteType.getFields().user.type).to.deep.equals(UserType)
    })
    
    it('Should have post field of type PostType', () => {
        expect(VoteType.getFields()).to.have.property('post')
        expect(VoteType.getFields().post.type).to.deep.equals(PostType)
    })
    
    it('Should have a non null isPretended field of type Boolean', () => {
        expect(VoteType.getFields()).to.have.property('isPretended')
        expect(VoteType.getFields().isPretended.type).to.deep.equals(new GraphQLNonNull(GraphQLBoolean))
    })

    it('Should have a non null created_at field of type String', () => {
        expect(VoteType.getFields()).to.have.property('created_at')
        expect(VoteType.getFields().created_at.type).to.deep.equals(new GraphQLNonNull(GraphQLString))
    })  
})