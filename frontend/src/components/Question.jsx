import React from 'react';
import "./Question.css"

export const Question = (prop) => {
    const upvoteImg = () => {
        return (prop.isLiked)? "/Upvote-Blue.png": "/Upvote-Black.png";
    }

    const likeColour = () => {
        return (prop.isLiked)? "liked": "";
    }

    return (
        <div className='question' >
            <div className='head' >
                <div className='name' >
                    <img className='user' src="/User.png" alt="user"/>
                    <h3>{prop.name}</h3>
                </div>
                <div className='like'>
                    <img src={upvoteImg()} alt="Upvote"/>
                    <h3 className={likeColour()}>{prop.likes}</h3>
                </div>
            </div>
            <p>{prop.question}</p>
        </div>
    )
}