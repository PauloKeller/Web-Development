import { UsersService } from './usersService'
import Client from './client'

export default class UserRepository implements UsersService {
  client: Client

  constructor() {
    this.client = new Client('http://users-service:8000')
  }

  async fetchUser(userId: string) {
    const response = await this.client.instance.get(`/users/${userId}`)

    if (response.data.error) {
      return response.data.error
    }

    return response.data
  }
}
