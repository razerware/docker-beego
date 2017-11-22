<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>开始使用layui</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css">
</head>
<body>
<form class="layui-form">
  <div class="layui-form-item">
    <label class="layui-form-label">名字</label>
    <div class="layui-input-block">
      <input type="json" name="username" required  lay-verify="required" placeholder="请输入名字" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">年龄</label>
    <div class="layui-input-inline">
      <input type="text" name="age" required lay-verify="required" placeholder="请输入密码" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">邮箱</label>
    <div class="layui-input-inline">
      <input type="text" name="text" required lay-verify="required" placeholder="请输入密码" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
      <button class="layui-btn" lay-submit lay-filter="formDemo">立即提交</button>
      <button type="reset" class="layui-btn layui-btn-primary">重置</button>
    </div>
  </div>
 </form>
<!-- 你的HTML代码 -->
 
<script src="../static/layui/layui.js"></script>
<!-- <script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script> -->
<script>
//一般直接写在一个js文件中
layui.use(['jquery','layer','form'],function(){
	var layer = layui.layer;
	var form = layui.form;
	var $ =layui.$;
    form.on('submit(formDemo)', function(data){
	    $.ajax({
	      url:"/lzy",
	      type:"post",
	      dataType:"json",
	      data:JSON.stringify(data.field),
	      success: function(msg){           
	         
	            console.log(msg);
	      }
	    });
	    return false;
    });
    
})
</script> 
</body>
</html>