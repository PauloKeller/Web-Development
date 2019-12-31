let { GraphQLSchema } = require ('graphql')

let Query = require('./query/Query')
let Mutation = require('./mutation/Mutation')

let Schema = new GraphQLSchema({
    query: Query,
    mutation: Mutation
})
  
module.exports =  Schema