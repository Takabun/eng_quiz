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
        :src="`${getImage()}`"
      />

    </div>
    <v-card-subtitle class="pb-0">{{item.user}}</v-card-subtitle>
    <v-card-text class="text--primary">
      <div style="height: 45px;" class="overflow-x-hidden">{{item.text}}</div>
    </v-card-text>
    <v-card-actions>
      <div v-for="item in item.tags" :key="item.Id" class="ma-1">
        <v-chip color="primary" >
          {{item.name}}
        </v-chip>
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
   },

  methods: {
    getImage() {
      console.log("aaa", this.item.images[0])  //ココを重点的に調査！
       this.item.images.forEach((element, index) => {
         if (index == 1) {
          console.log("keke", element.url)
          return element.url
        }
       });
      // return this.item.images[0]
    }
  },

  mounted() {
    console.log("item!", this.item)
  }
};
</script>