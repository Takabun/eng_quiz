<template>
  <v-layout>
    <v-flex >
      <div class="ma-6">
        <h3 class="mb-2">タグから絞り込み<small class="body-2 primary--text ml-1">＊</small></h3>
        <div class="tags-wrap">
          <v-chip-group >
            <v-chip 
              v-for="item in this.$accessor.tags.tags" 
              :key="item.id"
              :color="` ${selectedTags.indexOf(item.id) > -1 ? 'primary': 'blue lighten-5'}`"
              @click="ApplyTags(item.id)"
            > {{item.name}}</v-chip>
          </v-chip-group>
        </div>
      </div>
      <v-container class="pb-8" >
        <v-row>
          <v-col  md="3"
                  sm="12"
                  v-for="item in selecteditems" :key="item.id">
            <item-card  :item="item"
                        elevation="2"
                        @click="$router.push(`/items/${item.id}`)"
                      />
          </v-col>
        </v-row>
      </v-container>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import ItemCard from '../components/items/ItemCard.vue'
import {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question} from "../types/types"

export default Vue.extend({
  // Vue.extendがあるおかげで型推論が効くようになる
  components: {
    ItemCard
  },

  data() {
    return {
      selectedTags: [] as number[]
    };
  },

  methods: {
    ApplyTags(tagid: number) {
      if (this.selectedTags.indexOf(tagid) > -1) {
        this.selectedTags.splice(this.selectedTags.indexOf(tagid), 1)
      } else {
        this.selectedTags.push(tagid)
      }
    }
  },

  computed: {
    selecteditems() : Question[] {
      if (this.selectedTags.length == 0) {
        return this.$accessor.questions.questions
      } else {
        let array = [] as Question[]
        this.selectedTags.forEach((selectedtag: number) => { //各選択タグ
          this.$accessor.questions.questions.filter(q => { //各クエスチョン(@Store)
            const qtagIDs = q.tags.map(tag => tag.id)  //各クエスチョンのタグIDの配列
            if (qtagIDs.indexOf(selectedtag) > -1 && array.indexOf(q) == -1) { array.push(q) }
          })
        });
        return array;
      }
    }
  },

  async mounted() {
    await this.$accessor.tags.GetTags();
    await this.$accessor.questions.GetQuestions();
  }
});
</script>
