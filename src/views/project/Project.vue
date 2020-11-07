<template>
    <v-container fluid onload="updateProjectTable()">
        <v-card class="pl-5 pr-5">
            <!-- project header with search and sort-->
            <v-card-title>
                <v-btn icon class="mr-4" @click="expandAll" large color="green">
                    <v-icon >mdi-bookmark</v-icon>
                </v-btn>

                PROJECTS
                <v-btn
                        color="green"
                        text
                        @click="addNewProject"
                >
                    NEW PROJECT
                </v-btn>
                <v-spacer></v-spacer>

                <!-- Sort -->
                <template>
                    <v-spacer></v-spacer>
                    <v-select
                            v-model="sortBy"
                            hide-details
                            :items="keys"
                            label="Sort by"
                            color="darkgray"
                    ></v-select>
                    <v-btn-toggle
                            v-model="sortDesc"
                            mandatory
                    >
                        <v-btn
                                text
                                color="darkgray"
                                :value="false"
                        >
                            <v-icon>mdi-arrow-up</v-icon>
                        </v-btn>
                        <v-btn
                                text
                                color="darkgray"
                                :value="true"
                        >
                            <v-icon>mdi-arrow-down</v-icon>
                        </v-btn>
                    </v-btn-toggle>
                </template>
                <v-spacer></v-spacer>
                <!-- Search -->
                <v-text-field
                        id="project_search"
                        type="search"
                        append-icon="mdi-magnify"
                        label="Search"
                        single-line
                        hide-details

                        v-model="search"

                ></v-text-field>
            </v-card-title>
            <!-- content -->
            <v-data-iterator
                    :items="Projects"
                    :items-per-page.sync="itemsPerPage"
                    :page="page"
                    :search="search"
                    :sort-by="sortBy.toLowerCase()"
                    :sort-desc="sortDesc"
                    hide-default-footer
            >
                <!-- Projects Info cards -->
                <template v-slot:default="props">
                    <v-row v-if="showCards">
                        <v-col
                                v-for="(item, index) in props.items"
                                :key="index"
                                cols="12"
                                sm="6"
                                md="4"
                                lg="3"
                        >
                            <v-card>
                                <v-img
                                        src="https://picsum.photos/1920/1080?random"
                                        height="150px"
                                ></v-img>
                                <v-list three-line>
                                    <v-list-item>
                                        <v-list-item-avatar color="green darken-1">
                                            <img v-if="item.avatar_url" :src="item.avatar_url" alt="">
                                            <span v-if="!item.avatar_url" class="white--text headline">{{ item.name.charAt(0).toUpperCase() }}</span>
                                        </v-list-item-avatar>
                                        <v-list-item-content>
                                            <v-list-item-title>
                                                <div v-text="item.name"></div>
                                            </v-list-item-title>

                                        </v-list-item-content>
                                    </v-list-item>
                                </v-list>

                                <!-- Project Menu -->
                                <v-card-actions>
                                    <v-btn
                                            color="green"
                                            text
                                            @click="editProject"
                                            :id="item.name"
                                    >
                                        Manage Project
                                    </v-btn>
                                    <v-spacer></v-spacer>
                                    <v-btn
                                            icon
                                            @click=" show === index ? show = -1 : show = index "
                                    >
                                        <v-icon>{{ show === index || show === -2 ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
                                    </v-btn>
                                </v-card-actions>
                                <!-- Project Expansion -->
                                <v-expand-transition>
                                    <div v-show="show === index || show === -2">
                                        <v-divider></v-divider>
                                        <v-card-text>
                                            <div v-text="item.description"></div>
                                            <div>
                                                <span v-text="'Received totally '"></span>
                                                <span class="blue--text" v-text="'$' + item.total"></span>
                                                <span v-text="' from '"></span>
                                                <span class="blue--text" v-text="!item.donations ? 0 : item.donations "></span>
                                                <span v-text="' patrons'"></span>
                                            </div>
                                        </v-card-text>
                                    </div>
                                </v-expand-transition>


                            </v-card>
                        </v-col>
                    </v-row>
                </template>

                <!-- Footer -->
                <template v-slot:footer>
                    <v-row class="mt-2 ml-1" align="center" justify="center">
                        <span class="grey--text">Projects per page</span>
                        <v-menu offset-y>
                            <template v-slot:activator="{ on, attrs }">
                                <v-btn
                                        dark
                                        text
                                        color="green"
                                        class="ml-2"
                                        v-bind="attrs"
                                        v-on="on"
                                >
                                    {{ itemsPerPage }}
                                    <v-icon>mdi-chevron-down</v-icon>
                                </v-btn>
                            </template>
                            <v-list>
                                <v-list-item
                                        v-for="(number, index) in itemsPerPageArray"
                                        :key="index"
                                        @click="updateItemsPerPage(number)"
                                >
                                    <v-list-item-title>{{ number }}</v-list-item-title>
                                </v-list-item>
                            </v-list>
                        </v-menu>

                        <v-spacer></v-spacer>

                        <span
                                class="mr-4
                                grey--text"
                        >
                            Page {{ page }} of {{ numberOfPages }}
                        </span>
                        <v-btn
                                text
                                fab
                                dark
                                color="green"
                                class="mr-1"
                                @click="formerPage"
                        >
                            <v-icon>mdi-chevron-left</v-icon>
                        </v-btn>
                        <v-btn
                                text
                                fab
                                dark
                                color="green"
                                class="ml-1"
                                @click="nextPage"
                        >
                            <v-icon>mdi-chevron-right</v-icon>
                        </v-btn>
                    </v-row>
                </template>
            </v-data-iterator>
        </v-card>

    </v-container>
</template>

<script>
    import projectService from "../../service/projectService";

    export default {
        name: "Project",
        data () {
            return {
                itemsPerPageArray: [8, 12, 16],
                search: '',
                filter: {},
                sortDesc: false,
                page: 1,
                itemsPerPage: 8,
                sortBy: 'name',
                keys: [
                    'Name',
                    'Total',
                    'Donations',
                ],

                show: false,
                showCards: false,
                Projects: [{
                    id: "",
                    name: "",
                    description: "",
                    avatar_url: "",
                    total: "",
                    month: "",
                    donations: "",
                }],
            }
        },
        mounted: function() {
            this.updateProjectTable();
        },
        computed: {
            numberOfPages () {
                return Math.ceil(this.Projects.length / this.itemsPerPage)
            },
        },
        methods: {
            // 进入页面更新项目表单
            updateProjectTable() {
                projectService.getProjectTable().then((res) => {
                    this.Projects = res.data.data.Projects;
                    this.showCards = true;
                }).catch((err) => {
                    alert("表单更新时出错！" + err);
                    return null;
                });

                // 配置input监听器, 实现搜索功能
                const s = document.createElement('style');
                const input = document.getElementById("project_search");
                document.head.appendChild(s);
                input.addEventListener('input', function () {
                    if(this.value !== '') {
                        s.innerHTML = '.list:not([data-project*=' + this.value + ']){display:none}';
                    } else {
                        s.innerHTML = '';
                    }
                })
            },
            // 进入项目详细界面
            editProject(event) {
                const projectName = event.currentTarget.id;
                // 携带项目参数
                this.$router.push({ name: 'project/detail', params: {'name': projectName}});
            },
            // 全部展开
            expandAll() {
                if (this.show === -2) {
                    this.show = -1;
                } else {
                    this.show = -2;
                }
            },
            // 添加新项目
            addNewProject() {
                this.$router.push({ name: 'addProject' });
            },
            nextPage () {
                if (this.page + 1 <= this.numberOfPages) this.page += 1
            },
            formerPage () {
                if (this.page - 1 >= 1) this.page -= 1
            },
            updateItemsPerPage (number) {
                this.itemsPerPage = number
            },
        },
    }
</script>

<style scoped>

</style>
