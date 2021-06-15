import {useAuth} from "../context/AuthContext";
import {useHistory} from "react-router-dom";
import "./css/ProfileIcon.css";

export const ProfileIcon = () => {
    const {account} = useAuth();
    const history = useHistory();

    const clickedProfile = () => {
        if (!account) {
            history.push("/login");
        } else {
            history.push("/account");
        }
    }

    return (
        <div className='profile-div'>
            {account? <h2 className='profile-username'>{account.username}</h2>:null}
            <img className='profile' src="/Profile.png" alt="profile" onClick={clickedProfile}/>
        </div>
    )
}