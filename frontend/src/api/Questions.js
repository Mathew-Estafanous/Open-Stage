import {DELETE, GET, POST, PUT, sendRequest} from "./Request";

const MultiQuestionResponse = {
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
    return sendRequest({
        endpoint: '/questions/' + code,
        response: {...MultiQuestionResponse},
        method: GET,
        onRespond: (res, [ok, data]) => {
            if(!ok) {
                res.error = data.message
                return res
            }

            res.body = data
            return res
        }
    })
}

const QuestionResponse = {
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
    let data = {
        associated_room: roomCode,
        question: question
    }
    if(name !== '') {
        data.questioner_name = name;
    }

    return sendRequest({
            endpoint: '/questions',
            response: {...QuestionResponse},
            method: POST,
            body: data
        }
    )
}

export const UpdateLikes = (increment, id) => {
    let data = {
        question_id: id,
        like_increment: increment
    }

    return sendRequest({
        endpoint: '/questions',
        response: {...QuestionResponse},
        method: PUT,
        body: data
    })
}

export const DeleteQuestion = (id, token) => {
    return sendRequest({
        endpoint: '/questions/' + id,
        method: DELETE,
        headers: {
            'Authorization': 'Bearer ' + token
        },
        onRespond: (_, [ok, data]) => {
            return ok? {status: 200} : {...data.json()}
        }
    })
}