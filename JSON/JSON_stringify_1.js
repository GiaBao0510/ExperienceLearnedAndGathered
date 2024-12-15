//Hàm stringify() của JSON dùng để chuyern 1 đối tượng sang dạng chuỗi 

const bear = {
    name: 'Ice Bear',
    hobbies: ['knitting', 'cooking', 'dancing']
}

const serializedBear = JSON.stringify(bear);
console.log(serializedBear)
console.log('Kiểu: '+typeof(serializedBear))