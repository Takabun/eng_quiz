<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm12 md8>

      <div class="ma-6">
        <h3 class="mb-2">既存のタグから絞り込み<small class="body-2 primary--text ml-1">＊</small></h3>
        <div class="tags-wrap">
          <v-chip-group >
            <v-chip 
              v-for="item in getTags()" 
              :key="item.id"
              :color="` ${selectedTags.indexOf(item.id) > -1 ? 'primary': 'blue lighten-5'}`"
              @click="ApplyTags(item.id)"
            > {{item.name}}</v-chip>
          </v-chip-group>
        </div>
      </div>
      <v-container class="pb-8" >
        <div class="d-flex  flex-wrap">
          <div v-for="item in selecteditems" :key="item.id"
               class="ma-4">
            <item-card  :item="item"
                        :width="$isMd() ? '250' : '80vw'"
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
import {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question} from "../types/types"

export default Vue.extend({
  // Vue.extendがあるおかげで型推論が効くようになる
  components: {
    ItemCard
  },

  data() {
    return {
      allitems: [] as Question[],
      selecteditems: [] as Question[],
      selectedTags: [] as number[]
    };
  },

  methods: {
    getTags() {
      return this.$accessor.tags.tags
    },
    getQuestions() {
      return this.$accessor.questions.questions
    },
    ApplyTags(tagid: number) {
      if (this.selectedTags.indexOf(tagid) > -1) {
        this.selectedTags.splice(this.selectedTags.indexOf(tagid), 1)
      } else {
        this.selectedTags.push(tagid)
      }

      // 絞り込み
      if (this.selectedTags.length = 0) {
        this.selecteditems = this.allitems
      } else {

console.log("aaa", this.selectedTags)
console.log("いいい", this.selecteditems)
        this.selecteditems = [] // 初期化

        this.selectedTags.forEach(selectedtag => { //各選択タグ
          this.allitems.filter(element => { //各クエスチョン
          const tagIDs = element.tags.map(tag => tag.id)  //各クエスチョンのタグIDの配列
          if (tagIDs.indexOf(selectedtag) > -1 && this.selecteditems.indexOf(element) == -1) {
            this.selecteditems.push(element)
          }

        })
          
        });
        
      }

    }
  },

  async mounted() {
    await this.$accessor.tags.GetTags();
    await this.$accessor.questions.GetQuestions();
    this.allitems = this.getQuestions()
    this.selecteditems = this.getQuestions()
  }
});
</script>
