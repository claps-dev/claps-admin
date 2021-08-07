<template>
    <div class="container-pd" onload="updateProjectTrans()">
        <!-- Navigation Bar -->
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
                        <div v-text="this.projectName"></div>
                    </v-toolbar-title>

                    <v-spacer></v-spacer>
                </v-app-bar>
            </div>
        </template>
        <!-- Table -->
        <v-card>
            <v-card-title>
                Transactions
                <v-spacer></v-spacer>
                <v-text-field
                        v-model="search"
                        append-icon="mdi-magnify"
                        label="Search"
                        single-line
                        hide-details
                ></v-text-field>
            </v-card-title>
            <v-data-table
                    :headers="headers"
                    :items="Transactions"
                    :search="search"
            ></v-data-table>
        </v-card>
    </div>
</template>

<script>
    import projectService from "../../service/projectService";
    import util from "../../utils/util";

    export default {
        name: "ProjectTransaction",
        data() {
            return {
                search: '',
                headers: [
                    {
                        text: 'Project ID',
                        align: 'start',
                        sortable: false,
                        value: 'project_id',
                    },
                    { text: 'Asset Id', value: 'asset_id' },
                    { text: 'Amount ($)', value: 'amount' },
                    { text: 'Sender', value: 'sender' },
                    { text: 'Receiver', value: 'receiver' },
                    { text: 'Created At', value: 'created_at' },
                ],
                Transactions: [{
                    id: "",
                    project_id: "",
                    asset_id: "",
                    amount: "",
                    created_at: "",
                    sender: "",
                    receiver: "",
                }],
                projectName: "",
                Project: {
                    id: "",
                    name: "",
                    description: "",
                    avatar_url: "",
                    total: "",
                },
            }
        },
        mounted: function() {
            this.updateProjectTrans();
        },
        methods: {
            updateProjectTrans() {
                this.projectName = this.$route.params.name;
                // 由项目id获取项目
                projectService.getProjectByName(this.projectName).then((res) => {
                    if (res.data.code !== 200) {
                        alert("获取项目失败！");
                    } else {
                        this.Project = res.data.data.Project;
                        this.Transactions = res.data.data.Transactions;
                        for(let i = 0; i < this.Transactions.length; i++) {
                            this.Transactions[i].created_at = util.renderTime(this.Transactions[i].created_at);
                        }
                    }
                }).catch((err) => {
                    alert("获取项目时出错！" + err);
                });

            },
        }
    }
</script>

<style scoped>

</style>
