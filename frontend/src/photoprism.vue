<template>
    <div id="photoprism">
        <p-loading-bar height="4"></p-loading-bar>

        <p-alert></p-alert>

        <v-app>
            <p-navigation v-show="loggedIn"></p-navigation>

            <v-content>
                <router-view></router-view>
            </v-content>
        </v-app>

        <p-photo-viewer></p-photo-viewer>
    </div>
</template>

<script>
    import "./css/app.css";

    export default {
        name: 'photoprism',
        data(){
            return {
                loggedIn: true
            }
        },
        created(){
            if (this.$router.currentRoute.name == "Home") {
                this.loggedIn= false
            }
        },
        computed: {},
        methods: {
            login() {
                // this.$refs.loginDialog.open();
            },

            logout() {
                this.$session.logout();
            },
        },
    };
</script>
