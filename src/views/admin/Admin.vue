<template>
    <div>
        <template>
            <div>
                <v-app-bar
                        color="darkgray"
                        src="https://picsum.photos/1920/1080?random"
                        dark
                >
                    <v-app-bar-nav-icon>
                    </v-app-bar-nav-icon>
                    <v-toolbar-title>
                        <div v-text="' Admin management '"></div>
                    </v-toolbar-title>
                    <v-spacer></v-spacer>
                    <v-btn
                            color="green"
                            @click="addAdmin"
                    >
                        ADD ADMIN
                    </v-btn>
                </v-app-bar>
            </div>
        </template>
        <v-card style="max-width: 80vw; margin: 40px auto" onload="updateAdminTable()">
            <v-list three-line>
                <v-list-item
                        v-for="item in Admins"
                        :key="item.user_id"
                >
                    <v-list-item-avatar>
                        <v-img :src="item.avatar_url"></v-img>
                    </v-list-item-avatar>
                    <v-list-item-content>
                        <v-list-item-title v-html="item.name">
                        </v-list-item-title>
                        <v-list-item-subtitle v-html="item.email + ' / ' + item.role"></v-list-item-subtitle>
                    </v-list-item-content>
                    <v-btn
                            v-if="item.role.charAt(0) !== 's'"
                            color="red"
                            text
                            dark
                            @click="deleteAdmin"
                            :id="item.email"
                    >
                        DELETE
                    </v-btn>
                </v-list-item>
            </v-list>
        </v-card>
        <v-snackbar
                v-model="snackbar"
                :timeout="2000"
        >
            管理员删除成功！
        </v-snackbar>
    </div>

</template>

<script>
    import adminService from "../../service/adminService";

    export default {
        name: "Admin",
        data() {
            return {
                snackbar: false,
                Admins: [{
                    user_id: "",
                    name: "",
                    email: "",
                    role: "",
                    avatar_url: "",
                    mixin_id: "",
                }],
            };
        },
        mounted: function() {
            this.updateAdminTable();
        },
        methods: {
            // 进入页面更新表单
            updateAdminTable() {
                adminService.getAdminTable().then((res) => {
                    this.Admins = res.data.data.Admins;
                }).catch((err) => {
                    alert("获取表单时出错了：" + err);
                });
            },

            // 删除一个管理员，同时更新数据库
            deleteAdmin(event) {
                let r = confirm("确定要删除吗？");
                if (!r) return;

                // 获取要删除管理员的邮箱
                const email = event.currentTarget.id;
                adminService.deleteAdmin(email).then((res) => {
                    if (res.data.code === 200) {
                        location.reload();
                        this.snackbar = true;
                        return null;
                    }
                    alert(res.data.msg);
                }).catch((err) => {
                    alert(err);
                });
            },

            // 添加一个管理员
            addAdmin() {
                this.$router.push({ name: 'addAdmin' });
            },
        },
    };
</script>

<style scoped>

</style>
