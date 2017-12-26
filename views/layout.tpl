<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>layout Swarm集群管理 - Layui</title>
  <link rel="stylesheet" href="../static/layui/css/layui.css">
  <link rel="stylesheet" href="../static/css/lzy.css">
</head>
<body class="layui-layout-body">
  <div class="layui-layout layui-layout-admin">
    <div class="layui-header">
      <div class="layui-logo">服务创新云平台</div>
      <!-- 头部区域（可配合layui已有的水平导航） -->
      <ul class="layui-nav layui-layout-left">
        <li class="layui-nav-item"><a href="javascript:;">首页</a></li>
        <li class="layui-nav-item"><a href="javascript:;">平台简介</a></li>
        <li class="layui-nav-item"><a href="javascript:;">服务资源库</a></li>
        <li class="layui-nav-item">
          <a href="javascript:;">个人中心</a></li>
          <dl class="layui-nav-child">
            <dd><a href="javascript:;">我的集群</a></dd>
            <dd><a href="javascript:;">我的虚拟机</a></dd>
            <dd><a href="javascript:;">我的应用</a></dd>
          </dl>
        </ul>
        <ul class="layui-nav layui-layout-right">
          <li class="layui-nav-item">
            <a href="javascript:;">
              <img src="http://t.cn/RCzsdCq" class="layui-nav-img">
              正寅
            </a>
            <dl class="layui-nav-child">
              <dd><a href="javascript:;">基本资料</a></dd>
              <dd><a href="javascript:;">安全设置</a></dd>
            </dl>
          </li>
          <li class="layui-nav-item"><a href="javascript:;">退出</a></li>
        </ul>
      </div>

      <div class="layui-side layui-bg-black">
        <div class="layui-side-scroll">
          <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
          <ul class="layui-nav layui-nav-tree"  lay-filter="test">
            <li class="layui-nav-item layui-nav-itemed">
              <a class="" href="javascript:;">我的集群</a>
              <dl class="layui-nav-child">
                <dd class="layui-this"><a href="javascript:;" onclick="changeIframe('cluster_detail')">查看集群</a></dd>
                <dd><a href="javascript:;" onclick="changeIframe('cluster_apply')">集群初始化</a></dd>
                <dd><a href="javascript:;" onclick="changeIframe('cluster_monitor')">集群监控</a></dd>
              </dl>
            </li>
            <li class="layui-nav-item">
              <a href="javascript:;">我的虚拟机</a>
              <dl class="layui-nav-child">
                <dd><a href="javascript:;" onclick="changeIframe('vm_detail')">查看虚拟机</a></dd>
                <dd><a href="javascript:;" onclick="changeIframe('vm_apply')">申请虚拟机</a></dd>
                <dd><a href="javascript:;" onclick="changeIframe('http://10.109.252.172:8888')">虚拟机监控</a></dd>
              </dl>
            </li>
            <li class="layui-nav-item">
              <a href="javascript:;">我的应用</a>
              <dl class="layui-nav-child">
                <dd><a href="javascript:;" onclick="changeIframe('service_detail')">查看应用</a></dd>
                <dd><a href="javascript:;" onclick="changeIframe('service_apply')">申请应用</a></dd>
                <dd><a href="javascript:;" onclick="changeIframe('service_monitor')">应用监控</a></dd>
              </dl>
            </li>
          </ul>
        </div>
      </div>

      <div class="layui-body">
        <!-- 内容主体区域 -->
        <br/>
        <iframe src="cluster_detail" frameborder="0" id="demoAdmin" style="width: 100%; height: 96%;"></iframe>
      </div>

      <div class="layui-footer">
        <!-- 底部固定区域 -->
        © 李正寅退了 - 网络服务基础研究中心
      </div>
    </div>
    <script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.js"></script>
    <script src="../static/layui/layui.js"></script>
    <script>
    //JavaScript代码区域
    layui.use(['element'], function(){
      var element = layui.element;
    });

    function changeIframe(url){
      $("#demoAdmin").attr("src", url);  
    }
    </script>
</body>
</html>
