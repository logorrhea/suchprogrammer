$(document).ready(function() {

    var items = [];

    $('#searchBtn').on('click', function() {
        $.ajax({
            url: '/github/search',
            data: {query: $('#query').val()},
            dataType: 'json',
            success: function(body, status) {
                items = JSON.parse(body).items;
                $.each(items, function(i, info) {
                    console.log(info.url);
                });
                $('#response').html(body);
                $('#response').show();
            }
        });
    });

});
