import {
    Route,
} from "react-router-dom";
import React from "react";
import {Account} from "../pages/Account";

export const Authenticated = () => {
    return (
        <>
            <Route path="/account" component={Account} />
        </>
    )
}