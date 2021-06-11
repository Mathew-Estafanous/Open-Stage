import {Login} from "../components/Login";
import "./css/LoginPage.css"

export const LoginPage = () => {
    return (
        <>
            <header>
                <img className='logo' src='./Logo.png' alt='Logo'/>
            </header>
            <section className='login-section'>
                <Login />
            </section>
        </>
    )
}