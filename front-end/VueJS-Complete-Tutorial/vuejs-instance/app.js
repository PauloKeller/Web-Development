var data = {
  title: 'The VueJS Instance',
  showParagraph: false
}

Vue.component('hello', {
  template: '<h1>Hello!</h1>'
});

var vm1 = new Vue({  
  data: {
    data: data  
  },
  methods: {
    show: function() {
      this.showParagraph = true;
      this.updateTitle('The VueJS Instance (Updated)');
      this.$refs.myButton.innerText = 'Test';
    },
    updateTitle: function(title) {
      this.title = title;
    } 
  },
  computed: {
    lowercaseTitle: function() {
      return this.title.toLowerCase();
    }
  },
  watch: {
    title: function(value) {
      alert('Title changed, new value: ' + value);
    }
  }
});

vm1.$mount('#app1');

vm1.$refs.heading.innerText = 'Something else';

setTimeout(function() {
  vm1.title = 'changed by timer';
  vm1.show();
}, 3000);

var vm2 = new Vue({
  el: '#app2',
  data: {
    title: 'The second Instance'
  },
  methods: {
    onChange: function() {
      vm1.title = "Changed!";
    }
  }
});

var vm3 = new Vue({
  el: '.hello',
  template: '<h1>Hello!</h1>'
});


// vm3.$mount('#app3');
vm3.$mount();
document.getElementById('app3').appendChild(vm3.$el);

// Vue Lifecycle

var vm4 = new Vue ({
  data: {
    title: 'The VueJS Instance'
  },
  beforeCreate: function() {
    console.log('beforeCreate()');
  },
  created: function() {
    console.log('created()');
  },
  beforeMount: function() {
    console.log('beforeMount()');
  },
  mounted: function() {
    console.log('mounted()');
  },
  beforeUpdate: function() {
    console.log('beforeUpdate()');
  },
  updated: function() {
    console.log('updated()');
  },
  beforeDestroy: function() {
    console.log('beforeDestroy()');
  },
  destroyed() {
    console.log('destroyed()');
  },
  methods: {
    destroy: function() {
      this.$destroy();
    }
  }
})

vm4.$mount('#app4')
