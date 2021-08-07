<template>
    <div style="max-width: 80vw; margin: 40px auto" onload="updateTransactions()">
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
    import transService from "../../service/transService";
    import util from "../../utils/util";

    export default {
        name: "Transactions",
        data() {
            return {
                search: '',
                headers: [
                    {
                        text: 'Project Name',
                        align: 'start',
                        sortable: false,
                        value: 'project_name',
                    },
                    { text: 'Asset Id', value: 'asset_id' },
                    { text: 'Amount ($)', value: 'amount' },
                    { text: 'Sender', value: 'sender' },
                    { text: 'Receiver', value: 'receiver' },
                    { text: 'Created At', value: 'created_at' },
                ],
                Transactions: [{
                    amount: "",
                    asset_id: "",
                    asset: "",
                    created_at: "",
                    id: "",
                    project_id: "",
                    project_name: "",
                    receiver: "",
                    sender: "",
                }],
            };
        },

        mounted: function() {
            this.updateTransactions();
        },

        methods: {
            // 进入页面装载捐献记录
            updateTransactions() {
                // 获取项目捐赠流水记录
                transService.getAllTransactions().then((res) => {
                    if (res.data.code !== 200) {
                        alert(res.data.msg);
                    } else {
                        console.log(res.data.data.Transactions);
                        this.Transactions = res.data.data.Transactions;
                        for(let i = 0; i < this.Transactions.length; i++) {
                            this.Transactions[i].created_at = util.renderTime(this.Transactions[i].created_at);
                        }
                    }
                }).catch((err) => {
                    alert("获取捐赠流水时出错！" + err);
                });
            },
        },

    }
</script>

<style scoped>

</style>
