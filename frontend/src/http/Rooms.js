const url = 'http://localhost:8080/v1';

const roomResponse = {
    body: {
        host: '',
        room_code: ''
    },
    notFound: false,
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
                    response.notFound = true
                    response.error = errResp.message
                    return response
                }

                response.error = errResp.message
                return response
            }

            response.body = {...data}
            return response
        })
}