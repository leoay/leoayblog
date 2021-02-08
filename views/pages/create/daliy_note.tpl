<div class="uk-container uk-container-xlarge main_page">
    <div class="uk-flex uk-width-1 uk-flex-center uk-margin-medium-top uk-margin-medium-bottom">

        <div class="uk-width-2-3 uk-flex-center main_pc_write_article">
            <div class="uk-flex uk-flex-center" style="color:#fff">
                <h1 class="uk-text-light">碎碎念</h2>
            </div>
            {{$is_block := .is_block}}
            {{range $daliy_info := .article_lists}}
                <div class="uk-card uk-card-default uk-card-body uk-margin-medium-top uk-width-1 daliy article">
                    <div class="uk-flex uk-flex-center uk-margin-medium-bottom">
                        <h2 class="uk-text-light">{{ $daliy_info.Title }}</h2>
                    </div>
                    {{str2html $daliy_info.Content }}

                    <div style="display: {{$is_block}}">
                        <button class="uk-button uk-button-default uk-flex uk-flex-right" type="button">修改</button>
                        <div uk-dropdown>
                            <ul class="uk-nav uk-dropdown-nav">
                                <li><a href="/writeTool/{{$daliy_info.Id}}/edit_article">重新编辑</a></li>
                                <li><a id="js-modal-confirm" class="{{$daliy_info.Id}}" href="#">删除</a></li>
                                <li><a href="#">导出pdf</a></li>
                            </ul>
                        </div>
                    </div>

                </div>
            {{end}}

            {{/*分页组件*/}}
            <div class="uk-margin-large-top">
                <ul class="uk-pagination uk-flex-center" uk-margin>
                    {{$pgindex := .pageindex}}
                    <li><a href="/creativeAbility/brokenThoughts?pageIndex={{.prepage}}"><span uk-pagination-previous></span></a></li>
                    {{if gt .pagecount 7}} 
                        {{range $i, $v := .pagetop_arr}}
                            {{if eq $v $pgindex}}
                                <li class="uk-active" style="font-size: 17px"><a href="/creativeAbility/brokenThoughts?pageIndex={{$v}}">{{$v}}</a></li>
                            {{else}}
                                <li class=""><a href="/creativeAbility/brokenThoughts?pageIndex={{$v}}">{{$v}}</a></li>
                            {{end}}
                        {{end}}
                        <li class="uk-disabled"><span>...</span></li>
                        {{range $i, $v := .pagelast_arr}}
                            {{if eq $v $pgindex}}
                                <li class="uk-active" style="font-size: 17px"><a href="/creativeAbility/brokenThoughts?pageIndex={{$v}}">{{$v}}</a></li>
                            {{else}}
                                <li class=""><a href="/creativeAbility/brokenThoughts?pageIndex={{$v}}">{{$v}}</a></li>
                            {{end}}
                        {{end}}
                    {{else}}
                        {{range $i, $v := .pagetop_arr}}
                            {{if eq $v $pgindex}}
                                <li class="uk-active" style="font-size: 17px"><a href="/creativeAbility/brokenThoughts?pageIndex={{$v}}">{{$v}}</a></li>
                            {{else}}
                                <li class=""><a href="/creativeAbility/brokenThoughts?pageIndex={{$v}}">{{$v}}</a></li>
                            {{end}}
                        {{end}}
                    {{end}}

                    <li><a href="/creativeAbility/brokenThoughts?pageIndex={{.nexpage}}"><span uk-pagination-next></span></a></li>
                </ul>
            </div>
        </div>
    </div>
</div>

<script>
    UIkit.util.on('#js-modal-confirm', 'click', function (e) {
        e.preventDefault();
        e.target.blur();
        UIkit.modal.confirm('确认删除？').then(function () {
            console.log(e.target.className)
             $.ajax({
                url: "/writeTool/delete_article",
                type:"post",
                contentType:"application/json",
                data: {id: e.target.className},
                datatype: 'json',
                success:function(msg){
                    console.log("Success!")
                    location.reload();
                },
                error:function(xhr,textstatus,thrown){
                    console.log("Falied!")
                }
            });
        }, function () {
            console.log('Rejected.')
            //拒绝删除
        });
    });
</script>