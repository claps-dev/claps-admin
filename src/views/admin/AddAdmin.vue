<template>
    <v-card
            color="darkgray darken-1"
            dark
            onload="loadAllUsers()"
            min-height="70vh"
            max-width="80vw"
            style="margin: 20px auto;"
    >
        <v-img
                height="200"
                src="https://picsum.photos/1920/1080?random"
        >
            <v-row>
                <v-row
                        class="pa-4"
                        align="center"
                        justify="center"
                >
                    <v-col class="text-center">
                        <h3 class="headline">Add New Admin for admin.claps.dev</h3>
                    </v-col>
                </v-row>
            </v-row>
        </v-img>
        <v-form class="mt-16">
            <v-container>
                <v-autocomplete
                        v-model="admins"
                        :items="Users"
                        filled
                        chips
                        color="blue-grey lighten-2"
                        label="Project Members"
                        item-text="name"
                        item-value="name"
                        multiple
                >
                    <template v-slot:selection="data">
                        <v-chip
                                v-bind="data.attrs"
                                :input-value="data.selected"
                                close
                                @click="data.select"
                                @click:close="remove(data.item)"
                        >
                            <v-avatar left>
                                <v-img :src="data.item.avatar_url"></v-img>
                            </v-avatar>
                            {{ data.item.name }}
                        </v-chip>
                    </template>
                    <template v-slot:item="data">
                        <template v-if="typeof data.item !== 'object'">
                            <v-list-item-content v-text="data.item"></v-list-item-content>
                        </template>
                        <template v-else>
                            <v-list-item-avatar>
                                <img :src="data.item.avatar_url" alt="">
                            </v-list-item-avatar>
                            <v-list-item-content>
                                <v-list-item-title v-html="data.item.name"></v-list-item-title>
                                <v-list-item-subtitle v-html="data.item.email"></v-list-item-subtitle>
                            </v-list-item-content>
                        </template>
                    </template>
                </v-autocomplete>
            </v-container>
        </v-form>
        <v-divider></v-divider>
        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
                    class="pr-12 pl-12"
                    color="success"
                    @click="addNewAdmin"
            >
                SUBMIT
            </v-btn>
            <v-spacer></v-spacer>
        </v-card-actions>
        <v-snackbar
                v-model="hasSaved"
                :timeout="2000"
                absolute
                bottom
                left
        >
            Please waiting for a while...
        </v-snackbar>
        <v-snackbar
                v-model="notCompleted"
                :timeout="2000"
                absolute
                bottom
                left
        >
            Please choose an admin!
        </v-snackbar>
    </v-card>
</template>

<script>
    import userService from "../../service/userService";
    import adminService from "../../service/adminService";

    export default {
        name: "AddAdmin",
        data () {
            return {
                notCompleted: false,
                hasSaved: false,
                projectName: '',
                admins: [],
                Users: [
                    {
                        id: "",
                        name: "",
                        display_name: "",
                        email: "",
                        avatar_url: "",
                        mixin_id: "",
                    },
                ],
            }
        },
        mounted: function() {
            this.loadAllUsers();
        },
        methods: {
            remove (item) {
                const index = this.admins.indexOf(item.name);
                if (index >= 0) this.admins.splice(index, 1)
            },
            loadAllUsers () {
                this.projectName = this.$route.params.name;
                userService.getAllUsers().then((res) => {
                    this.Users = res.data.data.Users;
                }).catch((err) => {
                    alert("获取用户信息失败！" + err);
                });
            },
            addNewAdmin() {
                const linear = document.getElementById("progress-linear");
                linear.style.display="block";
                if (this.admins.length === 0) {
                    linear.style.display="none";
                    this.notCompleted = true;
                    return null;
                }

                adminService.addNewAdmins(this.admins).then((res) => {
                    alert(res.data.msg);
                    if (res.data.code === 200) {
                        linear.style.display="none";
                        this.hasSaved = true;
                        this.$router.go(-1);
                    }
                }).catch(() => {
                    linear.style.display="none";
                    alert("添加管理员失败！")
                });
            },
        },
    }
</script>

<style scoped>

</style>
