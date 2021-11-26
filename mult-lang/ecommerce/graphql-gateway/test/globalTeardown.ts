const globalAny: any = global

export = async () => {
  await globalAny.httpServer.close()
}
