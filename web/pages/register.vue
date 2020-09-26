<template>
  <!-- <v-content> -->
    <v-container>
      <v-row>
        <v-col
               md="12"
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
              <button id="upload-button" @click="onCreateImage()">Upload to S3</button>
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
import secrets from "../_secret/aws_info"
import AWS from "aws-sdk"
import uuid from "uuid"

export default Vue.extend({
  // head() {  // headを適用するにはexport default {}とする必要がある
  //   return {
  //     script: [
  //       { src:'https://sdk.amazonaws.com/js/aws-sdk-2.0.0-rc1.min.js' }
  //     ]
  //   }
  // },

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

    onCreateImage() {
      AWS.config.update({accessKeyId: secrets.accessKeyId, secretAccessKey: secrets.secretAccessKey});
      const bucket = new AWS.S3({params: {Bucket: secrets.Bucket}});
      const fileChooser= <HTMLInputElement>document.getElementById('file-chooser');
      const button = document.getElementById('upload-button');
      const results = <HTMLElement>document.getElementById('results');
      const file = fileChooser.files![0]; //!の位置でハマった
      
// before!
      // if (file) {
      //   results.innerHTML = '';
      //   const params: any = {Key: file.name, ContentType: file.type, Body: file};  //before
      //   // var params = {Key: file.name, Bucket: file.type, Body: file};  //後で確認 after

      //   const uploadPromise = bucket.putObject(params).promise();
      //   uploadPromise
      //     .then(function(data) {
      //       console.log("uploaded!", data);
      //       results.innerHTML = 'UPLOADED';
      //     })
      //     .catch(function(err) {
      //       console.error("error!", err, err.stack);
      //       results.innerHTML = 'ERROR!';
      //     });;
      // } else {
      //   results.innerHTML = 'Nothing to upload.';
      // }

// after!
      const bucketName = 'node-sdk-sample-' + uuid.v4();
      const keyName = 'hello_world.txt';
      const bucketPromise = new AWS.S3({apiVersion: '2006-03-01'}).createBucket({Bucket: bucketName}).promise();
      bucketPromise.then(
        function(data) {
          const objectParams = {Bucket: bucketName, Key: keyName, Body: 'Hello World!'};
          const uploadPromise = new AWS.S3({apiVersion: '2006-03-01'}).putObject(objectParams).promise();
          uploadPromise.then(
            function(data) {
              console.log("Successfully uploaded data to " + bucketName + "/" + keyName);
            });
      }).catch(
        function(err) {
          console.error(err, err.stack);
      });
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