import {
    Route, Switch,
} from "react-router-dom";
import React from "react";
import {Account} from "../pages/Account";
import {NotFound} from "../pages/NotFound";

export const Authenticated = () => {
    return (
        <Switch>
            <Route path="/account" component={Account} />
            <Route component={NotFound} />
        </Switch>
    )
}