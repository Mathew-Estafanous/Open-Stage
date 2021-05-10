import React from "react";
import "./NotFound.css"

export const NotFound = () => {
    return (
        <div className="notFound">
            <img className='logo' src='/Logo.png' alt="Logo"/>
            <h1>Not Found</h1>
            <h3>Sorry, We couldn't find the page you were looking for!</h3>
        </div>
    )
}