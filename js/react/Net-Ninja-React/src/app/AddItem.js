import React from 'react'
require('./css/additem.css')

export default class AddItem extends React.Component {
    constructor(props){
        super(props)
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    render(){
        return(
            <form id="add-todo" onSubmit={this.handleSubmit}>
                <input type="text" required ref="newItem" />
                <input type="submit" value="Hit me" />
            </form>
        )
    }

    // Custom functions
    handleSubmit(e){
        e.preventDefault()
        this.props.onAdd(this.refs.newItem.value)
    }
}