import React, {useState} from 'react';
import AskQuestion from "../components/AskQuestion";
import "./Room.css"

const Room = () => {
    const [roomName] = useState("GopherCon");

    return (
        <>
        <header>
            <img className='logo' src='/Logo.png' alt="Logo"/>

            <div className='roomInfo'>
                <h2 className='title'>Current Room</h2>
                <h3 className='name'>{roomName}</h3>
            </div>
        </header>

        <AskQuestion />
        </>
    )
}

export default Room;