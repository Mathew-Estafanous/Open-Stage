import React  from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
} from "react-router-dom";
import { RoomSelect } from './pages/RoomSelect'
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
                {account? <Authenticated />: <Unauthenticated />}
            </Switch>
        </Router>
        </>
    )
}

export default App;