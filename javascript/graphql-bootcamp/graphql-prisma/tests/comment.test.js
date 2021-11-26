import 'cross-fetch/polyfill'
import seedDatabase, { userOne, postOne, commentOne, commentTwo } from './utils/seedDatabase'
import { deleteComment, subscribeToComments } from './utils/operations'
import getClient from './utils/getClient'
import prisma from '../src/prisma'

const client = getClient()

jest.setTimeout(15000);
beforeEach(seedDatabase)

test('Should delete own comment', async () => {
  const client = getClient(userOne.jwt)
  const variables = {
    id: commentTwo.comment.id
  }

  await client.mutate({ mutation: deleteComment, variables }) 
  const exists = await prisma.exists.Comment({ id: commentTwo.comment.id })
  
  expect(exists).toBe(false) 
})

test('Should not delete other users comment', async () => {
  const client = getClient(userOne.jwt)
  const variables = {
    id: commentOne.comment.id
  }

  await expect(
    client.mutate({ mutation: deleteComment, variables })
  ).rejects.toThrow()
})

test('Should subscribe to comments for a post', async (done) => {
  const variables = {
    postId: postOne.post.id
  }
  client.subscribe({ query: subscribeToComments, variables }).subscribe({
    next(response) {
      expect(response.data.comment.mutation).toBe('DELETED')
      done()
    }
  })

  await prisma.mutation.deleteComment({ where: { id: commentOne.comment.id }})
})

