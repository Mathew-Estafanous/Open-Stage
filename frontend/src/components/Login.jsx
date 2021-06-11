import {useState} from "react";
import {Link} from "react-router-dom";
import "./css/Login.css"

export const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault()
        console.log("SUBMITTED!")
    }

    return (
        <form className='login-form' onSubmit={handleSubmit}>
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