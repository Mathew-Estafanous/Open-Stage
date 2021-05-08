import React from "react";
import "./AskQuestion.css"

const AskQuestion = () => {
    return (
        <form className='askform' >
            <input placeholder='Write your question here...' />
            <div>
                <img src="/User.png" alt="User" />
                <button>Post</button>
            </div>
        </form>
    )
}

export default AskQuestion;