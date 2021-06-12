import {useState} from "react";
import {Link, useHistory} from "react-router-dom";
import {LoginAccount} from "../http/Accounts";
import {Error} from "./Error";
import {useAuth} from "../context/AuthContext";
import "./css/Login.css"

export const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(false);

    const { setAccount } = useAuth();
    const history = useHistory();

    const handleSubmit = async (e) => {
        e.preventDefault();
        let result = await LoginAccount(username, password);
        if (result.error.status !== 200) {
            setError(true);
            return;
        }

        setAccount(result.body);
        history.push('/account');
    }

    return (
        <form className='login-form' onSubmit={handleSubmit}>
            {error?
                <Error msg={"Invalid username or password!"}/>: null
            }

            <div className='login-form-wrapper'>
                <div className='login-form-header'>
                    <h2>LOGIN</h2>
                </div>
                <div className='login-form-field'>
                    <label htmlFor='username'>Username</label>
                    <input type='text'
                           value={username}
                           onChange={e => setUsername(e.target.value)}
                           maxLength={30}
                           required />
                </div>
                <div className='login-form-field'>
                    <label htmlFor='password'>Password</label>
                    <input type='password'
                           value={password}
                           onChange={e => setPassword(e.target.value)}
                           maxLength={30}
                           required />
                </div>
                <button className='login-form-btn' type='submit'>Login</button>
                <hr className='login-form-divider'/>
                <p className='have-an-account'>Don't have an account? <Link to="/signup">Signup!</Link></p>
            </div>
        </form>
    )
}