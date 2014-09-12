(function() {
'use strict';

angular.module('Gourcey').controller('SocketController', ['$scope', 'WebSocket',
    function($scope, WebSocket) {

        $scope.loading = false;
        $scope.percentComplete = 0;

        $scope.openSocketConnection = function() {
            $scope.loading = true;
            WebSocket.onopen(function() {
                WebSocket.send('message');
            });
            WebSocket.onmessage(function(event) {
                $scope.percentComplete = event.data + "%";
            });
            WebSocket.onclose(function() {
                $scope.loading = false;
            });
        };

    }
]);
})();
