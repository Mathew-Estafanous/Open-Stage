import {Login} from "../components/Login";
import "./css/LoginPage.css"
import React, {useState} from "react";
import {Error} from "../components/Error";

export const LoginPage = () => {
    const [error, setError] = useState(false);
    return (
        <>
            <header>
                <img className='logo' src='./Logo.png' alt='Logo'/>
            </header>
            <section className='login-section'>
                {error?
                    <Error msg={"Invalid username or password!"}/>: null
                }
                <Login />
            </section>
        </>
    )
}