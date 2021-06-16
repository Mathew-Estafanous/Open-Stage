import {useHistory} from "react-router-dom";
import {ProfileIcon} from "../components/ProfileIcon";
import {useAuth} from "../context/AuthContext";
import {useEffect, useState} from "react";
import {GetAccountInfo} from "../http/Accounts";
import {RoomClip} from "../components/RoomClip";
import './css/AccountPage.css';
import {CreateRoom} from "../components/CreateRoom";

export const AccountPage = () => {
    const [username, setUsername] = useState('');
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [isOpen, setOpen] = useState(false);

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

    let tempRoomInfo = [
        {
            host: 'Mathew',
            room_code: 'goto',
        },
        {
            host: 'Elijah',
            room_code: 'GOlMkh6uME',
        },
        {
            host: 'Mathew',
            room_code: 'gto',
        },
        {
            host: 'Elijah',
            room_code: 'GOlMfsakh6uME',
        },
    ]

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
                <h3 className='account-rooms-title'>Your Rooms</h3>
                <div className='account-room-list'>
                    {tempRoomInfo.map(r => {
                        return <RoomClip key={r.room_code} {...r} />;
                    })}
                </div>
                <hr className='account-rooms-hr' />
                <img className='account-rooms-create' src="/Create.png"
                     alt="Create" onClick={() => setOpen(true)} />
                <CreateRoom trigger={isOpen} close={() => setOpen(false)} />
            </div>
        </div>
        </>
    )
}