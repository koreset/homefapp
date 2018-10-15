$(document).ready(function () {
    console.log("Ready for you");
    // $(window).scroll(function () {
    //     if($(window).scrollTop() >= 120){
    //         $('img.homef-logo').addClass('scrolled');
    //         $('nav.navbar').removeClass('bg-dark');
    //         $('nav.navbar').addClass('bg-dark-scrolled');
    //
    //     }else{
    //         $('img.homef-logo').removeClass('scrolled');
    //         $('nav.navbar').addClass('bg-dark');
    //         $('nav.navbar').removeClass('bg-dark-scrolled');
    //     }
    // })

    $.getJSON("/api/get-tweets", function (result, status, xhr) {
        console.log("We got kicked off");
        $.each(result, function (index, value) {
            console.log(value.text);
            $(".tweets").append(
                "<div>" +
                "<img src='"+ value.user.profile_image_url_https +"'/>" +
                "<p>" + value.text + "</p>"
                + "</div>"
            );
        });
    });

    $.getJSON("/api/get-flickr", function (result, status, xhr) {
        console.log(result);
        $.each(result.photos.photo, function (index, value) {
            console.log("Value: " + value.url_sq);
            $("div.flicker-images").append(
            '<img width="75" height="75" src='+ value.url_t +' alt="">'
            );
        });


    });

});