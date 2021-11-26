import { GraphQLServer } from 'graphql-yoga'
import { UserRepository } from './repositories'
import { resolvers } from './resolvers'

const options = {
  port: 5000,
  endpoint: '/graphql',
  playground: '/playground',
  getEndpoint: true,
}

const repository = new UserRepository()

const server = new GraphQLServer({
  typeDefs: './src/schema.graphql',
  resolvers,
  context: (request) => ({
    ...request,
    userRepository: repository,
  }),
})

export { server, options }
