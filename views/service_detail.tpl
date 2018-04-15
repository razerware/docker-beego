<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>table模块快速使用</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css" media="all">
  <link rel="stylesheet" href="../static/css/lzy.css">
</head>
<body>
 
<table id="demo" lay-filter="test"></table>
 
<script src="../static/layui/layui.js"></script>
<script>
layui.use('table', function(){
  var table = layui.table;
  
  //第一个实例
  table.render({
    elem: '#demo'
    ,height: 315
    // ,width: 800
    ,url: '/list_service' //数据接口
    ,page: false //开启分页
    ,cols: [[ //表头
      {field: 'name', title: '服务名称', width:100}
      ,{field: 'swarm_id', title: '集群ID', width:100}
      ,{field: 'service_id', title: '服务ID', width:100}
      ,{field: 'address', title: '服务地址'}
      ,{field: 'image', title: '镜像'}
      ,{field: 'replication', title: '副本数量', width:100}
      ,{field: 'desire_replica', title: '副本预设', width:100}
      ,{field: 'lower_limit', title: '节点下限', width:100}
      ,{field: 'upper_limit', title: '节点上限', width:100}
      ,{field: 'step', title: '步长', width:60}
      ,{field: 'cpu_lower', title: 'CPU阈值下限', width:120}
      ,{field: 'cpu_upper', title: 'CPU阈值上限', width:120}
      ,{field: 'mem_lower', title: 'MEM阈值下限', width:120}
      ,{field: 'mem_upper', title: 'MEM阈值上限', width:120}       
    ]]
  });

});
</script>
<script type="text/html" id="barDemo">
  <a class="layui-btn layui-btn-mini edit" lay-event="edit">编辑</a>
  <a class="layui-btn layui-btn-danger layui-btn-mini" lay-event="del">删除</a>
</script>
</body>
</html>
