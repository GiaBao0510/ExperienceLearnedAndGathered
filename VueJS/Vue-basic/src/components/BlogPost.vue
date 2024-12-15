
<script>
    export default {
        //Có thể liệt kê props như 1 đối tượng thay vì mảng
        props: {
            title: String,
            likes: Number,
            isPublished: Boolean,
            conmmentIds: Array,
            callback: Function,
            author: Object,
            contactsPromise: Promise //HOặc bất kỳ cấu trúc xây dựng khác
        },
        emits: ['enlarge-text'],
    };
</script>
 
<template>
    <h3>Gia Bảo phá đảo Vue.js</h3>
    <div class="blog-post">
        <h4> {{ title }}</h4>
        <button @click="$emit('enlarge-text')">
            Enlarge text
        </button>
    </div>
</template>

<style scoped>

</style>