import {url} from "./Rooms";

let accountResponse = {
    body: {
        id: 0,
        email: '',
        name: '',
        username: '',
    },
    error: {
        message: '',
        status: 200,
    }
}

export const GetAccountInfo = (username, token) => {
    let path = url + '/accounts/' + username;
    let response = {...accountResponse}
    let request = {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
    }
    return makeFetchRequest(response, request, path)
}

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
    let request = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }
    return makeFetchRequest(response, request, path)
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
    let request = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }
    return makeFetchRequest(response, request, path)
}

const makeFetchRequest = (response, request, path) => {
    return fetch(path, request)
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