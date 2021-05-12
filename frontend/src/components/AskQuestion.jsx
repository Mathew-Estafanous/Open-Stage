import React, {useState} from "react";
import "./AskQuestion.css"
import { PostQuestion } from "../http/Questions";

export const AskQuestion = ({code}) => {
    const [question, setQuestion] = useState('');
    const [name, setName] = useState('');

    const postQuestion = async () => {
        await PostQuestion(code, question, name);
    }

    return (
        <div className='askdiv'>
            <form className='askform' onSubmit={postQuestion}>
                <textarea rows='2' placeholder='Write your question here...'
                          value={question} onChange={e => setQuestion(e.target.value)} />
                <div>
                    <img src="/User.png" alt="User" />
                    <input maxLength={20} placeholder='Name (Optional)'
                           value={name} onChange={e => setName(e.target.value)} />
                    <button>Post</button>
                </div>
            </form>
        </div>
    )
}