<script>
    export default {
        methods: {
            gotoAbout(){
                this.$router.push('/about')
            },
            gotoUser(){
                this.$router.push('/user/giabao')
            },
        }
    }
</script>
<template>
    <h1>Home page: &nbsp; <button @click="gotoAbout"> Go to About</button></h1>
    <h1>User page: &nbsp; <button @click="gotoUser"> Go to User</button></h1>
</template>

