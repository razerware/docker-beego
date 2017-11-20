<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>开始使用layui</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css">
</head>
<body>
 <form method="post" id="user">
    名字：<input name="username" type="text" />
    年龄：<input name="age" type="text" />
    邮箱：<input name="Email" type="text" />
    <input type="submit" value="提交" />
</form>
<!-- 你的HTML代码 -->
 
<script src="../static/layui/layui.js"></script>
<script>
//一般直接写在一个js文件中
layui.use(['layer', 'form'], function(){
  var layer = layui.layer
  ,form = layui.form;
  form.on('submit(user)', function(data){
    layer.alert(JSON.stringify(data.field), {
      title: '最终的提交信息'
    })
    return false;
  });
});
</script> 
</body>
</html>