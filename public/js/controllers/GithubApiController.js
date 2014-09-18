(function() {
'use strict';

// This module also relies on D3.js
angular.module('Gourcey').controller('GithubApiController', ['$scope', '$http',
    function($scope, $http) {
        
        $scope.query = '';
        $scope.repo = null;
        $scope.repos = [];
        $scope.commits = [];

        $scope.search = function() {
            $http.post('/github/search', {query: $scope.query})
                .success(function(data, status) {
                    var jsonData = angular.fromJson(JSON.parse(data));
                    $scope.repos = jsonData.items;
                })
                .error(function(data, status) {
                    console.log(status);
                    console.log(data);
                });
        };

        $scope.getCommits = function(repo) {
            $scope.repos = [];
            $scope.commits = [];
            $scope.repo = repo;
            $http.post('/github/commits', {repo: repo.full_name})
                .success(function(data, status) {
                    var jsonData = angular.fromJson(JSON.parse(data));
                    $scope.commits = jsonData;
                })
                .error(function(data, status) {
                    console.log(status);
                    console.log(data);
                });
        };
        
        $scope.getFileTree = function(sha) {
            $http.post('/github/tree', { repo: $scope.repo.full_name, sha: sha })
                .success(function(data, status) {
                    var jsonData = angular.fromJson(JSON.parse(data));
                    console.log(jsonData);
                })
                .error(function(data, status) {
                    console.log(status);
                    console.log(data);
                });
        };

        $scope.testing = function() {
            console.log('testing for ' + $scope.query);
            $http.get('/github/testing', {repo: $scope.query});
        };

        $scope.regenerate = function() {
            regenerate(false);
        };

    }
]);

})();
