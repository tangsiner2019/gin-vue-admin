<template>
  <div class="bottom-info">
    <!-- <div>
      <span>Powered by</span>
      <span>
      <a href="https://www.bomqi.com" target="_blank">Bomqi Tec. Stu.</a>
      </span>
      <el-divider direction="vertical"></el-divider>
      <span>© 2016-2020. 黑麦设计 All Rights Reserved.</span>
      <span>
      <a href="http://www.ryedesign.cn" target="_blank">成都黑麦设计顾问有限公司</a>
      </span>
    </div> -->
  </div>
</template>
<script>
export default {
    name:"BottomInfo"
}
</script>
<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span{
      margin: 0 3px;
    }
  }
}
</style>
