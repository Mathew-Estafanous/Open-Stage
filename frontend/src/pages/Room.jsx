import React, {useState} from 'react';
import { AskQuestion } from "../components/AskQuestion";
import "./Room.css"
import { Question } from "../components/Question";

const Room = () => {
    const [roomName] = useState("GopherCon");

    let props = {
        likes: 3,
        name: "Anonymous",
        question: "If 1 + 1 = 2, then what does 2 + 2 equal?"
    };

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
        <Question {...props} />
        </>
    )
}

export default Room;