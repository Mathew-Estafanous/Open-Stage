import {
    Switch,
    Route,
    Redirect, BrowserRouter as Router
} from "react-router-dom";
import React from "react";
import {Login} from "./Login";
import {Signup} from "../pages/Signup";
import {LoginPage} from "../pages/LoginPage";

export const Unauthenticated = () => {
    return (
        <>
            <Route exact path="/login" component={LoginPage} />
            <Route exact path="/signup" component={Signup} />
            <Redirect to="/404" />
        </>
    )
}