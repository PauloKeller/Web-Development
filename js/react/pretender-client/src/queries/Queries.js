import { gql } from 'apollo-boost'

let getPosts = gql`
{
  getPosts{
    id
    created_at
    title
    body
    isPretended
    votes{
      id
      isPretended
      created_at
      user{
        name
        id
        email
      }
    }
  }
}
`
let getPost = gql`
query($id: ID!){
  getPost(id: $id) {
    id
    created_at
    title
    body
    author{
      id
      name
      points
      posts{
        id
        title
      }
    }
    isPretended
  }
}
`
let getPostsFromOtherUser = gql`
query($userID: ID!){
  getPostsFromOtherUser(userID: $userID) {
    id
    created_at
    title
    body
    author{
      id
      name
      points
      posts{
        id
        title
      }
    }
    isPretended
    votes{
      id
      isPretended
      created_at
      user{
        name
        id
        email
      }
    }
  }
}
`

export { getPosts, getPost, getPostsFromOtherUser }