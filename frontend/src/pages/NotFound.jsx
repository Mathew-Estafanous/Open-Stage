import React from 'react';
import './css/NotFound.css';
import { useHistory } from 'react-router-dom';

export const NotFound = () => {
  const history = useHistory();
  return (
    <div className='notFound'>
      <img
        className='logo' src='/Logo.png' alt='Logo'
        onClick={_ => history.push('/')}
      />
      <h1>Not Found</h1>
      <h3>Sorry, We couldn't find the page you were looking for!</h3>
    </div>
  );
};
