
<script>
  import BlogPost from './components/BlogPost.vue';
 import RenderlessCounter from './components/RenderlessCounter.vue';
  
  //export default{} dùng để xuất đối tượng mặc định từ 1 module. Để các file khác có thể sử dụng được
  export default {
    components:{    //Thuộc tính này chứa danh sách thành phần
      BlogPost,
      RenderlessCounter,
           
    },
    data(){         //Thuộc tính này dùng để định nghĩa các dữ liệu
      return{
        posts:[
          {
            id: 1, 
            title: "My journey with Vue", 
            likes:45, 
            isPublished: true,
            conmmentIds:[],
          },
          {
            id: 2, 
            title: "Blogging with Vue", 
            likes: 30,
            isPublished: true,
            conmmentIds:[],
          },
          {
            id: 3, 
            title: "Why Vue is so fun", 
            likes:55,
            isPublished: true,
            conmmentIds:[],
          },
        ],
        postFontSize: 1,  //Kích thước phông chữ cho bài đăng
      };
    },
  };
</script>

<template>
  <header>

    
    <div :style=" {fontSize: postFontSize+'em'}">
      <!-- Có thể sử dụng PascalCase cho thẻ thành phần trong tệp Vue-->
      <BlogPost
        v-for="post in posts"
        :key="post.id"
        :title="post.title"
        @enlarge-text="postFontSize+=0.1"
      />

      <!--
        1.Truyền Props:
      -->

      <h2> ---- Truyền 1 giá trị tĩnh ----- </h2>
      <BlogPost title="My journey with Vue 1"/>

      <h2> ------ Truyền 1 giá trị của biến -------</h2>
        <!--Gán giá trị dộng cho 1 biến-->
      <BlogPost v-for="post in posts" :title="post.title" />

        <!--Gán giá trị dộng cho 1 biểu thức phức tạp-->
      <!-- <BlogPost :title="posts.title + 'by' + posts.author.name"/> -->

      <h2> ------ Truyền 1 con số -------</h2>
        <!--Mặc dù `42` là tĩnh, chúng ta cần v-bind để kể cho Vue rằng đây là 1 Biểu thức JS hơn là 1 chuỗi-->
      <BlogPost v-bind="likes" :likes="42"/>
        <!--Gán giá trị dộng cho 1 biến-->
      <BlogPost  :likes="posts.likes" />

      <h2> ------ Truyền 1 boolean -------</h2>
      <!--Bao gồm Props không có giá trị gì nó cũng có nghĩa là True-->
      <BlogPost isPublished/>

      <!--Mặc dù 'false' là tĩnh, chúng ta cần v-bind nói với Vue rằng đây là 1 biểu thức JS hơn là 1 chuỗi-->
      <BlogPost v-bind="isPublished"  :isPublished="false"/>

      <!--Gán giá trị động trong 1 biến-->
      <BlogPost v-for="post in posts" :isPublished="post.isPublished"/>

      <h2> ------ Truyền 1 mảng -------</h2>
      <!--Mặc dù mảng là tĩnh, chúng ta cần v-bind để kể Vue rằng đây là 1 biểu thức JS hơn là String-->
      <BlogPost :conmmentIds="[234, 72 , 58]"/>
      <!-- Gán giá trị đọng đến 1 biến-->
      <BlogPost v-for="post in posts" :conmmentIds="post.conmmentIds" />

      <h2> ------ Truyền 1 đối tượng -------</h2>
      <!--Mặc dù đối tượng là tĩnh, chúng ta cần v-bind để kể Vue rằng đây là 1 biểu thức JS hơn là String-->
      <BlogPost :author="{
          name:'Baook',
          age: 22
        }"
      />
        <!-- Gán giá trị đọng đến 1 biến-->
      <BlogPost v-for="post in posts" :author="post.author" />

      <!--/Lấy từ tệp tin src\components\RenderlessCounter.vue-->
    <br><hr><br>
    <renderless-counter>
      <p>Tăng: {{ count }}</p>
      <button @click="increment">Increment</button>
    </renderless-counter>
    </div>
    
  </header>
  
</template>



<style scoped>

</style>
