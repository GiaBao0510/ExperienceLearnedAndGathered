// Phần mềm trung gian cấp độ bộ định tuyến: được liên kết đến 1 phiên bản của express.Router()
var express = require('express');
var app = express();
var router = express.Router();
var port = 3000;

router.use(function(req, res, next){
    console.log("Time: ",Date.now())
    next();
})

router.use('/user/:id',function(req, res, next){
    console.log("Request type: ",req.method)
    next();
})

router.use('/user/:id',function(req, res, next){
    res.send('USER')
})

app.use('/', router)
app.listen(port, function(){
    console.log(`Example app listening on post: ${port}!`);
});