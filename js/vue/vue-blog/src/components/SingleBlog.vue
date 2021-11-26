<template>
    <div id="single-blog">
        <h1>{{ blog.title }}</h1>
        <article>{{ blog.content }}</article>
        <p>Author: {{ blog.author }}</p>
        <ul>
            <li v-for="category in blog.categories"> {{ category }}</li>
        </ul>
    </div>
    
</template>

<script>

export default {
    data() {
        return{
            id: this.$route.params.id,
            blog: {}
        }
    },
    created() {
        this.$http.get('https://vue-blog-8d5b4.firebaseio.com/posts/' + this.id + '.json')
        .then(function(data){
            return data.json()
        }).then(function(data){
            this.blog = data
        })
    }
}
</script>

<style scoped>
#single-blog{
    max-width: 960px;
    margin: 0 auto;
    padding: 20px;
    background: #eee;
    border: 1px dotted #aaa;
}
</style>