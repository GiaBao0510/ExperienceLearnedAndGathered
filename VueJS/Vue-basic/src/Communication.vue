<template>
    <header>
        <div class="component">
            <p>Tên tôi là: {{ usename }}</p>
            <hr>
            <account-info :username="usename" @changeUser="ChangeName"></account-info>
        </div>
    </header>
    
</template>

<script>
    import AccountInfoVue from './components/AccountInfo.vue'

    export default {
        components:{
            accountInfo: AccountInfoVue
        },
        data: function(){
            return{
                usename: 'Pham Gia Bao'
            };
        },
        methods:{
            ChangeName(s){
                this.usename = s;
            }
        },
    }; 
</script>

