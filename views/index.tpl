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
          <ul class="nav navbar-nav navbar-right">
            <li><a href="/logout">退出登录</a></li>
          </ul>
        </div>
      </div>
    </nav>
    <!-- NAVBAR -->
  </header>
  <!-- HEADER -->

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
</body>

</html>