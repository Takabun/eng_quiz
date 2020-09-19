<template>
  <a>
  <v-card
    class="mx-auto"
    :width="width"
  >
    <v-img
      class=""
      height="200px"
      src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
    >
    </v-img>
    <v-card-subtitle class="pb-0">{{item.User}}</v-card-subtitle>
    <v-card-text class="text--primary">
      <div style="height: 45px;" class="overflow-x-hidden">{{item.Text}}</div>
    </v-card-text>
    <v-card-actions>
      <div v-for="item in item.Tags" :key="item.Id" class="ma-1">
      <v-btn color="primary" >
        {{item.Name}}
      </v-btn>
      </div>
    </v-card-actions>
  </v-card>
  </a>
</template>

<script>
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
    },
   }
};
</script>