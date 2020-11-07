<template>
    <div style="min-height: 99vh; background-image: url('https://picsum.photos/1920/1080');">
        <v-card class="mx-auto" max-width="500" style="position: relative; top: 10px;">
            <template>
                <v-container fluid>
                    <v-row>
                        <v-col
                                :cols="12"
                        >
                            <v-card @click="manageProject">
                                <v-img
                                        :src="'https://i.picsum.photos/id/692/1920/1080.jpg?hmac=Oq1BbdEMm2D6VOmB_sNxcxilNNFjY3cje_xZTn5onTk'"
                                        class="white--text align-end"
                                        gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
                                        height="200px"
                                >
                                    <v-card-title v-text="'Manage Projects'"></v-card-title>
                                </v-img>

                                <v-card-actions>
                                    <v-card-text class="text--primary">
                                        <div>Checking project members, add new projects or delete projects,
                                            you can learn which project is favoured most by others.</div>
                                    </v-card-text>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                        <v-col
                                :cols="6"
                        >
                            <v-card @click="viewTrans">
                                <v-img
                                        :src="'https://cdn.vuetifyjs.com/images/carousel/sky.jpg'"
                                        class="white--text align-end"
                                        gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
                                        height="200px"
                                >
                                    <v-card-title v-text="'Transactions & transfers'"></v-card-title>
                                </v-img>

                                <v-card-actions>
                                    <v-card-text class="text--primary">
                                        <div>Show recent bills of each project and transfers of users</div>
                                    </v-card-text>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                        <v-col
                                :cols="6"
                        >
                            <v-card @click="manageAdmin">
                                <v-img
                                        :src="'https://i.picsum.photos/id/197/1920/1080.jpg?hmac=5cTGvM5pMaQ9SrKjHQC7Y8Et1XC51pZM9x_IRfVZSlo'"
                                        class="white--text align-end"
                                        gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
                                        height="200px"
                                >
                                    <v-card-title v-text="'Manage Admins'"></v-card-title>
                                </v-img>

                                <v-card-actions>
                                    <v-card-text class="text--primary">
                                        <div>super administrator has all authority including delete other administrator.</div>
                                    </v-card-text>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                    </v-row>
                </v-container>
            </template>
        </v-card>
    </div>
</template>

<script>
    export default {
        name: "Home",
        data () {
            return {
                //active: false,
            }
        },
        // 在vuex中获取状态
        computed: {
            userInfo() {
                return this.$store.state.userModule.userInfo;
            },
        },
        mounted: function() {
            this.stopAnimation();
        },
        methods: {
            // 登出
            logout() {
                // 跳转登录页
                this.$router.replace({ name: 'login' });
                // 清除用户信息
                this.$store.dispatch('userModule/logout').then(() => {
                }).catch((err) => {
                    console.log("清除用户缓存失败，错误信息：", err);
                });
            },
            // 查看管理员列表
            manageAdmin() {
                // 跳转
                this.$router.push({ name: 'admin' });
            },
            // 查看项目列表
            manageProject() {
                this.$router.push({ name: 'project' });
            },
            // 查看捐赠流水
            viewTrans() {
                this.$router.push({ name: 'trans' });
            },
            // 停止动画
            stopAnimation() {
                const linear = document.getElementById("progress-linear");
                linear.style.visibility="hidden";
            },
        },
    };
</script>

<style scoped>
</style>
