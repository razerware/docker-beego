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
      <label class="layui-form-label">选择集群</label>
      <div class="layui-input-inline" style="width: 300px">
        <select name="manager_ip" lay-verify="required" lay-filter="manager_ip" id="manager_ip">
          <option value=""></option>
        </select>
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label" >服务名称</label>
      <div class="layui-input-inline">
        <input type="text" name="name" required lay-verify="required" placeholder="请输入虚拟机名称" autocomplete="off" class="layui-input" >
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">容器镜像</label>
      <div class="layui-input-inline" style="width: 300px">
        <select name="image" lay-verify="required" >
          <option value=""></option>
          <option value="10.109.252.163:5000/emilevauge/whoami">10.109.252.163:5000/emilevauge/whoami</option>
        </select>
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">副本个数</label>
      <div class="layui-input-inline">
        <input type="number" name="replicas" required lay-verify="required" placeholder="副本个数" autocomplete="off" class="layui-input" >
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">扩容上限</label>
      <div class="layui-input-inline">
        <input type="number" name="upper_limit" required lay-verify="required" placeholder="副本个数不高于该值" autocomplete="off" class="layui-input" >
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">扩容下限</label>
      <div class="layui-input-inline">
        <input type="number" name="lower_limit" required lay-verify="required" placeholder="副本个数不低于该值" autocomplete="off" class="layui-input"  >

      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label" >步长</label>
      <div class="layui-input-block" style="width: 200px">
        <select name="step" lay-verify="required" >
          <option value="0">0</option>
          <option value="1">1</option>
        </select>
      </div>
    </div>
    <div class="layui-form-item">
      <div class="layui-inline">
        <label class="layui-form-label">CPU阈值</label>
        <div class="layui-input-inline">
          <input type="tel" name="cpu_lower" lay-verify="required" autocomplete="off" class="layui-input" placeholder="下限%">
        </div>
      </div>
      <div class="layui-inline">
        <div class="layui-input-inline">
          <input type="text" name="cpu_upper" lay-verify="required" autocomplete="off" class="layui-input" placeholder="上限%">
        </div>
      </div>
    </div>
    <div class="layui-form-item">
      <div class="layui-inline">
        <label class="layui-form-label">内存阈值</label>
        <div class="layui-input-inline">
          <input type="tel" name="mem_lower" lay-verify="required" autocomplete="off" class="layui-input" placeholder="下限%">
        </div>
      </div>
      <div class="layui-inline">
        <div class="layui-input-inline">
          <input type="text" name="mem_upper" lay-verify="required" autocomplete="off" class="layui-input" placeholder="上限%">
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
  var swarm_id = '';

  $.get("/cluster_list",function(rs){
    for(let i in rs.data){
      $("#manager_ip").append(`<option value=${rs.data[i].manager_ip}>${rs.data[i].swarm_id}</option>`);
    }
    form.render('select');
  })

  form.on('select(manager_ip)',function(data){
    var sel = document.getElementById("manager_ip");    
    for(var i in sel.childNodes) {
        if(sel.childNodes[i].nodeType == 1 && sel.childNodes[i].selected) {
            swarm_id = sel.childNodes[i].innerText;
            break;
        }
    }
  })

  form.on('submit(formDemo)', function(data){
    var postData = {};
    postData.Name = data.field.name;
    postData.Image = data.field.image;
    postData.SwarmId = swarm_id;
    postData.Constraints = "node.role==worker";
    postData.Target = "traefik-net";
    postData.Replicas = parseInt(data.field.replicas);;
    postData['traefik.port'] = "80";
    postData['traefik.frontend.rule'] = "Host:test.com";
    postData.UpperLimit = parseInt(data.field.upper_limit);
    postData.LowerLimit = parseInt(data.field.lower_limit);
    postData.Step = parseInt(data.field.step);
    postData.CpuLower = parseInt(data.field.cpu_lower);
    postData.CpuUpper = parseInt(data.field.cpu_upper);
    postData.MemLower = parseInt(data.field.mem_lower);
    postData.MemUpper = parseInt(data.field.mem_upper);
    console.log(postData)
    $.ajax({
      url:"/service_apply?manager_ip="+data.field.manager_ip,
      type:"post",
      dataType:"json",
      data:JSON.stringify(postData),
      success: function(msg){           
        layer.open({
          title: '提示'
          ,content: '申请成功！'
          ,yes: function(index, layero){
            layer.close(index); //如果设定了yes回调，需进行手工关闭
            $("#service_detail", window.parent.document).click();
            $("#service_detail", window.parent.document).parent().siblings('dd').removeClass('layui-this');
            $("#service_detail", window.parent.document).parent().addClass('layui-this')
          }
        });     
      }
    });
    return false;
  });

});
</script> 
</body>
</html>
