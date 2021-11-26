import { server, options } from './server'

server.start(options, ({ port }) =>
  console.log(`Server is running on http://localhost:${port}`)
)
