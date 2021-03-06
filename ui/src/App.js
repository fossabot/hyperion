import React from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

import Home from "./pages/static/Home";
import LightList from "./pages/LightList";
import CueList from "./pages/CueList";
import Nav from "./components/Nav";
import WS from "./components/WS";
import { Container } from "semantic-ui-react";
import ReduxToastr from "react-redux-toastr";

const App = () => (
  <Router>
    <div>
      <ReduxToastr
        timeOut={4000}
        newestOnTop={false}
        preventDuplicates
        position="top-left"
        transitionIn="fadeIn"
        transitionOut="fadeOut"
        progressBar
      />
      <Nav />
      <Container
        fluid
        style={{ marginTop: "7em", width: "95%", minHeight: "100vh" }}
      >
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/lights" component={LightList} />
          <Route path="/cues" component={CueList} />
        </Switch>
        <WS />
      </Container>
    </div>
  </Router>
);
export default App;
