import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
export const state = () => ({
  showModal: true,
})

export const mutations = {
  changeStateShowModal(state) {
    state.showModal=!state.showModal
  }
}
 export const getters = {
  getShowModal(state) {
    return state.showModal
  },
 }
