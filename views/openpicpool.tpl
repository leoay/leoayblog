<!DOCTYPE html>

<html>
<head>
  <title>leoay简易图片池</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <script src="https://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
  <link rel="stylesheet" href="/static/layui-v2.5.7/layui/css/layui.css">
  <script src="/static/layui-v2.5.7/layui/layui.all.js"></script>
  <script src="js/echo.min.js"></script>
</head>

<body>
  <header>
    <fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
      <legend>图片或文件上传</legend>
    </fieldset> 
    <div style="width:100%">
      <div class="layui-upload-drag" id="test10" style="width:100%">
        <i class="layui-icon"></i>
        <p>点击上传，或将文件拖拽到此处</p>
        <div class="layui-hide" id="uploadDemoView">
          <hr>
          <img src="" alt="上传成功后渲染" style="max-width: 196px">
        </div>
      </div>
    </div>


    <fieldset class="layui-elem-field layui-field-title" style="margin-top: 50px;">
      <legend>系统中已存在的图片</legend>
    </fieldset>
    <ul class="flow-default grid-demo grid-demo-bg1" style="width:12%" id="LAY_demo2"></ul>

  </header>

  <footer>
    
  </footer>

  <script>
  layui.use('upload', function(){
    var upload = layui.upload;
    
    //执行实例
    var uploadInst = //拖拽上传
    upload.render({
      elem: '#test10'
      ,url: '/open/picpool/upload' //改成您自己的上传接口
      ,done: function(res){
        //alert(JSON.stringify(res))
        layer.msg('上传成功');
        var obj = JSON.parse(JSON.stringify(res));
        //console.log(obj);
        layui.$('#uploadDemoView').removeClass('layui-hide').find('img').attr('src', obj.url);
        console.log(res)
      }
    }); 
  });

  ////////////////////////////////////////////////////////

  layui.use('flow', function(){
    var flow = layui.flow;

    flow.load({
      elem: '#LAY_demo2' //流加载容器
      ,scrollElem: '#LAY_demo2' //滚动条所在元素，一般不用填，此处只是演示需要。
      ,isAuto: false
      ,isLazyimg: true
      ,done: function(page, next){ //加载下一页
        //模拟插入
        setTimeout(function(){
          var lis = [];
          var maxPageNum;
          var imageList;
          $.post("/open/picpool/queryimages", 
          {
              msg: "sdsdsd",
              patch_id: "sdsddd"
          },function(result) {
              alert(result);
              maxPageNum = 1;
              imageList = ["http://leoay.xyz:8085/static/picture_pool/1614759006205.jpg", "http://leoay.xyz:8085/static/picture_pool/1614759006205.jpg", "http://leoay.xyz:8085/static/picture_pool/1614759006205.jpg"];
          });

          for(var i = 0; i < 3; i++){
            lis.push('<li><img style="width:100%" lay-src='+ imageList[i] +'"><button type="button" class="layui-btn layui-btn-primary layui-btn-radius">复制图片地址</button></li>')
          }

          //for(var i = 0; i < 6; i++){
          //  lis.push('<li><img style="width:100%" lay-src="http://leoay.xyz:8085/static/picture_pool/1614759006205.jpg?v='+ ( (page-1)*6 + i + 1 ) +'"><button type="button" class="layui-btn layui-btn-primary layui-btn-radius">复制图片地址</button></li>')
          // }
          next(lis.join(''), page < maxPageNum); //假设总页数为 6
        }, 500);
      }
    });
  });
  
</script>


</body>
</html>
