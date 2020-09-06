import { getAccessorType, mutationTree , actionTree} from "typed-vuex";
import * as questions from "~/store/questions";

export const state = () => ({});
export const getters = {};
export const mutations = mutationTree(state, { })
export const actions =  actionTree({ state, getters, mutations }, {} )
export const accessorType = getAccessorType({
  state,
  getters,
  mutations,
  actions,
  modules: {
    questions
  }
});

// サブモジュールを読み込む際、これらを反映する事でうまく行った！
// ①state = () => ({})へと変更
// ②mutationTree, actionTreeを使う