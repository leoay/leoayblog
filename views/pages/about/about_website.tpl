<div class="uk-flex uk-flex-center" style="">
<div class="uk-width-5-5 uk-margin">

    <div class="uk-card uk-card-default uk-card-body uk-width-1">

        <div class="uk-flex uk-flex-center" style="">
            <h4>本站介绍</h4>
        </div>

    <h6 class="uk-card-title">这个网站的基本信息及技术介绍</h6>
    <p>这个网站主要用来记录leoay学习与实践的历程，顺便赚点小钱！</p>

    {{/* 时间线 */}}
    <div class="uk-section uk-section uk-preserve-color">
    <div class="uk-container">
        <fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
            <legend>leoay实验室组建时间线</legend>
        </fieldset>
        <ul class="layui-timeline">
        
            <li class="layui-timeline-item">
                <i class="layui-timeline-axis"><span uk-icon="tag"></span></i>
                <div class="layui-timeline-content layui-text">
                    <h3 class="layui-timeline-title">时间点： 8月18日</h3>
                    <p>
                        事件描述
                    </p>
                </div>
            </li>

            <li class="layui-timeline-item">
                <i class="layui-timeline-axis"><span uk-icon="tag"></span></i>
                <div class="layui-timeline-content layui-text">
                    <h3 class="layui-timeline-title">时间点： 8月18日</h3>
                    <p>
                        事件描述
                    </p>
                </div>
            </li>

            <li class="layui-timeline-item">
                <i class="layui-timeline-axis"><span uk-icon="tag"></span></i>
                <div class="layui-timeline-content layui-text">
                    <h3 class="layui-timeline-title">时间点： 8月18日</h3>
                    <p>
                        事件描述
                    </p>
                </div>
            </li>
        </ul>
    </div>
</div>

</div>

{{/* 时间线 */}}

</div>
</div>


<div id="id_test_video1" style="width:100%; height:auto;" style="display:block"></div>

<script src="https://imgcache.qq.com/open/qcloud/video/vcplayer/TcPlayer-2.3.3.js" charset="utf-8"></script>;

<script>
  var player =  new TcPlayer('id_test_video', {
"m3u8": "/static/video/out.m3u8",
"flv": "/static/video/out.flv", //增加了一个 flv 的播放地址，用于PC平台的播放 请替换成实际可用的播放地址
"autoplay" : true,      //iOS 下 safari 浏览器，以及大部分移动端浏览器是不开放视频自动播放这个能力的
"poster" : "http://www.test.com/myimage.jpg",
"width" :  '480',//视频的显示宽度，请尽量使用视频分辨率宽度
"height" : '320'//视频的显示高度，请尽量使用视频分辨率高度
});
</script>