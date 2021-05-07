import React  from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect
} from "react-router-dom";
import RoomSelect from './pages/RoomSelect'
import NotFound from "./pages/NotFound";

const App = () => {
    return (
        <Router>
            <Switch>
                <Route exact path="/" component={RoomSelect}/>
                <Route exact path="/404" component={NotFound} />
                <Redirect to="/404" />
            </Switch>
        </Router>
    )
}

export default App;