import { getAccessorType, mutationTree , actionTree} from "typed-vuex";
import * as age from "~/store/age";

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
    age
  }
});

// サブモジュールを読み込む際、これらを反映する事でうまく行った！
// ①state = () => ({})へと変更
// ②mutationTree, actionTreeを使う



// import {
//   getAccessorType,
//   getterTree,
//   mutationTree,
//   actionTree
// } from "typed-vuex";

// export const state = () => ({
//   message: "hello"
// });

// export const getters = getterTree(state, {
//   reversedMessage(state) {
//     return state.message
//       .split("")
//       .reverse()
//       .join("");
//   }
// });

// export const mutations = mutationTree(state, {
//   updateMessage(state, newMessage: string) {
//     state.message = newMessage;
//   }
// });

// export const actions = actionTree(
//   { state, getters, mutations },
//   {
//     // 戻り値の型を明示的にしないとthis.app.$accessor経由でmutationsやactionsを
//     // 呼び出そうとしたときに型推論が効かなくなってしまう
//     updateMessageAction(context, newMessage: string): void {
//       this.app.$accessor.anotherUpdateMessageAction(newMessage);
//     },

//     anotherUpdateMessageAction(context, newMessage: string): void {
//       context.commit("updateMessage", newMessage);
//       // または this.app.$accessor.updateMessage(newMessage)
//     }
//   }
// );

// export const accessorType = getAccessorType({
//   state,
//   getters,
//   mutations,
//   actions
// });
