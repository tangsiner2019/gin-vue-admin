<template>
  <div>
    <el-row :gutter="15">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-col :span="12">
          <el-form-item label="分类" prop="cate_id">
            <el-select v-model="formData.cate_id" placeholder="请选择分类" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in cate_idOptions" :key="index" :label="item.label"
                :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="标题" prop="title">
            <el-input v-model="formData.title" placeholder="请输入标题" :maxlength="50" show-word-limit clearable
              :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="副标题" prop="sub_title">
            <el-input v-model="formData.sub_title" placeholder="请输入副标题" :maxlength="50" show-word-limit
              clearable :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="题图" prop="img_path" required>
            <el-upload ref="img_path" :file-list="img_pathfileList" :action="img_pathAction"
              :before-upload="img_pathBeforeUpload" list-type="picture-card" accept="image/*">
              <i class="el-icon-plus"></i>
              <div slot="tip" class="el-upload__tip">只能上传不超过 200KB 的image/*文件</div>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="是否禁用" prop="status" required>
            <el-switch v-model="formData.status" :active-value='-1' :inactive-value='0'></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="是否发布" prop="field111" required>
            <el-switch v-model="formData.field111"></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="排序" prop="sort">
            <el-input-number v-model="formData.sort" placeholder="序号"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="内容图片" prop="imgs" required>
            <el-upload ref="imgs" :file-list="imgsfileList" :action="imgsAction" multiple
              :before-upload="imgsBeforeUpload" list-type="picture-card" accept="image/*">
              <i class="el-icon-plus"></i>
              <div slot="tip" class="el-upload__tip">只能上传不超过 200KB 的image/*文件</div>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="标签" prop="tags">
            <el-input v-model="formData.tags" placeholder="多个标签请使用英文逗号（ , ）分隔" clearable
              :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="内容" prop="content">
            <el-input v-model="formData.content" type="textarea" placeholder="请输入内容"
              :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item size="large">
            <el-button type="primary" @click="submitForm">提交</el-button>
            <el-button @click="resetForm">重置</el-button>
          </el-form-item>
        </el-col>
      </el-form>
    </el-row>
  </div>
</template>
<script>
export default {
  components: {},
  props: [],
  data() {
    return {
      formData: {
        cate_id: undefined,
        title: '',
        sub_title: undefined,
        img_path: null,
        status: 0,
        field111: true,
        sort: 0,
        imgs: null,
        tags: undefined,
        content: undefined,
      },
      rules: {
        cate_id: [{
          required: true,
          message: '请选择分类',
          trigger: 'change'
        }],
        title: [{
          required: true,
          message: '请输入标题',
          trigger: 'blur'
        }],
        sub_title: [{
          required: true,
          message: '请输入副标题',
          trigger: 'blur'
        }],
        sort: [{
          required: true,
          message: '序号',
          trigger: 'blur'
        }],
        tags: [],
        content: [{
          required: true,
          message: '请输入内容',
          trigger: 'blur'
        }],
      },
      img_pathAction: 'https://jsonplaceholder.typicode.com/posts/',
      img_pathfileList: [],
      imgsAction: 'https://jsonplaceholder.typicode.com/posts/',
      imgsfileList: [],
      cate_idOptions: [{
        "label": "商品",
        "value": "product"
      }, {
        "label": "资讯",
        "value": "news"
      }],
    }
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
    img_pathBeforeUpload(file) {
      let isRightSize = file.size / 1024 < 200
      if (!isRightSize) {
        this.$message.error('文件大小超过 200KB')
      }
      let isAccept = new RegExp('image/*').test(file.type)
      if (!isAccept) {
        this.$message.error('应该选择image/*类型的文件')
      }
      return isRightSize && isAccept
    },
    imgsBeforeUpload(file) {
      let isRightSize = file.size / 1024 < 200
      if (!isRightSize) {
        this.$message.error('文件大小超过 200KB')
      }
      let isAccept = new RegExp('image/*').test(file.type)
      if (!isAccept) {
        this.$message.error('应该选择image/*类型的文件')
      }
      return isRightSize && isAccept
    },
  }
}

</script>
<style>
.el-upload__tip {
  line-height: 1.2;
}

</style>
