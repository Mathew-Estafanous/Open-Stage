import React  from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect
} from "react-router-dom";
import { RoomSelect } from './pages/RoomSelect'
import { NotFound } from "./pages/NotFound";
import { Room } from "./pages/Room";
import {useAuth} from "./context/AuthContext";
import {Unauthenticated} from "./components/Unauthenticated";
import {Authenticated} from "./components/Authenticated";

const App = () => {
    const {account} = useAuth()

    return (
        <>
        <Router>
            <Switch>
                <Route exact path="/" component={RoomSelect} />
                <Route path="/room/:code" component={Room} />
                <Route exact path="/404" component={NotFound} />
                {account? <Authenticated />: <Unauthenticated />}
                <Redirect to="/404" />
            </Switch>
        </Router>
        </>
    )
}

export default App;