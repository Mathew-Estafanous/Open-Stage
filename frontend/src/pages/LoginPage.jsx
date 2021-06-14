import {Login} from "../components/Login";
import "./css/AuthPage.css"
import React from "react";
import {useHistory} from "react-router-dom";

export const LoginPage = () => {
    const history = useHistory()
    return (
        <>
            <header>
                <img className='auth-logo' src='./Logo.png' alt='Logo'
                     onClick={() => history.push('/')} />
            </header>
            <section className='auth-section'>
                <Login />
            </section>
        </>
    )
}