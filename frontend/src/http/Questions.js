function test() {
    let env = process.env.REACT_APP_ENV;
    return (env === 'production')? 'https://open-stage-api.herokuapp.com/v1' :'http://localhost:8080/v1';
}

const url = test();
let questionsResponse = {
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
    let response = {...questionsResponse}
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

let postQuestionResponse = {
    body: {
        question_id: 0,
        associated_room: '',
        question: '',
        questioner_name: '',
        total_likes: 0
    },
    err: ''
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

    let response = {...postQuestionResponse}
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

export const UpdateLikes = (likes, id) => {
    let path = url + '/questions'
    let data = {
        question_id: id,
        total_likes: likes
    }

    let result = {
        status: 200,
        error: ''
    };
    return fetch(path, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(resp => {
            if(!resp.ok) {
                let error = resp.json();
                result.status = error.status;
                result.error = error.message;
            }
            return result;
        })
}