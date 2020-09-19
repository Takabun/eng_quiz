import { getterTree, mutationTree, actionTree } from "typed-vuex";

interface Tag {
  id: number,
  name: string
}

interface Image{
  url: string;
}

interface Question {
  id: number,
  created_at: Date,
  updated_at: Date,
  user: string,
  text: string,
  default_image: number,
  tags: Tag[],
  images: Image[]
}

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

// 戻り値の型を明示的にしないとthis.app.$accessor経由でmutationsやactionsを呼び出そうとしたときに型推論が効かなくなってしまう
export const actions = actionTree({ state, getters, mutations }, {
    actionQuestions({ getters, commit }): void {
      const currentQuestions = getters.questions;
      commit("setQuestions", currentQuestions);
    }
  }
);