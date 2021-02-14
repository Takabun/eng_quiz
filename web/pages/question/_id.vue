<template>
  <!-- <v-content> -->
    <v-container>
      <v-row>
        <v-col
               md="12"
               lg="8">

          <!-- ユーザー名 -->
          <div class="mb-3">
            <h3 class="mb-2">ユーザー名<small class="body-2 primary--text ml-1">＊</small></h3>
            <v-text-field v-model="question.user"
                          outlined=""
                          readonly
                          maxlength="100" />
          </div>

          <!-- 投稿本文 -->
          <div class="mb-3">
            <h3 class="mb-2">クエスチョン内容<small class="body-2 primary--text ml-1">＊</small></h3>
            <v-textarea v-model="question.text"
                        outlined=""
                        readonly
                        maxlength="100" />
          </div>

          <!-- タグ -->
          <div class="mb-3">
            <h3 class="mb-2">タグ<small class="body-2 primary--text ml-1">＊</small></h3>
            <div class="tags-wrap">
              <v-chip-group >
                  <v-chip 
                    v-for="item in question.tags" 
                    :key="item.id"
                    color="primary"
                  > {{item.name}}</v-chip>
              </v-chip-group>
            </div>
          </div>

          <!-- 画像 -->
          <div v-if='question.default_image !== 0'>
            <v-img
              class=""
              width="200px"
              :src='`../quesition_default_${question.default_image}.jpg`'
            >
            </v-img>
          </div>
          <div v-else>
            <v-img  v-for="item in question.images" 
                    :key="item.id"
                    class=""
                    width="400px"
                    :src="item.url"
            />
          </div>

          <!-- Answer本文 -->
          <div class="mb-3 mt-12">
            <h3 class="mb-2">回答<small class="body-2 primary--text ml-1">＊</small></h3>
            <v-textarea v-model="answer.text"
                        outlined=""
                        readonly
                        maxlength="100" />
          </div>
          <v-img  v-for="item in answer.images" 
                    :key="item.id"
                    class=""
                    width="400px"
                    :src="item.url"
            />
        </v-col>
      </v-row>
    </v-container>
</template>

<script lang="ts">
import axios from "axios";
import {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question} from "../../types/types"

export default {

  data() {
    return {
      question: {},
      answer: {}
    }
  },

  async mounted() {
    const q = await axios.get(`api/question/${this.$route.params.id}`, {id: this.$route.params.id})
    console.log("res(question)", q)
    const payloadQ = {
      id: q.data.ID,
      created_at: q.data.CreatedAt,
      user: q.data.User,
      text: q.data.Text,
      default_image: q.data.DefaultImage,
      tags: q.data.Tags.map(obj => {const robj: Tag = {id: obj.ID, name: obj.Name}; return robj} ),
      images: q.data.QuestionImages.map(obj => {const robj: Image = {url: obj.Url, name: obj.Name}; return robj} ),
    }
    this.question = payloadQ

    const a = await axios.get(`api/answer/${this.$route.params.id}`, {id: this.$route.params.id})
    console.log("res(answer)", a)
    const payloadA = {
      id: a.data.ID,
      created_at: a.data.CreatedAt,
      text: a.data.Text,
      images: a.data.AnswerImage.map(obj => {const robj: Image = {url: obj.Url, name: obj.Name}; return robj} ),
    } 
    this.answer = payloadA
  }
}
</script>