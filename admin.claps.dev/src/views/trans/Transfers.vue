<template>
    <div style="max-width: 80vw; margin: 40px auto" onload="updateTransfers()">
        <v-card>
            <v-card-title>
                Transfers
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
                    :items="Transfers"
                    :search="search"
            ></v-data-table>
        </v-card>
    </div>
</template>

<script>
    import transService from "../../service/transService";
    import util from "../../utils/util";

    export default {
        name: "Transfers",
        data() {
            return {
                search: '',
                headers: [
                    {
                        text: 'Mixin ID',
                        align: 'start',
                        sortable: false,
                        value: 'mixin_id',
                    },
                    { text: 'Asset ID', value: 'asset_id' },
                    { text: 'Amount ($)', value: 'amount' },
                    { text: 'Trace ID', value: 'trace_id' },
                    { text: 'Memo', value: 'memo' },
                    { text: 'Snapshot ID', value: 'snapshot_id' },
                    { text: 'Created At', value: 'created_at' },
                ],
                Transfers: [{
                    bot_id: "",
                    snapshot_id: "",
                    mixin_id: "",
                    trace_id: "",
                    asset_id: "",
                    amount: "",
                    memo: "",
                    created_at: "",
                }],
            };
        },

        mounted: function() {
            this.updateTransfers();
        },

        methods: {
            // 进入页面装载捐献记录
            updateTransfers() {
                // 获取项目捐赠流水记录
                transService.getAllTransfers().then((res) => {
                    if (res.data.code !== 200) {
                        alert(res.data.msg);
                    } else {
                        console.log(res.data.data.Transfers);
                        this.Transfers = res.data.data.Transfers;
                        for(let i = 0; i < this.Transfers.length; i++) {
                            this.Transfers[i].created_at = util.renderTime(this.Transfers[i].created_at);
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
