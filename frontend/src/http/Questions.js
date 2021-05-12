const url = 'http://localhost:8080/v1';

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

export const GetAllQuestions = async (code) => {
    let path = url + '/questions/' + code;
    console.log(path)
    let response = {...questionsResponse}
    return await fetch(path)
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

export const PostQuestion = async (roomCode, question, name) => {
    let path = url + '/questions';
    let data = {
        associated_room: roomCode,
        question: question
    }
    if(name !== '') {
        data.questioner_name = name;
    }

    let response = {...postQuestionResponse}
    return await fetch(path, {
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