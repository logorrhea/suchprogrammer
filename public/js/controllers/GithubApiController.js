(function() {
'use strict';

angular.module('Gourcey').controller('GithubApiController', ['$scope', '$http',
    function($scope, $http) {
        
        $scope.query = '';
        $scope.repo = null;
        $scope.commits = null;
        $scope.repos = [];

        $scope.search = function() {
            $http.post('/github/search', {query: $scope.query})
                .success(function(data, status) {
                    $scope.repos = angular.fromJson(data).items;
                    console.log($scope.repos);
                })
                .error(function(data, status) {
                    //console.log('Error?');
                });
        };

    }
]);

})();
