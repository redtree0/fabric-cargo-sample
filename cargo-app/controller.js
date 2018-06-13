
var mysql=require('mysql');
var helper = require('./helper.js')();
var cahelper =require('./cahelper.js')();

module.exports = (function() {
return{
	get_all_cargo: function(req, res){

		var tx_id = null;

		const request = helper.getChaincodeRequest('cargo-app', tx_id, 'queryAllCargo', 'mychannel', ['']);
		helper.query(request, handler);

		function handler(query_responses){
			// res.header('Cache-Control', 'no-cache, no-store, must-revalidate')
			console.log("Query has completed, checking results");
			// query_responses could have more than one  results if there multiple peers were used as targets
			if (query_responses && query_responses.length == 1) {
				if (query_responses[0] instanceof Error) {
					console.error("error from query = ", query_responses[0]);
				} else {
					console.log("Response is ", query_responses[0].toString());
					res.json(JSON.parse(query_responses[0].toString()));
				}
			} else {
				console.log("No payloads were returned from query");
			}
		}

	},
	get_cargo: function(req, res){
		var tx_id = null;
		var key = req.params.id
		const request = helper.getChaincodeRequest('cargo-app', tx_id, 'queryCargo', 'mychannel', [key]);

		helper.query(request, handler);
		function handler(query_responses){
			// res.header('Cache-Control', 'no-cache, no-store, must-revalidate')
			    if (query_responses && query_responses.length == 1) {
			        if (query_responses[0] instanceof Error) {
			            console.error("error from query = ", query_responses[0]);
			            res.send("ERROR")
						// 에러처리 귀찮...
			        } else {
			            console.log("Response is ", query_responses[0].toString());
			            res.send(query_responses[0].toString())
			        }
			    } else {
			        console.log("No payloads were returned from query");
			        res.send("ERROR")
			    }
		}

	},
	add_cargo: function(req, res){
		console.log("submit recording of a Cargo: ");

		try{req.body = JSON.parse(Object.keys(req.body)[0])}catch(err){req.body = req.body}
		// console.log(req.body);
		var key = req.body.Key;
		var weight = req.body.weight;
		var registrant = req.body.Registrant;
        var driver = req.body.Driver;
        var date = req.body.date;
		var status = req.body.Status;
		var distance = req.body.distance;
		var money = req.body.money;
		var tx_id = null;

		const request = helper.getChaincodeRequest('cargo-app', tx_id, 'createContract', 'mychannel', [key, weight, distance, money, date, registrant, driver, status]);
		// helper.transaction(request, txHandler, resHandler);
		helper.transaction(request, resHandler);

		function resHandler(results, tx_id) {
			console.log('Send transaction promise and event listener promise have completed');
			// check the results in the order the promises were added to the promise all list
			function isAvailalbe(data, data_index, data_key, statement){
				if(data  && data[data_index] && data[data_index][data_key] === statement){
					return true;
				}
				return false;
			}
			if (isAvailalbe(results, 0, "status", "SUCCESS")) {
				// if (results && results[0] && results[0].status === 'SUCCESS') {
				console.log('Successfully sent transaction to the orderer.');
				// res.send(tx_id.getTransactionID());
			} else {
				// console.error('Failed to order the transaction. Error code: ' + response.status);
			}

			if(isAvailalbe(results, 1, "event_status", "VALID")) {
				// if(results && results[1] && results[1].event_status === 'VALID') {
				console.log('Successfully committed the change to the ledger by the peer');
				res.send(tx_id.getTransactionID());
			} else {
				console.log('Transaction failed to be committed to the ledger due to ::'+results[1].event_status);
			}
		}

	},

	change_status: function(req, res){

		try{req.body = JSON.parse(Object.keys(req.body)[0])}catch(err){req.body = req.body}
		console.log(req.body);
		// res.send("ok");
		var key = req.body.Key;

		var status = req.body.Status;

		var tx_id;

		// helper.transaction(request, txHandler, resHandler);
		const request = helper.getChaincodeRequest('cargo-app', tx_id, 'changeStatus', 'mychannel', [key, status]);

		helper.transaction(request, resHandler);

		function resHandler(results, tx_id) {
			console.log('Send transaction promise and event listener promise have completed');
			// check the results in the order the promises were added to the promise all list
			function isAvailalbe(data, data_index, data_key, statement){
				if(data  && data[data_index] && data[data_index][data_key] === statement){
					return true;
				}
				return false;
			}
			if (isAvailalbe(results, 0, "status", "SUCCESS")) {
				// if (results && results[0] && results[0].status === 'SUCCESS') {
				console.log('Successfully sent transaction to the orderer.');
				// res.send(tx_id.getTransactionID());
			} else {
				// console.error('Failed to order the transaction. Error code: ' + response.status);
			}

			if(isAvailalbe(results, 1, "event_status", "VALID")) {
				// if(results && results[1] && results[1].event_status === 'VALID') {
				console.log('Successfully committed the change to the ledger by the peer');
				res.send(tx_id.getTransactionID());
			} else {
				console.log('Transaction failed to be committed to the ledger due to ::'+results[1].event_status);
			}
		}


	},

	registeruser: function(req,res){

		try{req.body = JSON.parse(Object.keys(req.body)[0])}catch(err){req.body = req.body}
		console.log(req.body);
		console.log('start register User');
		var userid=req.body.userid,
		    password=req.body.password,
				dcert=req.body.dcert,
				bnum=req.body.bnum,
				phone=req.body.phone,
				tel=req.body.tel,
				cnum=req.body.cnum,
				anum=req.body.anum,
				uname=req.body.uname;
				console.log(userid);
	  if(userid!="" && password!="" && dcert!="" && bnum!="" && phone!="" && tel!="" && cnum!="" && anum!="" && uname!=""){
    //mysql 회원db에 회원등록

		var connection = mysql.createConnection({
	  host : 'localhost',
   	user : 'root',
  	password : 'konyang',
		port:3306,
  	database : 'blockchain'});

  	var member = {
     'id':userid,
     'pw':password,
     'dcert':dcert,
		 'bnum':bnum,
		 'phone':phone,
		 'tel':tel,
		 'cnum':cnum,
		 'anum':anum,
		 'uname':uname
	  };

		connection.connect();
		new Promise(function(resolve,reject){
		var dcheck=id_check(connection,userid,resolve,reject);
}).then(function(){
	//resolve호출시 (아이디 중복 없음)
	//mysql member Insert
	connection.query('INSERT INTO users set ?',member ,function(err, rows,fields){

	if(err) throw err;

	console.log('[mysql] member Insert success ');
	//fabric ca에 사용자 등록
	cahelper.registerCaUser(userid,password,handler,errhandler);
	connection.end();
	});//connection query end
},function(){
	//reject호출시 (아이디 중복)
	console.log('[fail] userid dup');
	res.send('[fail] userid dup');
	connection.end();
}).catch((err)=>{
	console.log("error");
	res.send('error');
});

    }else{
			res.send("[fail] arguments error")
    }

		function handler(){
			console.log("ca user register complete");
			res.send("ca user register complete");
		}

		function errhandler(){
			console.log("ca user register failed");
			res.send("ca user register failed");
		}

		function id_check(conn,userid,resolve,reject){
			console.log("idcheck start")
			conn.query("SELECT * FROM users WHERE id='"+userid+"'",function(err,rows,fields){
				if(err) throw err;

        var dup = rows[0];
        if(dup!=undefined){
					console.log('double-check fail');
					reject();
        }else{
				console.log('double-check clear');
				  resolve();
			  }
			});
		}

	},
	loginuser:function(req,res){
		try{req.body = JSON.parse(Object.keys(req.body)[0])}catch(err){req.body = req.body}
		console.log("loginuser start")
		var userid=req.body.userid,
		    password=req.body.password;

    if(userid!="" && password!=""){
			var conn = mysql.createConnection({
			host : 'localhost',
			user : 'root',
			password : 'konyang',
			port:3306,
			database : 'blockchain'});

		conn.query("SELECT pw FROM users WHERE id='"+userid+"'",function(err,rows,fields){
			if(err) throw err;

			if(rows[0]!=undefined){
        if(password==rows[0].pw){
					console.log("auth success");
					cahelper.enrollCaUser(userid,password,handler,errhandler);
				}else{
					console.log("auth fail");
					res.send("auth fail");
				}
			}else{
				res.send("[fail] user not found")
			}
		});
		conn.end();
	 }else{
			res.send("[fail] arguments error")
   }

		function handler(user){
			console.log("handler");
			req.session.userid=userid;
			req.session.userobj=JSON.stringify(JSON.parse(user));
			console.log("user login sucess");
			res.send("user login success");
		}

		function errhandler(err){
			console.log(err);
			res.send(err);
		}
	}//loginuser end

}
})();
