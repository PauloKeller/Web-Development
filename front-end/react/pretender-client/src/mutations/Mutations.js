import { gql } from 'apollo-boost'

let createUser = gql`
  mutation($name: String!, $email: String!) {
    createUser(name: $name, email: $email){
      name
      id
      points
      posts {
        title
        body
      }
    }
  }
`


let createPost = gql`
  mutation($title: String!, $body: String!, $isPretended: Boolean!, $authorID: ID!) {
    createPost(title: $title, body: $body, isPretended: $isPretended, authorID: $authorID){
      id
      title
      body
      isPretended
      author {
        name
        id
        points
      }
    }
  }
`

let login = gql`
  mutation($email: String!) {
    login(email: $email){
      name
      id
      email
      points
    }
  }
`

let createVote = gql`
  mutation($isPretended: Boolean!, $userID: ID!, $postID: ID!) {
    createVote(isPretended: $isPretended, userID: $userID, postID: $postID){
      id
      created_at
    }
  }
`

let updateUsersPointsWhenRightHint = gql`
  mutation($postID: ID!) {
    updateUsersPointsWhenRightHint(postID: $postID) {
      id
    }
  }
`


let updateUsersPointsWhenWrongHint = gql`
  mutation($postID: ID!) {
    updateUsersPointsWhenWrongHint(postID: $postID) {
      id
    }
  }
`

export { 
  createUser,
  createPost,
  login,
  createVote,
  updateUsersPointsWhenRightHint,
  updateUsersPointsWhenWrongHint
}