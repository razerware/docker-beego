<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Document</title>
</head>
<body>
<form action="">
  <input type="text" name="a" value="" id="input"/>
  <input type="submit" id="submit">
</form>

<script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
<script>
$("#submit").on('click',function(e){
  var data = $("#input").val();
  // $.ajax({
  //     async : true,
  //     url : "https://api.douban.com/v2/book/search",
  //     type : "GET",
  //     dataType : "jsonp", // 返回的数据类型，设置为JSONP方式
  //     jsonp : 'callback', //指定一个查询参数名称来覆盖默认的 jsonp 回调参数名 callback
  //     jsonpCallback: 'handleResponse', //设置回调函数名
  //     data : {
  //         q : "javascript", 
  //         count : 1
  //     }, 
  //     success: function(response, status, xhr){
  //         console.log('状态为：' + status + ',状态是：' + xhr.statusText);
  //         console.log(response);
  //     }
  // });
  
  // $.getJSON('https://api.douban.com/v2/book/search?q=javascript&count=1&callback=?', function(json, textStatus) {
  //     console.log(json, textStatus);
  // });

  $.post('/containers', data, function(res){
    console.log(res)
  });
  return false;
})
</script>
</body>
</html>