**Viết hàm tính độ đo F1-Score thường được sử dụng để đánh**  
**giá mô hình phân loại.**  

• Tn:
• Tp:
• Fn:
• Fp:
• Precision = TP /(TP + FP)  
- Ý nghĩa: của độ chích xác: Chỉ số này đo lường độ tin cậy khi mô hình đưa ra dự đoán Positive.
- Khi nào cần Precision cao? Khi chi phí cho việc báo sai (False Positive) quá lớn.

• Recall = TP /(TP + FN)  
- Ý nghĩa của recall: Chỉ số này đo lường khả năng "không bỏ sót" của mô hình.
- Khi nào cần Recall cao? Khi việc bỏ sót (False Negative) gây hậu quả nghiêm trọng.

• F1-score = 2 ∗ (Precision ∗ Recall)/  (Precision + Recall ) 
- Precision và Recall thường đánh đổi nhau (Precision tăng thì Recall thường giảm và ngược lại). F1 Score là trung bình điều hòa (harmonic mean) của hai chỉ số này.

• Input: hàm nhận 3 giá trị tp, fp, fn  
• Output: trả về và in ra kết quả của Precision, Recall, và F1-  
score

```python
def F1Score(tp, fp, fn):
    if(isinstance(tp, int) == False or
        isinstance(fp, int) == False or
        isinstance(fp, int) == False
    ):
        print("Input value must be an integer")
    
    #Tất cả phải điều có giá trị lớn hơn hoặc bằng 0
    if tp < 0 or fp < 0 or fn <0:
        print("Input value must be greater than or equal zero")
        
    #Kiểm tra nếu cả 3 cộng lại mà bằng 0 thì không thể tìm ra được F1    
    if(tp + fp + fn) == 0:
        print("Cannot calculate F1-score: all values are zero")
        
    precision = tp/(tp + fp) if (tp + fp) > 0 else 0
    recall = tp/(tp + fn) if(tp + fn) > 0 else 0
    
    F1_Score = 2 * (precision*recall)/(precision+recall)
    
    print(f"precision: {precision}")
    print(f"recall: {recall}")
    print(f"F1_Score: {F1_Score}")
```