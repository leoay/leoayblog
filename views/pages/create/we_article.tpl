<div class="uk-container uk-container-xlarge main_page">
    <div class="uk-flex uk-width-1 uk-flex-center uk-margin-medium-top uk-margin-medium-bottom">

        <div class="uk-width-2-3 uk-flex-center main_article">
            <div class="uk-flex uk-flex-center" style="color:#fff">
                <h2 class="uk-text-light uk-margin-small-bottom" style="font-weight:700;">公众号文章合集</h3>
            </div>
            
            <ul class="uk-list">
                {{/*一个文章列表*/}}
                {{range $daliy_info := .article_lists}}
                <li>
                    <a href="/show_article/{{$daliy_info.Id}}">
                        <div class="uk-card uk-card-default uk-card-hover uk-card-body article-card">
                            <div class="uk-grid-small uk-flex-middle uk-padding-remove" uk-grid>
                                <div class="uk-width-auto">
                                    <img class="uk-comment-avatar" src="{{ $daliy_info.MainPicPath }}" alt="" style="border-radius:4px;width:235px;height:100px">
                                </div>
                                <div class="uk-width-expand@l">
                                    <h3 class="uk-card-title article-title">{{ $daliy_info.Title }}</h3>
                                    <p class="uk-text uk-text-small uk-text-light uk-text-truncate uk-width-5-6 summary" style="margin-top: 13px">{{$daliy_info.Summary }}</p>
                                </div>
                                <div class="uk-flex uk-width-auto uk-flex-right uk-flex-middle count_info">
                                    <span class="count_set" style="font-size:13px">
                                        <span uk-icon="icon: star;ratio:0.8" uk-tooltip="点赞 0 + 评论 4 + 收藏 0"></span><span uk-tooltip="点赞 0 + 评论 4 + 收藏 0">12</span>
                                        <span class="count_seperator">|</span>
                                        <span class="created_time" title="{{$daliy_info.TimeUTC}}">{{$daliy_info.TimeSub}}</span>
                                    </span>
                                </div>
                            </div>
                        </div>
                    </a>
                </li>
                {{end}}
            </ul>
            
            {{/*分页组件*/}}
            <div class="uk-margin-large-top">
                <ul class="uk-pagination uk-flex-center" uk-margin>
                    {{$pgindex := .pageindex}}
                    <li><a href="/creativeAbility/wePubWriting?pageIndex={{.prepage}}"><span uk-pagination-previous></span></a></li>
                    {{if gt .pagecount 7}} 
                        {{range $i, $v := .pagetop_arr}}
                            {{if eq $v $pgindex}}
                                <li class="uk-active" style="font-size: 17px"><a href="/creativeAbility/wePubWriting?pageIndex={{$v}}">{{$v}}</a></li>
                            {{else}}
                                <li class=""><a href="/creativeAbility/wePubWriting?pageIndex={{$v}}">{{$v}}</a></li>
                            {{end}}
                        {{end}}
                        <li class="uk-disabled"><span>...</span></li>
                        {{range $i, $v := .pagelast_arr}}
                            {{if eq $v $pgindex}}
                                <li class="uk-active" style="font-size: 17px"><a href="/creativeAbility/wePubWriting?pageIndex={{$v}}">{{$v}}</a></li>
                            {{else}}
                                <li class=""><a href="/creativeAbility/wePubWriting?pageIndex={{$v}}">{{$v}}</a></li>
                            {{end}}
                        {{end}}
                    {{else}}
                        {{range $i, $v := .pagetop_arr}}
                            {{if eq $v $pgindex}}
                                <li class="uk-active" style="font-size: 17px"><a href="/creativeAbility/wePubWriting?pageIndex={{$v}}">{{$v}}</a></li>
                            {{else}}
                                <li class=""><a href="/creativeAbility/wePubWriting?pageIndex={{$v}}">{{$v}}</a></li>
                            {{end}}
                        {{end}}
                    {{end}}

                    <li><a href="/creativeAbility/wePubWriting?pageIndex={{.nexpage}}"><span uk-pagination-next></span></a></li>
                </ul>
            </div>

        </div>
    </div>
</div>

{{/*
<script>
    //时间转换
    var time_created = $(".created_time").attr("title")
    var unixTimestamp = new Date(time_created* 1000); 
    //console.log(unixTimestamp.toLocaleString());
    $(".created_time").attr("title", unixTimestamp.toLocaleString());


    //计算时间差
    var dateDiff =  new Date().getTime()-unixTimestamp;//时间差的毫秒数
    //console.log(dateDiff)
    var days = Math.floor(dateDiff / (24 * 3600 * 1000));//计算出相差天数
    //console.log(dayDiff)
     //计算出小时数
    var leave1=dateDiff%(24*3600*1000)    //计算天数后剩余的毫秒数
    var hours=Math.floor(leave1/(3600*1000))
    //计算相差分钟数
    var leave2=leave1%(3600*1000)        //计算小时数后剩余的毫秒数
    var minutes=Math.floor(leave2/(60*1000))
    //计算相差秒数
    var leave3=leave2%(60*1000)      //计算分钟数后剩余的毫秒数
    var seconds=Math.round(leave3/1000)
    console.log(" 相差 "+days+"天 "+hours+"小时 "+minutes+" 分钟"+seconds+" 秒")
    //$(".created_time").text("sdfsdfsdf");
    //以第一个不为0的时间差单位算起
    if(days >0) {
        $(".created_time").text(days + "天前");
    }else if(hours > 0) {
        $(".created_time").text(hours + "小时前");
    }else if(minutes > 0) {
        $(".created_time").text(minutes + "分钟前");
    }else if(seconds >0 ) {
        $(".created_time").text(seconds + "秒前");
    }


</script>
*/}}