{{/*PC端页面500篇技术文章+500篇非技术文章*/}}
    <div class="uk-container uk-container-xlarge main_page">
        <div class="uk-flex uk-flex-center uk-visible@m uk-margin-medium-top content">
        {{/*=====================================================左边栏*/}}
            <!-- <div class="uk-width-1-4 uk-visible@m uk-margin-small-right left_content">
                <div class="uk-card uk-card-default uk-card-body">
                    <div class="uk-width-expand">
                        <h4 class="uk-card-title">数据结构与算法系列文章</h3>
                        <hr>
                        <h5 class="uk-text-large uk-text-light">数据结构小白</h5>
                        <h5 class="uk-text-large uk-text-light">数据结构高手</h5>
                        <h5 class="uk-text-large uk-text-light">玩转LeetCode</h5>
                    </div>
                </div>

                <div class="uk-card uk-card-default uk-card-body">
                    <div class="uk-width-expand">
                        <h3 class="uk-card-title">Go语言系列文章</h3>
                        <hr>
                        <h5 class="uk-text-large uk-text-light">Go语言基础</h5>
                        <h5 class="uk-text-large uk-text-light">Go语言进阶</h5>
                        <h5 class="uk-text-large uk-text-light">Go语言书籍推荐</h5>
                        <h5 class="uk-text-large uk-text-light">Go语言框架之beego</h5>
                        <h5 class="uk-text-large uk-text-light">Go语言框架之iris</h5>
                        <h5 class="uk-text-large uk-text-light">Go语言框架之区块链开发</h5>
                        <h5 class="uk-text-large uk-text-light">Go语言之嵌入式开发</h5>
                    </div>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-margin-medium-top uk-height-medium uk-text-center uk-margin-medium-bottom show_card">
                    <h5 class="uk-card-title-side uk-text-bold">空白</h5>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-margin-small-top uk-height-medium uk-text-center show_card">
                    <h5 class="uk-card-title-side uk-text-bold">空白</h5>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-height-medium uk-text-center uk-margin-medium-bottom show_card">
                    <br>
                   
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-margin-medium-top uk-height-medium uk-text-center uk-margin-medium-bottom show_card">
                    <h5 class="uk-card-title-side uk-text-bold">友情链接</h5>
                    <hr>
                </div>
            </div> -->

            {{/*===================================================中间栏*/}}

            <div class="uk-width-2-3 main_pc_div uk-margin-large-bottom">

                <div class="uk-flex uk-flex-center" style="color:#fff">
                    <h2 class="uk-text-light uk-margin-small-top" style="font-weight:700;">技术速览</h3>
                </div>

                {{/*图片轮播*/}}
                <div class="uk-position-relative uk-visible-toggle uk-light" tabindex="-1" uk-slideshow="finite: false;ratio: 2:1;animation: scale" autoplay-interval=3000 autoplay="true" pause-on-hover="true" style="display:none">
                <ul class="uk-slideshow-items">
                    <li>
                        <img src="http://getuikit.dev-tang.com/skin/ukv3/images/photo.jpg" alt="" uk-cover>
                    </li>
                    <li>
                        <img src="http://getuikit.dev-tang.com/skin/ukv3/images/dark.jpg" alt="" uk-cover>
                    </li>
                    <li>
                        <img src="http://getuikit.dev-tang.com/skin/ukv3/images/light.jpg" alt="" uk-cover>
                    </li>
                </ul>
                <a class="uk-position-center-left uk-position-small uk-hidden-hover" href="#" uk-slidenav-previous uk-slideshow-item="previous"></a>
                <a class="uk-position-center-right uk-position-small uk-hidden-hover" href="#" uk-slidenav-next uk-slideshow-item="next"></a>
                </div>

                {{/*文章列表*/}}
                <div class="uk-margin-small-top">
                    <!-- <div class="uk-width-auto">
                        <span class="uk-badge">Go</span>
                        <span class="uk-badge">Php</span>
                        <span class="uk-badge">Python</span>
                        <span class="uk-badge">商业</span>
                    </div> -->
                    <hr>

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
                                            <p class="uk-text uk-text-small uk-text-light uk-text-truncate uk-width-5-6" style="margin-top: 13px">{{$daliy_info.Summary }}</p>
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
                            <li><a href="/codeTech/tecCol?pageIndex={{.prepage}}"><span uk-pagination-previous></span></a></li>
                            {{if gt .pagecount 7}} 
                                {{range $i, $v := .pagetop_arr}}
                                    {{if eq $v $pgindex}}
                                        <li class="uk-active" style="font-size: 17px"><a href="/codeTech/tecCol?pageIndex={{$v}}">{{$v}}</a></li>
                                    {{else}}
                                        <li class=""><a href="/codeTech/tecCol?pageIndex={{$v}}">{{$v}}</a></li>
                                    {{end}}
                                {{end}}
                                <li class="uk-disabled"><span>...</span></li>
                                {{range $i, $v := .pagelast_arr}}
                                    {{if eq $v $pgindex}}
                                        <li class="uk-active" style="font-size: 17px"><a href="/codeTech/tecCol?pageIndex={{$v}}">{{$v}}</a></li>
                                    {{else}}
                                        <li class=""><a href="/codeTech/tecCol?pageIndex={{$v}}">{{$v}}</a></li>
                                    {{end}}
                                {{end}}
                            {{else}}
                                {{range $i, $v := .pagetop_arr}}
                                    {{if eq $v $pgindex}}
                                        <li class="uk-active" style="font-size: 17px"><a href="/codeTech/tecCol?pageIndex={{$v}}">{{$v}}</a></li>
                                    {{else}}
                                        <li class=""><a href="/codeTech/tecCol?pageIndex={{$v}}">{{$v}}</a></li>
                                    {{end}}
                                {{end}}
                            {{end}}

                            <li><a href="/codeTech/tecCol?pageIndex={{.nexpage}}"><span uk-pagination-next></span></a></li>
                        </ul>
                    </div>

                </div>
            </div>

            {{/*===================================================右边栏*/}}

            <!-- <div class="uk-width-1-6 uk-visible@m uk-margin-small-left">
                <div class="uk-card uk-card-default uk-card-body">
                     <div class="uk-width-auto">
                        <img class="uk-comment-avatar" src="/static/picture_pool/1606292056122.png" width="30" height="3" alt="" style="border-radius:2px; ">
                    </div>
                    <div class="uk-width-expand">
                        <h3 class="uk-card-title">Go Modules 终极入门</h3>
                    </div>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-margin-small-top uk-height-medium uk-text-center show_card">
                    <h5 class="uk-card-title-side uk-text-bold">站价值观</h5>
                    <hr>
                    <span class="uk-text-medium uk-text-light">让技术给生活添加乐趣</span><br>
                    <span class="uk-text-medium uk-text-light">让技术给别人带来价值</span><br>
                    <span class="uk-text-medium uk-text-light">让技术实现自我的价值</span><br>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-margin-medium-top uk-height-medium uk-text-center uk-margin-medium-bottom show_card">
                    <h5 class="uk-card-title-side uk-text-bold">服务器推荐</h5>
                    <hr>
                    <a href="https://www.aliyun.com/1111/new?userCode=r43iclk0"><img class="uk-comment-avatar uk-background-cover" src="/static/picture_pool/1606967329408.jpg" alt="" style="border-radius:10px;border:solid 1px #d5d8d3"></a><br><br>
                    <a href="https://www.aliyun.com/1111/new?userCode=r43iclk0"><img class="uk-comment-avatar uk-background-cover" src="/static/picture_pool/1606967329408.jpg" alt="" style="border-radius:10px;border:solid 1px #d5d8d3"></a><br><br>
                    <a href="https://www.aliyun.com/1111/new?userCode=r43iclk0"><img class="uk-comment-avatar uk-background-cover" src="/static/picture_pool/1606967329408.jpg" alt="" style="border-radius:10px;border:solid 1px #d5d8d3"></a><br>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-margin-small-top uk-height-medium uk-text-center show_card">
                    <h5 class="uk-card-title-side uk-text-bold">站公众号</h5>
                    <hr>
                    <img class="uk-comment-avatar uk-background-cover" src="/static/siteImgAssert/leoay_labgh.jpg" width="180" height="180" alt="" style="border-radius:25px;border:solid 8px #d5d8d3"><br>
                    <span class="uk-text-middle uk-text-bold">leoay实验室</span><br>
                    <span class="uk-text-small uk-text-lighter uk-text-italic uk-text-success">专注于技术文章的写作与分享，关于算法，编程以及AI</span>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-height-medium uk-text-center uk-margin-medium-bottom show_card">
                    <br>
                    <img class="uk-comment-avatar uk-background-cover" src="/static/siteImgAssert/leoay_labgh.jpg" width="180" height="180" alt="" style="border-radius:25px;border:solid 8px #d5d8d3"><br>
                    <span class="uk-text-middle uk-text-bold">leoay</span><br>
                    <span class="uk-text-small uk-text-lighter uk-text-italic uk-text-success">专注于个人软技能文章的写作，关于赚钱，创作和阅读</span>
                </div>

                <div class="uk-card uk-card-default uk-card-body uk-margin-medium-top uk-height-medium uk-text-center uk-margin-medium-bottom show_card">
                    <h5 class="uk-card-title-side uk-text-bold">友情链接</h5>
                    <hr>
                </div>
            </div> -->

        </div>
    </div>
{{/*======================================================================================================================================*/}}

