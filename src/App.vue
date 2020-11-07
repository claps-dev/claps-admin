<template>
  <v-app>

    <div v-if="this.$route.name!=='login' && this.$route.name!=='app'">
      <v-app-bar
              app
              clipped-right
              color="darkgrey"
              dark
      >
        <!-- progress linear -->
        <v-progress-linear
                id="progress-linear-2"
                indeterminate
                color="cyan"
                style="display: none"
        ></v-progress-linear>
        <v-spacer></v-spacer>
        <v-toolbar-title id="mainTitle" style="display: block">Welcome to Admin.claps.dev</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-app-bar-nav-icon @click.stop="drawerRight = !drawerRight"></v-app-bar-nav-icon>
      </v-app-bar>
      <!-- right -->
      <v-navigation-drawer
              v-model="drawerRight"
              app
              clipped
              right
      >
        <v-list dense>
          <v-list-item @click.stop="returnHome">
            <v-list-item-action>
              <v-icon>mdi-exit-to-app</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Home</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item @click.stop="manageProject">
            <v-list-item-action>
              <v-icon>mdi-exit-to-app</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Project</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item @click.stop="manageAdmin">
            <v-list-item-action>
              <v-icon>mdi-exit-to-app</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Admin</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item @click.stop="transDrawer = !transDrawer">
            <v-list-item-action>
              <v-icon>mdi-exit-to-app</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Trans</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item @click.stop="logout">
            <v-list-item-action>
              <v-icon>mdi-exit-to-app</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Logout</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

        </v-list>
      </v-navigation-drawer>

      <!-- Temporary drawer of Trans -->
      <v-navigation-drawer
              v-model="transDrawer"
              fixed
              right
              temporary
      >
        <v-list dense>
          <v-list-item @click="viewTransactions">
            <v-list-item-action>
              <v-icon>mdi-exit-to-app</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Transactions</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item @click="viewTransfers">
            <v-list-item-action>
              <v-icon>mdi-exit-to-app</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Transfers</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <!-- left -->
      <!--<v-navigation-drawer
              v-model="drawer"
              app
              clipped
              left
      >
        <v-list dense>

        </v-list>

      </v-navigation-drawer>-->
    </div>
    <v-main>
      <!-- progress linear -->
      <v-progress-linear
              id="progress-linear"
              indeterminate
              color="cyan"
              style="display: none"
      ></v-progress-linear>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  props: {
    source: String,
  },
  data: () => ({
    drawer: null,
    drawerRight: null,
    drawerLeft: null,
    right: false,
    transDrawer: false,
  }),
  components: {

  },
  methods: {
    // 查看管理员列表
    manageAdmin() {
      // 跳转
      this.$router.push({ name: 'admin' });
    },
    // 查看项目列表
    manageProject() {
      this.$router.push({ name: 'project' });
    },
    // 查看捐赠流水
    viewTransactions() {
      this.$router.push({ name: 'transactions' });
    },
    // 查看提现记录
    viewTransfers() {
      this.$router.push({ name: 'transfers' });
    },
    // 返回主页
    returnHome() {
      this.$router.push({ name: 'home' });
    },
    // 登出
    logout() {
      // 跳转登录页
      this.$router.replace({ name: 'login' });
      // 清除用户信息
      this.$store.dispatch('userModule/logout').then(() => {
      }).catch((err) => {
        console.log("清除用户缓存失败，错误信息：", err);
      });
    },

    progressStart() {
      //this.servi
      const linear = document.getElementById("progress-linear-2");
      linear.style.display = "block";
      const title = document.getElementById("mainTitle");
      title.style.display = "none";
    },

    progressStop() {
      const linear = document.getElementById("progress-linear-2");
      linear.style.display = "none";
      const title = document.getElementById("mainTitle");
      title.style.display = "block";
    },
  },
};
</script>
<style scoped>

</style>
