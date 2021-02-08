<header class="header">
    <div class="uk-flex uk-flex-center" style="margin-top:0px">
        <div class="uk-width-1">
        <div style="height:10px;background:#234"></div>
    {{/* 移动端导航栏 
    */}}
        <!-- <div uk-sticky="show-on-up: true; animation: uk-animation-slide-top; sel-target: .uk-navbar-container; cls-active: uk-navbar-sticky; cls-inactive: uk-navbar-transparent;" class="uk-sticky">
            <nav class="uk-navbar uk-navbar-container uk-hidden@m uk-navbar-transparent">
                <div class="uk-navbar-left">
                    <a class="uk-navbar-item uk-logo" href="/">leoay 实验室</a>
                </div>
                <div class="uk-navbar-right">
                    <a class="uk-navbar-toggle uk-icon uk-navbar-toggle-icon" uk-navbar-toggle-icon="" href="#m-nav" uk-toggle=""><svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" data-svg="navbar-toggle-icon"><rect y="9" width="20" height="2"></rect><rect y="3" width="20" height="2"></rect><rect y="15" width="20" height="2"></rect></svg></a>
                </div>
            </nav>
        </div>
        <div class="uk-sticky-placeholder" style="height: 80px; margin: 0px;" hidden=""></div> -->

        {{/* 移动端侧边栏 */}}
        <!-- <div id="m-nav" uk-offcanvas="overlay: true" class="uk-offcanvas">
            <div class="uk-offcanvas-bar">
                <button class="uk-offcanvas-close uk-icon uk-close" type="button" uk-close=""></button>

                <ul class="uk-nav-default uk-nav-parent-icon uk-nav" uk-nav="">
                    
                        <div class="uk-flex uk-flex-center">
                            <div class="uk-flex uk-flex-center uk-margin-medium-top" style="display: {{.is_block}}">
                                <img class="uk-comment-avatar" src="http://leoay.xyz:8085/static/picture_pool/1609840506814.svg" width="80" height="80" alt="" style="border-radius:50%; ">
                            </div>
                        </div>

                    <a class="uk-navbar-item uk-logo" href="/">leoay 实验室</a>
                    <li class="uk-active"><a href="/"><span class="uk-margin-small-right uk-icon" uk-icon="icon: home"></span>首页</a></li>
                    <li class="uk-parent">
                        <a href="/codeTec/teccol"><span class="uk-margin-small-right uk-icon" uk-icon="icon: file-text"></span>代码技术</a>
                        <ul class="uk-nav-sub" aria-hidden="true" hidden="">
                            <li class="uk-active"><a href="#"><span uk-icon="icon:  album; ratio: 1"></span>技术专栏</a></li>
                            <li><a href="#"><span uk-icon="icon:  github; ratio: 1"></span>开源项目</a></li>
                            <li><a href="#"><span uk-icon="icon:  cog; ratio: 1"></span>实用工具</a></li>
                            <li><a href="#"><span uk-icon="icon:  happy; ratio: 1"></span>漫画技术</a></li>
                            <li><a href="#"><span uk-icon="icon: tripadvisor; ratio: 1"></span>名人故事</a></li>
                            <li><a href="#"><span uk-icon="icon: tripadvisor; ratio: 1"></span>航模实验室</a></li>
                        </ul>
                    </li>

                    <li class="uk-parent">
                        <a href="#"><span class="uk-margin-small-right uk-icon" uk-icon="icon: file-text"></span>创作笔记</a>
                        <ul class="uk-nav-sub" aria-hidden="true" hidden="">
                            <li class="uk-active"><a href="/creativeAbility/wePubWriting">公众号写作</a></li>
                            <li><a href="#">知乎写作</a></li>
                            <li><a href="#">画画的baby</a></li>
                            <li><a href="#">布鲁斯，布鲁斯</a></li>
                            <li><a href="#">视频创作记录</a></li>
                            <li><a href="#">我和摄影</a></li>
                            <li><a href="/creativeAbility/brokenThoughts">碎碎念</a></li>
                        </ul>
                    </li>

                    <li class="uk-parent">
                        <a href="#"><span class="uk-margin-small-right uk-icon" uk-icon="icon:  cog"></span>商业研习</a>
                        <ul class="uk-nav-sub" aria-hidden="true" hidden="">
                            <li class="uk-active"><a href="#">基金与股票研究</a></li>
                            <li><a href="#">商业故事</a></li>
                            <li><a href="#">商业人的基本功</a></li>
                            <li><a href="#">地产研究</a></li>
                        </ul>
                    </li>
                    
                    <li class="uk-parent">
                        <a href="#"><span class="uk-margin-small-right uk-icon" uk-icon="icon: file-text"></span>关于</a>
                        <ul class="uk-nav-sub" aria-hidden="true" hidden="">
                            <li class="uk-active"><a href="/about/leoaylab"><span uk-icon="icon:  world; ratio: 1"></span>网站介绍</a></li>
                            <li><a href="/about/leoay"><span uk-icon="icon: info; ratio: 1"></span>作者介绍</a></li>
                            <li><a href="#"><span uk-icon="icon: social; ratio: 1"></span>分享本站</a></li>
                        </ul>
                    </li>

                    <li class="uk-nav-divider"></li>
                    <div style="display: {{.is_load_block}}">
                        <li><a class="uk-button uk-button-primary uk-margin-small-top uk-width-2-5 uk-border-pill" href="/user/login" style="line-height:18px"><span uk-icon="icon: sign-in; ratio: 1"></span>登陆</a></li>
                        <li><a class="uk-button uk-button-default uk-width-2-5 uk-border-pill uk-margin-small-top" href="/user/signup" style="line-height:18px"><span uk-icon="icon: sign-in; ratio: 1"></span>注册</a></li>
                    </div>
                    <div style="display: {{.is_block}}">
                        <li><a class="uk-button uk-button-default uk-margin-small-top uk-width-2-5 uk-border-pill" href="/user/logout" style="line-height:18px"><span uk-icon="icon: sign-in; ratio: 1"></span>登出</a></li>
                    </div>
                    
                    <div id="modal-center" class="uk-flex-top" uk-modal>
                        <div class="uk-modal-dialog uk-modal-body uk-margin-auto-vertical">
                            <button class="uk-modal-close-default" type="button" uk-close></button>
                            <p>对不起，此功能暂时未开放！</p>
                        </div>
                    </div>

                    <li class="uk-nav-divider"></li>

                    <button class="uk-button uk-margin-small-top uk-button-default uk-width-2-5 uk-border-pill" type="button" style="line-height:18px">主题选择</button>
                    <div uk-dropdown="animation: uk-animation-slide-top-small; duration: 1000" style="background:#222;box-shadow: 0 5px 12px rgb(255 255 255 / 15%)">
                        <ul class="uk-nav uk-dropdown-nav">
                            <li><a href="#">Dark(Visual Studio)</a></li>
                            <li class="uk-nav-divider"></li>
                            <li><a href="#">Light(Visual Studio)</a></li>
                            <li class="uk-nav-divider"></li>
                            <li><a href="#">Red(Visual Studio)</a></li>
                        </ul>
                    </div>

                    <li class="uk-nav-divider"></li>

                    <button class="uk-button uk-margin-small-top uk-button-default uk-width-2-5 uk-border-pill" type="button" style="line-height:18px">站公众号</button>
                    <div uk-dropdown="animation: uk-animation-slide-top-small; duration: 1000" style="background:#222;box-shadow: 0 5px 12px rgb(255 255 255 / 15%)">
                        <ul class="uk-nav uk-dropdown-nav">
                            <li>
                                <span class="uk-badge">ID: leoay_lab</span>
                                <li class="uk-nav-divider"></li>
                                <img class="uk-comment-avatar" src="/static/picture_pool/1606214880708640.png" width="180" alt="">
                                <li class="uk-nav-divider"></li>
                                <span class="uk-badge uk-margin-small-top">ID: leoay_think</span>
                                <li class="uk-nav-divider"></li>
                                <img class="uk-comment-avatar" src="/static/picture_pool/1606214880708640.png" width="180" alt="">
                            </li>
                            <li class="uk-nav-divider"></li>
                            
                        </ul>
                    </div>

                    <li class="uk-nav-divider"></li>

                    <div>
                    <div class='homepage__find-wordcloud-canvas-container'>
                        <canvas height='800' id='homepage__find-wordcloud-canvas' width='800'></canvas>
                    </div>
                    <div class='homepage__find-wordcloud-list' id='homepage__find-wordcloud-list'>
                        <ul>
                            <li>
                                <a href="#" target="_blank">Go语言</a>
                            </li>
                            <li>
                                <a href="#" target="_blank">红黑树</a>
                            </li>
                            <li>
                                <a href="#" target="_blank">二叉树</a>
                            </li>
                            <li>
                                <a href="#" target="_blank">Python</a>
                            </li>
                            <li>
                                <a href="#" target="_blank">Tensorflow</a>
                            </li>
                            <li>
                                <a href="#" target="_blank">Pytorch</a>
                            </li>
                        </ul>
                    </div>
                </ul>
            </div>
        </div> -->

        {{/* PC端导航栏
            */}}
        <div uk-sticky="show-on-up: true; animation: uk-animation-slide-top; sel-target: .uk-navbar-container; cls-active: uk-navbar-sticky; cls-inactive: uk-navbar-transparent;" class="uk-sticky">
            <nav class="uk-navbar-container uk-visible@m uk-navbar uk-navbar-transparent" uk-navbar="">
                <div class="uk-navbar-left uk-width-1">
                    <a class="uk-navbar-item uk-logo" href="/">leoay 技术圈</a>
                    <ul class="uk-navbar-nav">
                        <li class="uk-active nav-topic"><a href="/">首页</a></li>
                            <li>
                                <li>
                                    <a href="#" class="nav-topic">代码技术</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="/codeTech/tecCol"><span uk-icon="icon:  album; ratio: 1"></span>技术速览</a></li>
                                           <li><a href="#"><span uk-icon="icon:  github; ratio: 1"></span>专栏系列</a></li>
                                           <!--  <li><a href="#"><span uk-icon="icon:  cog; ratio: 1"></span>实用工具</a></li>
                                            <li><a href="#"><span uk-icon="icon:  happy; ratio: 1"></span>漫画技术</a></li>
                                            <li><a href="#"><span uk-icon="icon: tripadvisor; ratio: 1"></span>名人故事</a></li>
                                            <li><a href="#"><span uk-icon="icon: tripadvisor; ratio: 1"></span>航模实验室</a></li> -->
                                        </ul>
                                    </div>
                                </li>

                                <li>
                                    <a href="javascript:void(0)">创作能力</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="/creativeAbility/wePubWriting?pageIndex=1">公众号</a></li>
                                  <!--      <li><a href="#">知乎</a></li>
                                            <li><a href="#">画画</a></li>
                                            <li><a href="#">布鲁斯</a></li>
                                            <li><a href="#">视频</a></li>
                                            <li><a href="#">摄影</a></li> -->
                                            <li><a href="/creativeAbility/brokenThoughts?pageIndex=1">碎碎念</a></li>
                                        </ul>
                                    </div>
                                </li>
