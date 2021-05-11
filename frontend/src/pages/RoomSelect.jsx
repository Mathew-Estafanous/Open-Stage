import React, {useState} from 'react'
import { useHistory } from 'react-router-dom';
import './RoomSelect.css'
import {GetRoom} from "../http/Rooms";

export const RoomSelect = () => {
    const [code, setCode] = useState("");
    const [isValid, setValid] = useState(true);
    const [error, setError] = useState('');

    const history = useHistory();

    const joinRoom = async () => {
        let result = await GetRoom(code);
        if(result.status !== 200) {
            setValid(false);
            setError(result.error);
            return;
        }

        history.push('/room/' + code)
    }

    return (
        <>
        <header>
            <img className='logo' src='/Logo.png' alt="Logo"/>
            <h1 className='title'>Open Stage</h1>

            <img className='profile' src="/Profile.png" alt=""/>
        </header>

        <form className='roomCode' >
            <h1>Join Room</h1>
            <hr/>
            <div>
                <img className='hashtag' src="/Hashtag-Symbol.png" alt="hashtag symbol"/>
                <input maxLength={20} placeholder='Enter Room Code'
                       onChange={e => setCode(e.target.value)} />
                <img className='btn'
                     src="/Select-Arrow.png" alt=""
                     onClick={joinRoom} />
            </div>
        </form>

        {isValid? null:
            <div className='errContainer' >
                <div className='error'>
                    <img src="/Warning.png" alt="Warning"/>
                    <p>{error}</p>
                </div>
            </div>
        }
        </>
    )
}