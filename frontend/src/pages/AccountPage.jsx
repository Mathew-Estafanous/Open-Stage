import {useHistory} from "react-router-dom";
import {ProfileIcon} from "../components/ProfileIcon";
import {useAuth} from "../context/AuthContext";
import {useEffect, useState} from "react";
import {GetAccountInfo} from "../http/Accounts";
import {RoomClip} from "../components/RoomClip";
import {CreateRoom} from "../components/CreateRoom";
import {AllRoomsAssociated} from "../http/Rooms";
import './css/AccountPage.css';

export const AccountPage = () => {
    const [username, setUsername] = useState('');
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');

    const [isOpen, setOpen] = useState(false);
    const [rooms, setRooms] = useState([]);

    const {account} = useAuth();
    const history = useHistory();

    useEffect( () => {
        let result = GetAccountInfo(account.username, account.access_token);
        result.then(res => {
            if(res.error.status !== 200) {
                history.push('/?error=' + res.error.message);
                return;
            }

            setUsername(res.body.username);
            setName(res.body.name);
            setEmail(res.body.email);
        })

        result = AllRoomsAssociated(account.access_token);
        result.then(res => {
            if(res.error.status !== 200) {
                history.push('/?error=' + res.error.message);
                return;
            }

            setRooms(res.body);
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
                <h3 className='account-rooms-title'>Your Rooms</h3>
                <div className='account-room-list'>
                    {rooms.length? (rooms.map(r => {
                        return <RoomClip key={r.room_code} {...r} />;
                    })): <NoRoom />}
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

const NoRoom = () => {
    return (
        <div className='no-room'>
            <h3>You don't have any rooms!</h3>
        </div>
    )
}