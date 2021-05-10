import React, { useState } from 'react';
import { AskQuestion } from "../components/AskQuestion";
import "./Room.css"
import { Question } from "../components/Question";

export const Room = () => {
    const [roomName] = useState("GopherCon");

    let questions = [
        {
            likes: 3,
            name: "Anonymous",
            question: "If 1 + 1 = 2, then what does 2 + 2 equal?",
            isLiked: true
        },
        {
            likes: 2,
            name: "Mathew Estafanous",
            question: "As as an employee, I am worried that my job may be in jeprody, how  would you explain your choices to us?",
            isLiked: true
        },
        {
            likes: 0,
            name: "Anonymous",
            question: "As as an employee, I am worried that my job may be in jeprody, how  would you explain your choices to us?",
            isLiked: false
        },
        {
            likes: 0,
            name: "Anonymous",
            question: "As as an employee, I am worried that my job may be in jeprody, how  would you explain your choices to us?",
            isLiked: false
        }
    ]

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
        {questions.map(question => {
            return <Question {...question}/>;
        })}
        </>
    )
}