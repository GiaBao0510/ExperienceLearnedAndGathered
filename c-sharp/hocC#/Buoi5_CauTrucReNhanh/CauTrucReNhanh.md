
### **Cấu trúc rẽ nhánh if_else**

**_Điều kiện khi thực hiện 1 câu lệnh._**
if(_Dieu_Kien_Logic_) Thuc_hien_mot_cau_lenh;

**_Điều kiện khi thực hiện 1 khối lệnh._**
if(_Dieu_Kien_Logic_) {
	Thuc_hien_khoi_lenh;
	....
}

**_Điều kiện khi thực hiện khối lệnh rẽ nhánh._**
if(_Dieu_Kien_Logic_) {
	Thuc_hien_khoi_lenh 1;
	....
}
else if(_Dieu_Kien_LogicA_){
	Thuc_hien_khoi_lenh 2;
	....
}
else if(_Dieu_Kien_LogicB){
	Thuc_hien_khoi_lenh 3;
	....
}
...
else{
	Thuc_hien_khoi_lenh n;
	....
}