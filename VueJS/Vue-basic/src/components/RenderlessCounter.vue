

<script>
    export default {
        data() {
            return{ count: 0} ;
        },
        methods: {
            increment(){
                this.count++;
            },
        },
        render(){
            return this.$slots.default({
                count: this.count,
                increment:this.increment,
            });
        }
    };
</script>

<template>
    <slot :count="count" :increment="increment" />
</template>