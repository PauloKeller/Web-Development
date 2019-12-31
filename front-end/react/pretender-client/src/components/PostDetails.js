import React, { Component } from 'react'
import { graphql, compose } from 'react-apollo'
import { getPost, getPostsFromOtherUser } from '../queries/Queries'
import { createVote, updateUsersPointsWhenRightHint, updateUsersPointsWhenWrongHint } from '../mutations/Mutations'

class PostDetails extends Component {

    makeQuery(postID) {
        this.props.client.query({
            query: getPost,
            variables: {
                postID: postID
            },
        }).then(({ data }) => {
            console.log(data)
        })
    }

    votePost(vote) {
        this.props.createVote({
            variables: { 
                isPretended: vote,
                userID: this.props.user.id,
                postID: this.props.postID
            },
            refetchQueries: [{ 
                query: getPostsFromOtherUser,
                variables: { userID: this.props.user.id }  
            }],
            
        }).then(({ data }) => {
            alert("Thanks for voting!")
        }).catch((error) => {
            console.log("error", error)
        });

        if (vote === this.props.data.getPost.isPretended) {
            this.props.updateUsersPointsWhenRightHint({
                variables: { 
                    postID: this.props.postID
                }
            }).then(({ data }) => {
                window.location.reload()
            }).catch((error) => {
                console.log("error", error)
            });
        } else {
            this.props.updateUsersPointsWhenWrongHint({
                variables: { 
                    postID: this.props.postID
                }
            }).then(({ data }) => {
                window.location.reload()
            }).catch((error) => {
                console.log("error", error)
            });
        }
    }

    displayPostDetails() {
        let { post } = this.props.data
        if (this.props.data.getPost != null) {
            post = this.props.data.getPost
        }
        if (post) {
            return(
                <div>
                    <h1>{ post.title }</h1>
                    <p>{ post.body }</p>
                    <p>{ post.author.name }</p>
                    <p>All posts by this atuhor:</p>
                    <ul className="other-posts">
                        { post.author.posts.map(item => {
                            return <li key={ item.id }>{ item.title }</li>
                        })}
                    </ul>
                    <div id="buttons-container">
                        <button id="lier-button" onClick={ 
                            (e) => { 
                                this.votePost(false)
                            }
                        }>LIE</button>

                        <button id="truth-button" onClick={
                            (e) => { 
                                this.votePost(true)
                            }
                        }>TRUTH</button>
                    </div>
                </div>   
            )
        } else {
            return (
                <div>No post selected...</div>
            )
        }
    }


    render() {
        if(this.props.data.loading) {
            return (<div>Loading...</div>)
        }
        return(
            <div id="post-details">
                {this.displayPostDetails()}
            </div>
        )
    }
}

export default compose(
    graphql(getPost, {
        options: (props) => {
            return {
                variables: {
                    id: props.postID
                }
            }
        }
    }),
    graphql(createVote, { name: 'createVote' }),
    graphql(updateUsersPointsWhenRightHint, { name: 'updateUsersPointsWhenRightHint' }),
    graphql(updateUsersPointsWhenWrongHint, { name: 'updateUsersPointsWhenWrongHint' }),
)(PostDetails);