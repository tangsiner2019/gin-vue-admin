<template>
  <div>
    <el-row :gutter="15">
      <el-form
        ref="elForm"
        :model="formData"
        :rules="rules"
        size="medium"
        label-width="100px"
      >
        <el-col :span="24">
          <el-form-item label="标题" prop="title">
            <el-input
              v-model="formData.title"
              placeholder="请输入标题"
              :maxlength="50"
              show-word-limit
              clearable
              :style="{ width: '100%' }"
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="链接" prop="link">
            <el-input
              v-model="formData.link"
              placeholder="请输入以http://或https://开头全地址，如果留空则不跳转"
              clearable
              :style="{ width: '100%' }"
            >
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="22">
          <el-form-item label="上传图片" prop="img_path">
            <el-upload
              ref="img_path"
              :action="img_pathAction"
              :before-upload="img_pathBeforeUpload"
              :on-success="uploadSuccess"
              list-type="picture-card"
              :show-file-list="false"
              accept="image/*"
            >
              <img
                v-if="formData.img_path"
                :src="formData.img_path | formatImgPath"
                style="width: 100%"
              />
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
              <!-- <i class="el-icon-plus"></i> -->
              <div slot="tip" class="el-upload__tip">
                只能上传不超过 200KB 的图片文件，参考尺寸：1920*680
              </div>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否禁用" prop="status">
            <el-switch
              v-model="formData.status"
              :active-value="-1"
              :inactive-value="0"
            ></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否发布" prop="is_pub">
            <el-switch
              v-model="formData.is_pub"
              :active-value="1"
              :inactive-value="0"
            ></el-switch>
          </el-form-item>
        </el-col>
      </el-form>
    </el-row>
    <!-- <slot></slot> -->
    <div class="dialog-footer">
      <el-button @click="close">取消</el-button>
      <el-button type="primary" @click="handelConfirm">确定</el-button>
    </div>
  </div>
</template>
<script>
const path = process.env.VUE_APP_BASE_API;
export default {
  inheritAttrs: false,
  components: {},
  props: {
    item: Object,
  },
  data() {
    return {
      formData: {
        title: "",
        link: undefined,
        img_path: null,
        status: 0,
        is_pub: 0,
      },
      rules: {
        title: [
          {
            required: true,
            message: "请输入标题",
            trigger: "blur",
          },
          {
            required: true,
            message: "请上传图片",
            trigger: "change",
          },
        ],
        img_path: [
          {
            required: true,
            message: "请上传图片",
            trigger: "change",
          },
        ],
        link: [],
      },
      img_pathAction: `${path}/fileUploadAndDownload/upload`,
    };
  },
  computed: {},
  filters: {
    formatImgPath: function (picSrc) {
      if (picSrc && picSrc.slice(0, 4) !== "http") {
        return path + picSrc;
      }
      return picSrc;
    },
  },
  watch: {},
  created() {},
  mounted() {
    if (this.item.ID > 0) {
      this.formData = {
        id: this.item.ID,
        title: this.item.title,
        link: this.item.link,
        img_path: this.item.imgPath,
        status: this.item.status,
        is_pub: this.item.isPub,
      };
    }
  },
  methods: {
    close() {
      this.$refs["elForm"].resetFields();
      this.$emit("close", null);
    },
    handelConfirm() {
      this.$refs["elForm"].validate((valid) => {
        if (!valid) return;
        this.$emit("submit", this.formData);
      });
    },
    img_pathBeforeUpload(file) {
      let isRightSize = file.size / 1024 < 200;
      if (!isRightSize) {
        this.$message.error("文件大小超过不能200KB");
      }
      let isAccept = new RegExp("image/*").test(file.type);
      if (!isAccept) {
        this.$message.error("应该选择图片类型的文件");
      }
      return isRightSize && isAccept;
    },
    uploadSuccess(res) {
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "上传成功",
        });
        this.formData.img_path = res.data.file.url;
      } else {
        this.$message({
          type: "warning",
          message: res.msg,
        });
      }
    },
  },
};
</script>
<style scoped>
.el-upload__tip {
  line-height: 1.2;
}
.dialog-footer {
  text-align: right;
}
</style>