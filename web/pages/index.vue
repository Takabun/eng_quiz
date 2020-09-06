<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      {{ this.questions }}
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import ItemCard from '../components/item/ItemCard.vue'

type sample = {
  aaa: string;
};

export default Vue.extend({
  // Vue.extendがあるおかげで型推論が効くようになる
  components: {
    ItemCard
  },

  data() {
    return {
      questions: Array
    };
  },

  methods: {
    ConsoleLogging() {
      console.log("できたよ！");
    }
  },

  mounted() {
    axios.get(`http://localhost:1323/questions`).then(res => {
      console.log("axios結果！！", res);
      this.questions = res.data;
      return;
    });
    // const age = this.$accessor.age.age;
    // this.$accessor.age.hoge();
  }
});
</script>
