$(function(){
    $.get("/db", function(data){
        $("#results").append(data);
    }, "json")
})
