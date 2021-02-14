import { getAccessorType, mutationTree , actionTree} from "typed-vuex";
import * as tags from "~/store/tags";
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
    tags,
    questions
  }
});