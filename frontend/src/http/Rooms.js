import {makeFetchRequest} from "./Accounts";
import jwtDecode from "jwt-decode";

export const url = (process.env.REACT_APP_ENV === 'production')?
    'https://open-stage-api.herokuapp.com/v1' :'http://localhost:8080/v1';

const roomResponse = {
    body: {
        host: '',
        room_code: ''
    },
    error: ''
}

export const GetRoom = (code) => {
    let path = url + '/rooms/' + code
    let response = {...roomResponse}

    return fetch(path)
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
            console.log(err);
            response.error = 'We encountered an error with our servers.';
            return response;
        })
}

const errorResponse = {
    message: '',
    status: 200
}

export const DeleteRoom = (code, token) => {
    let path = url + '/rooms/' + code;
    let request = {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
    }
    let response = {...errorResponse}
    return fetch(path, request)
        .then(res => Promise.all([res.ok , res]))
        .then(([ok, data]) => {
            if(!ok) {
                response = {...data.json()}
            }
            return response
        })
}

let createRoomResponse = {
    body: {
        host: '',
        room_code: '',
        account_id: 1
    },
    error: {
        message: '',
        status: 201
    }
}

export const CreateTheRoom = (host, room_code, token) => {
    let path = url + '/rooms'
    let data = {
        host: host,
        room_code: room_code,
        account_id: token.id
    }

    let request = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify(data)
    }

    let response = {...createRoomResponse}
    return makeFetchRequest(response, request, path)
}

let allRoomResponse = {
    body: [
        {
            room_code: '',
            host: '',
            account_id: 0,
        }
    ],
    error: {
        message: '',
        status: 200
    }
}

export const AllRoomsAssociated = (token) => {
    let path = url + '/rooms/all'
    let request = {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
    }
    let response = {...allRoomResponse}
    return fetch(path, request)
        .then(resp => Promise.all([resp.ok, resp.json()]))
        .then(([ok, data]) => {
            if (!ok) {
                response.error = {...data}
                return response
            }

            response.body = data
            return response
        })
}