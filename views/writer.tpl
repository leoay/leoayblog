<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="" type="image/x-icon" />
  <link rel="stylesheet" href="/static/editormd/css/editormd.css" />
	<script src="https://cdn.jsdelivr.net/npm/uikit@3.2.2/dist/js/uikit.min.js"></script>
  <script type="text/javascript" src="static/js/app.js" charset="utf-8"></script>
  <script src="https://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
  <script src="static/editormd/editormd.min.js"></script>
   <script src="static/js/jquery.cookie.js"></script>

</head>

<body>
  <header>

<div class="uk-text-center" uk-grid style="background: #444444;">
    <div class="uk-width-1-5" style="overflow:auto;margin:0px;height:750px;">
      <h3 style="color: #cccccc">历史文章（共 {{.totalNum}}篇）</h3>
      <div class="uk-child-width-auto@s" uk-grid>
        <div>
            <ul class="uk-list uk-list uk-list-muted" style="text-align: center;font-size:15px">
                {{range $filelist := .historyFile}}
                  <div class="uk-card uk-card-small uk-card-Primary uk-card-hover uk-card-body uk-width-1-1@m" id="{{$filelist.Name}}">
                      <p class="uk-card-title" style="font-size:15px">{{$filelist.Name}} style="color: #cccccc"</p>
                  </div>
                  <hr>
                {{end}}
            </ul>
        </div>
      </div>
    </div>

    <div class="uk-width-expand@m">
        <div class="uk-card uk-card-default uk-card-body" style="overflow:no;background: #444444">
        
                  <form id="article_editor" method="POST">
                    <div class="uk-flex uk-flex-center" style="margin-top:-30px">
                    <div class="uk-width-3xlarge uk-margin">
                    <div class="uk-width-1-2 uk-text-left" uk-grid>
                      <div>
                        <input class="uk-input" type="text" name="attitle" placeholder="输入文章标题" style="background: #444444;color: #cccccc;border: 4px solid #e5e5e5;">
                      </div>
                      <div></div>
                      
                    </div>
                    <div id="leoaylog-editor">

                    </div>
                    <div class="uk-child-width-1-5 uk-text-left" uk-grid>
                      <div>
                        <div class="uk-button-group">
                            <button class="uk-button uk-button-secondary" type="button" id="publish">发布</button>
                            <button class="uk-button uk-button-secondary" type="button" id="savetemp">暂存</button>
                        </div>
                      </div>
                      <div></div>
                    </div>
                  </form>
        
        </div>
    </div>
</div>

</header>



  <script type="text/javascript">
      $(function() {
          var editor = editormd("leoaylog-editor", {
              width  : "100%",
              height : 590,
              path   : "static/editormd/lib/",
              saveHTMLToTextarea  : true,
              theme : "dark",
              previewTheme : "dark",
              editorTheme : "3024-night",
              emoji : true,
              placeholder : "请用leoaylog开始愉快的写作吧！",
              toolbarIcons : function() {
                return ["h2", "h3", "h4", "code-block", "|", "bold", "hr", "datetime", "|", "preview", "watch", "|", "fullscreen", "emoji"]
        },
          });
      });
  </script>

  <script>
    document.addEventListener("keydown", function(e) {
        //可以判断是不是mac，如果是mac,ctrl变为花键
        //event.preventDefault() 方法阻止元素发生默认的行为。
        if (e.keyCode == 83 && (navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey)) {
            e.preventDefault();
            // Process event...
            console.log("保存成功！")
        }
    }, false);
  </script>

<script>
  //点击按钮，发布文章
  $("#publish").click(function() {
      //alert("sdsdsd")
      //ToDo: 使用ajax 发送post请求
      $("#article_editor").attr( 'action' , '/writer/publish');
      $("#article_editor").submit();
  });

  //点击按钮，暂存文章
  $("#savetemp").click(function() {
      //alert("sdsdsd")
      //ToDo: 使用ajax 发送post请求
      $("#article_editor").attr( 'action' , '/writer/savetemp');
      $("#article_editor").submit();
  });
</script>

  <script>
    $(document).ready(function() {
      let status = $.cookie('status');
      let message_info = $.cookie('message');
      console.log(status);
      if(status == 'success') {
        console.log(message_info);
        UIkit.notification({message: message_info})
      }
    });
  </script>
</body>
</html>