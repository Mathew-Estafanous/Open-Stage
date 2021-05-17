const url = (process.env.REACT_APP_ENV === 'production')?
    'https://open-stage-api.herokuapp.com/v1'
    :'http://localhost:8080/v1';

const roomResponse = {
    body: {
        host: '',
        room_code: ''
    },
    error: ''
}

export const GetRoom = async (code) => {
    let path = url + '/rooms/' + code
    let response = {...roomResponse}

    return await fetch(path)
        .then(res => Promise.all([res.ok , res.json()]))
        .then(([ok, data]) => {
            if(!ok) {
                let errResp = {...data}
                if(errResp.status === 404) {
                    response.error = errResp.message
                    return response
                }

                response.error = "Internal error, please try again."
                return response
            }

            response.body = {...data}
            return response
        })
        .catch(err => {
            response.error = 'We were unable to connect to our servers.';
            return response;
        })
}