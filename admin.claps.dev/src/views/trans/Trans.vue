<template>
    <div>
        <v-card class="mx-auto" max-width="700" style="position: relative; top: 15vh;">
            <template>
                <v-container fluid>
                    <v-row>
                        <v-col
                                :cols="6"
                        >
                            <v-card @click="viewTransaction">
                                <v-img
                                        :src="'https://cdn.vuetifyjs.com/images/cards/house.jpg'"
                                        class="white--text align-end"
                                        gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
                                        height="200px"
                                >
                                    <v-card-title v-text="'Transaction Table'"></v-card-title>
                                </v-img>

                                <v-card-actions>
                                    <v-card-text class="text--primary">
                                        <div>Checking project members, add new projects or delete projects.</div>
                                    </v-card-text>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                        <v-col
                                :cols="6"
                        >
                            <v-card @click="viewTransfer">
                                <v-img
                                        :src="'https://i.picsum.photos/id/251/1920/1080.jpg?hmac=x90F-RX-Ki7tHJFBhuOimBr_znHKQY1Ui3JFu6q8sDE'"
                                        class="white--text align-end"
                                        gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
                                        height="200px"
                                >
                                    <v-card-title v-text="'Transfer Table'"></v-card-title>
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
        name: "Trans",
        data() {
            return {

            }
        },
        methods: {
            // 进入用户提现详情界面
            viewTransfer() {
                this.$router.push({ name: 'transfers' });
            },

            // 进入项目详细流水界面
            viewTransaction() {
                this.$router.push({ name: 'transactions' });
            },
        },
    }
</script>

<style scoped>

</style>
