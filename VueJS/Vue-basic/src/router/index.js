/*  --- Router
 - Ánh xạ đường dẫn URL đến các thành phần
 - Hiển thị các thành phần được cải thiện trong vị trí được chỉ định
*/

// createWebHistory từ thư viện vue-router: Hàm này tạo ra lịch sử định tuyến (history) sử dụng API Web History.
// - createWebHistory tương tự như URL phổ biến và sử dụng history.pushState() để tránh việc tải lại trang
//createRouter từ thư viện vue-router: Hàm này tạo ra một đối tượng router để quản lý các tuyến (routes) trong ứng dụng Vue.js.
import { createWebHistory, createRouter } from "vue-router";
import Home from "../views/Home.vue";
import About from "../views/About.vue";
import Users from "@/User/Users.vue";
import NotFound from "@/views/NotFound.vue";
import UserGeneric from "@/User/UserGeneric.vue";
import UserProfile from "@/User/UserProfile.vue";
import UserPost from "@/User/UserPost.vue";
import UserHome from "@/User/UserHome.vue";
import AdminHome from "@/admin/AdminHome.vue";
import WorkAdmin from "@/admin/WorkAdmin.vue";
import listEmployee from "@/admin/listEmployee.vue";
import UserSettings from "@/settingUser/UserSettings.vue";
import UserEmailsSubscriptions from "@/settingUser/UserEmailsSubscriptions.vue";
import UserProfilePreview from "@/settingUser/UserProfilePreview.vue";

//một Thành phần có thể chứa chính nó, được lồng <router-view>
const User = {
    template:`
    <div">
        <h2>user {{ $route.params.id }}</h2>
        <router-view></router-view>
    </div>
    `,
}

//Định nghĩa các tuyến, Tạo ra một mảng lưu trữ các tuyến, Mỗi đối tượng là 1 tuyến trong ứng dụng
const routers=[
    {
        path: '/',
        name: 'Home',       //
        component: Home,    //Thành phần của vue ,hiển thị khi người dùng truy cập vào tuyến này
        
        //alias: Có thể đường dẫn đến trang chủ có thể có nhiều tên gọi khác
        alias: '/homepage',
        
    },
    {
        path: '/about',
        name: 'About',
        component: About,
    },
    {
        path: '/user/:id',
        component: Users,
        /*
            - Các tuyến đường lồng nhau:
                + Ứng dụng ULs thực sự thường bao gồm các thành phần được lồng với nhiều cấp độ sâu,
                + để thực hiện điều nayfm 1 tùy chọn chidren được thêm để tuyến đường có 1 mảng chứa các
            đối tượng tuyến đường con.
        */
        children: [
            {
                path: "homeuser",
                component: UserHome,
            },
            {
                path: "profile",
                component: UserProfile,
            },
        ],
    },
    /*
     Các tham số thông thường sẽ hợp các ký tự ở giữa các đoạn url, được phân tách bằng /,
    Nếu chúng ta muốn hợp với bất kỳ thứ gì , chúng ta có thể sử dụng 1 biểu thức chính quy
    tham số tùy chỉnh bằng cách thêm biểu thức chính quy bên trong cặp dấu ngoặt đơn sau tham số
    */
   {
    //Sẽ khớp với mọi thứ và đặt nó dưới $route.params.pathMath
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound
   },
   {
    //Sẽ khớp với bất kỳ mọi thứ bắt đầu với /user và đặt nó dưới $route.params.afterUser
        path: '/user-:afterUser(.*)',
        component: UserGeneric
   },
   {
        path: '/admin',
        component: AdminHome,
        children: [
            {path: 'work', component: WorkAdmin},
            {path: 'list', component: listEmployee},
        ],
   },
   {
        path: "/settings",
        component: UserSettings,

        /*
        Bảo vệ mỗi router
            - beforeEnterL bảo vệ này chỉ kích hoạt khi đang đi từ tuyến này sang tuyến khác
        */
        beforeEnter: (to, from) =>{
            return true;
        },
        children: [
            {
                path: "emails",
                component: UserEmailsSubscriptions,
            },
            {
                path: "profile",
                component: UserProfilePreview,
            },
        ]
   },
];
export const router = createRouter({  //để xuất khẩu đối tượng router ra khỏi module hiện tại.
    history: createWebHistory(),    //Thiết lập thuộc tính history của router bằng cách gọi createWebHistory(). Hàm này tạo ra lịch sử định tuyến sử dụng Web History API.
    routes: routers,                        //Thiết lập thuộc tính routes của router bằng mảng routers đã định nghĩa ở trên.
});
