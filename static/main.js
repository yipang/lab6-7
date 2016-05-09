$(function(){
    $.get("/ping", function(data){
        if(data.error == "true"){
            $("#results").prepend("<div class='alert alert-danger'><strong>Error!</strong> "+ data.message +"</div>");
        }
    }, "json")

    $.get("/query1", function(data){
        $("#results").append(data.message);
    }, "json")

    $.get("/query2", function(data){
        $("#results").append(data.message);
    }, "json")

    $.get("/query3", function(data){
        $("#results").append(data.message);
    }, "json")

    $.get("/query4", function(data){
        $("#results").append(data.message);
    }, "json")
})
