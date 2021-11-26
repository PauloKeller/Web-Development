import { UsersService } from './usersService'

describe('across entire suite', function () {
  test('Expect to return a user if id is Right', async () => {
    const repo = new UserMockRepository()
    const user = repo.fetchUser('1234')

    expect(user.balance).toBe(1000)
    expect(user.username).toBe('username')
  })

  test('Expect to return null if id is Wrong', async () => {
    const repo = new UserMockRepository()
    const user = repo.fetchUser('wrong')

    expect(user).toBe(null)
  })
})

const stub = {
  id: '1234',
  username: 'username',
  balance: 1000,
}

class UserMockRepository implements UsersService {
  constructor() {}

  fetchUser(userId: string) {
    if (userId == '1234') {
      return stub
    }
    return null
  }
}
