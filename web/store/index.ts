// import { getAccessorType } from "typed-vuex";

// // 例えば、store/age.ts のようなサブモジュールが存在する場合ここで import しておきます。
// import * as age from "~/store/age";

// // ここでは、state, getters, mutations, actions の記法は省略しています。
// // 記法については、後ほど記述する store/age.ts を参照してください。
// // これらは、たとえ必要なくても、以下のように空でいいので、必ず記述してください。
// export const state = () => {
//   //
// };

// export const getters = {
//   //
// };

// export const mutations = {
//   //
// };

// export const actions = {
//   //
// };

// export const accessorType = getAccessorType({
//   state,
//   getters,
//   mutations,
//   actions,
//   modules: {
//     // import したサブモジュールはここに記述します。
//     age
//   }
// });

import {
  getAccessorType,
  getterTree,
  mutationTree,
  actionTree
} from "typed-vuex";

export const state = () => ({
  message: "hello"
});

export const getters = getterTree(state, {
  reversedMessage(state) {
    return state.message
      .split("")
      .reverse()
      .join("");
  }
});

export const mutations = mutationTree(state, {
  updateMessage(state, newMessage: string) {
    state.message = newMessage;
  }
});

export const actions = actionTree(
  { state, getters, mutations },
  {
    // 戻り値の型を明示的にしないとthis.app.$accessor経由でmutationsやactionsを
    // 呼び出そうとしたときに型推論が効かなくなってしまう
    updateMessageAction(context, newMessage: string): void {
      this.app.$accessor.anotherUpdateMessageAction(newMessage);
    },

    anotherUpdateMessageAction(context, newMessage: string): void {
      context.commit("updateMessage", newMessage);
      // または this.app.$accessor.updateMessage(newMessage)
    }
  }
);

export const accessorType = getAccessorType({
  state,
  getters,
  mutations,
  actions
});
