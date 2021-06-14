<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal">
          <el-menu-item index="1">服务</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="18">
            <el-input v-model="input" placeholder="请输入内容"></el-input>
          </el-col>
          <el-col :span="6">
            <el-button type="primary">重新载入</el-button>
          </el-col>
        </el-row>
        <el-table style="width: 100%"
          :data="tableData"
          stripe>
          <el-table-column
            prop="id"
            label="ID"
            width="180">
          </el-table-column>
          <el-table-column
            prop="name"
            label="名字"
            width="180">
          </el-table-column>
          <el-table-column
            prop="endpoints"
            label="地址">
          </el-table-column>
          <el-table-column
            prop="update_at"
            label="最后更新时间">
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>
  </div>
</template>

<script>

import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      input: '',
      activeIndex: '',
      tableData: []
    }
  },
  created() {
    this.fetchService()
  },
  methods: {
    fetchService: function() {
      let mythis = this
      axios.get("http://127.0.0.1:8081/api/service").then(function(response){
        mythis.tableData = response.data.data
      }).catch(function(err) {
        console.log(err)
      })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  width: 70%;
  margin-left: 15%;
}
</style>
