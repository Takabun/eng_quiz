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
                    :color="` ${question.Tags.indexOf(item) > -1 ? 'primary': 'blue lighten-5'}`"
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
                        :style="`outline: ${question.DefaultImage == i+1 ? 3 : 0}px solid #1976d2; 
                                opacity: ${question.DefaultImage == i+1 ? 1 : 0.4};
                                 cursor: pointer `"
                        aspect-ratio="1.5"
                        @click="question.DefaultImage = i+1"
                         />
                </v-col>
              </v-row>
              {{question.DefaultImage}}
            </div>

            <!-- 画像アップロード(Question) -->
            <div class="mb-3">
              <h3 class="mb-2">画像アップロード(Question)<small class="body-2 primary--text ml-1">＊</small></h3>
              <input type="file" id="file-chooser" />
              <button id="upload-button">Upload to S3</button>
              <div id="results"></div>
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
import {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question} from "../types/types";
import AWS from "aws-sdk"


export default Vue.extend({
  data() {
    return {
      question: {
        User: "Test",
        Text: "ケケ",
        Tags: [] as Tag[],
        DefaultImage: 0,
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

      const button = document.getElementById('upload-button');
      const results = <HTMLElement>document.getElementById('results');
      // const file = fileChooser.files[0];
      const file = fileChooser.files !== null ? fileChooser.files[0] : null;  //後で確認！
      
      if (file) {
        results.innerHTML = '';
        // var params = {Key: file.name, ContentType: file.type, Body: file};  //before
        var params = {Key: file.name, Bucket: file.type, Body: file};  //後で確認！
        bucket.putObject(params, function (err, data) { 
          results.innerHTML = err ? 'ERROR!' : 'UPLOADED';
          console.log("data", data)
        });
      } else {
        results.innerHTML = 'Nothing to upload.';
      }
    }
  },
  
  computed: {
    isDisabled() {
      if (this.question.User !== "" && this.question.Text !=="" && this.answer.Text !==""
        && (this.question.DefaultImage !== 0 || this.questionImages.length !== 0) ) return false;
      return true;
    }
  },
  mounted() {
    this.$accessor.tags.GetTags();
  }
})
</script>