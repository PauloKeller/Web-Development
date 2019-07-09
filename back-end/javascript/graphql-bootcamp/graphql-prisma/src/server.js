import { GraphQLServer, PubSub } from 'graphql-yoga'

import db from './db'
import prisma from './prisma'
import { resolvers, fragmentReplacements }  from './resolvers/index'

const pubsub = new PubSub()

const server = new GraphQLServer({
	typeDefs: './src/schema.graphql',
	resolvers,
	fragmentReplacements,
	context(request) {
		return {
			db,
			pubsub,
			prisma,
			request
		}
	},
})

export { server as default }