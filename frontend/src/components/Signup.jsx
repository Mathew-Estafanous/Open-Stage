import {useState} from "react";
import {Error} from "./Error";
import {Link} from "react-router-dom";
import "./css/AuthForm.css";
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
        <form className='auth-form' onSubmit={submitSignup} >
            {error?
                <Error msg='There was an error while signing up.'/>: null
            }

            <div className='auth-form-wrapper'>
                <div className='auth-form-header'>
                    <h2>Create An Account</h2>
                </div>
                <div className='auth-form-field'>
                    <label htmlFor="username">Username</label>
                    <input type="text"
                            value={username}
                            onChange={e => setUsername(e.target.value)}
                            maxLength={25}
                            required />
                </div>
                <div className='auth-form-field'>
                    <label htmlFor="name">Name</label>
                    <input type="text"
                           value={name}
                           onChange={e => setName(e.target.value)}
                           required />
                </div>
                <div className='auth-form-field'>
                    <label htmlFor="email">Email</label>
                    <input type="text"
                           value={email}
                           onChange={e => setEmail(e.target.value)}
                           required />
                </div>
                <div className='auth-form-field'>
                    <label htmlFor="password">Password</label>
                    <input type="password"
                           value={password}
                           onChange={e => setPassword(e.target.value)}
                           minLength={8}
                           required />
                </div>
                <hr className='auth-form-divider' />
                <button className='auth-form-btn' type='submit'>Create Account</button>
                <p className='have-an-account'>Have an account? <Link to="/signup">Login!</Link></p>
            </div>
        </form>
    )
}