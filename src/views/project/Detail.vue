<template>
    <div class="container-pd" onload="updateProjectDetail()">
        <!-- Navigation Bar -->
        <template>
            <div>
                <v-app-bar
                        color="darkgray"
                        src="https://picsum.photos/1920/1080?random"
                        dark
                        min-width="700"
                >
                    <!-- dialog -->
                    <v-dialog v-model="dialog" persistent max-width="600px">
                        <template v-slot:activator="{ on, attrs }">
                            <v-btn
                                    class="ma-4"
                                    icon
                                    v-bind="attrs"
                                    v-on="on"
                            >
                                <v-avatar v-if="Project.avatar_url">
                                    <img :src="Project.avatar_url" alt="">
                                </v-avatar>
                                <v-icon v-if="!Project.avatar_url">mdi-pencil</v-icon>
                            </v-btn>
                        </template>
                        <v-card>
                            <v-card-title>
                                <span class="headline">Editing Project Information</span>
                            </v-card-title>
                            <v-card-text>
                                <v-container>
                                    <v-row>
                                        <v-col cols="12" sm="10" md="6">
                                            <v-text-field label="Project Name(optional)" id="e-editedName" autocomplete="off"></v-text-field>
                                        </v-col>
                                        <v-col cols="12" sm="10">
                                            <v-text-field
                                                    label="Project Avatar URL(optional)"
                                                    id="e-avatarUrl"
                                                    autocomplete="off"
                                                    :disabled="defaultUrl"
                                            ></v-text-field>
                                        </v-col>
                                        <v-checkbox v-model="defaultUrl" :label="'Default'"></v-checkbox>
                                        <v-col cols="12">
                                            <v-text-field label="Project Description(optional)" id="e-description" autocomplete="off"></v-text-field>
                                        </v-col>
                                    </v-row>
                                </v-container>
                            </v-card-text>
                            <v-card-actions>
                                <v-spacer></v-spacer>
                                <v-btn color="blue darken-1" text @click="dialog = false">Close</v-btn>
                                <v-btn color="blue darken-1" text @click="editProjectInfo">Update</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-dialog>

                    <!--<v-list-item-avatar color="green darken-1" class="ma-3">
                        <img v-if="Project.avatar_url" :src="Project.avatar_url" alt="">
                        <span v-if="!Project.avatar_url" class="white&#45;&#45;text headline">{{ Project.name.charAt(0).toUpperCase() }}</span>
                    </v-list-item-avatar>-->

                    <v-toolbar-title>
                        <div v-text="this.projectName"></div>
                    </v-toolbar-title>

                    <v-spacer></v-spacer>

                    <!-- menu -->
                    <v-menu
                            bottom
                            origin="center center"
                            transition="scale-transition"
                            class="float-right"
                    >
                        <template v-slot:activator="{ on, attrs }">
                            <v-btn
                                    color="green"
                                    dark
                                    v-bind="attrs"
                                    v-on="on"
                            >
                                Manage Project
                            </v-btn>
                        </template>

                        <v-list>
                            <v-list-item @click.stop="addNewRepository">
                                <v-list-item-content>
                                    <v-list-item-title>Add Repository</v-list-item-title>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item @click.stop="addNewMember">
                                <v-list-item-content>
                                    <v-list-item-title>Add Member</v-list-item-title>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item @click.stop="deleteProject">
                                <v-list-item-content>
                                    <v-list-item-title class="red--text">Delete Project</v-list-item-title>
                                </v-list-item-content>
                            </v-list-item>

                        </v-list>
                    </v-menu>
                </v-app-bar>
            </div>
        </template>
        <br>
        <div style="float: left; max-width: 450px; display: inline-block; margin: 5px;">
            <!-- Member List -->
            <div>
                <template>
                    <v-card
                            min-width="450"
                            class="mx-auto float-left mr-6"
                            v-if="showMembers"

                    >
                        <v-list three-line>
                            <template v-for="(member, index) in Members">
                                <v-subheader
                                        v-if="index === 0"
                                        :key="index"
                                        v-text="'MEMBERS'"
                                ></v-subheader>

                                <v-divider
                                        v-if="index !== 0"
                                        :key="member.name"
                                        :inset="true"
                                ></v-divider>
                                <v-list-item
                                        :key="member.id"
                                >
                                    <v-list-item-avatar>
                                        <v-img :src="member.avatar_url"></v-img>
                                    </v-list-item-avatar>

                                    <v-list-item-content>
                                        <v-list-item-title v-html="member.name"></v-list-item-title>
                                        <v-list-item-subtitle v-html="member.email"></v-list-item-subtitle>
                                        <v-list-item-subtitle v-html="member.mixin_id"></v-list-item-subtitle>
                                    </v-list-item-content>
                                    <v-menu
                                            bottom
                                            origin="center center"
                                            transition="scale-transition"
                                            class="float-right"
                                    >
                                        <template v-slot:activator="{ on, attrs }">
                                            <v-btn
                                                    color="green"
                                                    text
                                                    dark
                                                    v-bind="attrs"
                                                    v-on="on"
                                            >
                                                MORE
                                            </v-btn>
                                        </template>

                                        <v-list>
                                            <v-list-item @click.stop="deleteMember" :id="member.id">
                                                <v-list-item-content>
                                                    <v-list-item-title class="red--text">Delete</v-list-item-title>
                                                </v-list-item-content>
                                            </v-list-item>

                                        </v-list>
                                    </v-menu>
                                </v-list-item>
                            </template>
                        </v-list>
                    </v-card>
                </template>
            </div>


            <!-- Repo Table -->
            <div>
                <template>
                    <v-card
                            min-width="450"
                            class="mx-auto float-left mr-6"
                            v-if="showRepos"
                            style="margin-top: 35px;"
                    >
                        <v-list three-line>
                            <template v-for="(repo, index) in Repos">
                                <v-subheader
                                        v-if="index === 0"
                                        :key="index"
                                        v-text="'REPOSITORY'"
                                ></v-subheader>

                                <v-divider
                                        v-if="index !== 0"
                                        :key="repo.id"
                                        :inset="true"
                                ></v-divider>
                                <v-list-item
                                        :key="repo.name"
                                >
                                    <!--<v-list-item-avatar v-if="this.Project.avatar_url">
                                        <v-img :src="this.Project.avatar_url"></v-img>
                                    </v-list-item-avatar>-->
                                    <v-list-item-content>
                                        <v-list-item-title class="blue--text" v-html="repo.name"></v-list-item-title>
                                        <v-list-item-subtitle v-html="repo.slug"></v-list-item-subtitle>
                                    </v-list-item-content>
                                    <v-menu
                                            bottom
                                            origin="center center"
                                            transition="scale-transition"
                                            class="float-right"
                                    >
                                        <template v-slot:activator="{ on, attrs }">
                                            <v-btn
                                                    color="green"
                                                    text
                                                    dark
                                                    v-bind="attrs"
                                                    v-on="on"
                                            >
                                                MORE
                                            </v-btn>
                                        </template>

                                        <v-list>
                                            <v-list-item @click.stop="deleteRepository" :id="repo.id">
                                                <v-list-item-content>
                                                    <v-list-item-title class="red--text">Delete</v-list-item-title>
                                                </v-list-item-content>
                                            </v-list-item>
                                        </v-list>
                                    </v-menu>
                                </v-list-item>
                            </template>
                        </v-list>
                    </v-card>
                </template>
            </div>

        </div>

        <!-- chart -->
        <v-card
                class="float-left text-center mt-1"
                min-width="700"
                min-height="500"
        >
            <v-sheet
                    class="v-sheet--offset mx-auto"
                    color="cyan"
                    elevation="12"
                    max-width="calc(100% - 32px)"
                    min-height="330"
            >
                <v-sparkline
                        :value="value"
                        :labels="labels"
                        color="white"
                        line-width="2"
                        padding="24"
                        height="180"
                        auto-draw
                        smooth
                        stroke-linecap="round"
                >
                    <template v-slot:label="item">
                        {{ item.value }}
                    </template>
                </v-sparkline>
            </v-sheet>

            <v-card-text class="pt-0">
                <div class="title font-weight-light mb-2">Project Transactions</div>
                <div class="subheading font-weight-light grey--text">Transactions of {{this.projectName}} in {{this.Month[this.currentMon]}}</div>
                <v-icon
                        class="mr-2"
                        small
                >
                    mdi-clock
                </v-icon>
                <span class="caption grey--text font-weight-light">
                    last update {{ time.getMonth()+1 }}-{{ time.getDate() }}-{{ time.getFullYear() }} {{time.getHours()}}:{{time.getMinutes()}}
                </span>
                <v-divider class="my-2"></v-divider>
                <v-card-actions class="justify-center">
                    <v-btn text color="green" class="float-left mr-12" id="last-m" @click="changeMonth"><v-icon>mdi-chevron-left</v-icon>Last month</v-btn>
                    <v-btn text @click="viewProjectTransactions">Go to Report</v-btn>
                    <v-btn text color="green" class="float-right ml-12" id="next-m" @click="changeMonth">Next month<v-icon>mdi-chevron-right</v-icon></v-btn>
                </v-card-actions>
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
    import projectService from "../../service/projectService";
    import memberService from "../../service/memberService";
    import repositoryService from "../../service/repositoryService";

    export default {
        name: "Detail",
        data() {
            return {
                defaultUrl: false,
                dialog: false,
                time: new Date(),
                showMembers: false,
                showRepos: false,
                isActive: false,
                projectName: '',
                currentMon: 0,
                Month: ["" ,"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
                Members: [
                    {
                        id: "",
                        name: "",
                        display_name: "",
                        email: "",
                        avatar_url: "",
                        mixin_id: "",
                    },
                ],
                Repos: [{
                    id: "",
                    type: "",
                    slug: "",
                    name: "",
                }],
                Project: {
                    id: "",
                    name: "",
                    description: "",
                    avatar_url: "",
                    total: "",
                },
                EditedProjectInfo: [{
                    editedName: "",
                    description: "",
                    avatarUrl: "",
                }],
                Transactions: {
                    id: "",
                    project_id: "",
                    asset_id: "",
                    amount: "",
                    created_at: "",
                    sender: "",
                    receiver: "",
                },
                labels: [
                    '*', '*', '*', '*', '5', '*', '*', '*', '*', '10',
                    '*', '*', '*', '*', '15', '*', '*', '*', '*', '20',
                    '*', '*', '*', '*', '25', '*', '*', '*', '*', '30', '*',
                ],
                TransMonthly: [
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                    { amounts: [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0],},
                ],
                value: [],
            };
        },
        mounted: function() {
            this.updateProjectDetail();
        },
        methods: {
            // 进入页面自动更新项目的成员、仓库和钱包
            updateProjectDetail() {
                // 获取项目id
                this.projectName = this.$route.params.name;
                // 由项目id获取项目
                projectService.getProjectByName(this.projectName).then((res) => {
                    if (res.data.code !== 200) {
                        alert(res.data.msg);
                    } else {
                        this.Project = res.data.data.Project;
                        this.Transactions = res.data.data.Transactions;
                        this.calculateTransMonthly(this.Transactions);
                        this.currentMon = this.time.getMonth() + 1;
                        this.value = this.TransMonthly[this.currentMon].amounts;
                    }
                }).catch((err) => {
                    alert("获取项目时出错！" + err);
                });

                // 获取项目成员表单
                memberService.getProjectMember(this.projectName).then((res) => {
                    if (res.data.code !== 200) {
                        return null;
                    } else {
                        this.Members = res.data.data.Members;
                        this.showMembers = true;
                    }
                }).catch((err) => {
                    alert("获取成员表单时出错！" + err);
                });

                // 获取项目仓库表单
                repositoryService.getProjectRepo(this.projectName).then((res) => {
                    if (res.data.code !== 200) {
                        return null;
                    } else {
                        this.Repos = res.data.data.Repos;
                        this.showRepos = true;
                    }
                }).catch((err) => {
                    alert("获取仓库表单时出错！" + err);
                });
            },

            // 删除当前项目
            deleteProject() {
                // 确认提示
                let r = confirm("确定要删除吗？");
                if (!r) return;

                const projectName = this.projectName;
                // 由项目id删除项目
                projectService.deleteProject(projectName).then((res) => {
                    alert(res.data.msg);
                    if (res.data.code === 200) {
                        this.$router.replace({ name: 'project' })
                    }
                }).catch((err) => {
                    alert("删除项目时出错！" + err);
                });
            },

            // 添加成员
            addNewMember() {
                this.$router.push({ name: 'addMember', params: {'name': this.projectName} });
            },

            // 删除成员
            deleteMember(event) {
                // 确认提示
                let r = confirm("确定要删除吗？");
                if (!r) return;

                // 获取项目id
                const projectName = this.projectName;
                // 获取成员id
                const userId = event.currentTarget.id;
                memberService.deleteProjectMember(projectName, userId).then((res) => {
                    alert(res.data.msg);
                    if (res.data.code === 200) {
                        location.reload();
                    }
                }).catch((err) => {
                    alert("删除成员时出错！" + err)
                });
            },

            // 添加项目
            addNewRepository() {
                this.$router.push({ name: 'addRepository', params: {'name': this.projectName} });
            },

            // 删除仓库
            deleteRepository(event) {
                // 确认提示
                let r = confirm("确定要删除吗？");
                if (!r) return;

                // 获取项目id
                const repoId = event.currentTarget.id;
                repositoryService.deleteProjectRepo(repoId).then((res) => {
                    alert(res.data.msg);
                    if (res.data.code === 200) {
                        location.reload();
                    }
                }).catch((err) => {
                    alert("删除仓库时出错！" + err)
                });
            },

            // 为图表计算每月受捐记录
            calculateTransMonthly(transactions) {
                  for (let i = 0; i < transactions.length; i++) {
                      let dateStr = transactions[i].created_at;
                      let month = parseInt(dateStr.substring(5, 7));
                      let date = parseInt(dateStr.substring(8, 10));
                      this.TransMonthly[month].amounts[Math.floor(date-1)] += parseFloat(transactions[i].amount);
                      //console.log(this.TransMonthly[month-1].amounts[0])
                  }
            },

            // 查看捐赠流水
            viewProjectTransactions() {
                this.$router.push({ name: 'projectTransaction', params: {'name': this.projectName} });
            },

            // 改变月份
            changeMonth(event) {
                const t = event.currentTarget.id;
                let m = parseInt(this.currentMon);
                if (t === 'last-m') {
                    m - 1 < 1 ? m = 12 : m = m - 1;
                }
                else {
                    m + 1 > 12 ? m = 1 : m = m + 1;
                }
                this.value = this.TransMonthly[m].amounts;
                this.currentMon = m;
            },

            // 编辑项目信息并更新
            editProjectInfo() {
                this.EditedProjectInfo.editedName = document.getElementById("e-editedName").value;
                // 判断任务图标是否默认
                if (this.defaultUrl) {
                    this.EditedProjectInfo.avatarUrl = "defaultUrl";
                } else {
                    this.EditedProjectInfo.avatarUrl = document.getElementById("e-avatarUrl").value;
                }

                this.EditedProjectInfo.description = document.getElementById("e-description").value;
                projectService.sendEditedProjectInfo(this.projectName, this.EditedProjectInfo).then((res) => {
                    if (res.data.code === 200) {
                        if (this.EditedProjectInfo.editedName.length !== 0) {
                            this.$router.push({ name: 'project/detail', params: {'name': this.EditedProjectInfo.editedName}});
                        }
                        location.reload();
                        this.dialog = false;
                    }
                }).catch((err) => {
                    alert("编辑项目失败！" + err);
                })

            },
        },
    }
</script>

<style scoped>
</style>
