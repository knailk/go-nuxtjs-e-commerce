import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
export const state = () => ({
  abc: 0,
})

export const mutations = {
  changeStateShowModal(state) {
    state.showModal=!state.showModal
  }
}
 export const getters = {
  getnumProductInCart(state) {
    return state.numProductInCart
  },
 }
