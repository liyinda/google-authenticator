(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0e4b0c"],{"90fe":function(e,t,i){"use strict";i.r(t);var l=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",{staticClass:"app-container"},[i("div",{staticClass:"line"}),i("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],attrs:{data:e.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[i("el-table-column",{attrs:{align:"center",label:"ID",width:"50"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.$index)+" ")]}}])}),i("el-table-column",{attrs:{label:"User Name",width:"300",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.Name)+" ")]}}])}),i("el-table-column",{attrs:{label:"Description",width:"400",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.Description)+" ")]}}])}),i("el-table-column",{attrs:{align:"center",prop:"created_at",label:"CreatedTime",width:"200"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("i",{staticClass:"el-icon-time"}),i("span",[e._v(e._s(t.row.CreatedTime))])]}}])}),i("el-table-column",{attrs:{label:"2FA QrCode",width:"300",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(i){return e.googleView(t.row.ID)}}},[e._v(" Show Qrcode ")]),i("el-dialog",{attrs:{title:"qrcode",visible:e.googledialogFormVisible},on:{"update:visible":function(t){e.googledialogFormVisible=t}}},[i("img",{staticStyle:{width:"60%"},attrs:{src:e.img_src}}),i("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[i("el-button",{on:{click:function(t){e.googledialogFormVisible=!1}}},[e._v(" Cancel ")])],1)])]}}])}),i("el-table-column",{attrs:{align:"center",prop:"delete",label:"Delete",width:"130"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(i){return e.delQrcode(t.row.ID)}}},[e._v(" deleted ")])]}}])})],1)],1)},o=[],n=i("b775");function a(e){return Object(n["a"])({url:"/home/list",method:"get"})}function r(e){return Object(n["a"])({url:"/home/qrcode",method:"get",params:{id:e}})}function s(e){return Object(n["a"])({url:"/home/delete",method:"post",params:{id:e}})}i("5c96");var c={filters:{statusFilter:function(e){var t={published:"success",draft:"gray",deleted:"danger"};return t[e]}},data:function(){return{list:null,listLoading:!0,visible:!1,id:"",googledialogFormVisible:!1,img_src:""}},created:function(){this.fetchData()},methods:{fetchData:function(){var e=this;this.listLoading=!0,a().then((function(t){e.list=t.data.items,e.listLoading=!1}))},googleView:function(e){var t=this;this.googledialogFormVisible=!0,console.log(this.img_src),r(e).then((function(e){t.img_src=e.data}))},delQrcode:function(e){var t=this;s(e).then((function(e){2e4===e.code&&(t.$message({message:"delete project successed!"}),location.reload())}))}}},d=c,u=i("2877"),g=Object(u["a"])(d,l,o,!1,null,null,null);t["default"]=g.exports}}]);