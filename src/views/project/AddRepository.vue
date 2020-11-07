<template>
    <v-card
            color="darkgray darken-1"
            dark
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
                        <h3 class="headline">{{ this.$route.params.name }}</h3>
                        <span class="grey--text text--lighten-1">Add new repository for {{ this.$route.params.name }}</span>
                    </v-col>
                </v-row>
            </v-row>
        </v-img>
        <v-form class="mt-16">
            <v-container>
                <v-form v-model="valid">
                    <v-card-text>
                        <v-text-field
                                color="white"
                                label="Repository Name"
                                id="repoName"
                                autocomplete="off"

                                :rules="repoNameRules"
                                :counter="50"
                        ></v-text-field>
                        <v-text-field
                                color="white"
                                label="Repository Url"
                                id="repoUrl"
                                autocomplete="off"

                                :rules="repoUrlRules"
                                :counter="100"
                        ></v-text-field>
                        <v-text-field
                                color="white"
                                label="Repository Description"
                                id="repoDescription"
                                autocomplete="off"

                                :rules="descriptionRules"
                                :counter="120"
                        ></v-text-field>
                        <v-autocomplete
                                :items="repoTypes"
                                :filter="customFilter"
                                color="white"
                                item-text="name"
                                label="Repository Type"
                                id="repoType"

                                :rules="repoTypeRules"
                        ></v-autocomplete>
                    </v-card-text>
                </v-form>
            </v-container>
        </v-form>
        <v-divider></v-divider>
        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
                    class="pr-12 pl-12"
                    color="success"
                    @click="addRepository"
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
            Project profile isn't completed!
        </v-snackbar>
    </v-card>
</template>

<script>
    import repositoryService from "../../service/repositoryService";

    export default {
        name: "AddRepository",
        data () {
            return {
                valid: false,
                notCompleted: false,
                hasSaved: false,
                projectName: '',
                members: [],
                repoNameRules: [
                    v => !!v || "Repository Name is required",
                    v => v.length <= 50 || 'Repository name must be less than 50 characters',
                ],
                repoTypeRules: [
                    v => !!v || "Repository type is required",
                ],
                descriptionRules: [
                    v => v.length <= 120 || 'Description must be less than 120 characters',
                ],
                repoUrlRules: [
                    v => !!v || "Repository URL is required",
                    //v => /.+http.+/.test(v) || 'Repository URL must be valid (Begin with \'http\' or \'https\')',
                    v => v.length <= 100 || 'Repository URL must be less than 100 characters',
                ],
                repoTypes: [
                    { name: 'GITHUB', abbr: 'GITHUB', id: 1 },
                    { name: 'BITBUCKET', abbr: 'BITBUCKET', id: 2 },
                    { name: 'GITLAB', abbr: 'GITLAB', id: 3 },
                    { name: 'GIT', abbr: 'GIT', id: 4 },
                ],
                Repository: {
                    description: "",
                    repoName: "",
                    repoType: "",
                    repoUrl: "",
                },
            }
        },
        methods: {
            customFilter (item, queryText) {
                const textOne = item.name.toLowerCase();
                const textTwo = item.abbr.toLowerCase();
                const searchText = queryText.toLowerCase();

                return textOne.indexOf(searchText) > -1 ||
                    textTwo.indexOf(searchText) > -1
            },
            addRepository() {
                if (!this.valid) {
                    this.notCompleted = true;
                    return null;
                }
                this.projectName = this.$route.params.name;
                // 动画
                const linear = document.getElementById("progress-linear");
                linear.style.display="block";
                this.hasSaved = true;
                this.Repository.repoName = document.getElementById("repoName").value;
                this.Repository.description = document.getElementById("repoDescription").value;
                this.Repository.repoType = document.getElementById("repoType").value;
                this.Repository.repoUrl = document.getElementById("repoUrl").value;
                repositoryService.addRepository(this.projectName, this.Repository).then((res) => {
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
