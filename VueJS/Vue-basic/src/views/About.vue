<script setup>
    import {computed} from 'vue'
    import {useRoute, useRouter} from 'vue-router'

    const router = useRouter()
    const route = useRoute()

    const search = computed({
        get(){
            return route.query.search ?? ''
        },
        set(search){
            route.replace({query: {search} })
        }
    })
</script>

<template>
    <h1>About</h1>
    <label >
        Search: <input v-model.trim="search" maxlength="20">
    </label>
</template>