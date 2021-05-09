import React from "react";
import "./AskQuestion.css"

export const AskQuestion = () => {
    return (
        <div className='askdiv'>
            <form className='askform' >
                <textarea rows='2' placeholder='Write your question here...' />
                <div>
                    <img src="/User.png" alt="User" />
                    <input maxLength={20} placeholder='Name (Optional)' />
                    <button>Post</button>
                </div>
            </form>
        </div>
    )
}