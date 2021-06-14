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
    return makeFetchRequest(response, data, path)
}

let signupResponse = {
    body: {
        username: '',
        password: '',
        email: '',
        name: '',
    },
    error: {
        message: '',
        status: 200
    }
}


export const CreateAccount = (username, password, name, email) => {
    let path = url + '/accounts/signup';
    let data = {
        username: username,
        password: password,
        email: email,
        name: name,
    }

    let response = { ...signupResponse}
    return makeFetchRequest(response, data, path)
}

const makeFetchRequest = (response, data, path) => {
    return fetch(path, {
        method: 'POST',
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