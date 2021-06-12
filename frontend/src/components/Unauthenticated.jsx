import {
    Route,
} from "react-router-dom";
import React from "react";
import {Signup} from "../pages/Signup";
import {LoginPage} from "../pages/LoginPage";
import {NotFound} from "../pages/NotFound";

export const Unauthenticated = () => {
    return (
        <>
            <Route exact path="/login" component={LoginPage} />
            <Route exact path="/signup" component={Signup} />
            <Route component={NotFound} />
        </>
    )
}