<!DOCTYPE html>

<html lang="zh-CN">

<head>
  <title>{{.Title}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <link href="//cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  <link rel="stylesheet" href="./static/css/pd_index.css">
  <script src="//cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
  <script src="//cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>

  <style type="text/css">
    *,
    body {
      margin: 0px;
      padding: 0px;
    }
    
    body {
      margin: 0px;
      font-family: "PingFang SC", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
      box-shadow: inset 0px 0px 100px #ddd;
    }
    /*header,
    footer {
      width: 960px;
      text-align: center;
      margin-left: auto;
      margin-right: auto;
    }
    
    header {
      padding: 100px 0;
    }*/
    
    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }
    
    .description {
      text-align: center;
      font-size: 16px;
    }
    
    a {
      color: #444;
      text-decoration: none;
    }
    /*.backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }*/
    
    section {
      margin-top: 20px;
      margin-bottom: 150px;
      text-align: center;
      margin-right: 200px;
    }
  </style>

</head>

<body>

  <!-- HEADER -->
  <header>
    <!-- NAVBAR -->
    <nav class="navbar navbar-default" role="navigation">
      <div class="container-fluid">
        <div class="navbar-header">
          <a class="navbar-brand" href="#">{{ .Title }}</a>
        </div>
        <div>
          <ul class="nav navbar-nav navbar-left">
            <li>
              <!-- Button trigger modal -->
              <a type="button" class="btn active" data-toggle="modal" data-target="#myModal">
  添加新词典
</a></li>
          </ul>
          <ul class="nav navbar-nav navbar-right">
            <li><a href="/logout">退出登录</a></li>
          </ul>
        </div>
      </div>
    </nav>
    <!-- NAVBAR -->
  </header>
  <!-- HEADER -->

  <section>
    <h4 id="add_pd_tips" style="color: crimson"></h4>
  </section>


  <!-- CONTENT -->

  <section>
    <h2><a href="#PD1" name="PD1">PD 1</a></h2>
    <article>
      <div class="feat">
        <h5 class="pd-date">
        <time datetime>2016-08-20 11:20</time>
        </h5>
      </div>
      This is PD 1.
    </article>
  </section>

  <section>
    <h2><a href="#PD2" name="PD2">PD 2</a></h2>
    <article>
      <div class="feat">
        <h5 class="pd-date">
        <time datetime>2026-08-20 2:20</time>
        </h5>
      </div>
      This is PD 2.
    </article>
  </section>

  <section>
    <h2><a href="#PD3" name="PD3">PD 3</a></h2>
    <article>
      <div class="feat">
        <h5 class="pd-date">
        <time datetime>2036-08-20 3:20</time>
        </h5>
      </div>
      This is PD 3.
    </article>
  </section>
  <!-- CONTENT -->


  <!-- ASIDE -->
  <aside>
    <nav class="navbar navbar-vertical-right">
      <h2>词条导航</h2>
      <ul>
        <li><a href="#PD1">PD 1</a></li>
        <li><a href="#PD2">PD 2</a></li>
        <li><a href="#PD3">PD 3</a></li>
      </ul>
    </nav>
  </aside>
  <!-- ASIDE -->


  <!-- FOOTER -->
  <footer>
    <div class="author">
      <p>{{.Welcome}}</p>
    </div>
  </footer>
  <!-- FOOTER -->

  <div class="backdrop"></div>

  <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">新词典</h4>
        </div>
        <div class="modal-body">
          <form id="pd-form" >
            <div class="form-group">
              <label for="recipient-name" class="control-label">词：</label>
              <input type="text" class="form-control" name="Keyword" id="pd_Keyword">
            </div>
            <div class="form-group">
              <label for="message-text" class="control-label">意义：</label>
              <textarea class="form-control" name="Content" id="pd_Content" data-pd-content=""></textarea>
            </div>
          </form>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          <button type="button" class="btn btn-primary add_pd_btn">保存</button>
        </div>
      </div>
    </div>
  </div>
</body>

<script>
  $(document).ready(function() {
    $(".add_pd_btn").click(function() {
      console.log("I am AJAX");
    $('#myModal').modal('hide')
      // AJAX 提交
      $.ajax({
        type: "POST",
        url: "/pd/add",
        data: $('#pd-form').serialize(),
        sync: false,
        error: function() {
          $("#add_pd_tips").text("网络异常，请重新登录后重试！");
        },
        success: function(data) {
          if (data.status) {
            $("#add_pd_tips").text("词典添加成功，请刷新页面查看！");
            $("#pd_Keyword").val("")
            $("#pd_Content").val("")
          } else {
            $("#add_pd_tips").text(data.info);
          }
        }
      });
    });
  });
</script>

</html>