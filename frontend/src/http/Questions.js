import {url} from "./Rooms";

let multipleQuestionResponse = {
    body: [
        {
            associated_room: '',
            question: '',
            question_id: 0,
            questioner_name: '',
            total_likes: 0
        }
    ],
    error: ''
}

export const GetAllQuestions = (code) => {
    let path = url + '/questions/' + code;
    console.log(path)
    let response = {...multipleQuestionResponse}
    return fetch(path)
        .then(resp => Promise.all([resp.ok, resp.json()]))
        .then(([ok, data]) => {
            if(!ok) {
                response.error = data.message
                return response
            }


            response.body = data
            return response
    })
}

let questionResponse = {
    body: {
        question_id: 0,
        associated_room: '',
        question: '',
        questioner_name: '',
        total_likes: 0
    },
    error: ''
}

export const PostQuestion = (roomCode, question, name) => {
    let path = url + '/questions';
    let data = {
        associated_room: roomCode,
        question: question
    }
    if(name !== '') {
        data.questioner_name = name;
    }

    let response = {...questionResponse}
    return fetch(path, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(resp => Promise.all([resp.ok, resp.json()]))
        .then(([ok, data]) => {
            if(!ok) {
                response.error = data.message
                return response
            }

            response.body = data
            return response
        })
}

export const UpdateLikes = (increment, id) => {
    let path = url + '/questions'
    let data = {
        question_id: id,
        like_increment: increment
    }

    let response = {...questionResponse}
    return fetch(path, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(resp => Promise.all([resp.ok, resp.json()]))
        .then(([ok, data]) => {
            if(!ok) {
                response.error = data.message
                return response
            }

            response.body = data
            return response
        })
}

export const DeleteQuestion = (id, token) => {
    let path = url + '/questions/' + id;
    let request = {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token,
        },
    }

    return fetch(path, request)
        .then(resp => Promise.all([resp.ok, resp]))
        .then(([ok, data]) => {
            if(!ok) {
                return {...data.json()}
            }

            return {status: 200}
        })
}