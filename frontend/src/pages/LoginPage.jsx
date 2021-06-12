import {Login} from "../components/Login";
import "./css/LoginPage.css"
import React from "react";
import {useHistory} from "react-router-dom";

export const LoginPage = () => {
    const history = useHistory()
    return (
        <>
            <header>
                <img className='login-logo' src='./Logo.png' alt='Logo'
                     onClick={() => history.push('/')} />
            </header>
            <section className='login-section'>
                <Login />
            </section>
        </>
    )
}