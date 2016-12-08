<template>
  <div class="top-nav-bar row border-bottom">
    <nav class="navbar navbar-static-top white-bg" role="navigation" style="margin-bottom: 0">
      <div class="navbar-header">
        <div class="navbar-minimalize">
          <a class="minimalize-styl-2 btn btn-primary">
            <i class="fa fa-bars"></i>
          </a>
        </div>
        <div role="search" class="navbar-form-custom" method="post" action="">
          <div class="form-group">
            <input type="text" placeholder="Search for something..." class="form-control" name="top-search" id="top-search">
          </div>
        </div>
      </div>
      <ul class="nav navbar-top-links navbar-right">
        <li>
          <a ui-sref="login">
            <i class="fa fa-sign-out"></i> 退出
          </a>
        </li>
      </ul>
    </nav>
  </div>
</template>

<style scoped>
.navbar-minimalize {
  float: left;
}
.navbar-minimalize a {
  padding-right: 0px;
  padding-left: 0px;
  margin-left: 15px;
}
.navbar-form-custom {
  margin-left: 15px;
}

</style>