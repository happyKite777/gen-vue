<template>
    <div class="app-container">
        <div class="filter-container">
            <!--FILTER-->
            <el-button v-waves class="filter-item" type="primary" size="mini" icon="el-icon-search" @click="handleFilter">搜索</el-button>
        </div>

        <el-table
                v-loading="listLoading"
                :data="list"
                :header-cell-style="{background:'#f4f4f5',color:'#606266'}"
                fit
                highlight-current-row
                height="648">
            <el-table-column
                    v-for="item in colNames"
                    :key="item.prop"
                    :label="item.label"
                    :prop="item.prop"
                    align="center"/>
            <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                    <el-button plain type="primary" size="mini" @click="detail(scope.row)">
                        详情
                    </el-button>
                </template>
            </el-table-column>
        </el-table>

        <pagination
                v-show="total > 0 "
                :total="total"
                :page.sync="listQuery.page"
                :limit.sync="listQuery.page_num"
                @pagination="getList"/>

    </div>
</template>

<script>
    import {dataList} from '@/api/dispatch-system/xxx'
    import waves from '@/directive/waves'
    import Pagination from '@/components/Pagination'

    export default {
        name: '<!--NAME-->',
        components: {Pagination},
        directives: {waves},
        filters: {},
        data() {
            const colNames = [
                <!--COLUMNS-->
            ]
            return {
                colNames,
                list: [],
                total: 0,
                listLoading: true,
                listQuery: {
                    page: 1,
                    page_num: 20,
                    <!--PARAM-->
                }
            }
        },
        created() {
            this.getList()
        },
        methods: {
            getList() {
                this.listLoading = true
                dataList(this.listQuery).then(response => {
                    this.list = response.data.list
                    this.total = response.data.total_num
                    this.listLoading = false
                })
            },
            handleFilter() {
                this.listQuery.page = 1
                this.getList()
            },
            detail() {
                // TODO
            }
        }
    }
</script>

<style scoped>
</style>
