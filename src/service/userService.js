import request from '../utils/request';

// 根据用户名字获取某个用户
const getUserByName = (userName) => {
    return request.get('user/' + userName + '/info' );
};

// 获取所有用户信息
const getAllUsers = () => {
    return request.get('users/all');
};

export default {
    getAllUsers,
    getUserByName,
};
