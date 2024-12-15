
const express = require('express');
const app = express();

app.get('/',(req, res)=>res.send("Hello world!"));

//Lang nghe tren cong 3000
app.listen(3000, function(){
    console.log('Example app listening on post 3000!');
});
