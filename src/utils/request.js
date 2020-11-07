import axios from 'axios';
import storageService from "../service/storageService";
//import Vue from 'vue';

const service = axios.create({
    // api的url地址
    baseURL: process.env.VUE_APP_API_URL,
    // 持续时间2分钟
    timeout: 1000 * 120,
    // token验证请求的有效性
    headers: { "Authorization": storageService.get(storageService.USER_TOKEN)},
});

export default service
