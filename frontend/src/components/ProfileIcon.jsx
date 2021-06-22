import {useAuth} from "../context/AuthContext";
import {useHistory} from "react-router-dom";
import "./css/ProfileIcon.css";
import {Logout} from "../http/Accounts";

export const ProfileIcon = () => {
    const {account, setAccount} = useAuth();
    const history = useHistory();

    const clickedProfile = () => {
        if (!account) {
            history.push("/login");
        } else {
            history.push("/account");
        }
    }

    const clickedLogout = () => {
        let result = Logout(account.access_token, account.refresh_token);
        result.then(res => {
            if(res.status !== 200) {
                console.log(res.message);
            }
            history.push('/');
            setAccount(null);
        })
    }

    return (
        <div className='profile-div'>
            {account? <h2 className='profile-username'>{account.username}</h2>:null}
            <img className='profile' src="/Profile.png" alt="profile" onClick={clickedProfile}/>
            {account? <img className='logout' src="/Logout.png" alt="Logout" onClick={clickedLogout} />: null}
        </div>
    )
}