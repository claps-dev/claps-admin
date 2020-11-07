// 本地缓存服务

const PREFIX = 'claps_';

// user 模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;
//let token = JSON.parse(localStorage.getItem(USER_INFO));


// 序列化存储
const set = (key, data) => {
    localStorage.setItem(key, JSON.stringify(data));
};

// 反序列化读取
const get = (key) => {
    return JSON.parse(localStorage.getItem(key));
};

export default {
    set,
    get,
    USER_TOKEN,
    USER_INFO,
};
