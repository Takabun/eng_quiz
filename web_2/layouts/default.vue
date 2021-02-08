<template>
  <v-app dark>
     <v-app-bar app=""
               elevation="1"
               grow>
      <v-toolbar-title class="font-weight-bold"
                       @click="$router.push('/')">
        <span class="primary--text">ENG_QUIZ</span>
      </v-toolbar-title>
      <v-spacer />
      <v-tabs centered
              style="max-width: 500px;">
        <v-tab :ripple="false"
               exact=""
               to="/">見る</v-tab>
        <v-tab :ripple="false"
               to="/register">投稿する</v-tab>
      </v-tabs>
      <v-spacer />
    </v-app-bar>
    
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer
      :absolute="true"
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>
<script>
export default {
  data () {
    return {
    }
  }
}
</script>
