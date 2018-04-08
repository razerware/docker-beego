<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>table模块快速使用</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css" media="all">
  <link rel="stylesheet" href="../static/css/lzy.css">
</head>
<body>

<h3 style="margin-left: 20px">集群信息表</h3>
<table id="swarm-table" lay-filter="swarm-table"></table>
<h3 style="display: none; margin-left: 20px" id="vm_title">虚拟机信息表</h3>
<table id="vm-table" lay-filter="vm-table"></table>
 
<script src="../static/layui/layui.js"></script>
<script>
layui.use(['jquery','table','form'], function(){
  var $ =layui.$;
  var table = layui.table;
  var form = layui.form;

  //第一个实例
  table.render({
    elem: '#swarm-table'
    ,height: 315
    // ,width: 800
    ,url: 'cluster_list' //数据接口
    ,page: false //开启分页
    ,cols: [[ //表头
      {field: 'manager_ip', title: '管理节点IP地址', width:160}
      ,{field: 'lower_limit', title: '节点下限', width:100}
      ,{field: 'upper_limit', title: '节点上限', width:100}
      ,{field: 'step', title: '步长', width:60}
      ,{field: 'cpu_lower', title: 'CPU阈值下限', width:120}
      ,{field: 'cpu_upper', title: 'CPU阈值上限', width:120}
      ,{field: 'mem_lower', title: 'MEM阈值下限', width:160}
      ,{field: 'mem_upper', title: 'MEM阈值上限', width:160}
      ,{field: 'token', title: '集群token'}
      ,{fixed: 'right', title: '查看', width:150, align:'center', toolbar: '#barDemo'} //这里的toolbar值是模板元素的选择器
    ]]
  });

  table.on('tool(swarm-table)', function(obj){
    $("#vm_title").show();
    var swarmId = obj.data.swarm_id;
    //第2个实例
    table.render({
      elem: '#vm-table'
      ,height: 315
      // ,width: 800
      ,url: 'vm_list?swarmId='+swarmId //数据接口
      ,page: false //开启分页
      ,cols: [[ //表头
        {field: 'instance_id', title: '虚拟机ID'}
        ,{field: 'ip', title: '虚拟机IP'}
        ,{field: 'cpu', title: 'CPU'}
        ,{field: 'mem', title: 'MEM'}
        ,{field: 'disk', title: 'DISK'}
      ]]
    });
  });



});
</script>
<script type="text/html" id="barDemo">
  <a class="layui-btn layui-btn-danger layui-btn-mini" lay-event="del">查看</a>
</script>
</body>
</html>
