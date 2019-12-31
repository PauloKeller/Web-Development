import React, { Component } from 'react';
import { graphql, compose } from 'react-apollo'
import { createPost } from '../mutations/Mutations'
import { getPosts } from '../queries/Queries';

class AddPost  extends Component {
    constructor(props) {
        super(props)
        this.state = {
            body: "",
            title: "",
            authorId: "",
            isPretended: "",
        }
    }

    submitForm(e) {
        e.preventDefault()  
        this.props.createPost({
            variables: {
                body: this.state.body,
                title: this.state.title,
                authorID: this.props.user.id,
                isPretended: this.state.isPretended
            },
            refetchQueries: [{ query: getPosts }]
        }).then(({ data }) => {
            alert('Post added')
        }).catch((error) => {
            console.log('there was an error sending the query', error);
        });
        
    }

    render() {
      return (
        <form id="add-post" onSubmit={ this.submitForm.bind(this) }>
            <div className="field">
                <label>Post Title:</label>
                <input type="text" onChange={ (e) => this.setState({ title: e.target.value })} />
            </div>
            <div className="field">
                <label>Body:</label>
                <input type="text" onChange={ (e) => this.setState({ body: e.target.value })} />
            </div>
            <div className="field">
                <label>Pretended:</label>
                <select onChange={ (e) => this.setState({ isPretended: e.target.value })} >
                    <option>Is this a pretended post?</option>
                    <option value={true}>Truth</option>
                    <option value={false}>Lie</option>
                </select>
            </div>
            <button>+</button>
        </form>
      );
    }
}
  
export default compose(
    graphql(createPost, { name: 'createPost' })
)(AddPost);