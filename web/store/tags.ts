import { getterTree, mutationTree, actionTree } from "typed-vuex";
import axios from "axios";
import {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question} from "../types/types"

export const state = () => ({
  tags: [] as Tag[]
});

export type RootState = ReturnType<typeof state>;

export const getters = getterTree(state, {
  tags: state => state.tags
});

export const mutations = mutationTree(state, {
  setTags(state, tags: Tag[]): void {
    state.tags = tags;
  }
});

// 戻り値の型を明示的にしないとthis.app.$accessor経由でmutationsやactionsを呼び出そうとしたときに型推論が効かなくなってしまう
export const actions = actionTree({ state, getters, mutations }, {
  GetTags({ getters, commit }): void {
    console.log("is process.env.todoApiUrl?", process.env.todoApiUrl)
    axios.get(`${process.env.todoApiUrl}/tags`).then((res) => {
      console.log("tags get respose", res)
      let list: Tag[] = [];
      res.data.forEach((element: Raw_Tag) => {
        const payload = {
          id: element.ID,
          name: element.Name
        }
        list.push(payload)
      });
      commit("setTags", list);
      });
    }
  }
);