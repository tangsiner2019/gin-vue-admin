<template>
  <div>
    <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
      <el-form-item label="网站名称" prop="title">
        <el-input v-model="formData.title" placeholder="请输入网站名称" :maxlength="50" show-word-limit clearable
          :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="网站域名" prop="domain">
        <el-input v-model="formData.domain" placeholder="请输入网站域名" show-word-limit clearable
          :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="备案号" prop="beian">
        <el-input v-model="formData.beian" placeholder="请输入备案号" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="公司名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入公司名称" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="公司地址" prop="address">
        <el-input v-model="formData.address" placeholder="请输入公司地址" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="邮编" prop="postcode">
        <el-input v-model="formData.postcode" placeholder="请输入邮编" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="联系电话" prop="phone">
        <el-input v-model="formData.phone" placeholder="请输入联系电话" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="联系手机号" prop="mobile">
        <el-input v-model="formData.mobile" placeholder="请输入联系手机号" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="传真" prop="fax">
        <el-input v-model="formData.fax" placeholder="请输入传真" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="电子邮箱" prop="email">
        <el-input v-model="formData.email" placeholder="请输入电子邮箱" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="客户QQ" prop="qq">
        <el-input v-model="formData.qq" placeholder="请输入客户QQ" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="网站关键词" prop="keywords">
        <el-input v-model="formData.keywords" placeholder="多个关键词请使用英文逗号（ , ）分隔" clearable
          :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="网站描述" prop="description">
        <el-input v-model="formData.description" type="textarea" placeholder="请输入网站描述"
          :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="关于我们" prop="about">
        <el-input v-model="formData.about" type="textarea" placeholder="请输入关于我们"
          :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="联系我们" prop="conact">
        <el-input v-model="formData.conact" type="textarea" placeholder="请输入联系我们"
          :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  components: {},
  props: [],
  data() {
    return {
      formData: {
        title: '',
        domain: undefined,
        beian: undefined,
        name: undefined,
        address: undefined,
        postcode: undefined,
        phone: undefined,
        mobile: undefined,
        fax: undefined,
        email: undefined,
        qq: undefined,
        keywords: undefined,
        description: undefined,
        about: undefined,
        conact: undefined,
      },
      rules: {
        title: [{
          required: true,
          message: '请输入网站名称',
          trigger: 'blur'
        }],
        domain: [{
          required: true,
          message: '请输入网站域名',
          trigger: 'blur'
        }],
        beian: [{
          required: true,
          message: '请输入备案号',
          trigger: 'blur'
        }],
        name: [{
          required: true,
          message: '请输入公司名称',
          trigger: 'blur'
        }],
        address: [{
          required: true,
          message: '请输入公司地址',
          trigger: 'blur'
        }],
        postcode: [{
          required: true,
          message: '请输入邮编',
          trigger: 'blur'
        }],
        phone: [{
          required: true,
          message: '请输入联系电话',
          trigger: 'blur'
        }],
        mobile: [{
          required: true,
          message: '请输入联系手机号',
          trigger: 'blur'
        }],
        fax: [{
          required: true,
          message: '请输入传真',
          trigger: 'blur'
        }],
        email: [{
          required: true,
          message: '请输入电子邮箱',
          trigger: 'blur'
        }],
        qq: [{
          required: true,
          message: '请输入客户QQ',
          trigger: 'blur'
        }],
        keywords: [{
          required: true,
          message: '多个关键词请使用英文逗号（ , ）分隔',
          trigger: 'blur'
        }],
        description: [{
          required: true,
          message: '请输入网站描述',
          trigger: 'blur'
        }],
        about: [{
          required: true,
          message: '请输入关于我们',
          trigger: 'blur'
        }],
        conact: [{
          required: true,
          message: '请输入联系我们',
          trigger: 'blur'
        }],
      },
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
  }
}

</script>
<style>
</style>
