(function(){
'use strict';

angular.module('Gourcey').config(['$routeProvider', '$locationProvider',
    function($routeProvider, $locationProvider) {
        $routeProvider
            .when('/', {
                controller: 'GithubApiController',
                templateUrl: '/public/js/templates/gource.html'
            })
            .otherwise({redirectTo: '/'});
    }
]);

})();
