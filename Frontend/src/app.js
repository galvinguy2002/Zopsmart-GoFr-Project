import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import AddStock from './components/AddStock';
import ViewStocks from './components/ViewStocks';
import UpdateStock from './components/UpdateStock';
import DeleteStock from './components/DeleteStock';

const App = () => {
  return (
    <Router>
      <div className="App">
        <Switch>
          <Route exact path="/" component={ViewStocks} />
          <Route exact path="/add" component={AddStock} />
          <Route exact path="/update" component={UpdateStock} />
          <Route exact path="/delete" component={DeleteStock} />
        </Switch>
      </div>
    </Router>
  );
};

export default App;
