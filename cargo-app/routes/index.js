var express = require('express');
var router = express.Router();


//SPDX-License-Identifier: Apache-2.0

var controller = require('../controller.js');

router.get('/get_cargo/:id', function(req, res){
  //check session
  var sess=req.session;
  if(sess.userid){
  controller.get_cargo(req, res);
  }else{
   var response={
	   result:"fail",
	   value:"login please"
   };
   console.log("request fail - login please");
   res.json(response);
  }

});

router.post('/add_cargo', function(req, res){
  //check session
  var sess=req.session;
  if(sess.userid){
  controller.add_cargo(req, res);
  }else{
   var response={
	   result:"fail",
	   value:"login please"
   };
   console.log("request fail - login please");
   res.json(response);
  }
});

router.get('/get_all_cargo', function(req, res){
  //check session
  var sess=req.session;
  if(sess.userid){
  controller.get_all_cargo(req, res);
  }else{
   var response={
	   result:"fail",
	   value:"login please"
   };
   console.log("request fail - login please");
   res.json(response);
  }

});

router.post('/change_status', function(req, res){
  //check session
  var sess=req.session;
  if(sess.userid){
  controller.change_status(req, res);
  }else{
   var response={
	   result:"fail",
	   value:"login please"
   };
   console.log("request fail - login please");
   res.json(response);
  }

});

/* GET home page. */

router.get('/get_point/', function(req, res){
  //check session
  var sess=req.session;
  if(sess.userid){
  controller.get_point(req, res);
   }else{
   var response={
	   result:"fail",
	   value:"login please"
   };
   console.log("request fail - login please");
   res.json(response);
  }
});


router.post('/subtract_point', function(req, res){
  //check session
  var sess=req.session;
  if(sess.userid){
  controller.subtract_point(req, res);
   }else{
   var response={
	   result:"fail",
	   value:"login please"
   };
   console.log("request fail - login please");
   res.json(response);
  }
});

router.post('/add_point', function(req, res){
  //check session
  var sess=req.session;
  if(sess.userid){
  controller.add_point(req, res);
   }else{
   var response={
	   result:"fail",
	   value:"login please"
   };
   console.log("request fail - login please");
   res.json(response);
  }
});

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index');
});

router.get('/login',(req, res, next) => {
  res.render('login');
});


module.exports = router;
