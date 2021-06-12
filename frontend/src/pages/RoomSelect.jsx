import React, {useState} from 'react'
import { useHistory, useLocation } from 'react-router-dom';
import { GetRoom } from "../http/Rooms";
import './css/RoomSelect.css';
import {Oval} from "@agney/react-loading";
import {useAuth} from "../context/AuthContext";
import {Error} from "../components/Error";


function useQuery() {
    return new URLSearchParams(useLocation().search);
}

export const RoomSelect = () => {
    const [code, setCode] = useState("");
    const [isLoading, setLoading] = useState(false);

    const query = useQuery();
    const history = useHistory();
    const {account} = useAuth()

    const joinRoom = async (e) => {
        e.preventDefault()
        setLoading(true);
        let result = await GetRoom(code);
        setLoading(false);
        if(result.error !== '') {
            history.push("/?error=" + result.error)
            return;
        }

        history.push("/room/" + code);
    }

    const clickedProfile = (e) => {
        if (!account) {
            history.push("/login");
        } else {
            history.push("/account");
        }
    }

    return (
        <>
        <header>
            <img className='logo' src='/Logo.png' alt="Logo"/>
            <h1 className='title'>Open Stage</h1>

            <img className='profile' src="/Profile.png" alt="profile" onClick={clickedProfile}/>
        </header>
        <form className='roomCode' onSubmit={joinRoom} >
            <h1>Join Room</h1>
            <hr/>
            <div className='selector'>
                <img className='hashtag' src="/Hashtag.png" alt="hashtag symbol"/>
                <input maxLength={20} placeholder='Enter Room Code'
                       onChange={e => setCode(e.target.value)} />
                <button type="submit">
                    <img className='btn'
                         src="/Select-Arrow.png" alt=""/>
                </button>
            </div>

            {isLoading?
                <div className='load'>
                    <Oval className='loader' />
                </div>: null
            }

            {query.get("error") && !isLoading?
                <Error msg={query.get("error")}/>
                :null
            }
        </form>


        </>
    )
}