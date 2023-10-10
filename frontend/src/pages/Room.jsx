import React, {useEffect, useState} from "react";
import {useHistory, useParams} from "react-router-dom";
import { AskQuestion } from "../components/AskQuestion";
import { Question } from "../components/Question";
import {GetRoom} from "../api/Rooms";
import { GetAllQuestions } from "../api/Questions";
import { Oval } from "@agney/react-loading";
import "./css/Room.css";
import {useAuth} from "../context/AuthContext";

export const Room = () => {
    const [room, setRoom] = useState("");
    const [questions, setQuestions] = useState([])
    const [isLoading, setLoading] = useState(true);

    const { code } = useParams();
    const { account } = useAuth();
    const history = useHistory();

    const updateAllQuestions = () => {
        let questionResult = GetAllQuestions(code);
        questionResult.then(res => {
            if(res.error !== '') {
                history.push("/?error=" + res.error);
                return;
            }

            res.body
                .sort((a, b) => {
                    return (a.total_likes < b.total_likes)? 1: (a.total_likes > b.total_likes)? -1: 0;
                });

            setQuestions(res.body);
        })
    }

    const isOwnerOfRoom = () => {
        if(account === null) {
            return false;
        }

        return room.account_id === account.id;
    }

    // Initial setup of the room that gets room information such as code,
    // and all the questions.
    useEffect(  () => {
        let roomResult = GetRoom(code);
        roomResult.then(res => {
            if(res.error !== '') {
                history.push("/?error=" + res.error);
                return;
            }
            setRoom(res.body);

            updateAllQuestions();
            setLoading(false);
        })
    }, [code, history])

    useEffect(() => {
        const interval = setInterval(async () => {
            await updateAllQuestions();
        }, 6000);
        return () => clearInterval(interval)
    }, [])

    return (
        <>
        <header>
            <img className='logo' src='/Logo.png' alt="Logo"
                 onClick={() => history.push('/')} />

            <div className='roomInfo'>
                <h2 className='title'>Current Room</h2>
                <h3 className='name'>{room.room_code}</h3>
            </div>
        </header>

        <AskQuestion code={room.room_code} onPost={updateAllQuestions} />
        {isLoading?
            <div className='load'>
                <Oval className='loader' />
            </div>: null
        }
        {questions.map(q => {
            return <Question key={q.question_id} {...q}
                             is_owner={isOwnerOfRoom()} account={account} />;
        })}
        </>
    )
}