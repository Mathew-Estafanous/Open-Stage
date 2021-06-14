import {useHistory} from "react-router-dom";
import './css/Account.css';

export const Account = () => {
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