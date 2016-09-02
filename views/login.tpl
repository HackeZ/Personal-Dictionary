<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8" />
    <title>个人词典 - 登录</title>
    <meta name="author" content="DeathGhost" />
    <link rel="stylesheet" type="text/css" href="./static/css/login.css" tppabs="css/style.css" />
    <style>
        body {
            height: 100%;
            background: #16a085;
            overflow: hidden;
        }
        
        canvas {
            z-index: -1;
            position: absolute;
        }
    </style>
    <script src="//cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
    <script src="./static/js/Particleground.js" tppabs="js/Particleground.js"></script>
    <script>
        $(document).ready(function() {
            //粒子背景特效
            $('body').particleground({
                dotColor: '#5cbdaa',
                lineColor: '#5cbdaa'
            });

            // 提交登录
            var URL = "/login"
            $(".submit_btn").click(function() {
                if ($("#id_username").val() == "" || $("#id_password").val() == "") {
                    $(".login_tips").text("请填写好信息再进行提交～")
                    return
                }
                // AJAX 提交
                $.ajax({
                    type: "POST",
                    url: URL + "?isajax=1",
                    data: $('#form').serialize(),
                    sync: false,
                    error: function() {
                        $(".login_tips").text("网络出问题了，请联系管理员！")
                        $("#id_username").val("").focus();
                        $("#id_password").val("");
                        $("input[name=captcha]").val("");
                    },
                    success: function(data) {
                        if (data.status) {
                            location.href = "/pd"
                        } else {
                            $(".login_tips").text(data.info)
                            $("#id_password").val("").focus();
                            $("input[name=captcha]").val("");
                            $(".captcha-img").click();
                        }
                    }
                });
            });
        });
    </script>
</head>

<body>
    <dl class="admin_login">
        <dt>
  <strong>{{ .Title }}</strong>
 </dt>
        <form id="form">
            <div class="form-item">
                <dd class="user_icon">
                    <input id="id_username" name="username" type="text" placeholder="用户名" class="login_txtbx" tabindex="1" required/>
                </dd>
            </div>
            <dd class="pwd_icon">
                <input id="id_password" name="password" type="password" placeholder="密码" class="login_txtbx" tabindex="2" required/>
            </dd>
            <dd class="val_icon">
                <div class="checkcode">
                    {{create_captcha}}
                    <input name="captcha" placeholder="请输入验证码" type="text" class="login_txtbx" tabindex="3" />
                </div>
            </dd>
            <br>
            <dd>
                <input type="button" value="立即登陆" class="submit_btn" />
            </dd>
            <dd>
                <a class="sign_btn" href="/signup">立即注册</a>
            </dd>
            <dd>
                <p class="login_tips"></p>
            </dd>
            <dd>
                <p>{{ .Copyright | pd_markdown | str2html }}</p>
            </dd>
        </form>

    </dl>
</body>

</html>