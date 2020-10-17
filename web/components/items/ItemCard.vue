<template>
    <!-- :width="width" -->
  <v-card
    class="mx-auto"
    :width="width"
    @click="$router.push(`/question/${item.id}`)"
  >

    <div v-if='item.default_image !== 0'>
      <v-img
        class=""
        height="200px"
        :src='`quesition_default_${item.default_image}.jpg`'
      >
      </v-img>
    </div>
    <div v-else>
      <v-img
        class=""
        height="200px"
        :src="getFirstImage"
      />

    </div>
    <v-card-subtitle class="pb-0">{{item.user}}</v-card-subtitle>
    <v-card-text class="text--primary">
      <div style="height: 45px;" class="overflow-x-hidden">{{item.text}}</div>
    </v-card-text>
    <div class="tags-wrap">
    <v-card-actions>
      
      <div v-for="item in item.tags" :key="item.Id" class="mr-2">
        <v-chip color="primary" >
          {{item.name}}
        </v-chip>
      </div>
      
    </v-card-actions>
    </div>
  </v-card>
</template>

<script>
import define from '../../plugins/define'
export default {
   props: {
    item: {
      type: Object,
      default: {}
    },
    width: {
      type: Number | String,
      default: null
    },
   },

   data() {
    return {
      picture: "https://cdn.vuetifyjs.com/images/cards/docks.jpg"
    }
   },

  computed: {
    getFirstImage() {
      let aaa = '';
      this.item.images.forEach((element, index) => {  // filter(reduce?)文を使うとよりスムーズ
        if (index == 0) {
          aaa = element.url
        }
      });
      return aaa;
    }
  },

};
</script>

<style scoped>
.tags-wrap {
  height: 48px !important;
  width: 100% !important;
  overflow-x: scroll !important;
}
</style>