<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>table模块快速使用</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css" media="all">
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
    ,width: 800
    ,url: 'vm_detail_api' //数据接口
    ,page: true //开启分页
    ,cols: [[ //表头
      {field: 'vm_name', title: '虚拟机名称', minWidth:120, sort: true, fixed: 'left'}
      ,{field: 'vm_ip', title: '虚拟机IP地址', minWidth:180}
      ,{field: 'swarm_id', title: '所属集群id', minWidth:120, sort: true}
      ,{field: 'vm_info', title: '虚拟机规格', minWidth:180} 
    ]]
  });
  
});
</script>
</body>
</html>
