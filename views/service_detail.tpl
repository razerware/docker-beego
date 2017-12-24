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
    <label class="layui-form-label">虚拟机名称</label>
    <div class="layui-input-inline">
      <input type="text" name="vm_name" required lay-verify="required" placeholder="请输入虚拟机名称" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">虚拟机镜像</label>
    <div class="layui-input-block">
      <select name="city" lay-verify="required">
        <option value=""></option>
        <option value="0">ubuntu16.04</option>
        <option value="1">ubuntu14.04</option>
      </select>
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">虚拟机类型</label>
    <div class="layui-input-block">
      <select name="city" lay-verify="required">
        <option value=""></option>
        <option value="0">4核心8G内存</option>
        <option value="1">2核心4G内存</option>
      </select>
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">硬盘大小</label>
    <div class="layui-input-block">
      <select name="city" lay-verify="required">
        <option value=""></option>
        <option value="0">40G</option>
        <option value="1">80G</option>
        <option value="1">160G</option>
      </select>
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