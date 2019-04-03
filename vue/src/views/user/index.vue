<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">{{ $t('table.add') }}</el-button>
    </div>
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="userName">
        <template slot-scope="scope">
          <span>{{ scope.row.user_name }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="realName">
        <template slot-scope="scope">
          <span>{{ scope.row.real_name }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="Phone">
        <template slot-scope="scope">
          <span>{{ scope.row.phone }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="Email">
        <template slot-scope="scope">
          <span>{{ scope.row.email }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="CreatTime">
        <template slot-scope="scope">
          <span>{{ scope.row.create_time }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="Actions" width="280">
        <template slot-scope="scope">
          <el-button type="primary" size="small" icon="el-icon-edit" @click="handleUpdate(scope.row)"></el-button>
          <el-button type="primary" size="small" icon="eye-open" @click="googleView(scope.row)">Code</el-button>
          <el-button type="primary" size="small" icon="el-icon-delete" @click="delUser(scope.row)"></el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="userRules" :model="temp" label-position="left" label-width="90px" style="width: 400px; margin-left:50px;">
        <el-form-item label="userName" prop="user_name">
          <el-input v-model="temp.user_name"/>
        </el-form-item>
        <el-form-item label="realName" prop="realName">
          <el-input v-model="temp.real_name"/>
        </el-form-item>
        <el-form-item label="Phone" prop="Phone">
          <el-input v-model="temp.phone"/>
        </el-form-item>
        <el-form-item label="Email" prop="Email">
          <el-input v-model="temp.email"/>
        </el-form-item>
        <el-form-item label="Password" prop="Password">
          <el-input v-model="temp.password"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button type="primary" @click="dialogStatus==='create'?addData():updateData()">{{ $t('table.confirm') }}</el-button>
      </div>
    </el-dialog>
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="googledialogFormVisible">
      <img :src="img_src" style="width:100%">
      <div slot="footer" class="dialog-footer">
        <el-button @click="googledialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {} from '@/utils/validate'
import { getUserList, updateUser, addUser, getGoogleCode, delUser } from '@/api/user'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import waves from '@/directive/waves' // Waves directive

export default {
  name: 'UserList',
  components: { Pagination },
  directives: { waves },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    const validateRequire = (rule, value, callback) => {
      if (value === '') {
        this.$message({
          message: rule.field + '为必传项',
          type: 'error'
        })
        callback(new Error(rule.field + '为必传项'))
      } else {
        callback()
      }
    }
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      importanceOptions: [1, 2, 3],
      sortOptions: [{ label: 'ID Ascending', key: '+id' }, { label: 'ID Descending', key: '-id' }],
      statusOptions: ['published', 'draft', 'deleted'],
      showReviewer: false,
      temp: {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        type: '',
        status: 'published'
      },
      dialogFormVisible: false,
      googledialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogPvVisible: false,
      pvData: [],
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      userRules: {
        user_name: [{ required: true, trigger: 'blur', validator: validateRequire }]
      },
      downloadLoading: false,
      googlepic: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif'
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getUserList(this.listQuery).then(response => {
        this.list = response.data.users
        this.total = response.data.count
        this.listLoading = false
      })
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        status: 'published',
        type: ''
      }
    },
    googleView(row) {
      this.resetTemp()
      this.dialogStatus = 'qCode'
      this.googledialogFormVisible = true
      const tempData = Object.assign({}, row)
      getGoogleCode(tempData.id).then(response => {
        this.googledialogFormVisible = false
        this.img_src = tempData.qrcode
        this.googledialogFormVisible = true
      })
    },
    delUser(row) {
      const tempData = Object.assign({}, row)
      delUser(tempData).then(response => {
        this.$notify({
          title: '成功',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    addData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          tempData.timestamp = +new Date(tempData.timestamp)
          addUser(tempData).then(() => {
            for (const v of this.list) {
              if (v.id === this.temp.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.temp)
                break
              }
            }
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row) // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
          updateUser(tempData).then(() => {
            for (const v of this.list) {
              if (v.id === this.temp.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.temp)
                break
              }
            }
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getList()
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.getList()
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
