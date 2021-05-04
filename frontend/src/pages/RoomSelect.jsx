import React from 'react'
import './RoomSelect.css'

const RoomSelect = () => {
    return (
        <div className='wrapper'>
            <div className='header'>
                <img className='logo' src='/Logo.png' alt="Logo"/>
                <h1 className='title'>Open Stage</h1>

                <img className='profile' src="/Profile.png" alt=""/>
            </div>

            <form><h1>Join Room</h1><input placeholder='#  Enter Room Code' />
            </form>
        </div>
    )
}

export default RoomSelect;