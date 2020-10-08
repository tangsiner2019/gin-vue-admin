<template>
  <div>
    <el-row :gutter="15">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-col :span="24">
          <el-form-item label="姓名" prop="title">
            <el-input v-model="formData.name" placeholder="请输入姓名" :maxlength="50" show-word-limit clearable
              :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="英文名" prop="en_name">
            <el-input v-model="formData.en_name" placeholder="请输入英文名" :maxlength="50" show-word-limit
              clearable :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="上传照片" prop="img_path" required="">
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
              <div slot="tip" class="el-upload__tip">
                只能上传不超过 200KB 的图片文件，参考尺寸：600*720
              </div>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="是否禁用" prop="status" required>
            <el-switch v-model="formData.status" :active-value='-1' :inactive-value='0'></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="是否发布" prop="is_pub" required>
            <el-switch v-model="formData.is_pub" :active-value='1' :inactive-value='0'></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="排序" prop="sort">
            <el-input-number v-model="formData.sort" placeholder="序号"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="内容" prop="content" >
            <quill-editor
              :options="editorOption"
              @blur="onEditorBlur($event)"
              @change="onEditorChange($event)"
              @focus="onEditorFocus($event)"
              ref="myQuillEditor"
              v-model="formData.content"
      ></quill-editor>
          </el-form-item>
        </el-col>
        
      </el-form>
    </el-row>
    <div class="dialog-footer">
      <el-button @click="close">取消</el-button>
      <el-button type="primary" @click="handelConfirm">确定</el-button>
    </div>
  </div>
</template>
<script>
const path = process.env.VUE_APP_BASE_API;
export default {
  components: {},
  props: {
    item: Object,
  },
  data() {
    return {
      formData: {
        name: '',
        en_name: '',
        img_path: null,
        status: 0,
        is_pub: 1,
        sort: 0,
        content: '',
      },
      rules: {
        name: [{
          required: true,
          message: '请输入姓名',
          trigger: 'blur'
        }],
        en_name: [{
          required: true,
          message: '请输入英文名',
          trigger: 'blur'
        }],
        sort: [{
          required: true,
          message: '序号',
          trigger: 'blur'
        }],
        content: [{
          required: true,
          message: '请输入内容',
          trigger: 'blur'
        }],
      },
      editorOption: {
        placeholder: "请输入内容"
      },
      img_pathAction: `${path}/fileUploadAndDownload/upload`,
    }
  },
  filters: {
    formatImgPath: function (picSrc) {
      if (picSrc && picSrc.slice(0, 4) !== "http") {
        return path + picSrc;
      }
      return picSrc;
    },
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {
    if (this.item.ID > 0) {
      this.formData = {
        id: this.item.ID,
        name: this.item.name,
        en_name: this.item.enName,
        content: this.item.content,

        sort: this.item.sort,
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
    onEditorReady() {},// 准备编辑器
    onEditorBlur() {}, // 失去焦点事件
    onEditorFocus() {}, // 获得焦点事件
    onEditorChange() {}, // 内容改变事件

  }
}

</script>
<style scoped>
.el-upload__tip {
  line-height: 1.2;
}
.dialog-footer {
  text-align: right;
}
</style>
