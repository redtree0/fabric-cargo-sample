
var helper = require('./helper.js')();

var controller = function(){};
var invokeHandler = function(res, results, tx_id) {
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

var queryhandler = function(res, query_responses){
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


controller.prototype.get_all_cargo = function(req, res){

	var tx_id = null;
	
	const request = helper.getChaincodeRequest('cargo-app', tx_id, 'queryAllCargo', 'mychannel', ['all']); 
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

}

controller.prototype.get_cargo = function(req, res){
	var tx_id = null;
	var key = req.params.id
	const request = helper.getChaincodeRequest('cargo-app', tx_id, 'queryCargo', 'mychannel', [key]); 

	helper.query(request, queryhandler.bind(this, res));

}

controller.prototype.add_cargo = function(req, res){
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
		helper.transaction(request, invokeHandler.bind(this, res));
}

controller.prototype.change_status = function(req, res){
	
	try{req.body = JSON.parse(Object.keys(req.body)[0])}catch(err){req.body = req.body}
	console.log(req.body);
	// res.send("ok");
	var key = req.body.Key;
	var tx_Id_org = req.body.TxId;
	var status = req.body.Status;
	
	var tx_id;

	// helper.transaction(request, txHandler, resHandler);
	const request = helper.getChaincodeRequest('cargo-app', tx_id, 'changeStatus', 'mychannel', [key, tx_Id_org,status]); 

	helper.transaction(request, invokeHandler.bind(this, res));

}

controller.prototype.get_point = function(req, res){
	var tx_id = null;
	var key = req.params.id
	const request = helper.getChaincodeRequest('cargo-app', tx_id, 'queryPoint', 'mychannel', [key]); 

	helper.query(request, queryhandler.bind(this, res));

}

controller.prototype.add_point = function(req, res){
	// try{req.body = JSON.parse(Object.keys(req.body)[0])}catch(err){req.body = req.body}
	// console.log(req.body);
	var key = req.body.Key;
	var point = req.body.point;
		var tx_id;
	const request = helper.getChaincodeRequest('cargo-app', tx_id, 'addPoint', 'mychannel', [key, point]); 
	// helper.transaction(request, txHandler, resHandler);
	helper.transaction(request, invokeHandler.bind(this, res));

}
controller.prototype.subtract_point = function(req, res){
	
	try{req.body = JSON.parse(Object.keys(req.body)[0])}catch(err){req.body = req.body}
	// console.log(req.body);
	var key = req.body.Key;
	var point = req.body.point;
	var tx_id;
	const request = helper.getChaincodeRequest('cargo-app', tx_id, 'subtractPoint', 'mychannel', [key, point]); 
	// helper.transaction(request, txHandler, resHandler);
	helper.transaction(request,  invokeHandler.bind(this, res));

}


module.exports = new controller() ;
