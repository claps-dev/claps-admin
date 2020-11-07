<template>
    <v-card
            color="darkgray darken-1"
            dark
            onload="loadAllUsers()"
            min-height="80vh"
    >
        <template v-slot:progress>
            <v-progress-linear
                    absolute
                    color="green lighten-3"
                    height="4"
                    indeterminate
            ></v-progress-linear>
        </template>
        <v-img
                height="220"
                src="https://picsum.photos/1920/1080?random"
        >
            <v-row>
                <v-row
                        class="pa-4"
                        align="center"
                        justify="center"
                >
                    <v-col class="text-center">
                        <h3 class="headline">{{ projectName }}</h3>
                        <span class="grey--text text--lighten-1">Add new member for {{ projectName }}</span>
                    </v-col>
                </v-row>
            </v-row>
        </v-img>
        <v-form class="mt-16">
            <v-container>
                <v-autocomplete
                        v-model="members"
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
                    @click="addMember"
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
    </v-card>
</template>

<script>
    import userService from "../../service/userService";
    import memberService from "../../service/memberService";

    export default {
        name: "AddMember",
        data () {
            return {
                hasSaved: false,
                projectName: '',
                members: [],
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
                const index = this.members.indexOf(item.name);
                if (index >= 0) this.members.splice(index, 1)
            },
            loadAllUsers () {
                this.projectName = this.$route.params.name;
                userService.getAllUsers().then((res) => {
                    this.Users = res.data.data.Users;
                }).catch((err) => {
                    alert("获取用户信息失败！" + err);
                });
            },
            addMember() {
                // 动画
                const linear = document.getElementById("progress-linear");
                linear.style.display="block";
                this.hasSaved = true;
                memberService.addMembers(this.projectName, this.members).then((res) => {
                    linear.style.display="none";
                    alert(res.data.msg);
                    if (res.data.code === 200) {
                        this.$router.go(-1);
                    }
                }).catch((err) => {
                    linear.style.display="none";
                    alert(err);
                })
            },
        },
    }
</script>

<style scoped>

</style>
