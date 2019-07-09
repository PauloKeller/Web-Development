import React from 'react'
import ReactDOM from 'react-dom'
import { Router, Route, browserHistory, Link} from 'react-router'
require('./css/index.css');

// Modules requires
import TodoItem from './TodoItem'
import AddItem from './AddItem'
import About from './About'

class App extends React.Component {
    render() {
        return (
            <Router history={browserHistory}>
                <Route path={'/'} component={TodoComponent}></Route>
                <Route path={'/about'} component={About}></Route>
            </Router>
        )
    }
}


class TodoComponent extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            todos:['wash up', 'eat some cheese', 'take a nap', 'buy flowers'],
            age: 30
        }
        this.onDelete = this.onDelete.bind(this);
        this.onAdd = this.onAdd.bind(this);
    }

    render(){
        var todos = this.state.todos
        todos = todos.map(function(item, index){
            return (
                <TodoItem item={item} key={index} onDelete={this.onDelete} />
            )
        }.bind(this))

        return(
            <div id="todo-list">
                <Link to={'/about'}>About</Link>
                <p>The busiest people have the most leisure...</p>
                <ul>{todos}</ul>
                <AddItem onAdd={this.onAdd} />
            </div>   
        )
    } // render

    // Custom functions
    onDelete(item){
        var updatedTodos = this.state.todos.filter(function(val, index){
            return item !== val
        })
        this.setState({
            todos: updatedTodos
        })
    }

    onAdd(item){
        var updatedTodos = this.state.todos;
        updatedTodos.push(item)
        this.setState({
            todos: updatedTodos
        })
    }

    // Life-cycle functions
    componentWillMount(){
        console.log('componentWillMount')
    }

    componentDidMount(){
        console.log('componentDidMount')
        // any grabbing of external data
    }

    componentWillUpdate(){
        console.log('componentWillUpdate')
    }

}

export default TodoComponent



// put component into html page
ReactDOM.render(<App/>, document.getElementById('todo-wrapper'))