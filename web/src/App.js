import React from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import './App.css';
import { SignUp } from "./components/SignUp";
import { Login } from "./components/Login";

function App() {
  return (
    <Router>
      <Switch>
          <Route exact path="/">
            <Login />
          </Route>
          <Route path="/sign-up">
            <SignUp />
          </Route>
        </Switch>
    </Router>
  );
}
 
 export default App;