<template>
  <v-layout>
    <v-flex >
      <div class="ma-6">
        <h3 class="mb-2">タグから絞り込み<small class="body-2 primary--text ml-1">＊</small></h3>
        <div class="tags-wrap">
          <v-chip-group >
            <v-chip 
              v-for="item in tags" 
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
  components: {
    ItemCard
  },

  async asyncData({axios}) {
   const resTags = await axios.get(`api/tags`)
    console.log("tags get response", resTags)
    const tags: Tag[] = [];
    resTags.data.forEach((element: Raw_Tag) => {
      const payload = {
        id: element.ID,
        name: element.Name
      }
      tags.push(payload)
    });

    const resQ = await axios.get(`api/questions`)
    console.log("questions get response", resQ)
    const items: Question[] = [];
    resQ.data.forEach((element: Raw_Question) => {
      const payload = {
        id: element.ID,
        created_at: element.CreatedAt,
        user: element.User,
        text: element.Text,
        default_image: element.DefaultImage,
        tags: element.Tags.map(obj => {const robj: Tag = {id: obj.ID, name: obj.Name}; return robj} ),
        images: element.QuestionImages.map(obj => {const robj: Image = {url: obj.Url, name: obj.Name}; return robj} ),
      }
      items.push(payload)
    });
    return { tags, items, selectedTags: [] }
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
        return this.items
      } else {
        let array = [] as Question[]
        this.selectedTags.forEach((selectedtag: number) => { 
          this.items.filter(q => { 
            const qtagIDs = q.tags.map(tag => tag.id) 
            if (qtagIDs.indexOf(selectedtag) > -1 && array.indexOf(q) == -1) { array.push(q) }
          })
        });
        
        return array;
      }
    }
  },

});
</script>
