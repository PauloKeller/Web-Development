let {
  GraphQLObjectType,
  GraphQLNonNull,
  GraphQLString,
  GraphQLList,
  GraphQLFloat,
  GraphQLBoolean,
  GraphQLID
}  = require('graphql')

let User = require('../../models/User')
let Post = require('../../models/Post')
let Vote = require('../../models/Vote')

let UserType = new GraphQLObjectType({
  name: 'User',
  description: 'A user',
  fields: () => ({
    id: {
      type: new GraphQLNonNull(GraphQLID),
      description: "A user's id",
    },
    name: {
      type: new GraphQLNonNull(GraphQLString),
      description: "A user's name",
    },
    email: {
      type: new GraphQLNonNull(GraphQLString),
      description: "A user's email",
    },
    points: {
      type: new GraphQLNonNull(GraphQLFloat),
      description: "A user's points",
    },
    posts: {
      type: new GraphQLList(PostType),
      description: "A user's posts",
      resolve(parent, args) {
        return Post.find({ authorID: parent.id })
      }
    }
  })
})

let PostType = new GraphQLObjectType({
  name: 'Post',
  description: 'A post',
  fields: () => ({
    id: {
      type: new GraphQLNonNull(GraphQLID),
      description: "A post's id",
    },
    created_at: {
      type: new GraphQLNonNull(GraphQLString),
      description: 'Time post was created'
    },
    title: {
      type: new GraphQLNonNull(GraphQLString),
      description: 'Title'
    },
    body: {
      type: new GraphQLNonNull(GraphQLString),
      description: 'Body text'
    },
    author: {
      type: UserType,
      description: "A user that create",
      resolve(parent, args) {
        return User.findById(parent.authorID)
      }
    },
    isPretended: {
      type: new GraphQLNonNull(GraphQLBoolean),
      description: 'Is a pretended post'
    },
    votes: {
      type: new GraphQLList(VoteType),
      description: "Users votes",
      resolve(parent, args) {
        return Vote.find({ postID: parent.id })
      }
    },
  })
})

let VoteType = new GraphQLObjectType({
  name: 'Vote',
  description: 'A vote',
  fields: () => ({
    id: {
      type: new GraphQLNonNull(GraphQLID),
      description: "A post's id",
    },
    user: {
      type: UserType,
      description: "A user that voted",
      resolve(parent, args) {
        return User.findById(parent.userID)
      }
    },
    post: {
      type: PostType,
      description: 'Post that was voted',
      resolve(parent, args) {
        return Post.findById(parent.postID)
      }
    },
    isPretended: {
      type: new GraphQLNonNull(GraphQLBoolean),
      description: 'Is a pretended post'
    },
    created_at: {
      type: new GraphQLNonNull(GraphQLString),
      description: 'Time vote was done'
    },
  })
})

module.exports = {
  UserType: UserType,
  PostType: PostType,
  VoteType: VoteType
}