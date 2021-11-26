import { server } from '../src/server'

const globalAny: any = global

export = async () => {
  globalAny.httpServer = await server.start({ port: 5000 })
}
