#### **switch case**:
Là câu lệnh kiểu rẽ nhãnh. Với nhiều nhánh để rẽ, với mỗi nhành là mỗi trường hợp.

**_cú pháp:_**
switch([Biến]){
	case giá trị A:
		//khối lệnh 1
	break;
	case giá trị B:
		//khối lệnh 2
	break;
	case giá trị C:
		//khối lệnh 3
	break;
	.....
	defalt: (_Chỗ này mặc định luôn thực hiên, Và chỗ này có cũng được mà không có cũng không sao_)
		//Khối lệnh n1
	break;
}

### **goto:**
goto câu lệnh này dùng để điều hướng đến nơi bắt đầu lại trong đoạn code nào đó, nhưng trước tiên phải đặt nhãn mà nơi nó hướng tới khi xảy ra đúng điều kiện.

**_Ví dụ:_**
L1:  //Đặt nhãn
/ *
	Khối lệnh nào đó
*  /
goto L1;      //Quay về vạch xuất phát