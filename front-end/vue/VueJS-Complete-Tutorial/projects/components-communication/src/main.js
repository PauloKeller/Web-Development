import Vue from 'vue';
import App from './App.vue';

// use like a service
// use like model and centralize things up
const eventBus = new Vue({
  data: {
    data: 'data',
  },
  methods: {
    changeAge(age) {
      this.$emit('ageWasEdited', age);
    },
  },
});

export { eventBus };

new Vue({
  el: '#app',
  render: h => h(App),
});
