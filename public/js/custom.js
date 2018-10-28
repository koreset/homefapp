$(document).ready(function () {
    console.log("Ready for you");

    $.getJSON("/api/get-tweets", function (result, status, xhr) {
        console.log("We got kicked off");
        $.each(result, function (index, value) {
            console.log(value.fulltext);
            if (value.entities.urls.length > 0) {
                console.log(value.entities.urls);
            }

            $('.tweets').append(
                '<div class="media"><div class="media-body">' +
                '<img class="mr-2 float-left" src="' + value.user.profile_image_url_https + '" />'
                + getFullText(value)
                + '</div></div>'
            )
        });
    });

    $.getJSON("/api/get-flickr", function (result, status, xhr) {
        console.log(result);
        $.each(result.photos.photo, function (index, value) {
            $("div.flicker-images").append(
                '<img width="75" height="75" src=' + value.url_t + ' alt="">'
            );
        });
    });

    $("#lightgallery").lightGallery({
        thumbnail:true,
        thumbMargin: 40
    });

});


function getFullText(value) {
    if (value.retweeted_status !== null){
        return '<p>' + value.retweeted_status.full_text + '</p>'
    }else {
        return '<p>' + value.full_text + '</p>'
    }
}