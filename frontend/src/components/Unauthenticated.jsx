import {
    Route, Switch,
} from "react-router-dom";
import React from "react";
import {Signup} from "./Signup";
import {LoginPage} from "../pages/LoginPage";
import {NotFound} from "../pages/NotFound";

export const Unauthenticated = () => {
    return (
        <Switch>
            <Route exact path="/login" component={LoginPage} />
            <Route exact path="/signup" component={Signup} />
            <Route component={NotFound} />
        </Switch>
    )
}