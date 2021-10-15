import React from 'react';
import './App.css';
import {Router} from "react-router-dom";
import Layout from "./Layouts/Layout";
import './style.css'
import * as hist from 'history'

export const history = hist.createHashHistory();

function App() {
  return (
    <Router history={history}>
      <Layout/>
  </Router>
  );
}

export default App;
