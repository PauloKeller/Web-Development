import { Context } from '../utils'
import { IResolverObject } from 'graphql-tools'

type ResolverFn = (parent: any, args: any, ctx: Context) => any

export interface QueryResolver extends IResolverObject {
  me: ResolverFn
}

export const Query: QueryResolver = {
  me: async (root, { id }, ctx: Context) => {
    const result = await ctx.userRepository.fetchUser(id)
    console.log(result)
    return result
  },
}
