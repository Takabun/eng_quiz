<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm12 md8>
      <!-- {{ questions }} -->
      <v-container v-if="questions.length"
                   class="pb-8"
                   >
        <div class="d-flex  flex-wrap">
          <!-- style="overflow-x:scroll" -->
          <div v-for="item in questions" :key="item.id"
               class="ma-4">
             <item-card :item="item"
                        :width="250"
                        elevation="2"
                        />
          </div>
        </div>
      </v-container>
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
      questions: Array,
      tags: Array
    };
  },

  methods: {
    ConsoleLogging() {
      console.log("できたよ！");
    }
  },

  mounted() {
    axios.get(`http://localhost:1323/questions`).then(res => {
      console.log("axios結果あ！！", res);
      this.questions = res.data;
      return;
    });
    axios.get(`http://localhost:1323/tags`).then(res => {
      console.log("tags！！", res);
      this.tags = res.data;
      return;
    });
    // const age = this.$accessor.age.age;
    // this.$accessor.age.hoge();
  }
});
</script>
