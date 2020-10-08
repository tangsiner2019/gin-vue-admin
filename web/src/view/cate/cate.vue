<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button
            @click="openDialog"
            plain
            type="primary"
            icon="el-icon-plus"
            >新增</el-button
          >
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text"
                >取消</el-button
              >
              <el-button @click="onDelete" size="mini" type="primary"
                >确定</el-button
              >
            </div>
            <el-button
              plain
              icon="el-icon-delete"
              size="mini"
              slot="reference"
              type="danger"
              >批量删除</el-button
            >
          </el-popover>
        </el-form-item>
      </el-form>
    </div>

    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column
        type="selection"
        width="40"
        align="center"
      ></el-table-column>

      <el-table-column label="类型" width="100" sortable sort-by="type" align="center">
        <template slot-scope="scope">{{
          scope.row.type | formatType
        }}</template>
      </el-table-column>
      <el-table-column label="名称" prop="name" width="200"></el-table-column>
      <el-table-column label="是否禁用" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 0" type="info">正常</el-tag>
          <el-tag v-if="scope.row.status === -1" type="warning">已禁用</el-tag>
        </template>
      </el-table-column>

      <el-table-column
        label="序号"
        prop="sort"
        width="60"
        align="center"
      ></el-table-column>

      <el-table-column label="创建日期" width="160">
        <template slot-scope="scope">{{
          scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>

      <el-table-column label="操作" width="">
        <template slot-scope="scope">
          <el-button
            @click="updateCate(scope.row)"
            size="small"
            icon="el-icon-edit"
            type="primary"
            plain
            >编辑</el-button
          >
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button
                size="mini"
                type="text"
                @click="scope.row.visible = false"
                >取消</el-button
              >
              <el-button
                type="primary"
                size="mini"
                @click="deleteCate(scope.row)"
                >确定</el-button
              >
            </div>
            <el-button
              plain
              type="danger"
              icon="el-icon-delete"
              size="mini"
              slot="reference"
              style="margin-left: 10px"
              >删除</el-button
            >
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog
      :before-close="closeDialog"
      :visible.sync="dialogFormVisible"
      :title="type === 'create' ? '新增分类' : '编辑分类'"
    >
      <add-item
        v-if="dialogFormVisible"
        @close="closeDialog"
        @submit="enterDialog"
        :item="formData"
      ></add-item>
    </el-dialog>
  </div>
</template>

<script>
import {
  createCate,
  deleteCate,
  deleteCateByIds,
  updateCate,
  findCate,
  getCateList,
} from "@/api/cate"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import AddItem from "./add";
export default {
  name: "Cate",
  mixins: [infoList],
  components: {
    AddItem,
  },
  data() {
    return {
      listApi: getCateList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name: null,
        sort: null,
        status: null,
        type: null,
      },
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
    formatType: function (str) {
      return str === "news" ? "资讯" : "商品";
    },
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    async onDelete() {
      const ids = [];
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.ID);
        });
      if (ids.length < 1) {
        this.$message({
          type: "warning",
          message: "请选择待删除的项",
        });
        return;
      }
      const res = await deleteCateByIds({ ids });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    async updateCate(row) {
      const res = await findCate({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.recate;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        name: null,
        sort: null,
        status: null,
        type: null,
      };
    },
    async deleteCate(row) {
      this.visible = false;
      const res = await deleteCate({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        this.getTableData();
      }
    },
    async enterDialog(item) {
      let res;
      this.formData.name = item.name;
      this.formData.type = item.type;
      this.formData.sort = item.sort;
      this.formData.status = item.status;
      switch (this.type) {
        case "create":
          res = await createCate(this.formData);
          break;
        case "update":
          this.formData.ID = item.id;
          res = await updateCate(this.formData);
          break;
        default:
          res = await createCate(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功",
        });
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
  },
  async created() {
    await this.getTableData();
  },
};
</script>

<style>
</style>