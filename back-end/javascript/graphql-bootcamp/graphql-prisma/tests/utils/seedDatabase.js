import bcrypt from 'bcryptjs'
import jwt from 'jsonwebtoken'
import prisma from '../../src/prisma'

// Create user one
const userOne = {
  input: {
    name: "Seed User",
    email: "seed@seed.com",
    password: bcrypt.hashSync("seed123456")
  },
  user: undefined,
  jwt: undefined
}

// Create user two
const userTwo = {
  input: {
    name: "User Two",
    email: "user@two.com",
    password: bcrypt.hashSync("seedcomment123")
  },
  user: undefined,
  jwt: undefined
}

// Create post one
const postOne = {
  input: {
    title: "Test Post",
    body: "Test body",
    published: true
  },
  post: undefined
}

const postTwo = {
  input: {
    title: "Draft Post",
    body: "Draft body",
    published: false
  },
  post: undefined
}

// Create comment one
const commentOne = {
  input: {
    text: "Great Post Thanks for sharing"
  },
  comment: undefined
}

// Create comment two
const commentTwo = {
  input: {
    text: "Thanks!"
  },
  comment: undefined
}

const seedDatabase =  async () => {
  // Delete test data
  await prisma.mutation.deleteManyComments()
  await prisma.mutation.deleteManyPosts()
  await prisma.mutation.deleteManyUsers()

  // Create user one
  userOne.user = await prisma.mutation.createUser({
    data: userOne.input
  })
  userOne.jwt = jwt.sign({ userId: userOne.user.id}, process.env.JWT_SECRET)

  // Create user two
  userTwo.user = await prisma.mutation.createUser({
    data: userTwo.input
  })
  userTwo.jwt = jwt.sign({ userId: userTwo.user.id}, process.env.JWT_SECRET)

  // Create post one
  postOne.post = await prisma.mutation.createPost({
    data: {
      ...postOne.input,
      author: {
        connect: {
          id: userOne.user.id
        }
      }
    }
  })

  // Create post two
  postTwo.post = await prisma.mutation.createPost({
    data: {
      ...postTwo.input,
      author: {
        connect: {
          id: userOne.user.id
        }
      }
    }
  })

  // Create comment one
  commentOne.comment = await prisma.mutation.createComment({
    data: {
      ...commentOne.input,
      author: {
        connect: {
          id: userTwo.user.id,
        }
      },
      post: {
        connect: {
          id: postOne.post.id,
        }
      }
    }
  })

  // Create comment two
  commentTwo.comment = await prisma.mutation.createComment({
    data: {
      ...commentTwo.input,
      author: {
        connect: {
          id: userOne.user.id,
        }
      },
      post: {
        connect: {
          id: postOne.post.id,
        }
      }
    }
  })
}

export { seedDatabase as default, userOne, userTwo, postOne, postTwo, commentOne, commentTwo }

