<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>table模块快速使用</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css" media="all">
  <link rel="stylesheet" href="../static/css/lzy.css">
</head>
<body>
 
<table id="demo" lay-filter="vm-table"></table>
 
<form class="layui-form" style="display: none;" id="form">
    <div class="layui-form-item">
      <label class="layui-form-label">选择集群</label>
      <div class="layui-input-inline" style="width: 300px">
        <select name="manager_ip" lay-verify="required" lay-filter="manager_ip" id="manager_ip">
          <option value=""></option>
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

<script src="../static/layui/layui.js"></script>
<script>
layui.use(['jquery','table','form'], function(){
  var $ =layui.$;
  var table = layui.table;
  var form = layui.form;

  //第一个实例
  table.render({
    elem: '#demo'
    ,height: 315
    // ,width: 800
    ,url: 'vm_list_all' //数据接口
    ,page: false //开启分页
    ,cols: [[ //表头
      {field: 'instance_id', title: '虚拟机ID'}
      ,{field: 'swarm_id', title: '所属集群'}
      ,{field: 'ip', title: '虚拟机IP'}
      ,{field: 'cpu', title: 'CPU'}
      ,{field: 'mem', title: 'MEM'}
      ,{field: 'disk', title: 'DISK'}
      ,{fixed: 'right', title: '操作', width:150, align:'center', toolbar: '#barDemo'} //这里的toolbar值是模板元素的选择器
    ]]
    ,done: function(res, curr, count){
      var disable_arr = [];
      for(let i in res.data){
        if(res.data[i].swarm_id.length){
          disable_arr.push(i);
        }
      }
      var $trs = $("tbody").children("tr");
      for(let i in disable_arr){
        $($trs[disable_arr[i]]).children("td:last-child").children("div").children("a").css("pointer-events","none").css("background-color","grey");
      }
    }
  });


  var manager_ip = '';
  var myip = '';
  table.on('tool(vm-table)', function(obj){
    $("#form").show();
    myip = obj.data.ip;
    $.get("/cluster_list",function(rs){
      for(let i in rs.data){
        $("#manager_ip").append(`<option value=${rs.data[i].manager_ip} data-token=${rs.data[i].token}>${rs.data[i].swarm_id}</option>`);
      }
      form.render('select');
    });

  });

  form.on('select(manager_ip)',function(data){
    var sel = document.getElementById("manager_ip");    
    for(var i in sel.childNodes) {
        if(sel.childNodes[i].nodeType == 1 && sel.childNodes[i].selected) {
            token = sel.childNodes[i].dataset.token;
            break;
        }
    }
  })

  form.on('submit(formDemo)', function(data){
    var postData = {};
    postData.ManagerIp = data.field.manager_ip;
    postData.AdvertiseAddr = myip;
    postData.JoinToken = token;
    console.log(postData)
    $.ajax({
      url:"/cluster_join",
      type:"post",
      dataType:"json",
      data:JSON.stringify(postData),
      success: function(msg){           
        layer.open({
          title: '提示'
          ,content: '申请成功！'
          ,yes: function(index, layero){
            layer.close(index); //如果设定了yes回调，需进行手工关闭
            // $("#service_detail", window.parent.document).click();
            // $("#service_detail", window.parent.document).parent().siblings('dd').removeClass('layui-this');
            // $("#service_detail", window.parent.document).parent().addClass('layui-this')
          }
        });     
      }
    });
    return false;
  });

});
</script>
<script type="text/html" id="barDemo">
  <a class="layui-btn layui-btn-danger layui-btn-mini" lay-event="del">加入集群</a>
</script>
</body>
</html>
