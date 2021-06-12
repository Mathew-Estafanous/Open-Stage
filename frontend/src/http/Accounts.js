import {url} from "./Rooms";

let loginResponse = {
    body: {
        access_token: '',
        response_token: '',
    },
    error: {
        message: '',
        status: 200
    }
}

export const LoginAccount = (username, pass) => {
    let path = url + '/accounts/login';
    let data = {
        password: pass,
        username: username
    }

    let response = {...loginResponse}
    return fetch(path, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(resp => Promise.all([resp.ok, resp.json()]))
        .then(([ok, data]) => {
            if (!ok) {
                response.error = {...data}
                return response
            }

            response.body = {...data}
            return response
        })
}