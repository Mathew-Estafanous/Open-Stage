import React from "react";
import "./css/Error.css"

export const Error = ({msg}) => {
    return (
        <div className='error'>
            <img src="/Warning.png" alt="Warning"/>
            <p>{msg}</p>
        </div>
    )
}