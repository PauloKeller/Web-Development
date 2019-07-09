const Query = {
  me() {
    return {
      id: "12313",
      name: "Paulo Keller",
      email: "test@test.com",
      age: 28
    }
  },
  post() {
    return {
      id: "1233",
      title: "quihsuihwuiahd",
      body: "adwodjoapwd",
      published: true,
    }
  },
  users(parent, args, { db }, info) {
    if (!args.query) {
      return db.users
    }

    return db.users.filter((user) => {
      return user.name.toLowerCase().includes(args.query.toLowerCase())
    })
  },
  posts(parent, args, { db }, info) {
    if (!args.query) {
      return db.posts
    }

    return db.posts.filter((post) => {
      const isTitleMatch = post.title.toLowerCase().includes(args.query.toLowerCase())
      const isBodyMatch = post.body.toLowerCase().includes(args.query.toLowerCase())
      
      if (isBodyMatch || isTitleMatch) {
        return post;
      }
    })
  },
  comments(parent, args, { db }, info) {
    return db.comments
  }
}

export { Query as default }