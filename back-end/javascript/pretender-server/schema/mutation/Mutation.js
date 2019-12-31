let {
  GraphQLObjectType,
  GraphQLString,
  GraphQLBoolean,
  GraphQLNonNull,
  GraphQLID
} = require('graphql')
let {UserType, PostType, VoteType} = require('../type/Types')

let User = require('../../models/User')
let Post = require('../../models/Post')
let Vote = require('../../models/Vote')

let Logger = require('../../Logger')

let Mutation = new GraphQLObjectType({
  name: 'RootMutationType',
  description: `In REST, any request might end up causing some side-effects on the server, but by convention it's suggested that one doesn't use GET requests to modify data. GraphQL is similar - technically any query could be implemented to cause a data write. However, it's useful to establish a convention that any operations that cause writes should be sent explicitly via a mutation.

  Just like in queries, if the mutation field returns an object type, you can ask for nested fields. This can be useful for fetching the new state of an object after an update. Let's look at a simple example mutation: `,
  fields: {
    createUser: {
      type: UserType,
      description: "Create's a user",
      args: {
        name: {
          name: 'name',
          description: "A user's name",
          type: new GraphQLNonNull(GraphQLString)
        },
        email: {
          name: 'email',
          description: "A user's email",
          type: new GraphQLNonNull(GraphQLString)
        },
      },
      resolve(parent, args){
        let user = new User({
          name: args.name,
          email: args.email,
          points: 0.0,
          created_at: new Date(),
        })

        Logger.log('info', 'Create User', {
          time: new Date().toLocaleString(),
          user: user,
        })
      
        return user.save()
      }
    },
    createPost: {
      type: PostType,
      description: "Create's a post",
      args: {
        title: {
          name: 'title',
          description: "A post's title",
          type: new GraphQLNonNull(GraphQLString)
        },
        body: {
          name: 'body',
          description: "A post's body",
          type: new GraphQLNonNull(GraphQLString)
        },
        isPretended: {
          name: 'isPretended',
          description: "Set if is a lier or truth post",
          type: new GraphQLNonNull(GraphQLBoolean)
        },
        authorID: {
          name: 'authorID',
          description: "The id from user that create's the post",
          type: new GraphQLNonNull(GraphQLID),
        },
      },
      resolve(parent, args){
        let post = new Post({
          title: args.title,
          body: args.body,
          created_at: new Date(),
          authorID: args.authorID,
          isPretended: args.isPretended  
        })

        Logger.log('info', 'Create Post', {
          time: new Date().toLocaleString(),
          post: post,
        })
      
        return post.save()
      }
    },
    createVote: {
      type: VoteType,
      description: "Create a vote from a post",
      args: {
        isPretended: {
          name: 'isPretended',
          description: "Vote if is a lier post or truth",
          type: new GraphQLNonNull(GraphQLBoolean)
        },
        userID: {
          name: 'userID',
          description: "Id from the user that are voting in the post",
          type: new GraphQLNonNull(GraphQLID)
        },
        postID: {
          name: 'postID',
          description: "A Post's id",
          type: new GraphQLNonNull(GraphQLID)
        },
      },
      resolve(parent, args){
        let vote = new Vote({
          userID: args.userID,
          postID: args.postID,
          isPretended: args.isPretended,
          created_at: new Date(),
        })
        
        Logger.log('info', 'Create Vote', {
          time: new Date().toLocaleString(),
          vote: vote,
        })

        return vote.save()
      }
    },
    login: {
      type: UserType,
      description: "Performe a login",
      args: {
        email: {
          name: 'email',
          description: "User's email to login",
          type: new GraphQLNonNull(GraphQLString)
        },
      },
      resolve(parent, args ) {
        Logger.log('info', 'Create Vote', {
          time: new Date().toLocaleString(),
          email: args.email,
        })

        return User.findOne({ email: args.email })
      }
    },
    updateUsersPointsWhenRightHint: {
      type: PostType,
      description: "Update the points from user's when a right hint are done",
      args: {
        postID: {
          name: 'postID',
          description: "A post's id that hint was done",
          type: new GraphQLNonNull(GraphQLID)
        },
      },
      resolve(parent, args){
        Post.findById(args.postID).lean().exec(function (err, post) {
          Vote.find({ postID: args.postID }).lean().exec(function (err, votes) {
        
            var rightVotes = []
            votes.map( vote => {
              if (vote.isPretended == post.isPretended) {
                rightVotes.push(vote)
              } 
            })

            let author
            User.findById(post.authorID).lean().exec(function (err, user) {
              author = user
            })

            var usersWithRightVotes = []
            var millisecondsToWait = 5000

            rightVotes.forEach(rightVote => {              
              User.findById(rightVote.userID).lean().exec(function (err, user) {
                usersWithRightVotes.push(user)
              })                          
            })
            setTimeout(function() {
              let rightVotesCount = rightVotes.length
              var points = 1

              if(rightVotesCount >= 3) {

                for ( i = 1; i < rightVotesCount-2; i++) {
                  points = 2 * points 
                }

                let updateAuthor = { points: author.points - points }
                User.findByIdAndUpdate(post.authorID, updateAuthor, { new: true}, (err, numAffected) => {
                  console.log(err)
                })

                console.log(points)
                points = points/rightVotesCount
                
                usersWithRightVotes.forEach(user => {
                  update = { points: points + user.points }, 
                  User.findByIdAndUpdate(user._id, update, {new: true}, (err, numAffected) => {
                    console.log(err)
                  })
                })

                Logger.log('info', "Update the points from user's", {
                  time: new Date().toLocaleString(),
                  postID: args.postID,
                  points: points,
                  usersWithRightVotes: usersWithRightVotes,
                  authorID: post.authorID
                })
              }
            }, millisecondsToWait);
          })
        })

        return Post.findById(args.postID)  
      }
    },
    updateUsersPointsWhenWrongHint: {
      type: PostType,
      description: "Update the points from user's when a wrong hint are done",
      args: {
        postID: {
          name: 'postID',
          description: "A post's id that hint was done",
          type: new GraphQLNonNull(GraphQLID)
        },
      },
      resolve(parent, args){
        Post.findById(args.postID).lean().exec(function (err, post) {
          Vote.find({ postID: args.postID }).lean().exec(function (err, votes) {
        
            var rightVotes = []
            votes.map( vote => {
              if (vote.isPretended != post.isPretended) {
                rightVotes.push(vote)
              } 
            })

            let author
            User.findById(post.authorID).lean().exec(function (err, user) {
              author = user
            })

            var usersWithRightVotes = []
            var millisecondsToWait = 5000

            rightVotes.forEach(rightVote => {              
              User.findById(rightVote.userID).lean().exec(function (err, user) {
                usersWithRightVotes.push(user)
              })                          
            })
            setTimeout(function() {
              let rightVotesCount = rightVotes.length
              var points = 1

              console.log(rightVotesCount)

              if(rightVotesCount >= 3) {

                console.log("Update")
                
                for ( i = 1; i < rightVotesCount-2; i++) {
                  points = 2 * points 
                }
                
                let updateAuthor = { points: author.points + points }
                User.findByIdAndUpdate(post.authorID, updateAuthor, { new: true}, (err, numAffected) => {
                  console.log(err)
                })
                
                console.log(points)
                points = points/rightVotesCount
                
                usersWithRightVotes.forEach(user => {
                  update = { points: user.points - points }, 
                  User.findByIdAndUpdate(user._id, update, {new: true}, (err, numAffected) => {
                    console.log(err)
                  })
                })

                Logger.log('info', "Update the points from user's", {
                  time: new Date().toLocaleString(),
                  postID: args.postID,
                  points: points,
                  usersWithRightVotes: usersWithRightVotes,
                  authorID: post.authorID
                })
              }
            }, millisecondsToWait);
          })
        })
        return Post.findById(args.postID)  
      }
    }
  }
});

module.exports = Mutation