<!--
                                <li>
                                    <a href="#" class="nav-topic">商业研习</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="#">基金与股票研究</a></li>
                                            <li><a href="#">商业故事</a></li>
                                            <li><a href="#">商业人的基本功</a></li>
                                            <li><a href="#">leoay私董会</a></li>
                                            <li><a href="#">地产研究</a></li>
                                        </ul>
                                    </div>
                                </li>
-->
<!--
                                <li>
                                    <a href="#" class="nav-topic">知识星球归档</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="#">Python星球</a></li>
                                            <li><a href="#">leoay</a></li>
                                            <li><a href="#"></a></li>
                                            <li><a href="#"></a></li>
                                        </ul>
                                    </div>
                                </li>
-->
<!-- 注释
                                <li>
                                   <a href="#" class="nav-topic">好物推荐</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="#"></a>leoay自营</li>
                                            <li><a href="#">leoay</a>其他平台商品推荐</li>
                                        </ul>
                                    </div>
                                </li>
-->
                                <li>
                                    <a href="#" class="nav-topic">关于</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="/about/leoaylab">网站介绍</a></li>
                                            <li><a href="/about/leoay">作者介绍</a></li>
                                            <li><a href="#"><span uk-icon="icon: social; ratio: 1"></span>分享本站</a></li>
                                        </ul>
                                    </div>
                                </li>

                                <li>
                                    <a href="#" class="nav-topic">快捷导航</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li><a href="https://www.iodraw.com/diagram/" target="_blank">思维导图工具</a></li>
                                            <li><a href="https://www.uupoop.com/" target="_blank">在线PS工具(稿定设计)</a></li>
                                            <li><a href="https://docs.qq.com/" target="_blank">腾讯文档</a></li>
                                            <li><a href="http://getuikit.com/" target="_blank">UIkit</a></li>
                                            <li><a href="https://www.mdnice.com/" target="_blank">mdnice</a></li>
                                            <li><a href="https://www.mdnice.com/" target="_blank">知乎</a></li>
                                            <li><a href="https://www.mdnice.com/" target="_blank">B站</a></li>
                                            <li><a href="https://www.mdnice.com/" target="_blank">知识星球</a></li>
                                            <li><a href="https://www.mdnice.com/" target="_blank">力扣</a></li>
                                            
                                        </ul>
                                    </div>
                                </li>

                                <li>
                                    <div class="uk-visible@l">
                                        <a class="uk-navbar-toggle uk-icon uk-search-icon" href="#" uk-search-icon="" aria-expanded="false">
                                            <svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" data-svg="search-icon"><circle fill="none" stroke="#000" stroke-width="1.1" cx="9" cy="9" r="7"></circle><path fill="none" stroke="#000" stroke-width="1.1" d="M14,14 L18,18 L14,14 Z"></path>
                                            </svg>
                                        </a>
                                            <div class="uk-drop" uk-drop="mode: click; pos: right-center; offset: 0">
                                            <div class="uk-grid-small uk-flex-middle uk-grid uk-grid-stack" uk-grid="">
                                        <div class="uk-width-expand">
                                            <form class="uk-search uk-search-navbar uk-width-1-1">
                                                <input class="uk-search-input" type="search" placeholder="搜索..." autofocus="">
                                            </form>
                                        </div>
                                        <div class="uk-width-auto">
                                            <a class="uk-navbar-dropdown-close uk-icon uk-close" href="#" uk-close=""><svg width="14" height="14" viewBox="0 0 14 14" xmlns="http://www.w3.org/2000/svg" data-svg="close-icon"><line fill="none" stroke="#000" stroke-width="1.1" x1="1" y1="1" x2="13" y2="13"></line><line fill="none" stroke="#000" stroke-width="1.1" x1="13" y1="1" x2="1" y2="13"></line></svg></a>
                                        </div>
                                            </div>
                                        </div>
                                    </div>
                                </li>
                            </li>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <div class="uk-margin-large-left"></div>
                            <!-- <img class="uk-comment-avatar" src="http://leoay.xyz:8085/static/picture_pool/1609839383316.svg" width="30" height="30" alt="" style="border-radius:50%; "> -->
                            <li>
                                <li class="uk-flex-right uk-margin-medium-left" style="display: {{.is_block}}">
                                    <a class="uk-icon" uk-icon="icon: triangle-down;ratio: 1"><span class="uk-icon" uk-icon="icon:  pencil;ratio: 1.5"></span></a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="/writeTool">写博文</a></li>
                                            <li><a href="#">写专栏</a></li>
                                            <li><a href="#">草稿</a></li>
                                        </ul>
                                    </div>
                                </li>
                                <li style="display: {{.is_block}}">
                                    <a>{{.user.Name}}</a>
                                    <div class="uk-navbar-dropdown">
                                        <ul class="uk-nav uk-navbar-dropdown-nav">
                                            <li class="uk-active"><a href="#">我的博客</a></li>
                                            <li class="uk-active"><a href="#">个人中心</a></li>
                                            <li><a href="#"><i class="fa fa-meetup" aria-hidden="true"></i>我的收藏</a></li>
                                            <li><a href="/user/logout"><i class="fa fa-sign-out" aria-hidden="true"></i>退出登陆</a></li>
                                        </ul>
                                    </div>
                                </li>
                                <li style="display: {{.is_block}}">
                                    <a><img class="uk-comment-avatar" src="http://leoay.xyz:8085/static/picture_pool/1609839383316.svg" width="30" height="30" alt="" style="border-radius:50%; "></a>
                                </li>
                            </li>
                            <li style="display: {{.is_load_block}}" class="uk-width-auto">
                                <a  href="#modal-login-box" uk-toggle><i class="fa fa-sign-in" aria-hidden="true"></i>登陆</a>
                            </li>
                            <li style="display: {{.is_load_block}}">
                                <a href=""><i class="fa fa-user-plus" aria-hidden="true"></i> 注册</a>
                            </li>
                        </li>
                    </ul>
                </div>
                <div class="uk-navbar-right uk-margin-large-right">
                    <div class="uk-navbar-item uk-visible@l">
                       
                    </div>
                </div>
            </nav>
        </div>
    </div>

    {{/*弹出登陆界面*/}}
    <div id="modal-login-box" class="uk-flex-top uk-animation-fade" uk-modal>
        <div class="uk-modal-dialog uk-modal-body uk-margin-auto-vertical">
            <button class="uk-modal-close-default uk-icon uk-close" type="button" uk-close=""><svg width="14" height="14" viewBox="0 0 14 14" xmlns="http://www.w3.org/2000/svg" data-svg="close-icon"><line fill="none" stroke="#000" stroke-width="1.1" x1="1" y1="1" x2="13" y2="13"></line><line fill="none" stroke="#000" stroke-width="1.1" x1="13" y1="1" x2="1" y2="13"></line></svg></button>
            <div class="uk-section uk-section-small uk-flex uk-flex-middle uk-text-center uk-form-login" style="min-height: calc(35vh);">
                <div class="uk-width-1-1">
                    <div class="uk-container" style="max-width: 330px;" uk-margin-medium-top uk-height-viewport="expand: true">
                    <p class="uk-text-lead">登录</p>
                    <form action="/user/login" method="POST">
                        <div class="uk-margin">
                            <div class="uk-inline uk-width-1-1">
                                <span class="uk-form-icon uk-icon" uk-icon="icon: user"><svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" data-svg="user"><circle fill="none" stroke="#000" stroke-width="1.1" cx="9.9" cy="6.4" r="4.4"></circle><path fill="none" stroke="#000" stroke-width="1.1" d="M1.5,19 C2.3,14.5 5.8,11.2 10,11.2 C14.2,11.2 17.7,14.6 18.5,19.2"></path></svg></span>
                                <input class="uk-input" name="username" type="text" placeholder="电子邮箱/手机号/用户名">
                            </div>
                        </div>
                        <div class="uk-margin">
                            <div class="uk-inline uk-width-1-1">
                                <span class="uk-form-icon uk-icon" uk-icon="icon: lock"><svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" data-svg="lock"><rect fill="none" stroke="#000" height="10" width="13" y="8.5" x="3.5"></rect><path fill="none" stroke="#000" d="M6.5,8 L6.5,4.88 C6.5,3.01 8.07,1.5 10,1.5 C11.93,1.5 13.5,3.01 13.5,4.88 L13.5,8"></path></svg></span>
                                <input class="uk-input" name="password" type="password" placeholder="密码">
                            </div>
                        </div>
                        <div class="uk-margin">
                            <div class="uk-inline">
                                <label><input class="uk-checkbox" type="checkbox"> 记住登录状态</label>
                            </div>
                        </div>
                        <button class="uk-button uk-button-primary uk-width-1-1">登录</button>
                        <br><br>
                        <a uk-icon="icon: github; ratio: 2" href="https://github.com/login/oauth/authorize?client_id=d0c10af3bba71a8d6753&redirect_uri=http://leoay.xyz:8085/api/github_oauth"></a>
                        <i class="fa fa-weixin" aria-hidden="true"></i>
                    </form>
                </div>
            </div>
        </div>
    </div>

</header>