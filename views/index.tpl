<!DOCTYPE html>

<html>

<head>
  <title>{{.Title}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,
    body {
      margin: 0px;
      padding: 0px;
    }
    
    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }
    
    header,
    footer {
      width: 960px;
      text-align: center;
      margin-left: auto;
      margin-right: auto;
    }
    
    header {
      padding: 100px 0;
    }
    
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
    
    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
  <link href="//cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  <script src="//cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
  <script src="//cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>

<body>
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


  <!-- HEADER -->
  <header>
    <div class="navbar">
      <div class="nav">
        <ul>
          <li>
            <a href="##">{{ .Title }}</a>
            <span class="cursor"></span>
          </li>
        </ul>
      </div>
    </div>
  </header>
  <!-- HEADER -->

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