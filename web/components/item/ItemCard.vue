<template>
  <a>
  <v-card
    class="mx-auto"
    :width="width"
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
        :src="item.images[0]"
      >
      </v-img>
    </div>
    <v-card-subtitle class="pb-0">{{item.user}}</v-card-subtitle>
    <v-card-text class="text--primary">
      <div style="height: 45px;" class="overflow-x-hidden">{{item.text}}</div>
    </v-card-text>
    <v-card-actions>
      <div v-for="item in item.Tags" :key="item.Id" class="ma-1">
      <v-btn color="primary" >
        {{item.name}}
      </v-btn>
      </div>
    </v-card-actions>
  </v-card>
  </a>
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
    thumbWidth: {
      type: Number | String,
      default: 150
    }
   },

   data() {
    return {
      picture: "https://cdn.vuetifyjs.com/images/cards/docks.jpg"
    }
   }
};
</script>