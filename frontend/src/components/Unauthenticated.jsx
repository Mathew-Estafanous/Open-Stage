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
            <Route path="/login" component={LoginPage} />
            <Route path="/signup" component={Signup} />
        </>
    )
}