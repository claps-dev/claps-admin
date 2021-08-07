import request from '../utils/request';

// 为项目添加成员
const addMembers = (projectName, members) => {
    return request.post('project/' + projectName + '/addMembers', {members})
};

// 请求项目成员信息
const getProjectMember = (projectName) => {
    return request.get('project/' + projectName + '/member');
};

// 根据成员id删除成员
const deleteProjectMember = (projectName, userId) => {
    return request.post('project/' + projectName + '/memberDelete', {userId})
};

export default {
    addMembers,
    getProjectMember,
    deleteProjectMember,
};
