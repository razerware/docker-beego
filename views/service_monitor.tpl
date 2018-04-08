<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>应用监控</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css" media="all">
</head>
<body>
 
先不管
 
<script src="../static/layui/layui.js"></script>
<script>
//获取cookie  
function getCookie(cname) {  
    var name = cname + "=";  
    var ca = document.cookie.split(';');  
    for(var i=0; i<ca.length; i++) {  
        var c = ca[i];  
        while (c.charAt(0)==' ') c = c.substring(1);  
        if (c.indexOf(name) != -1) return c.substring(name.length, c.length);  
    }  
    return "";  
}
console.log(getCookie('url'))
location.href = 'http://' + getCookie('url');

layui.use(['jquery','table','form'], function(){
	var $ =layui.$;

});
</script>
</body>
</html>
