<template>
  <!-- <v-content> -->
    <v-container>
      <v-row>
        <v-col cols="7"
               md="7"
               lg="8">
            <v-alert class="">
              <small class="primary--text">＊</small>は必須項目です
            </v-alert>
            <!-- ユーザー名 -->
            <div class="mb-3">
              <h3 class="mb-2">ユーザー名<small class="body-2 primary--text ml-1">＊</small></h3>
              <v-text-field v-model="editdata"
                            outlined=""
                            maxlength="100" />
            </div>

            <!-- 投稿本文 -->
            <div class="mb-3">
              <h3 class="mb-2">投稿文<small class="body-2 primary--text ml-1">＊</small></h3>
              <v-text-field v-model="editdata"
                            outlined=""
                            maxlength="100" />
            </div>

            <!-- タグ選択 -->
            <div class="mb-3">
              <h3 class="mb-2">タグ選択<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- 新規タグ追加 -->
            <div class="mb-3">
              <h3 class="mb-2">新規にタグ追加<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- デフォルト画像 -->
            <div class="mb-3">
              <h3 class="mb-2">サムネイル画像選択(画像をアップロードしない場合)<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- 画像アップロード(Question) -->
            <div class="mb-3">
              <h3 class="mb-2">画像アップロード(Question)<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- 答え -->
            <div class="mb-3">
              <h3 class="mb-2">答え<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- 画像アップロード(Answer) -->
            <div class="mb-3">
              <h3 class="mb-2">画像アップロード(Answer)<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- 投稿ボタン -->
            <div class="text-center">
                <v-btn  color="primary"
                        depressed=""
                        width="255"
                        large=""
                        @click="onCreate()">上記の内容で投稿する</v-btn>
              </div>
          </v-col>
      </v-row>
    </v-container>
  <!-- </v-content> -->
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
  data() {
    return {
      editdata: {

      }
    }
  },

  methods: {
    onCreate() {
      
    }
  }
})
</script>