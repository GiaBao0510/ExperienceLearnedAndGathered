//Hàm parse() của JSON dùng để chuyern 1 chuỗi đối tượng sang dạng đối tượng 

const bear = '{ "name": "Ice Bear", "hobbies": ["knitting", "cooking", "dancing"] }';

const serializedBear = JSON.parse(bear);
console.log(serializedBear)
console.log("Kiểu: "+typeof(serializedBear))