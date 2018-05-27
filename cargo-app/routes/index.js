var express = require('express');
var router = express.Router();


//SPDX-License-Identifier: Apache-2.0

var controller = require('../controller.js');

router.get('/get_cargo/:id', function(req, res){
  controller.get_cargo(req, res);
});
router.post('/add_cargo', function(req, res){
  controller.add_cargo(req, res);
});
router.get('/get_all_cargo', function(req, res){
  controller.get_all_cargo(req, res);
});
router.post('/change_status', function(req, res){
  controller.change_status(req, res);
});
/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index');
});

module.exports = router;
