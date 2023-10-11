import { AuthHeader, GET, POST, sendRequest } from './Request';

const AccountResponse = {
  body: {
    id: 0,
    email: '',
    name: '',
    username: ''
  },
  error: {
    message: '',
    status: 200
  }
};

export const GetAccountInfo = (username, token) => {
  return sendRequest({
    endpoint: '/accounts/' + username,
    response: { ...AccountResponse },
    method: GET,
    headers: AuthHeader(token)
  });
};

const AccessResponse = {
  body: {
    access_token: '',
    refresh_token: ''
  },
  error: {
    message: '',
    status: 200
  }
};

export const LoginAccount = (username, pass) => {
  const data = {
    password: pass,
    username
  };
  return sendRequest({
    endpoint: '/accounts/login',
    response: { ...AccessResponse },
    method: POST,
    body: data
  });
};

export const Logout = (access) => {
  const data = {
    access_token: access
  };

  return sendRequest({
    endpoint: '/accounts/logout',
    method: POST,
    body: data
  });
};

const SignupResponse = {
  body: {
    username: '',
    password: '',
    email: '',
    name: ''
  },
  error: {
    message: '',
    status: 200
  }
};

export const CreateAccount = (username, password, name, email) => {
  const data = {
    username,
    password,
    email,
    name
  };

  return sendRequest({
    endpoint: '/accounts/signup',
    response: { ...SignupResponse },
    method: POST,
    body: data
  });
};

export const RefreshToken = () => {
  return sendRequest({
    endpoint: '/accounts/refresh',
    response: { ...AccessResponse },
    method: POST
  });
};
