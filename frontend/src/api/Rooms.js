import { AuthHeader, DELETE, GET, POST, sendRequest } from './Request';

const RoomResponse = {
  body: {
    host: '',
    room_code: ''
  },
  error: ''
};

export const GetRoom = (code) => {
  return sendRequest({
    endpoint: '/rooms/' + code,
    response: { ...RoomResponse },
    method: GET,
    onRespond: (resp, [ok, data]) => {
      if (ok) {
        resp.body = { ...data };
        return resp;
      }

      const errResp = { ...data };
      if (errResp.status === 404) {
        resp.error = errResp.message;
        return resp;
      }

      resp.error = 'Internal error, please try again.';
      return resp;
    }
  });
};

export const DeleteRoom = (code, token) => {
  return sendRequest({
    endpoint: '/rooms/' + code,
    headers: AuthHeader(token),
    method: DELETE
  });
};

const CreateRoomResponse = {
  body: {
    host: '',
    room_code: '',
    account_id: 1
  },
  error: {
    message: '',
    status: 201
  }
};

export const CreateTheRoom = (host, roomCode, token) => {
  const data = {
    host,
    room_code: roomCode,
    account_id: token.id
  };

  return sendRequest({
    endpoint: '/rooms',
    response: { ...CreateRoomResponse },
    method: POST,
    headers: AuthHeader(token),
    body: data
  });
};

const AllRoomResponse = {
  body: [
    {
      room_code: '',
      host: '',
      account_id: 0
    }
  ],
  error: {
    message: '',
    status: 200
  }
};

export const AllRoomsAssociated = (token) => {
  return sendRequest({
    endpoint: '/rooms/all',
    response: { ...AllRoomResponse },
    method: GET,
    headers: AuthHeader(token),
    onRespond: (res, [ok, data]) => {
      if (!ok) {
        res.error = { ...data };
      } else {
        res.body = data;
      }
      return res;
    }
  });
};