{{/*移动端页面*/}}
<div class="uk-flex uk-flex-center uk-hidden@m">
    <div class="uk-width-1-1">
        {{/*图片轮播*/}}
        <div class="uk-position-relative uk-visible-toggle uk-light" tabindex="-1" uk-slideshow="finite: false;ratio: 2:1;animation: scale" autoplay-interval=3000 autoplay="true" pause-on-hover="true">
            <ul class="uk-slideshow-items">
                <li>
                    <img src="http://getuikit.dev-tang.com/skin/ukv3/images/photo.jpg" alt="" uk-cover>
                </li>
                <li>
                    <img src="http://getuikit.dev-tang.com/skin/ukv3/images/dark.jpg" alt="" uk-cover>
                </li>
                <li>
                    <img src="http://getuikit.dev-tang.com/skin/ukv3/images/light.jpg" alt="" uk-cover>
                </li>
            </ul>
            <a class="uk-position-center-left uk-position-small uk-hidden-hover" href="#" uk-slidenav-previous uk-slideshow-item="previous"></a>
            <a class="uk-position-center-right uk-position-small uk-hidden-hover" href="#" uk-slidenav-next uk-slideshow-item="next"></a>
        </div>

        {{/*文章列表*/}}
        <div class="uk-margin-media-top">

            <ul class="uk-list">
                <li>
                    <div class="uk-card uk-card-default uk-card-hover uk-card-body">
                        <a href="#"><h4 class="article-list-title">Go Modules 终极入门</h4></a>
                        <p>这是一个测试文章</p>
                        <p>2020年11月26日</p>
                    </div>
                </li>
                <li>
                    <div class="uk-card uk-card-default uk-card-hover uk-card-body">
                        <h4 class="article-list-title">文章1</h4>
                        <p>这是一个测试文章</p>
                        <p>2020年11月26日</p>
                    </div>
                </li>
                <li>
                    <div class="uk-card uk-card-default uk-card-hover uk-card-body">
                        <h4 class="article-list-title">文章1</h4>
                        <p>这是一个测试文章</p>
                        <p>2020年11月26日</p>
                    </div>
                </li>
            </ul>

            其他展示
            <hr>
        </div>
    </div>
</div>