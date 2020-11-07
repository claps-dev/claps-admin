import request from '../utils/request';

// 为项目添加仓库
const addRepository = (projectName, {description, repoName, repoType, repoUrl}) => {
    return request.post('project/' + projectName + '/addRepository', {description, repoName, repoType, repoUrl})
};

// 请求项目仓库信息
const getProjectRepo = (projectName) => {
    return request.get('project/' + projectName + '/repo');
};

const deleteProjectRepo = (repoId) => {
    return request.post('projects/repoDelete', {repoId})
};

export default {
    addRepository,
    getProjectRepo,
    deleteProjectRepo,
};
