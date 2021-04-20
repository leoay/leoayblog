<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">
  <legend>常规时间线</legend>
</fieldset>
    <ul class="layui-timeline">
  <li class="layui-timeline-item">
    <i class="layui-icon layui-timeline-axis"></i>
    <div class="layui-timeline-content layui-text">
      <h3 class="layui-timeline-title">8月18日</h3>
      <p>
        layui 2.0 的一切准备工作似乎都已到位。发布之弦，一触即发。
        <br>不枉近百个日日夜夜与之为伴。因小而大，因弱而强。
        <br>无论它能走多远，抑或如何支撑？至少我曾倾注全心，无怨无悔 <i class="layui-icon"></i>
      </p>
    </div>
  </li>
  <li class="layui-timeline-item">
    <i class="layui-icon layui-timeline-axis"></i>
    <div class="layui-timeline-content layui-text">
      <h3 class="layui-timeline-title">8月16日</h3>
      <p>杜甫的思想核心是儒家的仁政思想，他有<em>“致君尧舜上，再使风俗淳”</em>的宏伟抱负。个人最爱的名篇有：</p>
      <ul>
        <li>《登高》</li>
        <li>《茅屋为秋风所破歌》</li>
      </ul>
    </div>
  </li>
  <li class="layui-timeline-item">
    <i class="layui-icon layui-timeline-axis"></i>
    <div class="layui-timeline-content layui-text">
      <h3 class="layui-timeline-title">8月15日</h3>
      <p>
        中国人民抗日战争胜利日
        <br>常常在想，尽管对这个国家有这样那样的抱怨，但我们的确生在了最好的时代
        <br>铭记、感恩
        <br>所有为中华民族浴血奋战的英雄将士
        <br>永垂不朽
      </p>
    </div>
  </li>
  <li class="layui-timeline-item">
    <i class="layui-icon layui-timeline-axis"></i>
    <div class="layui-timeline-content layui-text">
      <div class="layui-timeline-title">过去</div>
    </div>
  </li>
</ul>

<div id="id_test_video" style="width:100%; height:auto;"></div>



<script src="https://imgcache.qq.com/open/qcloud/video/vcplayer/TcPlayer-2.3.3.js" charset="utf-8"></script>;

<script>
  var player =  new TcPlayer('id_test_video', {
"m3u8": "http://2157.liveplay.myqcloud.com/2157_358535a.m3u8",
"flv": "http://2157.liveplay.myqcloud.com/live/2157_358535a.flv", //增加了一个 flv 的播放地址，用于PC平台的播放 请替换成实际可用的播放地址
"autoplay" : true,      //iOS 下 safari 浏览器，以及大部分移动端浏览器是不开放视频自动播放这个能力的
"poster" : "http://www.test.com/myimage.jpg",
"width" :  '480',//视频的显示宽度，请尽量使用视频分辨率宽度
"height" : '320'//视频的显示高度，请尽量使用视频分辨率高度
});
</script>