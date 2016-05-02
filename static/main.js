$(function(){
    $.get("/db", function(data){
        $("#results").append(data.message);
    }, "json")
})
