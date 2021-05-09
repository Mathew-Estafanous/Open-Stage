import React from  'react';
import "./Question.css"

export const Question = (prop) => {
    return (
        <div className='question' >
            <div className='head' >
                <div className='name' >
                    <img className='user' src="/User.png" alt="user"/>
                    <h3>{prop.name}</h3>
                </div>
                <div className='like'>
                    <img src="/Upvote-Black.png" alt="Upvote"/>
                    <h3>{prop.likes}</h3>
                </div>
            </div>
            <p>{prop.question}</p>
        </div>
    )
}