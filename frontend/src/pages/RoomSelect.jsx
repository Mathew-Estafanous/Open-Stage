import React, { useState } from 'react';
import { useHistory, useLocation } from 'react-router-dom';
import { GetRoom } from '../api/Rooms';
import './css/RoomSelect.css';
import { Oval } from '@agney/react-loading';
import { Error } from '../components/Error';
import { ProfileIcon } from '../components/ProfileIcon';

function useQuery () {
  return new URLSearchParams(useLocation().search);
}

export const RoomSelect = () => {
  const [code, setCode] = useState('');
  const [isLoading, setLoading] = useState(false);

  const query = useQuery();
  const history = useHistory();

  const joinRoom = async (e) => {
    e.preventDefault();
    setLoading(true);
    const result = await GetRoom(code);
    setLoading(false);
    if (result.error !== '') {
      history.push('/?error=' + result.error);
      return;
    }

    history.push('/room/' + code);
  };

  return (
    <>
      <header>
        <img className='logo' src='/Logo.png' alt='Logo' />
        <h1 className='title'>Open Stage</h1>

        <ProfileIcon />
      </header>
      <form className='roomCode' onSubmit={joinRoom}>
        <h1>Join Room</h1>
        <hr />
        <div className='selector'>
          <img className='hashtag' src='/Hashtag.png' alt='hashtag symbol' />
          <input
            className='form-input' maxLength={20} placeholder='Enter Room Code'
            onChange={e => setCode(e.target.value)}
          />
          <button type='submit'>
            <img className='btn' src='/Select-Arrow.png' alt='' />
          </button>
        </div>

        {isLoading
          ? <div className='load'>
            <Oval className='loader' />
            </div>
          : null}

        {query.get('error') && !isLoading
          ? <Error msg={query.get('error')} />
          : null}
      </form>

    </>
  );
};
