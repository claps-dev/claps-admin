import request from '../utils/request';


// 获取管理员列表信息
const getAdminTable = () => {
    return request.get('admin/table');
};

// 发送要删除的管理员信息
const deleteAdmin = (email) => {
    return request.post('admin/delete', {email});
};

// 发送新创建的管理员信息
const addNewAdmins = (admins) => {
    return request.post('addAdmins', {admins});
};

export default {
    getAdminTable,
    deleteAdmin,
    addNewAdmins,
};
