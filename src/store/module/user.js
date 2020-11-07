import storageService from "../../service/storageService";
import loginService from "../../service/loginService";

// 存储用户状态和信息
const userModule = {
    // 开启namespaced, 加上命名空间
    namespaced: true,
    // 定义状态
    state: {
        token: storageService.get(storageService.USER_TOKEN),
        userInfo: storageService.get(storageService.USER_INFO),
    },

    // 定义方法
    mutations: {
        // 当前状态+参数, 且状态变更只能在mutations里面进行，需要启动严格模式
        SET_TOKEN(state, token) {
            // 更新本地缓存
            storageService.set(storageService.USER_TOKEN, token);
            // 更新state
            state.token = token;
        },

        SET_USER_INFO(state, userInfo) {
            // 更新本地缓存
            storageService.set(storageService.USER_INFO, userInfo);
            // 更新state
            state.userInfo = userInfo;
        },
    },

    actions: {
        // 存储用户信息和Token
        storeUserInfo(context, {User, Token}) {
            return new Promise(() => {
                // 结构体序列化
                JSON.stringify(User);
                context.commit('SET_USER_INFO', User);
                context.commit('SET_TOKEN', Token);
            });
        },

        login(context, {Code}) { // 链式调用，传送code->检查token->存储用户信息
            return new Promise((resolve, reject) => {
                loginService.sendCode( {Code} ).then((res) => {
                    if (res.data.code !== 200) {
                        alert(res.data.msg);
                        return null;
                    }
                    console.log(res.data);
                    return loginService.getLoginInfo();
                }).then((res) => {
                    // 检查token是否存在
                    if (!res.data.data.Token) {
                        alert("Token不存在！");
                        return null;
                    }
                    // 检查token是否失效

                    // 存储用户名称和账号, (json序列化，能以字符串显示出来)
                    context.commit('SET_USER_INFO', res.data.data.User);
                    // 存储token
                    context.commit('SET_TOKEN', res.data.data.Token);
                    resolve(res);
                }).catch((err) => {
                    console.log("页面出错了！" + err.toString());
                    reject(err);
                });
            });
        },

        // 登出清理本地用户信息缓存
        logout(context) {
            return new Promise(() => {
                context.commit('SET_USER_INFO', "");
                context.commit('SET_TOKEN', "");
            });
        }
    },
};

export default userModule
