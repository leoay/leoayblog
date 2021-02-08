<div class="uk-container uk-container-xlarge main_page">
    <div class="uk-flex uk-width-1 uk-flex-center uk-margin-medium-top uk-margin-medium-bottom">

        <div class="uk-width-2-3 uk-flex-center main_article">
            <div class="uk-flex uk-flex-left" style="color:#fff;">
                <p class="uk-text-light uk-width-1 title uk-text-break">{{ .articleInfo.ArticleTitle }}</p>
            </div>

            <div class="uk-margin-medium-top"> 
                <a href="{{.articleInfo.ArticleFirstClassRoute}}">
                    <span class="uk-label">{{.articleInfo.ArticleFirstClass}}</span>
                </a>
                <a href="{{.articleInfo.ArticleSecondClassRoute}}">
                    <span class="uk-label">{{.articleInfo.ArticleSecondClass}}</span>
                </a><br><br>
                <span class="uk-label">
                    作者： {{.articleInfo.ArticleAuthor}} 
                    <label title="{{.articleInfo.ArticleCreateTimeUTC}}">
                        发布于{{.articleInfo.ArticleCreateTimeSub}}
                    </label> | 
                    <label title="{{.articleInfo.ArticleUpdateTimeUTC}}">
                        更新于{{.articleInfo.ArticleUpdateTimeSub}}
                    </lable>

                    <label style="display: {{.is_block}}">
                       <a style="color: #fff" href="/writeTool/{{.articleInfo.Id}}/edit_article">编辑</a>
                    </lable>
                </span>

            </div>
            <hr class="uk-divider-icon">
            <div>
                {{str2html .articleInfo.ArticleContent }}
            </div>

            <div class="uk-child-width-1-6 uk-flex-center uk-flex uk-flex-top uk-grid-small" uk-grid>
                <hr class="uk-width-2-3 uk-divider-small">
            </div>
        </div>
    </div>
</div>