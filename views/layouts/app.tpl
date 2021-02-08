<!DOCTYPE html>

<html>
<head>
  <title>leoay lab</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="shortcut icon" href="" type="image/x-icon" />
    <script>
		var now = new Date().getTime();
		document.write('<link href="/static/editormd/css/editormd.css" rel="stylesheet" type="text/css"/>');
    document.write('<script src="/static/editormd/editormd.js"><\/script>');
  </script>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/uikit@3.2.2/dist/css/uikit.min.css" />
	<script src="https://cdn.jsdelivr.net/npm/uikit@3.2.2/dist/js/uikit.min.js"></script>
  <link rel="stylesheet" type="text/css" href="https://www.layuicdn.com/layui/css/layui.css" />
  <script src="https://www.layuicdn.com/layui/layui.js"></script>
  <script src="https://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
  <script src="http://cdn.staticfile.org/uikit/3.2.7/js/uikit-icons.min.js"></script>
  <link href="/static/font-awesome-4.7.0/css/font-awesome.css" rel="stylesheet">

  <link href="http://cdn.bootcss.com/highlight.js/8.0/styles/monokai_sublime.min.css" rel="stylesheet">  
  <script src="http://cdn.bootcss.com/highlight.js/8.0/highlight.min.js"></script>  
  <script >hljs.initHighlightingOnLoad();</script>  

  <link href="https://fonts.font.im/css?family=Open+Sans:300,400,600,700,800%7CShadows+Into+Light" rel="stylesheet" type="text/css">

  <style type="text/css">
        .uk-navbar-sticky {
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.09) !important;
        }
  </style>

</head>

<body>
    <pre style="display:none">
        _                          _       _       ___   ___ ___   ___  
        | |                        | |     | |     |__ \ / _ \__ \ / _ \ 
        | | ___  ___   __ _ _   _  | | __ _| |__      ) | | | | ) | | | |
        | |/ _ \/ _ \ / _` | | | | | |/ _` | '_ \    / /| | | |/ /| | | |
        | |  __/ (_) | (_| | |_| | | | (_| | |_) |  / /_| |_| / /_| |_| |
        |_|\___|\___/ \__,_|\__, | |_|\__,_|_.__/  |____|\___/____|\___/ 
                            __/ |                                       
                            |___/                                        
    </pre>


    {{template "layouts/header.tpl" .}}
    {{/*分页页面模板*/}}

      {{.LayoutContent}}
    <br>
    <br>
    <br>
    <br>
    <br>
    {{template "layouts/footer.tpl" .}}

</body>

  <script>
		var now = new Date().getTime();
		document.write('<link href="/static/editormd/css/editormd.css" rel="stylesheet" type="text/css"/>');
    document.write('<script src="/static/editormd/editormd.js"><\/script>');
    document.write('<link href="/static/css/theme.css?v=' + now + '" rel="stylesheet" type="text/css"/>');
  </script>

<script src="https://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
<script src="/static/js/jquery.cookie.js"></script>
  <script>
    $(document).ready(function() {
      let status = $.cookie('status');
      let message_info = $.cookie('message');
      //console.log(status);
      if(status == 'success') {
        console.log(message_info);
        UIkit.notification({
          message: message_info,
          timeout: 1000
          })
      }
    });
  </script>

  <script type="text/javascript" language=JavaScript charset="UTF-8">
    document.onkeydown=function(event){
        var e = event || window.event || arguments.callee.caller.arguments[0];
        if(e && e.keyCode==123){ // 按 Esc 
          //要做的事情
          //alert("别看源码啦！这个是真模糊！");
          UIkit.notification({
          message: '',
          status: 'primary',
          pos: 'top-center',
          timeout: 5000
        });
      }
    }; 
  </script>

<script src="https://www.goat1000.com/tagcanvas.js" type="text/javascript"></script>
<script>
  TagCanvas.Start("homepage__find-wordcloud-canvas", "homepage__find-wordcloud-list", {
      initial: [.05, 0],
      shape: "sphere",
      lock: "y",
      textHeight: 45,
      radiusX: 2.5,
      radiusY: .6,
      radiusZ: 1.5,
      textColour: "#fff",
      frontSelect: true,
      outlineColour: "#3e97f5",
      reverse: !0,
      depth: .8,
      centreImage: !0,
      imageAlign: !0,
      wheelZoom: !1,
      minSpeed: .01,
      maxSpeed: .05
  });
</script>