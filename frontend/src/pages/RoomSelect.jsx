import React from 'react'
import './RoomSelect.css'

const RoomSelect = () => {
    const joinRoom = () => {
        console.log("HEYYYY");
    }

    return (
        <div className='wrapper'>
            <div className='header'>
                <img className='logo' src='/Logo.png' alt="Logo"/>
                <h1 className='title'>Open Stage</h1>

                <img className='profile' src="/Profile.png" alt=""/>
            </div>

            <form>
                <h1>Join Room</h1>
                <hr/>
                <div>
                    <img className='hashtag' src="/Hashtag-Symbol.png" alt="hashtag symbol"/>
                    <input maxLength={20} placeholder='Enter Room Code' />
                    <img className='btn'
                         src="/Select-Arrow.png" alt=""
                         onClick={joinRoom} />
                </div>
            </form>
        </div>
    )
}

export default RoomSelect;