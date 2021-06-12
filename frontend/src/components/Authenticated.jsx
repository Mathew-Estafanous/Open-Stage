import {
    Route,
} from "react-router-dom";
import React from "react";
import {Account} from "../pages/Account";
import {NotFound} from "../pages/NotFound";

export const Authenticated = () => {
    return (
        <>
            <Route path="/account" component={Account} />
            <Route component={NotFound} />
        </>
    )
}