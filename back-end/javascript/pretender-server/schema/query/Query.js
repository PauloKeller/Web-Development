let {
  GraphQLObjectType,  
  GraphQLList,
  GraphQLID,
  GraphQLNonNull
} = require('graphql')

let { UserType, VoteType, PostType  } = require('../type/Types')
let User = require('../../models/User')
let Post = require('../../models/Post')
let Vote = require('../../models/Vote')
let Logger = require('../../Logger')

const Query = new GraphQLObjectType({
  name: 'RootQueryType',
  description: "Server Queries",
  fields: {
    getUser: {
      type: UserType,
      description: "Retrive a user by id",
      args: {
        id: {
          name: 'id',
          description: "The id of a user",
          type: new GraphQLNonNull(GraphQLID)
        }
      },
      resolve(parent, args) {
        Logger.log('info', 'Get User', {
          id: args.id,
          time: new Date().toLocaleString(),
        })
        return User.findById(args.id)
      }
    },
    getUsers: {
      description: "Retrive all users",
      type: new GraphQLList(UserType),
      resolve(parent, args){
        Logger.log('info', 'Get Users', {
          time: new Date().toLocaleString(),
        })
        return User.find()
      }
    },
    getPost: {
      type: PostType,
      description: "Retrive a post by id",
      args: {
        id: {
          name: 'id',
          description: "The id of a post",
          type: new GraphQLNonNull(GraphQLID)
        }
      },
      resolve(parent, args){
        Logger.log('info', 'Get Post', {
          time: new Date().toLocaleString(),
          id: args.id,
        })
        return Post.findById(args.id)
      }
    },
    getPosts: {
      type: new GraphQLList(PostType),
      description: "Retrive all posts",
      resolve(parent, args){
        Logger.log('info', 'Get Posts', {
          time: new Date().toLocaleString(),
        })
        return Post.find()
      }
    },
    getVote: {
      type: VoteType,
      description: "Retrive a vote by id",
      args: {
        id: {
          name: 'id',
          description: "The id of a vote",
          type: new GraphQLNonNull(GraphQLID)
        }
      },
      resolve(parent, args){
        Logger.log('info', 'Get Vote', {
          time: new Date().toLocaleString(),
          id: args.id,
        })
        return Vote.findById(args.id)
      }
    },
    getVotes: {
      type: new GraphQLList(VoteType),
      description: "Retrive all votes",
      resolve(parent, args){
        Logger.log('info', 'Get Votes', {
          time: new Date().toLocaleString(),
        })
        return Vote.find()
      }
    },
    getPostsFromOtherUser: {
      type: new GraphQLList(PostType),
      description: "Retrive all posts from other users",
      args: {
        userID: {
          name: 'userID',
          description: "The id of a user that post are not retrived",
          type: new GraphQLNonNull(GraphQLID)
        }
      },
      resolve(parent, args){
        Logger.log('info', 'Get Posts from other users', {
          time: new Date().toLocaleString(),
          id: args.userID,
        })
        return Post.find( { authorID: { $ne: args.userID }})
      }
    },
  }
});

module.exports = Query