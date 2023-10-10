export const API_URL = process.env.REACT_APP_API_URL;

export const POST = 'POST';
export const GET = "GET";
export const PUT = "PUT";
export const DELETE = "DELETE"

export const sendRequest = ({
    endpoint,
    response = {},
    method = GET,
    headers = null,
    body = null,
    onRespond = null,
}) => {
    const path = API_URL + endpoint;
    const request = {
        method: method,
        headers: {
            'Content-Type': 'application/json',
            ...headers
        },
        body: body? JSON.stringify(body) : null,
    }

    return fetch(path, request)
        .then(resp => Promise.all([resp.ok, resp.json()]))
        .then(([ok, data]) => {
            if (onRespond) {
                return onRespond(response, [ok, data])
            }

            if (!ok) {
                response.error = {...data}
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

export const AuthHeader = (token) => {
    return {
        'Authorization': 'Bearer ' + token
    }
}