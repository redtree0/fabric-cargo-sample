var express = require('express');
var router = express.Router();
var controller = require('../controller.js');

//로그인처리
router.post('/login', (req, res) => {
 controller.loginuser(req,res);
});

//로그아웃처리
router.get('/logout',(req,res) => {
  sess=req.session;
  if(sess.userid){
    req.session.destroy(function(err){
              if(err){
                  console.log(err);
              }else{
                  res.send("logout success");
              }
          })
  }else{
    console.log("fail login please");
    res.send("fail login please");
  }
});

//회원가입처리
router.post('/register',(req,res)=>{
 controller.registeruser(req,res);
});

router.get('/logininfo',(req,res)=>{
  sess=req.session;
  if(sess.userid){
    var info={
      userid:sess.userid,
      userobj:sess.userobj
    };
    res.json(info);
  }else{
    res.send("logininfo not found, login please!~");
  }
});

module.exports = router;
