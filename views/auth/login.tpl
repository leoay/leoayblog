<div class="uk-section cdsc uk-section-small uk-flex uk-flex-middle uk-text-center uk-form-login" style="min-height: calc(35vh);">
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
