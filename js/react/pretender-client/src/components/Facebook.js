import React, { Component } from 'react'
import FacebookLogin from 'react-facebook-login'
import { graphql, compose } from 'react-apollo'

// Graphql
import { login, createUser } from '../mutations/Mutations'

// Components
import PostsList from './PostsList'


class Facebook extends Component {
    state = {
        isLoggedIn: false,
        userID: '',
        name: '',
        email: ''
    }

    responseFacebook = response => {
        this.setState({
            isLoggedIn: true,
            userID: response.userID,
            name: response.name,
            email: response.email,
            makeLogin: true,
        })
    }

    login() {
        this.props.login({
            variables: { email: this.state.email }
        }).then(({ data }) => {
            let user = {
                id: data.login.id,
                name: data.login.name,
                points: data.login.points,
                email: data.login.email,
            }
            this.setState({
                user: user,
                makeLogin: false
            })
        }).catch((error) => {
            this.createUser()
        });
    }

    createUser() {
        this.props.createUser({
          variables: { 
            email: this.state.email,
            name: this.state.name 
          }
        }).then(({ data }) => {
          let user = {
            id: data.createUser.id,
            name: data.createUser.name,
            points: data.createUser.points,
            email: data.createUser.email,
          }
          this.setState({
            user: user
          })
            return(
                <div>
                    <PostsList user={this.state.user} />
                </div> 
            )
        }).catch((error) => {
          return (
              <div>
                  <h3>Cannot create the user please reload the page!</h3>
              </div>
          )
        });
      }

    render() {
        let fbContent

        if (this.state.isLoggedIn) {
            if (this.state.makeLogin) {
                this.login()
            }
            if (this.state.user) {
                fbContent = (
                    <div>
                       <PostsList user={this.state.user} />
                    </div> 
                )
            } 
        } else {
            fbContent = (
                <div id="facebook-button">
                    <FacebookLogin
                        appId='1722380314549985'
                        autoLoad={true}
                        fields="name,email"
                        onClick={this.componentClicked}
                        callback={this.responseFacebook}
                    />
                </div>
            )
        }

        return (
            <div>
                {fbContent}
            </div>
        )
    }
}

export default compose(
    graphql(login, { name: 'login' }),
    graphql(createUser, {name: 'createUser' })
)(Facebook)