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
        <el-col :span="8">
          <el-form-item label="分类" prop="cate_id">
            <el-select
              v-model="formData.cate_id"
              placeholder="请选择分类"
              clearable
              :style="{ width: '100%' }"
            >
              <el-option
                v-for="(item, index) in cates"
                :key="index"
                :label="item.name"
                :value="item.ID"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="标题" prop="title">
            <el-input
              v-model="formData.title"
              placeholder="请输入标题"
              :maxlength="20"
              show-word-limit
              clearable
              :style="{ width: '100%' }"
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="副标题" prop="sub_title">
            <el-input
              v-model="formData.sub_title"
              placeholder="请输入副标题"
              :maxlength="20"
              show-word-limit
              clearable
              :style="{ width: '100%' }"
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="题图" prop="img_path" required>
            <el-upload
              ref="img_path"
              :file-list="img_pathfileList"
              :action="img_pathAction"
              :before-upload="img_pathBeforeUpload"
              list-type="picture-card"
              accept="image/*"
            >
              <i class="el-icon-plus"></i>
              <div slot="tip" class="el-upload__tip">
                只能上传不超过 200KB 的image/*文件
              </div>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="是否禁用" prop="status" required>
            <el-switch
              v-model="formData.status"
              :active-value="-1"
              :inactive-value="0"
            ></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="是否发布" prop="field111" required>
            <el-switch v-model="formData.field111"></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="排序" prop="sort">
            <el-input-number
              v-model="formData.sort"
              placeholder="序号"
            ></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="内容图片" prop="imgs" required>
            <el-upload
              ref="imgs"
              :file-list="imgsfileList"
              :action="imgsAction"
              multiple
              :before-upload="imgsBeforeUpload"
              list-type="picture-card"
              accept="image/*"
            >
              <i class="el-icon-plus"></i>
              <div slot="tip" class="el-upload__tip">
                只能上传不超过 200KB 的image/*文件
              </div>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="标签" prop="tags">
            <el-input
              v-model="formData.tags"
              placeholder="多个标签请使用英文逗号（ , ）分隔"
              clearable
              :style="{ width: '100%' }"
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="内容" prop="content">
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
    <div style="text-align:right;">
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
    cates: Array,
  },
  data() {
    return {
      formData: {
        cate_id: null,
        title: "",
        sub_title: "",
        img_path: null,
        status: 0,
        field111: true,
        sort: 0,
        imgs: null,
        tags: "",
        content: "",
      },
      rules: {
        cate_id: [
          {
            required: true,
            message: "请选择分类",
            trigger: "change",
          },
        ],
        title: [
          {
            required: true,
            message: "请输入标题",
            trigger: "blur",
          },
        ],
        sub_title: [
          {
            required: true,
            message: "请输入副标题",
            trigger: "blur",
          },
        ],
        sort: [
          {
            required: true,
            message: "序号",
            trigger: "blur",
          },
        ],
        tags: [],
        content: [
          {
            required: true,
            message: "请输入内容",
            trigger: "blur",
          },
        ],
      },
      editorOption: {
        placeholder: "请输入内容",
      },
      img_pathAction: `${path}/fileUploadAndDownload/upload`,
    };
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
  mounted() {},
  methods: {
    close() {
      this.$refs["elForm"].resetFields();
      this.$emit("close", null);
    },
    handelConfirm() {
      this.$refs["elForm"].validate((valid) => {
        if (!valid) return;
        // TODO 提交表单
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
    onEditorReady() {}, // 准备编辑器
    onEditorBlur() {}, // 失去焦点事件
    onEditorFocus() {}, // 获得焦点事件
    onEditorChange() {}, // 内容改变事件
  },
};
</script>
<style scoped>
.el-upload__tip {
  line-height: 1.2;
}
</style>
