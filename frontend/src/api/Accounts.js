import {AuthHeader, GET, POST, sendRequest} from "./Request";

let AccountResponse = {
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
    return sendRequest({
        endpoint: '/accounts/' + username,
        response: {...AccountResponse},
        method: GET,
        headers: AuthHeader(token)
    })
}

let AccessResponse = {
    body: {
        access_token: '',
        refresh_token: '',
    },
    error: {
        message: '',
        status: 200
    }
}

export const LoginAccount = (username, pass) => {
    let data = {
        password: pass,
        username: username
    }
    return sendRequest({
        endpoint: '/accounts/login',
        response: {...AccessResponse},
        method: POST,
        body: data,
    })
}

export const Logout = (access, refresh) => {
    let data = {
        access_token: access,
        refresh_token: refresh,
    }

    return sendRequest({
        endpoint: '/accounts/logout',
        method: POST,
        body: data
    })
}

let SignupResponse = {
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
    let data = {
        username: username,
        password: password,
        email: email,
        name: name,
    }

    return sendRequest({
        endpoint: '/accounts/signup',
        response: {...SignupResponse},
        method: POST,
        body: data
    })
}

export const RefreshToken = (refresh) => {
    let data = {
        refresh_token: refresh
    }

    return sendRequest({
        endpoint: '/accounts/refresh',
        response: {...AccessResponse},
        method: POST,
        body: data
    })
}
