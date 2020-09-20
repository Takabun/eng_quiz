<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm12 md8>
      <!-- {{ getTags() }}
      {{ getQuestions() }} -->
      <v-container v-if="getQuestions().length"
                   class="pb-8"
                   >
        <div class="d-flex  flex-wrap">
          <!-- style="overflow-x:scroll" -->
          <div v-for="item in getQuestions()" :key="item.id"
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

export default Vue.extend({
  // Vue.extendがあるおかげで型推論が効くようになる
  components: {
    ItemCard
  },

  data() {
    return {
      // questions: Array,
      // tags: this.$accessor.tags.tags
    };
  },

  methods: {
    getTags() {
      return this.$accessor.tags.tags
    },
    getQuestions() {
      return this.$accessor.questions.questions
    }
  },

  mounted() {
    this.$accessor.tags.GetTags();
    this.$accessor.questions.GetQuestions();
  }
});
</script>
