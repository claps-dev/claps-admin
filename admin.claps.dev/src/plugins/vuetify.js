import Vue from 'vue';
import Vuetify from 'vuetify/lib';

import colors from "vuetify/lib/util/colors";

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        light: true,
        primary: colors.brown.darken1,
        secondary: colors.brown.darken2,
    },
    icons: {
        iconfont: 'mdi',
    },
});
