import { getterTree, mutationTree, actionTree } from "typed-vuex";

interface Tag {
  id: number,
  name: string
}

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
    actionTags({ getters, commit }): void {
      const currentTags = getters.tags;
      commit("setTags", currentTags);
    }
  }
);