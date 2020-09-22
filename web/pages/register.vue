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
              <v-text-field v-model="question.User"
                            outlined=""
                            maxlength="100" />
            </div>

            <!-- 投稿本文 -->
            <div class="mb-3">
              <h3 class="mb-2">クエスチョン内容<small class="body-2 primary--text ml-1">＊</small></h3>
              <v-textarea v-model="question.Text"
                            outlined=""
                            maxlength="100" />
            </div>

            <!-- 既存タグ選択 -->
            <div class="mb-3">
              <h3 class="mb-2">既存のタグから選択<small class="body-2 primary--text ml-1">＊</small></h3>
              <v-chip-group >
                  <v-chip 
                    v-for="item in getTags()" 
                    :key="item.id"
                    @click="ApplyExistingTag(item)"
                    color="blue lighten-5"
                  > {{item.name}}</v-chip>
              </v-chip-group>
              <v-chip-group >
                  <v-chip 
                    v-for="item in question.Tags" 
                    :key="item.id"
                    color="primary"
                  > {{item.name}}</v-chip>
              </v-chip-group>
            </div>

            <!-- 新規タグ追加 -->
            <div class="mb-3">
              <h3 class="mb-2">新規にタグ追加<small class="body-2 primary--text ml-1">＊</small></h3>
              <div class="d-flex flex-row align-baseline">
                <v-text-field v-model="newTagDraft"
                  outlined=""
                  maxlength="100" />
                  <v-btn 
                      @click="AddNewTag()"
                      color="primary"
                      class="ml-5">
                      新規追加
                  </v-btn>
              </div>
              <v-chip-group >
                <v-chip 
                  v-for="item in newTags" 
                  :key="item.id"
                  color="primary"
                > {{item}}</v-chip>
              </v-chip-group>
            </div>

            <!-- デフォルト画像 -->
            <div class="mb-3">
              <h3 class="mb-2">サムネイル画像選択(画像をアップロードしない場合)<small class="body-2 primary--text ml-1">＊</small></h3>
              <v-row dense=""
                   class="mb-5 mt-2">
                <v-col v-for="(item,i) in $define.SAMPLE_IMAGES"
                      :key="i"
                      cols="4">
                  <v-img :src="item.value"
                        :style="`outline: ${question.DefaultImage == i ? 3 : 0}px solid #1976d2;cursor: pointer `"
                        aspect-ratio="1.5"
                        @click="question.DefaultImage = i"
                         />
                </v-col>
              </v-row>
              {{question.DefaultImage}}
            </div>

            <!-- 画像アップロード(Question) -->
            <div class="mb-3">
              <h3 class="mb-2">画像アップロード(Question)<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- 答え -->
            <div class="mb-3">
              <h3 class="mb-2">答え<small class="body-2 primary--text ml-1">＊</small></h3>
              <v-textarea v-model="answer.Text"
                            outlined=""
                            maxlength="100" />
            </div>

            <!-- 画像アップロード(Answer) -->
            <div class="mb-3">
              <h3 class="mb-2">画像アップロード(Answer)<small class="body-2 primary--text ml-1">＊</small></h3>
            </div>

            <!-- 投稿ボタン -->
            <div class="text-left">
                <v-btn  color="primary"
                        depressed=""
                        width="255"
                        large=""
                        :disabled='isDisabled'
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
import {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question} from "../types/types"

export default Vue.extend({
  data() {
    return {
      question: {
        User: "Test",
        Text: "ケケ",
        Tags: [] as Tag[],
        DefaultImage: 1,
        // QuestionImages: [],
      },
      newTagDraft: '' as string,
      newTags: [] as string[],
      questionImages: [],
      answer: {
        Text: "答えは",
        // AnswerImage: []
      },
      answerImages: [],

    }
  },

  methods: {
    ApplyExistingTag(item: Tag) {
      if (this.question.Tags.indexOf(item) > -1) {
        this.question.Tags.splice(this.question.Tags.indexOf(item), 1)
      } else {
        this.question.Tags.push(item)
      }
    },

    AddNewTag() {
      this.newTags.push(this.newTagDraft)
      this.newTagDraft = ''
    },

    getTags() {
      return this.$accessor.tags.tags
    },

    onCreate() {
      
    }
  },
  
  computed: {
    isDisabled() {
      if (this.question.User !== "" && this.question.Text !=="" && this.answer.Text !=="" ) return false;
      return true;
    }
  },
  mounted() {
    this.$accessor.tags.GetTags();
  }
})
</script>