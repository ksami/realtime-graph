var data = [];
var myChart;
var count=0;

window.onload = function(){
  var ctx = $("#canvas").get(0).getContext("2d");
  myChart = new Chart(ctx).Line(lineChartData, {
    responsive: false
  });
};

var timer = setInterval(function(){
  $.get("details", function(data){
    var json = $.parseJSON(data);
    var content = json["with"][0]["content"];
    myChart.addData([content["memory"]["total_alloc"]], content["timestamp"]);
    if(count > 20){
      myChart.removeData();
    }
    if(count <= 20){
      count++;
    }
  });
}, 1000);

$("button").click(function(){
  clearInterval(timer);
})

var lineChartData = {
    labels : [],
    datasets : [
                   {
                       label: "My First dataset",
                       fillColor : "rgba(220,220,220,0.2)",
                       strokeColor : "rgba(220,220,220,1)",
                       pointColor : "rgba(220,220,220,1)",
                       pointStrokeColor : "#fff",
                       pointHighlightFill : "#fff",
                       pointHighlightStroke : "rgba(220,220,220,1)",
                       data : []
                   },
    ]
};

