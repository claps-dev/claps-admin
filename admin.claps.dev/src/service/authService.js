import request from '../utils/loginRequest';


// 验证token
const tokenAuth = (token) => {
    return request.post('tokenAuth', {token});
};

export default {
    tokenAuth,
};
