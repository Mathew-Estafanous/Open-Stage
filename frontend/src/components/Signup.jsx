import {useState} from "react";
import {Error} from "./Error";
import {Link, useHistory} from "react-router-dom";
import {CreateAccount} from "../http/Accounts";
import "./css/AuthForm.css";

export const Signup = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');

    const [error, setError] = useState('');
    const history = useHistory();

    const submitSignup = async (e) => {
        e.preventDefault()
        let resp = await CreateAccount(username, password, name, email);
        if (resp.error.status !== 201) {
            if(resp.error.status === 409) {
                setError('The username is already taken.');
            } else if(resp.error.status === 400 ) {
                setError('Please ensure all your fields are valid.');
            } else {
                setError('We encountered an error, please try again.');
            }
            return;
        }

        history.push('/login');
    }

    return (
        <form className='auth-form' onSubmit={submitSignup} >
            {error?
                <Error msg={error} />: null
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
                    <input type="text"
                           value={password}
                           onChange={e => setPassword(e.target.value)}
                           minLength={8}
                           required />
                </div>
                <hr className='auth-form-divider' />
                <button className='auth-form-btn' type='submit'>Create Account</button>
                <p className='have-an-account'>Have an account? <Link to="/login">Login!</Link></p>
            </div>
        </form>
    )
}