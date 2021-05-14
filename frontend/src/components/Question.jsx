import React, {useState} from 'react';
import "./Question.css"
import {UpdateLikes} from "../http/Questions";

export const Question = (prop) => {
    const [liked, setLiked] = useState(prop.isLiked);
    const [totalLikes, setTotalLikes] = useState(prop.total_likes);

    const clickLike = async () => {
        let likes = totalLikes + ((!liked)? 1: -1);
        let result = await UpdateLikes(likes, prop.question_id);
        if(result.status !== 200) {
            console.log(result.error);
            return;
        }
        setTotalLikes(likes);
        setLiked(!liked);
    }

    return (
        <div className='question' >
            <div className='head' >
                <div className='name' >
                    <img className='user' src="/User.png" alt="user"/>
                    <h3>{prop.questioner_name}</h3>
                </div>
                <div className='like'>
                    <img src={(liked)? "/Upvote-Blue.png": "/Upvote-Black.png"} alt="Upvote"
                        onClick={clickLike}/>
                    <h3 className={(liked)? "liked": ""}>{totalLikes}</h3>
                </div>
            </div>
            <p>{prop.question}</p>
        </div>
    )
}