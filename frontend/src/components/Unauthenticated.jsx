import {
    Route, Switch, Redirect
} from "react-router-dom";
import React from "react";
import {LoginPage} from "../pages/LoginPage";
import {NotFound} from "../pages/NotFound";
import {SignupPage} from "../pages/SignupPage";

export const Unauthenticated = () => {
    return (
        <Switch>
            <Route exact path="/login" component={LoginPage} />
            <Route exact path="/signup" component={SignupPage} />
            <Redirect from='/account' to='/login'/>
            <Route component={NotFound} />
        </Switch>
    )
}