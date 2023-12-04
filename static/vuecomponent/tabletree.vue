<template>
    <el-table :data="data" row-key="id" border style="width: 100%" :height="height" :row-style="showTr" size="mini">
        <el-table-column v-for="(column,index) in columns" :key="column.dataIndex" :label="column.text" show-overflow-tooltip>
            <template slot-scope="scope">
                <span v-if="spaceIconShow(index)" v-for="(space, levelIndex) in scope.row._level" class="ms-tree-space"></span>
                    <i v-if="!scope.row._expanded && toggleIconShow(index,scope.row)" class="el-icon-caret-right" aria-hidden="true" @click="toggle(scope.$index)"></i>
                    <i v-if="scope.row._expanded && toggleIconShow(index,scope.row)" class="el-icon-caret-bottom" aria-hidden="true" @click="toggle(scope.$index)"></i>
                <span>
                    {{scope.row[column.dataIndex]}}
                </span>
            </template>
        </el-table-column>
    </el-table>
</template>

<script>
module.exports = {
    name: 'TreeGrid',
    props: {
        // 该属性是确认父组件传过来的数据是否已经是树形结构了，如果是，则不需要进行树形格式化
        treeStructure: {
            type: Boolean,
            default: function () {
                return false
            }
        },
        height: {
            type: Number,
            default: function () {
                return 500
            }
        },
        more: {
            type: Boolean,
            default: function () {
                return false
            }
        },
        // 这是相应的字段展示
        columns: {
            type: Array,
            default: function () {
                return []
            }
        },
        // 这是数据源
        dataSource: {
            type: Array,
            default: function () {
                return []
            }
        },

        // 这个作用是根据自己需求来的，比如在操作中涉及相关按钮编辑，删除等，需要向服务端发送请求，则可以把url传过来
        requestUrl: {
            type: String,
            default: function () {
                return ''
            }
        },
        // 这个是是否展示操作列
        treeType: {
            type: String,
            default: function () {
                return 'normal'
            }
        },
        // 是否默认展开所有树
        defaultExpandAll: {
            type: Boolean,
            default: function () {
                return false
            }
        }
    },
    data() {
        return {}
    },
    computed: {
        // 格式化数据源
        data: function () {
            let me = this

            if (me.treeStructure) {
                let data = tool.treeToArray(me.dataSource, null, null, me.defaultExpandAll)
                return data
            }
            
            return me.dataSource
        }
    },
    methods: {
        // 显示行
        showTr(row, index) {
                let show = (row.row._parent ? (row.row._parent._expanded && row.row._parent._show) : true)
                row.row._show = show
                return show ? '' : 'display:none;'
        },
        // 展开下级
        toggle(trIndex) {
            let me = this
            let record = me.data[trIndex]
            record._expanded = !record._expanded
        },
        // 显示层级关系的空格和图标
        spaceIconShow(index) {
            let me = this
            if (me.treeStructure && index === 0) {
                return true
            }
            return false
        },
        // 点击展开和关闭的时候，图标的切换
        toggleIconShow(index, record) {
            let me = this
            if (me.treeStructure && index === 0 && record.children && record.children.length > 0) {
                return true
            }
            return false
        },
    }
}

</script>

<style>
.ms-tree-space {
	position:relative;
	top:1px;
	display:inline-block;
	font-family:"Glyphicons Halflings";
	font-style:normal;
	font-weight:400;
	line-height:1;
	width:18px;
	height:14px;
}
.ms-tree-space::before {
	content:"";
}
table td {
	line-height:26px;
}
</style>