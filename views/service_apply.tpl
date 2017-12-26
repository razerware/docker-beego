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
      <label class="layui-form-label" >服务名称</label>
      <div class="layui-input-inline">
        <input type="text" name="vm_name" required lay-verify="required" placeholder="请输入虚拟机名称" autocomplete="off" class="layui-input" >
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">swarm集群</label>
      <div class="layui-input-block" style="width: 200px">
        <select name="city" lay-verify="required">
          <option value=""></option>
          <option value="0">test</option>
        </select>
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">容器镜像</label>
      <div class="layui-input-inline" style="width: 300px">
        <select name="city" lay-verify="required" >
          <option value=""></option>
          <option value="0">10.109.252.163:5000/test11</option>
        </select>
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">副本个数</label>
      <div class="layui-input-inline">
        <input type="text" name="vm_name" required lay-verify="required" placeholder="初始副本个数" autocomplete="off" class="layui-input" >
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">扩容上限</label>
      <div class="layui-input-inline">
        <input type="number" name="vm_name" required lay-verify="required" placeholder="副本个数不高于该值" autocomplete="off" class="layui-input" >
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">扩容下限</label>
      <div class="layui-input-inline">
        <input type="number" name="vm_name" required lay-verify="required" placeholder="副本个数不低于该值" autocomplete="off" class="layui-input"  >

      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label" >步长</label>
      <div class="layui-input-block" style="width: 200px">
        <select name="city" lay-verify="required" >
          <option value=""></option>
          <option value="0">1</option>
        </select>
      </div>
    </div>
    <div class="layui-form-item">
      <div class="layui-inline">
        <label class="layui-form-label">CPU阈值</label>
        <div class="layui-input-inline">
          <input type="tel" name="phone" lay-verify="required" autocomplete="off" class="layui-input" placeholder="下限%">
        </div>
      </div>
      <div class="layui-inline">
        <div class="layui-input-inline">
          <input type="text" name="email" lay-verify="required" autocomplete="off" class="layui-input" placeholder="上限%">
        </div>
      </div>
    </div>
    <div class="layui-form-item">
      <div class="layui-inline">
        <label class="layui-form-label">内存阈值</label>
        <div class="layui-input-inline">
          <input type="tel" name="phone" lay-verify="required" autocomplete="off" class="layui-input" placeholder="下限%">
        </div>
      </div>
      <div class="layui-inline">
        <div class="layui-input-inline">
          <input type="text" name="email" lay-verify="required" autocomplete="off" class="layui-input" placeholder="上限%">
        </div>
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
