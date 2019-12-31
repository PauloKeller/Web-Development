import React, { Component } from 'react'

// Graphql
import { graphql, compose } from 'react-apollo'
import { getPostsFromOtherUser } from '../queries/Queries'

// Components
import Header from './Header'
import AddPost from './AddPost'
import PostDetails from './PostDetails'

class PostsList  extends Component {
  state = {
    selected: ""
  }

  filterPosts(posts) {
    let userID = this.props.user.id
    let freeVotesPostsArray = []
    posts.map( post => {
      var freeVotePost = true
      post.votes.map( vote => {
        if (vote.user.id === userID) {
          freeVotePost = false
        }
      }) 
      if (freeVotePost){
        return freeVotesPostsArray.push(post)
      }
    })
    return freeVotesPostsArray
  }


  getPosts() {
    var data = this.props
    if (data.loading){
      return (<div>Loading posts...</div>)
    } else {
      var posts = data.data.getPostsFromOtherUser
      if (posts != null) {
        posts = this.filterPosts(posts)
        return posts.map( post => {
          return (
            <div id="post-container" key={ post.id } onClick={ (e) => { this.setState({ selected: post.id })} }>
              <h3>{ post.title }</h3>
              <p>{ post.body }</p>
            </div>
          )
        })
      }
    }
  }

  render() {
    return(
      <div>
        <div id="header-container">
          <Header user={this.props.user}/>
        </div>
        <div id="posts-list">
          {this.getPosts()}
        </div>
        <div id="post-details-container">
          <PostDetails postID={ this.state.selected } user={ this.props.user }/>
        </div>
        <div id="add-post-container">
          <AddPost user={this.props.user}/>
        </div>
      </div>
    )
  }
}

export default compose(
  graphql(getPostsFromOtherUser, { 
    options: (props) => {
      return {
        variables: {
          userID: props.user.id
        }
      }
    } 
  }),
)(PostsList)