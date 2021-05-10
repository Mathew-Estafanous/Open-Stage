import React, {useState} from 'react'
import { useHistory } from 'react-router-dom';
import './RoomSelect.css'
import {GetRoom} from "../http/Rooms";

export const RoomSelect = () => {
    const [code, setCode] = useState("");
    const history = useHistory();

    const joinRoom = async () => {
        let result = await GetRoom(code);
        if(result.notFound === true) {
            alert(result.error);
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
        </>
    )
}