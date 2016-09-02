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
            var URL = "/signup"
            $(".submit_btn").click(function() {
                if ($("#id_username").val() == "" || $("#id_password").val() == "") {
                    $(".signup_tips").text("请填写好信息再进行提交～")
                    return
                }
                if($("#id_repass").val() !== $("#id_password").val()) {
                    $(".signup_tips").text("密码不匹配，请重新输入！")
                    return
                }
                // AJAX 提交
                $.ajax({
                    type: "POST",
                    url: URL + "?isajax=1",
                    data: $('#form').serialize(),
                    sync: false,
                    error: function() {
                        $(".signup_tips").text("网络出问题了，请联系管理员！")
                        $("#id_username").val("").focus();
                        $("#id_password").val("");
                        $("#id_repass").val("");
                        $("#id_email").val("");
                        $("input[name=captcha]").val("");
                    },
                    success: function(data) {
                        if (data.status) {
                            location.href = "/pd"
                            $(".signup_tips").text(data.info)
                        } else {
                            $(".signup_tips").text(data.info)
                            $("#id_username").val("").focus();
                            $("#id_password").val("");
                            $("#id_repass").val("");
                            $("#id_email").val("");
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
            <dd class="pwd_icon">
                <input id="id_repass" name="repass" type="password" placeholder="重复你的密码" class="login_txtbx" tabindex="3" required/>
            </dd>
            <dd class="email_icon">
                <input id="id_email" name="email" type="email" placeholder="常用邮箱" class="login_txtbx" tabindex="4" required/>
            </dd>
            <dd class="val_icon">
                <div class="checkcode">
                    {{create_captcha}}
                    <input name="captcha" placeholder="请输入验证码" type="text" class="login_txtbx" tabindex="5" />
                </div>
            </dd>
            <br>
            <dd>
                <input type="button" value="立即注册" class="submit_btn" />
            </dd>
            <dd>
                <p class="signup_tips"></p>
            </dd>
            <dd>
                <p>{{ .Copyright | pd_markdown | str2html }}</p>
            </dd>
        </form>

    </dl>
</body>

</html>