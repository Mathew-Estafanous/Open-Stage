import {useHistory} from "react-router-dom";
import './css/AccountPage.css';

export const AccountPage = () => {
    const history = useHistory();
    return (
        <>
        <header>
            <img className='logo' src='./Logo.png' alt='Logo'
                 onClick={() => history.push('/')} />
        </header>
        <h1>IM THE ACCOUNT PAGE</h1>
        </>
    )
}