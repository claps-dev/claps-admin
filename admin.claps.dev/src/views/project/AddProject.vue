<template>
    <v-card
            class="overflow-hidden"
            color="darkgray lighten-1"
            dark
            onload="loadAllUsers()"
    >
        <v-toolbar
                flat
                color="darkgray"
        >
            <v-toolbar-title class="font-weight-light">NEW PROJECT PROFILE</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn
                    color="green"
                    fab
                    small
                    @click="isEditing = !isEditing"
            >
                <v-icon v-if="isEditing">mdi-close</v-icon>
                <v-icon v-else>mdi-pencil</v-icon>
            </v-btn>
        </v-toolbar>


        <v-form v-model="valid">
            <!-- input items -->
            <v-card-text>
                <v-text-field
                        :disabled="!isEditing"
                        color="white"
                        label="Project Name"
                        id="v-projectName"
                        autocomplete="off"

                        :rules="nameRules"
                        :counter="50"
                        value=""
                ></v-text-field>
                <v-text-field
                        :disabled="!isEditing"
                        color="white"
                        label="Project Description"
                        id="v-description"
                        autocomplete="off"

                        :rules="descriptionRules"
                        :counter="120"
                        value=""
                ></v-text-field>
                <v-text-field
                        :disabled="!isEditing"
                        color="white"
                        label="Project Avatar URL(optional)"
                        id="v-avatarUrl"
                        autocomplete="off"

                        :rules="avatarRules"
                        :counter="100"
                        value=""
                ></v-text-field>
                <v-text-field
                        :disabled="!isEditing"
                        color="white"
                        label="Repository Name"
                        id="v-repoName"
                        autocomplete="off"

                        :rules="repoNameRules"
                        :counter="50"
                        value=""
                ></v-text-field>
                <v-text-field
                        :disabled="!isEditing"
                        color="white"
                        label="Repository Url"
                        id="v-repoUrl"
                        autocomplete="off"

                        :rules="repoUrlRules"
                        :counter="100"
                        value=""
                ></v-text-field>
                <v-autocomplete
                        :disabled="!isEditing"
                        :items="repoTypes"
                        :filter="customFilter"
                        color="white"
                        item-text="name"
                        label="Repository Type"
                        id="v-repoType"

                        :rules="repoTypeRules"
                ></v-autocomplete>
                <v-autocomplete
                        :disabled="!isEditing"
                        :items="distributions"
                        :filter="customFilter"
                        color="white"
                        item-text="name"
                        label="Distribution"
                        id="v-distribution"

                        :rules="distributionRules"
                ></v-autocomplete>
                <v-autocomplete
                        :disabled="!isEditing"
                        v-model="ProjectInfo.members"
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
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                        :disabled="!isEditing"
                        color="success"
                        @click="addNewProject"
                >
                    SUBMIT
                </v-btn>
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
                Project profile isn't completed!
            </v-snackbar>
        </v-form>
    </v-card>
</template>

<script>
    import userService from "../../service/userService";
    import projectService from "../../service/projectService";
    import App from "../../App";

    export default {
        name: "AddProject",
        data () {
            return {
                hasSaved: false,
                notCompleted: false,
                isEditing: null,
                model: null,
                valid: false,
                nameRules: [
                    v => !!v || 'Project Name is required',
                    v => v.length <= 50 || 'Name must be less than 50 characters',
                ],
                repoNameRules: [
                    v => !!v || 'Repository Name is required',
                    v => v.length <= 50 || 'Name must be less than 50 characters',
                ],
                descriptionRules: [
                    v => !!v || 'Description is required',
                    v => v.length <= 120 || 'Description must be less than 120 characters',
                ],
                avatarRules: [
                    v => v.length <= 100 || 'Avatar Url must be less than 100 characters',
                ],
                repoUrlRules: [
                    v => !!v || "Repository URL is required",
                    //v => /.+http.+/.test(v) || 'Repository URL must be valid (Begin with \'http\' or \'https\')',
                    v => v.length <= 100 || 'Repository URL must be less than 100 characters',
                ],
                repoTypeRules: [
                    v => !!v || "Repository type is required",
                ],
                distributionRules: [
                    v => !!v || "Distribution type is required",
                ],

                distributions: [
                    { name: 'PersperAlgorithm', abbr: 'PA', id: 1 },
                    { name: 'Commits', abbr: 'C', id: 2 },
                    { name: 'ChangedLines', abbr: 'CL', id: 3 },
                    { name: 'IdenticalAmount', abbr: 'IA', id: 4 },
                ],
                repoTypes: [
                    { name: 'GITHUB', abbr: 'GITHUB', id: 1 },
                    { name: 'BITBUCKET', abbr: 'BITBUCKET', id: 2 },
                    { name: 'GITLAB', abbr: 'GITLAB', id: 3 },
                    { name: 'GIT', abbr: 'GIT', id: 4 },
                ],
                ProjectInfo: {
                    projectName: "",
                    avatarUrl: "",
                    description: "",
                    members: [],
                    repoName: "",
                    repoUrl: "",
                    repoType: "",
                    distribution: "",
                },
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
            loadAllUsers () {
                userService.getAllUsers().then((res) => {
                    this.Users = res.data.data.Users;
                    console.log(this.Users);
                }).catch((err) => {
                    alert("获取用户信息失败！" + err);
                });
            },
            customFilter (item, queryText) {
                const textOne = item.name.toLowerCase();
                const textTwo = item.abbr.toLowerCase();
                const searchText = queryText.toLowerCase();

                return textOne.indexOf(searchText) > -1 ||
                    textTwo.indexOf(searchText) > -1
            },
            addNewProject() {
                if (!this.valid) {
                    this.notCompleted = true;
                    return null;
                }
                this.isEditing = !this.isEditing;
                this.hasSaved = true;
                // 动画
                App.methods.progressStart();

                this.ProjectInfo.projectName = document.getElementById("v-projectName").value;
                this.ProjectInfo.description = document.getElementById("v-description").value;
                this.ProjectInfo.avatarUrl = document.getElementById("v-avatarUrl").value;
                this.ProjectInfo.repoType = document.getElementById("v-repoType").value;
                this.ProjectInfo.repoUrl = document.getElementById("v-repoUrl").value;
                this.ProjectInfo.repoName = document.getElementById("v-repoName").value;
                this.ProjectInfo.distribution = document.getElementById("v-distribution").value;
                projectService.sendNewProjectInfo(this.ProjectInfo).then((res) => {
                    App.methods.progressStop();
                    alert(res.data.msg);
                    if (res.data.code === 200) {
                        this.$router.push( {name: 'project'} );
                    }
                }).catch((err) => {
                    App.methods.progressStop();
                    alert("添加项目失败！" + err);
                });
            },
            remove (item) {
                const index = this.members.indexOf(item.name);
                if (index >= 0) this.members.splice(index, 1)
            },
        },
    }
</script>

<style scoped>

</style>
