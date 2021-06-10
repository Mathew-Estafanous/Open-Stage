import {
    Switch,
    Route,
    Redirect, BrowserRouter as Router
} from "react-router-dom";
import React from "react";
import {Login} from "../pages/Login";
import {Signup} from "../pages/Signup";

export const Unauthenticated = () => {
    return (
        <>
            <Route path="/login" component={Login} />
            <Route path="/signup" component={Signup} />
        </>
    )
}