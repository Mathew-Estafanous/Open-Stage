import { useHistory } from 'react-router-dom';
import { Signup } from '../components/Signup';
import './css/AuthPage.css';

export const SignupPage = () => {
  const history = useHistory();
  return (
    <>
      <header>
        <img
          className='logo' src='./Logo.png' alt='Logo'
          onClick={() => history.push('/')}
        />
      </header>
      <section className='auth-section'>
        <Signup />
      </section>
    </>
  );
};
