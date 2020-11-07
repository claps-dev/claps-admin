<template>
    <div onload="checkCodeExist()" style="min-height: 99vh; background-image: url('https://picsum.photos/1920/1080');">
        <v-alert
                dismissible
                type="warning"
                v-model="tipAble"
                transition="slide-y-transition"
        >
            An error occurred! Please try again later.
        </v-alert>
        <v-card
                class="mx-auto my-auto"
                max-width="400"
                min-height="500"
                style="position: relative; top: 15vh;"
        >
            <v-img
                    class="white--text align-end"
                    height="270px"
                    color="darkgray"
                    src="https://i.picsum.photos/id/683/1920/1080.jpg?hmac=EewMJbxDanIFuo8ZJCRiasurjxJAfION35RjWUKqIwM"
            >
                <v-card-title>Welcome to admin.claps.dev!</v-card-title>
            </v-img>

            <v-card-subtitle class="pb-0">Sign in with your Github account.</v-card-subtitle>

            <v-card-text class="text--primary">
                <div>Click the button below,</div>

                <div>you will be redirected to github authorization page</div>
            </v-card-text>

            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                        color="green"
                        class="py-5 pr-13 pl-13 mt-8"
                        id="loginBtn"
                        style="position: relative; bottom: 10px; height: 50px"
                        @click="thirdPartyAuth()"
                        text
                        :disabled="btnDisabled"
                >
                    LOGIN WITH GITHUB
                </v-btn>

                <v-spacer></v-spacer>
            </v-card-actions>
        </v-card>
    </div>
</template>

<script>
    import loginService from "../../service/loginService";
    import store from '../../store';
    export default {
        name: "Login",

        data() {
            return {
                tipAble: false,
                btnDisabled: false,
                AuthorizeUrl: "",
                Token: "",
                User: {
                    Id: "",
                    Name: "",
                    Account: "",
                    Role: "",
                    AvatarUrl: "",
                },
            };
        },

        mounted: function() {
            this.checkCodeExist();
        },

        methods: {

            // 请求登录，转向github授权
            thirdPartyAuth() {
                // 信息存在则直接登录
                if (store.state.userModule.token) {
                    this.$router.replace({ name: 'home' });
                    return null;
                }

                // 加载动画
                this.btnDisabled = true;
                const linear = document.getElementById("progress-linear");
                linear.style.display="block";

                // 请求api
                loginService.getThirdAuthUrl().then((res) => {
                    // 跳转到github登录界面
                    window.location.href = res.data.data.AuthorizeUrl;
                    // 恢复按钮
                    linear.style.display="none";
                    this.btnDisabled = false;
                }).catch(() => {
                    linear.style.display="none";
                    this.tipAble = true;
                    this.btnDisabled = false;
                });
            },

            checkCodeExist() {
                const Code = window.location.search;
                if (Code === "") return null;

                // 加载动画
                const linear = document.getElementById("progress-linear");
                linear.style.display="block";
                // 设置按钮不可点击
                this.btnDisabled = true;

                // 发送code并验证用户身份
                loginService.sendCode( {Code} ).then((res) => {
                    if (res.data.code !== 200) {
                        this.tipAble = true;
                        return null;
                    }
                    // 获取用户信息
                    return loginService.getLoginInfo();
                }).then((res) => {
                    if (res.data.code !== 200) {
                        alert("登录失败！请稍后重试...");
                        return null;
                    }
                    let User = res.data.data.User;
                    let Token = res.data.data.Token;
                    // 本地储存用户信息和Token
                    this.$store.dispatch('userModule/storeUserInfo', {User, Token});
                    // 恢复按钮
                    linear.style.display="none";
                    this.btnDisabled = false;
                    return this;
                }).then(() => {
                    this.$router.replace({ name: 'home' });
                }).catch(() => {
                    // 恢复按钮
                    linear.style.display="none";
                    this.btnDisabled = false;
                });
            },
        },
    };
</script>

<style scoped>
</style>
