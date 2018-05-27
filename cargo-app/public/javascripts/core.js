// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllCragos = function(){

		appFactory.queryAllCragos(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				// parseInt(data[i].Key);
				data[i].Record.Key = (data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			// console.log(array);
			$scope.all_cargo = array;
		});
	}

	$scope.queryCargo = function(){

		var id = $scope.cargo_id;

		appFactory.queryCargo(id, function(data){
			$scope.query_cargo = data;
			// console.log(data);
			if ($scope.query_cargo == "ERROR"){
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordCargo = function(){

		appFactory.recordCargo($scope.cargo, function(data){
			$scope.create_cargo = data;
			$("#success_create").show();
		});
	}

	$scope.changeStatus = function(){

		appFactory.changeStatus($scope.cargo, function(data){
			$scope.change_status = data;
			if ($scope.change_status == "ERROR"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllCragos = function(callback){

    	$http.get('/get_all_cargo/').success(function(output){
			callback(output)
		});
	}

	factory.queryCargo = function(id, callback){
    	$http.get('/get_cargo/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordCargo = function(data, callback){


		$http.post('/add_cargo/', data).success(function(output){
			callback(output)
		});
	
	}

	factory.changeStatus = function(data, callback){

		$http.post('/change_status/', data).success(function(output){
			callback(output)
		});
	}

	return factory;
});


