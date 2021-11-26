import React, { Component } from 'react';
import ApolloClient from 'apollo-boost'
import { ApolloProvider } from 'react-apollo'
import { BrowserRouter as Router } from 'react-router-dom'
import Route from 'react-router-dom/Route'

import Facebook from './components/Facebook'
import PrivacyPolicy from './components/PrivacyPolicy'

let client = new ApolloClient({
  uri:'https://pretender-server.herokuapp.com/graphql'
})

class App extends Component {

  render() {
    return (
      
      <ApolloProvider client={client}>
        <Router>
          <div>
            <Route path="/" exact strict component={Facebook} />
            <Route path="/privacy" exact strict component={PrivacyPolicy} />
          </div>
        </Router>
      </ApolloProvider>

    );
  }
}

export default App;
