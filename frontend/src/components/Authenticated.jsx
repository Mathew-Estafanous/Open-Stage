import {
    Route, Switch,
} from "react-router-dom";
import React from "react";
import {AccountPage} from "../pages/AccountPage";
import {NotFound} from "../pages/NotFound";

export const Authenticated = () => {
    return (
        <Switch>
            <Route path="/account" component={AccountPage} />
            <Route component={NotFound} />
        </Switch>
    )
}