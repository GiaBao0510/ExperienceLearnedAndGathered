<template>
    <p>Tên tôi là: {{ username }}</p>
    <button v-on:click="changeUsername"> Change Username</button>
</template>

<script>
    export default {
        props:{
            username: String
        },
        methods: {
            changeUsername(){
                this.$emit('changeUser', "[ --- Pham Gia Bao --- ]")
            }
        }
    }
</script>

