import './css/CreateRoom.css';

export const CreateRoom = (props) => {
    const submitRoom = (e) => {
        e.preventDefault();
    }

    return (
        (props.trigger)? (
            <div className='popup'>
                <div className='popup-inner'>
                    <h3 className='popup-title'>Create A Room</h3>
                    <form className="popup-form" onSubmit={submitRoom}>
                        <div className='form-field'>
                            <label htmlFor="host">Host</label>
                            <input type="text" placeholder='Enter host name.' required />
                        </div>
                        <div className='form-field'>
                            <label htmlFor="code">Room Code</label>
                            <input type="text" placeholder='Enter room code (Optional).' />
                        </div>
                    </form>
                    <img className='popup-close' src="/Close.png" alt="Close" onClick={props.close}/>
                </div>
            </div>
        ) : ''
    )
}