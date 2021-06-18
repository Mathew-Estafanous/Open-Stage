import React, {useEffect} from 'react';
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

    useEffect(() => {
        // Route to HTTPS if the current protocol is not secured. SSL routing is not possible
        // through heroku, so the app itself must route.
        if((process.env.REACT_APP_ENV === 'production') && (window.location.protocol !== "https:")) {
            window.location.protocol = "https";
        }
    }, [])

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