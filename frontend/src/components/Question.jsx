import React, {useEffect, useState} from 'react';
import "./Question.css"
import {UpdateLikes} from "../http/Questions";

export const Question = (prop) => {
    const isLiked = () => {
        let likeData = JSON.parse(localStorage.getItem('like.data'));
        if(likeData === null) {
            return false;
        }

        let roomArr = likeData[prop.associated_room] || [];
        let result = roomArr.indexOf(prop.question_id);
        return (result !== -1);
    }

    const [liked, setLiked] = useState(isLiked());
    const [totalLikes, setTotalLikes] = useState(prop.total_likes);

    const updateLocalStorage = () => {
        let likeData = JSON.parse(localStorage.getItem('like.data') || '{}');
        let roomArr = likeData[prop.associated_room] || [];
        if(liked) {
            let index = roomArr.indexOf(prop.question_id)
            roomArr.splice(index, 1);
        } else {
            roomArr.push(prop.question_id);
        }
        likeData[prop.associated_room] = roomArr;
        localStorage.setItem('like.data', JSON.stringify(likeData))
    }

    const clickLike = async () => {
        let increment = (!liked)? 1: -1;
        let result = await UpdateLikes(increment, prop.question_id);
        if(result.error !== '') {
            console.log(result.error);
            return;
        }
        setTotalLikes(result.body.total_likes);
        setLiked(!liked);

        updateLocalStorage();
    }

    // Update the total likes when the prop passed in is updated.
    useEffect(() => {
        setTotalLikes(prop.total_likes);
        setLiked(isLiked());
    }, [prop])

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