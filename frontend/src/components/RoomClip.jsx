import './css/RoomClip.css';
import {useHistory} from "react-router-dom";
import {DeleteRoom} from "../http/Rooms";
import {useAuth} from "../context/AuthContext";

export const RoomClip = (prop) => {
    const history = useHistory();
    const { account } = useAuth();

    const enterRoom = () => {
        history.push('/room/' + prop.room_code);
    }

    const deleteRoom = () => {
        let result = DeleteRoom(prop.room_code, account.access_token);
        result.then(resp => {
            if(resp.status !== 200) {
                alert('We encountered an error while deleting the room.');
            }
        })
    }

    return (
        <div className='room-clip-wrapper'>
            <h4 className='room-clip-name'>{prop.host}</h4>
            <img className='hashtag' src="/Hashtag.png" alt="Hashtag"/>
            <p className='room-clip-code'>{prop.room_code}</p>
            <img className='enter' src="/Enter.png" alt="Enter" onClick={enterRoom}/>
            <img className='delete' src="/Trash-Can.png" alt="Trash" onClick={deleteRoom} />
        </div>
    )
}