import "./css/Signup.css"
import {useState} from "react";
import {Error} from "./Error";
import {Link} from "react-router-dom";

export const Signup = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');

    const [error, setError] = useState(false);

    const submitSignup = (e) => {
        e.preventDefault()
        console.log("Submitted signup form");
    }

    return (
        <form className='signup-form' onSubmit={submitSignup} >
            {error?
                <Error msg='There was an error while signing up.'/>: null
            }

            <div className='signup-form-wrapper'>
                <div className='signup-form-header'>
                    <h2>Create An Account</h2>
                </div>
                <div className='signup-form-field'>
                    <label htmlFor="username">Username</label>
                    <input type="text"
                            value={username}
                            onChange={e => setUsername(e.target.value)}
                            maxLength={25}
                            required />
                </div>
                <div className='signup-form-field'>
                    <label htmlFor="name">Name</label>
                    <input type="text"
                           value={name}
                           onChange={e => setName(e.target.value)}
                           required />
                </div>
                <div className='signup-form-field'>
                    <label htmlFor="email">Email</label>
                    <input type="text"
                           value={email}
                           onChange={e => setEmail(e.target.value)}
                           required />
                </div>
                <div className='signup-form-field'>
                    <label htmlFor="password">Password</label>
                    <input type="password"
                           value={password}
                           onChange={e => setPassword(e.target.value)}
                           minLength={8}
                           required />
                </div>
                <button className='signup-form-btn' type='submit'>Create Account</button>
                <Link to="/signup">Login to your account.</Link>
            </div>
        </form>
    )
}