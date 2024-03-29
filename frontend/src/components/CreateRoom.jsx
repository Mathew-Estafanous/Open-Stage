import './css/CreateRoom.css';
import { useAuth } from '../context/AuthContext';
import React, { useState } from 'react';
import { CreateTheRoom } from '../api/Rooms';

export const CreateRoom = ({ update, trigger, close }) => {
  const [host, setHost] = useState('');
  const [code, setCode] = useState('');

  const { account } = useAuth();

  const submitRoom = (e) => {
    e.preventDefault();
    const result = CreateTheRoom(host, code, account.access_token);
    result.then(res => {
      if (res.error.status !== 201) {
        console.log(res.error.message);
        return;
      }

      update();
      close();
    });
  };

  return (
    (trigger)
      ? (
        <div className='popup'>
          <div className='popup-inner'>
            <h3 className='popup-title'>Create A Room</h3>
            <hr className='popup-hr' />
            <form className='popup-form' onSubmit={submitRoom}>
              <div className='form-field'>
                <label htmlFor='host'>Host</label>
                <input
                  type='text'
                  value={host}
                  placeholder='Enter host name.'
                  onChange={e => setHost(e.target.value)}
                  required
                />
              </div>
              <div className='form-field'>
                <label htmlFor='code'>Room Code</label>
                <input
                  type='text'
                  value={code}
                  placeholder='Create a room code. (Optional)'
                  onChange={e => setCode(e.target.value)}
                />
              </div>
              <button className='form-btn'>Create</button>
            </form>
            <img className='popup-close' src='/Close.png' alt='Close' onClick={close} />
          </div>
        </div>
        )
      : ''
  );
};
