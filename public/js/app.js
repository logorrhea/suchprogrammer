(function(){
'use strict';

angular.module('Gourcey',  ['ngRoute',
                            'angular-websocket'])
    .config(function(WebSocketProvider) {
        WebSocketProvider
            .prefix('')
            .uri('ws://localhost:9000/sockets/testing');
    });

})();
