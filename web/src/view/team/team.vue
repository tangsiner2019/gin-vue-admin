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
      <el-table-column label="照片" width="100">
        <template slot-scope="scope">
          <CustomPic picType="file" :picSrc="scope.row.imgPath" />
        </template>
      </el-table-column>

      <el-table-column label="姓名" prop="name" width="300"></el-table-column>

      <el-table-column
        label="英文名"
        prop="enName"
        width=""
      ></el-table-column>

      <el-table-column label="是否发布" prop="isPub" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.isPub === 1" type="success">已发布</el-tag>
          <el-tag v-if="scope.row.isPub === 0" type="info">待发布</el-tag>
        </template>
      </el-table-column>
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

      <el-table-column label="操作" width="200">
        <template slot-scope="scope">
          <el-button
            @click="updateTeam(scope.row)"
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
                @click="deleteTeam(scope.row)"
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
      :title="type === 'create' ? '新增团队成员' : '编辑团队成员'"
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
  createTeam,
  deleteTeam,
  deleteTeamByIds,
  updateTeam,
  findTeam,
  getTeamList,
} from "@/api/team"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import CustomPic from "@/components/customPic";
import AddItem from "./add";

export default {
  name: "Team",
  mixins: [infoList],
  components: {
    AddItem,
    CustomPic,
  },
  data() {
    return {
      listApi: getTeamList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name: null,
        imgPath: null,
        enName: null,
        content: null,
        status: null,
        sort: null,
        isPub: null,
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
      const res = await deleteTeamByIds({ ids });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    async updateTeam(row) {
      const res = await findTeam({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reteam;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        name: null,
        imgPath: null,
        enName: null,
        content: null,
        status: null,
        sort: null,
        isPub: null,
      };
    },
    async deleteTeam(row) {
      this.visible = false;
      const res = await deleteTeam({ ID: row.ID });
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
      this.formData.imgPath = item.img_path;
      this.formData.enName = item.en_name;
      this.formData.status = item.status;
      this.formData.content = item.content;
      this.formData.isPub = item.is_pub;
      this.formData.sort = item.sort;
      switch (this.type) {
        case "create":
          res = await createTeam(this.formData);
          break;
        case "update":
          this.formData.ID = item.id;
          res = await updateTeam(this.formData);
          break;
        default:
          res = await createTeam(this.formData);
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