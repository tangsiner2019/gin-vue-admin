<template>
  <div>
    <el-row :gutter="15">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-col :span="24">
          <el-form-item label="姓名" prop="name">
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
          <el-form-item label="头像" prop="img_path" required>
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
          <el-form-item label="个人介绍" prop="content">
            <el-input v-model="formData.content" type="textarea" placeholder="请输入个人介绍"
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
        name: '',
        en_name: undefined,
        img_path: null,
        status: 0,
        field111: true,
        sort: 0,
        content: undefined,
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
          message: '请输入个人介绍',
          trigger: 'blur'
        }],
      },
      img_pathAction: 'https://jsonplaceholder.typicode.com/posts/',
      img_pathfileList: [],
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
  }
}

</script>
<style>
.el-upload__tip {
  line-height: 1.2;
}

</style>
