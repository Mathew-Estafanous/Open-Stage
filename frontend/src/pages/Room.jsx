import React, {useEffect, useState} from "react";
import {useHistory, useParams} from "react-router-dom";
import { AskQuestion } from "../components/AskQuestion";
import { Question } from "../components/Question";
import { GetRoom } from "../http/Rooms";
import { GetAllQuestions } from "../http/Questions";
import "./Room.css";

export const Room = () => {
    const [room, setRoom] = useState("");
    const [questions, setQuestions] = useState([])
    const { code } = useParams();
    const history = useHistory();

    // INITIAL setup of the room that gets room information such as code,
    // and all the questions.
    useEffect( () => {
        async function callAPIs() {
            let roomResult = await GetRoom(code);
            if(roomResult.error !== '') {
                history.push("/?error=" + roomResult.error);
                return;
            }
            setRoom(roomResult.body.room_code);

            let questionResult = await GetAllQuestions(code);
            if(questionResult.error !== '') {
                history.push("/?error=" + questionResult.error);
                return;
            }

            questionResult.body
                .sort((a, b) => {
                    return (a.total_likes < b.total_likes)? 1: (a.total_likes > b.total_likes)? -1: 0;
                }).map(q => q.isLiked = false)

            setQuestions(questionResult.body);
        }
        callAPIs();
    }, [code])

    return (
        <>
        <header>
            <img className='logo' src='/Logo.png' alt="Logo"/>

            <div className='roomInfo'>
                <h2 className='title'>Current Room</h2>
                <h3 className='name'>{room}</h3>
            </div>
        </header>

        <AskQuestion code={room} />
        {questions.map(q => {
            return <Question key={q.question_id} {...q}/>;
        })}
        </>
    )
}