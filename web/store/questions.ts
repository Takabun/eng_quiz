import { getterTree, mutationTree, actionTree } from "typed-vuex";
import axios from "axios";
import {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question} from "../types/types"

export const state = () => ({
  questions: [] as Question[]
});

export type RootState = ReturnType<typeof state>;

export const getters = getterTree(state, {
  questions: state => state.questions
});

export const mutations = mutationTree(state, {
  setQuestions(state, questions: Question[]): void {
    state.questions = questions;
  }
});


export const actions = actionTree({ state, getters, mutations }, {
  GetQuestions({ getters, commit }): void {
    axios.get(`${process.env.todoApiUrl}/questions`).then((res) => {
    console.log("questions get response", res)
    let list: Question[] = [];
    res.data.forEach((element: Raw_Question) => {
      const payload = {
        id: element.ID,
        created_at: element.CreatedAt,
        user: element.User,
        text: element.Text,
        default_image: element.DefaultImage,
        tags: element.Tags.map(obj => {const robj: Tag = {id: obj.ID, name: obj.Name}; return robj} ),
        images: element.QuestionImages.map(obj => {const robj: Image = {url: obj.Url, name: obj.Name}; return robj} ),
      }
      list.push(payload)
    });
    commit("setQuestions", list);
    });
  }
  }
);