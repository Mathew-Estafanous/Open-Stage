import React, { useState } from 'react';
import './css/AskQuestion.css';
import { PostQuestion } from '../api/Questions';

export const AskQuestion = ({ code, onPost }) => {
  const [question, setQuestion] = useState('');
  const [name, setName] = useState('');

  const postQuestion = async (e) => {
    e.preventDefault();
    if (question.length === 0) {
      console.log('questions should not be empty');
      return;
    }
    const result = await PostQuestion(code, question, name);
    if (result.error) {
      console.log(result.error);
      return;
    }

    setQuestion('');
    setName('');
    onPost();
  };

  return (
    <div className='ask-div'>
      <form className='askform' onSubmit={postQuestion}>
        <textarea
          rows='2' placeholder='Write your question here...'
          value={question} onChange={e => setQuestion(e.target.value)}
        />
        <div>
          <img src='/User.png' alt='User' />
          <input
            maxLength={20} placeholder='Name (Optional)'
            value={name} onChange={e => setName(e.target.value)}
          />
          <button type='submit'>Post</button>
        </div>
      </form>
    </div>
  );
};
