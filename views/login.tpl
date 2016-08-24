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
    <script src="./static/js/verificationNumbers.js" tppabs="js/verificationNumbers.js"></script>
    <script src="./static/js/Particleground.js" tppabs="js/Particleground.js"></script>
    <script>
        $(document).ready(function() {
            //粒子背景特效
            $('body').particleground({
                dotColor: '#5cbdaa',
                lineColor: '#5cbdaa'
            });
            //验证码
            createCode();
            // 提交登录

            var URL = "/login"
            $(".submit_btn").click(function() {
                // AJAX 提交
                $.ajax({
                    type: "POST",
                    url: URL + "?isajax=1",
                    data: {
                        username:$("#id_username").val(),
                        password:$("#id_password").val()
                    },
                    sync: false,
                    error: function() {
                        $(".login_tips").text("网络出问题了，请联系管理员！")
                        $("#id_username").val("").focus();
                        $("#id_password").val("");
                    },
                    success: function(data) {
                        if (data.status) {
                            location.href = "/pb"
                        } else {

                            $(".login_tips").text(data.info)
                            $("#id_username").val("").focus();
                            $("#id_password").val("");
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
                    <input id="id_username" type="text" placeholder="账号" class="login_txtbx" required/>
                </dd>
            </div>
            <dd class="pwd_icon">
                <input id="id_password" type="password" placeholder="密码" class="login_txtbx" required/>
            </dd>
            <dd class="val_icon">
                <div class="checkcode">
                    <input type="text" id="J_codetext" placeholder="验证码" maxlength="4" class="login_txtbx">
                    <canvas class="J_codeimg" id="myCanvas" onclick="createCode()">对不起，您的浏览器不支持canvas，请下载最新版浏览器!</canvas>
                </div>
                <input type="button" value="Check" class="ver_btn" onClick="validate();">
            </dd>
            <dd>
                <input type="button" value="立即登陆" class="submit_btn" />
            </dd>
            <dd>
                <p class="login_tips"></p>
            </dd>
            <dd>
                <p>{{ .Copyright }}</p>
            </dd>
        </form>

    </dl>
</body>

</html>