import { QueryResolver } from './query'
import { Context } from '../utils'

test('Should return a user if userId is Right', async () => {
  const me = await MockQuery.me(null, { id: '1234' }, null)

  expect(me.id).toBe('1234')
  expect(me.username).toBe('username')
  expect(me.birthdate).toBe(null)
})

test('Should return null or error if userId is Wrong', async () => {
  const me = await MockQuery.me(null, { id: 'wrong' }, null)

  expect(me).toBe(null)
})

const MockQuery: QueryResolver = {
  me: async (root, { id }, ctx: Context) => {
    if (id == '1234') {
      return stub
    }

    return null
  },
}

const stub = {
  id: '1234',
  username: 'username',
  firstName: 'Paulo',
  lastName: 'Keller',
  email: 'test@example.com',
  birthdate: null,
}
