import React, { Component } from 'react'

export default class Header extends Component {
    render() {
        return(
            <div id="header">
                <h1>User: {this.props.user.name}</h1>
                <h2>Points: {this.props.user.points}</h2>
            </div>
        )
    }
}