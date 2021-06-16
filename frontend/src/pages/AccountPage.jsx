import {useHistory} from "react-router-dom";
import './css/AccountPage.css';
import {ProfileIcon} from "../components/ProfileIcon";
import {useAuth} from "../context/AuthContext";
import {useEffect, useState} from "react";
import {GetAccountInfo} from "../http/Accounts";

export const AccountPage = () => {
    const [username, setUsername] = useState('');
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');

    const {account} = useAuth();
    const history = useHistory();

    useEffect( () => {
        let result = GetAccountInfo(account.username, account.access_token);
        result.then(res => {
            if(res.error.status !== 200) {
                history.push('/?error=' + result.error.message);
                return;
            }

            setUsername(res.body.username);
            setName(res.body.name);
            setEmail(res.body.email);
        })
    }, [account])
    return (
        <>
        <header>
            <img className='logo' src='./Logo.png' alt='Logo'
                 onClick={() => history.push('/')} />

            <ProfileIcon />
        </header>
        <div className='account-wrapper'>
            <div className='account-info'>
                <h3 className='account-info-title'>Account Information</h3>
                <div className='account-info-field'>
                    <p className='account-info-name'>Username: </p>
                    <p className='account-info-result'>{username}</p>
                </div>
                <div className='account-info-field'>
                    <p className='account-info-name'>Name: </p>
                    <p className='account-info-result'>{name}</p>
                </div>
                <div className='account-info-field'>
                    <p className='account-info-name'>Email: </p>
                    <p className='account-info-result'>{email}</p>
                </div>
            </div>
            <div className='account-rooms'>
                <h3>Your Rooms</h3>
            </div>
        </div>
        </>
    )
}