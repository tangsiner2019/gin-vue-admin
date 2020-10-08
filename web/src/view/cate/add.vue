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
        <el-col :span="12">
          <el-form-item label="分类" prop="type">
            <el-select
              v-model="formData.type"
              placeholder="请选择分类"
              clearable
              :style="{ width: '100%' }"
            >
              <el-option
                v-for="(item, index) in types"
                :key="index"
                :label="item.label"
                :value="item.value"
                :disabled="item.disabled"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="名称" prop="name">
            <el-input
              v-model="formData.name"
              placeholder="请输入名称"
              :maxlength="20"
              show-word-limit
              clearable
              :style="{ width: '100%' }"
            ></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否禁用" prop="status" required>
            <el-switch
              v-model="formData.status"
              :active-value="-1"
              :inactive-value="0"
            ></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="排序" prop="sort">
            <el-input-number
              v-model="formData.sort"
              placeholder="序号"
            ></el-input-number>
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
export default {
  inheritAttrs: false,
  components: {},
  props: {
    item: Object,
  },
  data() {
    return {
      formData: {
        type: "",
        name: "",
        status: 0,
        sort: 0,
      },
      rules: {
        type: [
          {
            required: true,
            message: "请选择分类",
            trigger: "change",
          },
        ],
        name: [
          {
            required: true,
            message: "请输入名称",
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
      },
      types: [
        {
          label: "商品",
          value: "product",
        },
        {
          label: "资讯",
          value: "news",
        },
      ],
    };
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {
    if (this.item.ID > 0) {
      this.formData = {
        id: this.item.ID,
        name: this.item.name,
        type: this.item.type,
        sort: this.item.sort,
        status: this.item.status,
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
