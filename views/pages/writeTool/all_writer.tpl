<div class="uk-container uk-container-xlarge main_page">
    <div class="uk-flex uk-width-1 uk-flex-center uk-margin-medium-top uk-margin-medium-bottom">

            <div class="uk-width-2-3 uk-flex-center main_pc_write_article">
                <form class="uk-width-1" id="article_editor" method="POST">
                    <div class="uk-flex uk-flex-center">
                        <h3 class="uk-margin-small-bottom"><span class="uk-icon" uk-icon="icon: pencil"></span>新建一篇文章</h3>
                    </div>
                    <div uk-grid class="uk-grid-small">
                        <div class="uk-width-1-4">

                            <div class="uk-text-center uk-grid-collapse" uk-grid>
                                <div class="uk-width-1-2">
                                   <select class="uk-input" id="article_class1" name="article_class1" type="text" uk-border-pill style="background: #fff;color: #111;border: 2px solid #e6d4d4;border-radius: 31px;">
                                    </select>
                                </div>
                                <div class="uk-width-1-2">
                                    <select class="uk-input" type="text" name="article_class2" id="article_class2" uk-border-pill style="background: #fff;color: #111;border: 2px solid #e6d4d4;border-radius: 31px;">
                                         
                                    </select>
                                </div>
                            </div>

                        </div>
                        <div class="uk-width-expand">
                            <input class="uk-input" type="text" id="article_title" name="title" uk-border-pill placeholder="输入文章标题，过长的标题会被隐藏的" style="background: #fff;color: #111;border: 2px solid #e6d4d4;border-radius: 31px;">
                        </div>
                        <div class="uk-width-expand">
                            <input class="uk-input" type="text" id="article_main_pic" name="main_pic_path" uk-border-pill placeholder="输入文章主图链接" style="background: #fff;color: #111;border: 2px solid #e6d4d4;border-radius: 31px;">
                        </div>
                    </div>
                
                    <div id="leoaylog-editor" class="uk-margin-medium-top">
                    </div>
                    <div class="uk-margin">
                        <textarea class="uk-textarea summary"  maxlength="120" name="summary" rows="5" placeholder="文章摘要(不超过120字)"></textarea>
                        <i class="uk-flex uk-flex-right "><span class="count-change">0</span>/120</i>
                    </div>
                </form>
                
                <button class="uk-button uk-button-default  uk-margin-medium-left uk-border-pill" id="publish">发布</button>
                <button class="uk-button uk-button-default uk-margin-small-left uk-border-pill" id="savetemp">保存</button>

            <button class="uk-button uk-button-default uk-margin-small-left uk-border-pill" type="button" id="genPic1" uk-toggle="target: #modal-close-default">生成分享图</button>

            <div id="modal-close-default" uk-modal>
                <div class="uk-modal-dialog uk-modal-body">
                    <button class="uk-modal-close-default" type="button" uk-close></button>
                    <div id="shareCard" style="padding:15px;background: #1b1c1f;color: #f4e1b9;border-radius: 0 30px 0 0;border: #fff 15px;border-style: solid;box-shadow: black;box-shadow: darkgrey 4px -3px 20px 3px;">
                        <br>
                        <h2 class="uk-modal-title" style="color: #f4e1b9;font-size: 2em;"></h2>
                        <hr>
                        <br>
                        <div class="cas_content"></div>
                        <br>
                        <br>
                        <div id="imgArea" style="display:none; "></div>
                        <br>
                        <div class="uk-panel uk-panel-box uk-text-truncate uk-flex uk-flex-center" style="color:#A4A4A4">【由leoay写作工具生成】</div>
                    </div>

                    <div class="uk-modal-footer uk-text-right">
                        <button class="uk-button uk-button-default uk-modal-close" type="button">取消</button>
                        <button class="uk-button uk-button-primary" id="saveGenCard" type="button">保存</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<script type="text/javascript">
    $(function() {
        var editor = editormd("leoaylog-editor", {
            width  : "100%",
            height : 600,
            path   : "/static/editormd/lib/",
            saveHTMLToTextarea  : true,
            watch  : true,
            showTrailingSpace : false,
            preview : true,
            emoji : true,
            placeholder : "写一篇",
            markdown : {{.markdown}},
            autoHeight : false,
            imageUpload : true,
            lineNumbers : true,
            imageFormats   : ["jpg", "jpeg", "gif", "png", "bmp", "webp", "svg"],
            imageUploadURL : "/picpool/upload",
            toolbarIcons : function() {
                return ["h2", "h3", "h6", "|", "bold", "code", "code-block", "|", "image", "watch", "fullscreen"]
            },
        });
    });
</script>

<script>

    $(".editormd-markdown-textarea").bind('input propertychange', function () {
        alert($(".editormd-html-textarea").val())
    });

    $("#genPic1").click(function() {
        $(".uk-modal-title").text($("#article_title").val());
        $(".cas_content").html($(".editormd-html-textarea").val());
    });
</script>
<script src="/static/js/html2canvas.js"></script>

<script>

function downloadFile(content, fileName) { //下载base64图片
                var base64ToBlob = function(code) {
                    let parts = code.split(';base64,');
                    let contentType = parts[0].split(':')[1];
                    let raw = window.atob(parts[1]);
                    let rawLength = raw.length;
                    let uInt8Array = new Uint8Array(rawLength);
                    for(let i = 0; i < rawLength; ++i) {
                        uInt8Array[i] = raw.charCodeAt(i);
                    }
                    return new Blob([uInt8Array], {
                        type: contentType
                    });
                };
                let aLink = document.createElement('a');
                let blob = base64ToBlob(content); //new Blob([content]);
                let evt = document.createEvent("HTMLEvents");
                evt.initEvent("click", true, true); //initEvent 不加后两个参数在FF下会报错  事件类型，是否冒泡，是否阻止浏览器的默认行为
                aLink.download = fileName;
                aLink.href = URL.createObjectURL(blob);
                aLink.click();
            };

