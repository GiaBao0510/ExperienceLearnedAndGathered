
import { query } from 'express';
<template>
    <h2>Single User page</h2>
    <h1>The user is {{ $route.params.name }}</h1>
    
    <!--
        Truy cập đến tham số truy vấn
    -->
    <p>({{ id }}) {{ name }}</p>

    <!--
    Có thể sử dụng nhiều <router-view> tại cùng 1 thời điểm.
    1 <router-view> không có tên sẽ tự mặc định lấy tên của nó. 
   -->
   <router-view  name="UserHome"></router-view>

    <router-link
        :to="{
            name: 'gia Bao',
            params:{
                id: 1234
            }
        }"
    >
        User with id: 1234
    </router-link>
    <!--
        Tham số truy vấn:
        -  Các tham số truy vấn được xây dựng giống như các tham số định tuyến,
        ngoại trừ chúng ta sử dụng tùy chọn truy vấn với 1 đối tượng để chứa cặp
        khóa: giá trị
    -->
   <router-link
    :to="{
        name:'routeName',
        query:{ query:'queryValue'}
    }"
   >
    Link text
   </router-link>

   
</template>

<script>
    export default {
        props:['id'],
        computed:{
            name(){ return this.$route.query.name}
        }
    }
</script>