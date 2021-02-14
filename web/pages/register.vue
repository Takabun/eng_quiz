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
              <div class="tags-wrap">
                <v-chip-group >
                    <v-chip 
                      v-for="item in tags" 
                      :key="item.id"
                      @click="ApplyExistingTag(item)"
                      :color="` ${question.Tags.indexOf(item) > -1 ? 'primary': 'blue lighten-5'}`"
                    > {{item.name}}</v-chip>
                </v-chip-group>
              </div>
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
            <div v-if='this.questionImages.length == 0'>
              <div class="mb-3">
                <h3 class="mb-2">サムネイル画像選択(画像をアップロードしない場合)<small class="body-2 primary--text ml-1">＊</small></h3>
                <v-row dense=""
                    class="mb-5 mt-2">
                  <v-col v-for="(item,i) in $define.SAMPLE_IMAGES"
                        :key="i"
                        cols="2">
                    <v-img :src="item.value"
                          :style="`outline: ${question.DefaultImage == i+1 ? 3 : 0}px solid #1976d2; 
                                  opacity: ${question.DefaultImage == i+1 ? 1 : 0.4};
                                  cursor: pointer `"
                          aspect-ratio="1.5"
                          @click="question.DefaultImage = i+1"
                          />
                  </v-col>
                </v-row>
                <!-- {{question.DefaultImage}} -->
              </div>
            </div>

            <!-- 画像アップロード(Question) -->
            <div class="mb-3">
              <h3 class="mb-2">画像アップロード(Question)<small class="body-2 primary--text ml-1">＊</small></h3>
              <input type="file" accept="image/jpeg, image/png" @change="onQuestionImageChange"/>
              <div v-if='this.questionImages.length > 0'>
                <img :src="questionImagePreview" alt="Avatar" class="image" width="150" style="display: block; margin-top: 30px;">
              </div>
              <button @click="onCreateQuestionImage()">Upload to S3</button>
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
              <input type="file" accept="image/jpeg, image/png" @change="onAnswerImageChange"/>
              <div v-if='this.answerImages.length > 0'>
                <img :src="answerImagePreview" alt="Avatar" class="image" width="150" style="display: block; margin-top: 30px;">
              </div>
              <button @click="onCreateQuestionImage()">Upload to S3</button>
            </div>

            <!-- 投稿ボタン -->
            <div class="text-left mt-10">
                <v-btn  color="primary"
                        depressed=""
                        width="255"
                        large=""
                        :disabled='isDisabled'
                        @click="CreatenewOne()">上記の内容で投稿する</v-btn>
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
// import uuid from "uuid"
import {uuid} from "vue-uuid";
import "array-foreach-async";

// const NAMESPACE = "65f9af5d-f23f-4065-ac85-da725569fdcd";
interface HTMLElementEvent<T extends HTMLElement> extends Event {
  target: T;
}

export default Vue.extend({

  async asyncData ({axios}) {
    const resTags = await axios.get(`api/tags`)
    let tags: Tag[] = [];
    resTags.data.forEach((element: Raw_Tag) => {
      const payload = {
        id: element.ID,
        name: element.Name
      }
      tags.push(payload)
    });
    return {
      tags
    }
  },
  data() {
    return {
      question: {
        User: "",
        Text: "",
        Tags: [] as Tag[],
        DefaultImage: 0,
      },
      newTagDraft: '' as string,
      newTags: [] as string[],
      questionImages: [] as File[],
      answer: {
        Text: "",
      },
      answerImages: [] as File[],
      questionImagePreview: '',
      answerImagePreview: '' 
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
      if (this.newTagDraft !=="") {
        this.newTags.push(this.newTagDraft)
        this.newTagDraft = ''
      }
    },

    onCreateQuestionImage(qid: number) {
      AWS.config.update({accessKeyId: secrets.accessKeyId, secretAccessKey: secrets.secretAccessKey});
      const bucket = new AWS.S3({params: {Bucket: secrets.Bucket}});
      if (this.questionImages[0]) {
        const params: any = {Key: uuid.v4() + "-" + this.questionImages[0].name, ContentType: this.questionImages[0].type, Body: this.questionImages[0]}; 
        const uploadPromise = bucket.putObject(params).promise();
        uploadPromise
          .then(function(data: any) {
            console.log("uploaded!", data.$response);  
            axios.post(`api/questionimage`, {
              QuestionId: qid,
              Name:  data.$response.request.params.Body.name,
              Url: "https://quiztest-kt.s3-ap-northeast-1.amazonaws.com/" + data.$response.request.params.Key
            }).then(function(res) {
              console.log("DBへのQuestionImage保存まで完了", res)
            })
          })
          .catch(function(err) {
            console.error("error!", err, err.stack);
          });;
      }
    },

    onCreateAnswerImage(aid: number) {
      AWS.config.update({accessKeyId: secrets.accessKeyId, secretAccessKey: secrets.secretAccessKey});
      const bucket = new AWS.S3({params: {Bucket: secrets.Bucket}});
      if (this.answerImages[0]) {
        const params: any = {Key: uuid.v4() + "-" + this.answerImages[0].name, ContentType: this.answerImages[0].type, Body: this.answerImages[0]};
        const uploadPromise = bucket.putObject(params).promise();
        uploadPromise
          .then(function(data: any) {
            console.log("uploaded!", data.$response); 
            axios.post(`api/answerimage`, {
              AnswerId: aid,
              Name:  data.$response.request.params.Body.name,
              Url: "https://quiztest-kt.s3-ap-northeast-1.amazonaws.com/" + data.$response.request.params.Key
            }).then(function(res) {
              console.log("DBへのAnswerImage保存まで完了", res)
            })

          })
          .catch(function(err) {
            console.error("error!", err, err.stack);
          });;
      }
    },

    onQuestionImageChange(e: HTMLElementEvent<HTMLInputElement>) {
      const images = e.target.files 
      this.question.DefaultImage = 0;
      this.questionImages = []  
      this.questionImages.push(images![0])
      this.getBase64(this.questionImages[0])
        .then((image: unknown | string) => typeof image == "string" ? this.questionImagePreview = image : '')
        .catch(error => console.log('画像のアップロードに失敗しました。'))
    },

    onAnswerImageChange(e: HTMLElementEvent<HTMLInputElement>) {
      const images = e.target.files 
      this.answerImages = []
      this.answerImages.push(images![0])
      this.getBase64(this.answerImages[0])
        .then((image: unknown | string) => typeof image == "string" ? this.answerImagePreview = image : '')
        .catch(error => console.log('画像のアップロードに失敗しました。'))
    },

    getBase64 (file: File) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onload = () => resolve(reader.result)
        reader.onerror = error => reject(error)
      })
    },

    async CreatenewOne() {
      for (let i = 0; i < this.newTags.length ; i ++ ) {
        const res = await axios.post(`api/tag`, { Name: this.newTags[i] })
        this.question.Tags.push(res.data)
      };
      const questionresponse = await axios.post(`api/question`,this.question)
      const questionimageresponse = await this.onCreateQuestionImage(questionresponse.data.ID)
      const answerresponse = await axios.post(`api/answer`, {
        Text: this.answer.Text,
        QuestionId: questionresponse.data.ID
      })
      const answerimageresponse = await this.onCreateAnswerImage(answerresponse.data.ID)
    },
  },
  
  computed: {
    isDisabled() {
      if (this.question.User !== "" && this.question.Text !=="" && this.answer.Text !==""
        && (this.question.DefaultImage !== 0 || this.questionImages.length !== 0) ) return false;
      return true;
    }
  },
  
  async mounted() {
    // const resTags = await axios.get(`api/tags`)
    // let tlist: Tag[] = [];
    // resTags.data.forEach((element: Raw_Tag) => {
    //   const payload = {
    //     id: element.ID,
    //     name: element.Name
    //   }
    //   tlist.push(payload)
    // });
    // this.tags = tlist
  }
})
</script>
<style scoped>
.tags-wrap {
  height: 48px !important;
  width: 100% !important;
  overflow-x: scroll !important;
}
</style>