//
$("#saveGenCard").click(function () {
        html2canvas(document.querySelector("#shareCard"), {
            allowTaint: true,
            useCORS: true, 
            taintTest: false, 
            scale: 3.5,
            y: document.querySelector("#shareCard").getBoundingClientRect().top + window.scrollY,
            async: false, // 是否异步解析和呈现元素
            // 以下字段必填
        }).then(canvas => {
            let oImg = new Image();
            oImg.src = canvas.toDataURL();  // 导出图片
            $("#imgArea").append(oImg);
        });

        var base64ToBlob = function(code) {
            let parts = code.split(';base64,');
            let contentType = parts[0].split(':')[1];
            let raw = window.atob(parts[1]);
            let rawLength = raw.length;
            let uInt8Array = new Uint8Array(rawLength);
            for(let i = 0; i < rawLength; ++i) {
                uInt8Array[i] = raw.charCodeAt(i);
            }
            return new Blob([uInt8Array], {
                type: contentType
            });
        };
        let aLink = document.createElement('a');
        let blob = base64ToBlob(imgData); //new Blob([content]);
        let evt = document.createEvent("HTMLEvents");
        evt.initEvent("click", true, true); //initEvent 不加后两个参数在FF下会报错  事件类型，是否冒泡，是否阻止浏览器的默认行为
        aLink.download = "下载.png";
        aLink.href = URL.createObjectURL(blob);
        aLink.click();

    })
</script>

<script>
    //监测从后台获取的可编辑文字、键盘输入的文字字数的变化，并赋值给span
    $(function () {
        var text = $('.summary').val();
        var len = text.length;
        $('.summary').next().find('span').html(len);
        $('.summary').keyup(function () {
            var text = $(this).val();
            len = text.length;
            if (len > 120) {
                return false;
            }
            $(this).next().find('span').html(len);
        })
    });
    </script>

<script>
    document.addEventListener("keydown", function(e) {
        //可以判断是不是mac，如果是mac,ctrl变为花键
        //event.preventDefault() 方法阻止元素发生默认的行为。
        if (e.keyCode == 83 && (navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey)) {
            e.preventDefault();
            // Process event...
            UIkit.notification({message: "保存成功"})
        }
    }, false);
</script>

<script>
    //点击按钮，发布文章
    $("#publish").click(function() {
        //alert("sdsdsd")
        //ToDo: 使用ajax 发送post请求
        $("#article_editor").attr( 'action' , '/writeTool/publish?newarticle={{.newarticle}}&articleid={{.articleid}}');
        $("#article_editor").submit();
    });

    //点击按钮，暂存文章
    $("#savetemp").click(function() {
        //alert("sdsdsd")
        //ToDo: 使用ajax 发送post请求
        $("#article_editor").attr( 'action' , '/writeTool/save_temp');
        $("#article_editor").submit();
    });
</script>

{{/*标题选择联动*/}}
<script>
    $("#article_class1").empty();
    $("#article_class2").empty();
    $(document).ready(function() {
        $.getJSON("/static/topic_class_conf.json", function(data) {
            data['mon_class'].forEach(function(ele1) {
                console.log(ele1.name);
                $("#article_class1").append("<option value=" +  ele1.name.class_name + ">" + ele1.name.class_name + "</option>");
            });
        });
    });

    $(document).ready(function() {
        $.getJSON("/static/topic_class_conf.json", function(data) {
            data['mon_class'].forEach(function(ele1) {
                if($("#article_class1").val() == ele1.name.class_name) {
                    ele1['son_class'].forEach(function(elem2) {
                        $("#article_class2").append("<option value=" + elem2.class_name + ">" + elem2.class_name + "</option>");
                    });
                }
            });
        });
    });

    //从json中加载选项信息
    $("#article_class1").change(function() {
        $("#article_class2").empty();
            $(document).ready(function() {
                $.getJSON("/static/topic_class_conf.json", function(data) {
                    data['mon_class'].forEach(function(ele1) {
                        if($("#article_class1").val() == ele1.name.class_name) {
                            ele1['son_class'].forEach(function(elem2) {
                                $("#article_class2").append("<option value=" + elem2.class_name + ">" + elem2.class_name + "</option>");
                            });
                        }
                    });
                });
            });
        }
    );
    
    //判断当前是否是碎碎念，保证每天只写一篇碎碎念
      $("#article_class2").change(function() {
        if($("#article_class2").val() == "碎碎念") {
            //获取当前日期
            var date = new Date();
            var year = date.getFullYear();
            var month = date.getMonth() + 1;
            var day = date.getDate();
            if (month < 10) {
                month = "0" + month;
            }
            if (day < 10) {
                day = "0" + day;
            }
            var nowDate = "碎碎念" + year + "年" + month + "月" + day + "日";
            $("#article_title").val(nowDate);
        }
    });
  
</script>

<script>
    //初始化标题
    $("#article_title").val({{.artTitle}});
    //初始化类目
     $("#article_class1").val({{.firstClass}});
     $("#article_class2").val({{.secondClass}});
</script>