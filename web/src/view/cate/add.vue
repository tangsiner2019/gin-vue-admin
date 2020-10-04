<template>
  <div>
    <el-dialog v-bind="$attrs" v-on="$listeners" @open="onOpen" @close="onClose" title="Dialog Titile">
      <el-row :gutter="15">
        <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
          <el-col :span="12">
            <el-form-item label="分类" prop="field107">
              <el-select v-model="formData.field107" placeholder="请选择分类" clearable :style="{width: '100%'}">
                <el-option v-for="(item, index) in field107Options" :key="index" :label="item.label"
                  :value="item.value" :disabled="item.disabled"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="名称" prop="name">
              <el-input v-model="formData.name" placeholder="请输入名称" :maxlength="20" show-word-limit clearable
                :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="是否禁用" prop="status" required>
              <el-switch v-model="formData.status" :active-value='-1' :inactive-value='0'></el-switch>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排序" prop="sort">
              <el-input-number v-model="formData.sort" placeholder="序号"></el-input-number>
            </el-form-item>
          </el-col>
        </el-form>
      </el-row>
      <div slot="footer">
        <el-button @click="close">取消</el-button>
        <el-button type="primary" @click="handelConfirm">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
export default {
  inheritAttrs: false,
  components: {},
  props: [],
  data() {
    return {
      formData: {
        field107: undefined,
        name: '',
        status: 0,
        sort: 0,
      },
      rules: {
        field107: [{
          required: true,
          message: '请选择分类',
          trigger: 'change'
        }],
        name: [{
          required: true,
          message: '请输入名称',
          trigger: 'blur'
        }],
        sort: [{
          required: true,
          message: '序号',
          trigger: 'blur'
        }],
      },
      field107Options: [{
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
    onOpen() {},
    onClose() {
      this.$refs['elForm'].resetFields()
    },
    close() {
      this.$emit('update:visible', false)
    },
    handelConfirm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        this.close()
      })
    },
  }
}

</script>
<style>
</style>
