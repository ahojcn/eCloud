<template>
  <div>
    <Layout>
      <Header>
        <Menu mode="horizontal" theme="dark" active-name="1">
          <div class="layout-logo"></div>
          <div class="layout-nav">
            <MenuItem name="1">
              <Icon type="ios-eye" size="large"></Icon>
              统一识别
            </MenuItem>
          </div>
        </Menu>
      </Header>
      <Content :style="{padding: '0 50px', minHeight: '100vh'}">
        <Tabs type="card">
          <TabPane label="多图识别">
            <Row :gutter="18">
              <Col :span="4">
                <Card title="上传图片">
                  <Button slot="extra" type="primary" @click="onCaiCai">全部识别</Button>
                  分类算法：
                  <Select v-model="selected23">
                    <Option v-for="item in select23" :value="item.value" :key="item.value">{{ item.label }}</Option>
                  </Select>
                  分割算法：
                  <Select v-model="selectedfenge">
                    <Option v-for="item in selectefenge" :value="item.value" :key="item.value">{{ item.label }}</Option>
                  </Select>
                  <Divider></Divider>
                  <div class="demo-upload-list" v-for="(item, index) in uploadList" :key="index">
                    <template v-if="item.status === 'finished'">
                      <img :src="item.url" style="width: 100px">
                      <div class="demo-upload-list-cover">
                        <Icon type="ios-eye-outline" @click.native="handleView(item.url)"></Icon>
                      </div>
                    </template>
                    <template v-else>
                      <Progress v-if="item.showProgress" :percent="item.percentage" hide-info></Progress>
                    </template>
                  </div>
                  <Upload
                      ref="upload"
                      :show-upload-list="false"
                      :default-file-list="defaultList"
                      :on-success="onMoreUploadSucceed"
                      :format="['jpg','jpeg','png']"
                      :max-size="2048"
                      multiple
                      type="drag"
                      :action="baseURL + 'image_upload'"
                      style="display: inline-block;width:58px;">
                    <div style="width: 58px;height:58px;line-height: 58px;">
                      <Icon type="ios-camera" size="20"></Icon>
                    </div>
                  </Upload>
                </Card>
              </Col>
              <Col :span="20">
                <Card title="识别结果">
                  <Table :columns="columns1" :data="recResults">
                    <template slot-scope="{ row }" slot="image">
                      <div v-if="selected23 === 2">
                        二分类：
                        <Tag :color="row.erfenlei.status">
                          {{ row.erfenlei.msg }}
                        </Tag>
                      </div>
                      <div v-else>
                        三分类：
                        <Tag :color="row.sanfenlei.status">
                          {{ row.sanfenlei.msg }}
                        </Tag>
                      </div>
                    </template>
                    <template slot="image_view" slot-scope="{ row }">
                      <img :src="baseURL + 'static/upload_image/' + row.image_name" style="width: 100px">
                      <Button icon="ios-eye-outline" type="text"
                              @click.native="handleView(baseURL + 'static/upload_image/' + row.image_name)"
                      ></Button>
                    </template>
                    <template slot="time">
                      {{ new Date().toLocaleString() }}
                    </template>
                  </Table>
                </Card>
              </Col>
            </Row>
          </TabPane>
          <TabPane label="统一识别">
            <Breadcrumb :style="{margin: '20px 0'}">
              <BreadcrumbItem>上传图像进行识别</BreadcrumbItem>
            </Breadcrumb>
            <Upload
                multiple
                type="drag"
                :on-success="onUploadSucceed"
                :on-error="onUploadError"
                :before-upload="beforeUpload"
                action="http://192.168.0.115:5000/cai">
              <div style="padding: 20px 0">
                <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                <p>拖拽或者点击此处上传</p>
              </div>
            </Upload>
            <Card>
              <div style="min-height: 100vh;" class="demo-split">
                <Split v-model="split1">
                  <div slot="left" class="demo-split-pane">
                    <img v-if="formData.imgSrc !== ''" :src="formData.imgSrc" style="max-width: 100%" alt="需要识别的图片">
                    <h2 v-else>图片展示位</h2>
                  </div>
                  <div slot="right" style="text-align: center">
                    <Alert name="1">
                      二分类识别结果
                      <p slot="desc">
                        <Tag :color="respData.erfenlei.status">{{ respData.erfenlei.msg }}</Tag>
                      </p>
                    </Alert>
                    <Alert name="2">
                      三分类识别结果
                      <p slot="desc">
                        <Tag :color="respData.sanfenlei.status">{{ respData.sanfenlei.msg }}</Tag>
                      </p>
                    </Alert>
                    <Alert type="error">
                      识别出错信息
                      <template slot="desc">
                        {{ respData.err_msg }}
                      </template>
                    </Alert>
                  </div>
                </Split>
              </div>
            </Card>
          </TabPane>
        </Tabs>
      </Content>
      <Footer class="layout-footer-center">202106 &copy; CaiZiGui</Footer>
    </Layout>

    <Modal title="View Image" v-model="visible">
      <img :src="imgName" v-if="visible" style="width: 100%">
    </Modal>
  </div>
</template>

<script>
import requests from "@/util/requests";

export default {
  name: "Cai",
  data() {
    return {
      baseURL: 'http://127.0.0.1:5000/',
      split1: 0.5,
      formData: {
        imgSrc: ''
      },
      imgName: '',
      visible: false,
      select23: [
        {value: 2, label: '二分类'}, {value: 3, label: '三分类'},
      ],
      selected23: '',
      defaultList: [],
      uploadList: [],
      recResults: [],
      respData: {
        err_msg: '未检测到图片，请先上传',
        erfenlei: {
          msg: '未检测到图片',
          status: 'blue',
        },
        sanfenlei: {
          msg: '未检测到图片',
          status: 'blue',
        },
      },
      columns1: [
        {title: '图片名', key: 'image_name'},
        {title: '图片', slot: 'image_view', align: 'center'},
        {title: '概览', slot: 'image', align: 'center'},
        {title: '识别时间', slot: 'time', align: 'center'},
      ],
      selectedfenge: '',
      selectefenge: [
        {value: 2, label: '分割算法-1'}, {value: 3, label: '分割算法-2'},
      ],
    }
  },
  methods: {
    handleView(name) {
      this.imgName = name;
      this.visible = true;
    },
    onMoreUploadSucceed(response, file, fileList) {
      console.log(response, file, fileList)
      file.url = this.baseURL + 'static/upload_image/' + response
      this.uploadList.push(file)
    },
    onUploadSucceed(response, file, fileList) {
      this.$Spin.hide()
      console.log(response, file, fileList)
      this.respData = response
    },
    onUploadError(response, file, fileList) {
      this.$Spin.hide()
      console.log(response, file, fileList)
    },
    beforeUpload(file) {
      this.$Spin.show()
      const self = this
      const reader = new FileReader()
      reader.readAsArrayBuffer(file)
      reader.onload = function (e) {
        e;
        const bf = this.result
        const blob = new Blob([bf], {type: 'text/plain'})
        const str = URL.createObjectURL(blob)
        self.formData.imgSrc = str
      }
    },
    onCaiCai() {
      requests({
        url: this.baseURL + 'caicai',
        data: {'images': this.uploadList, 'selected': this.selected23},
        method: 'POST',
      }).then(res => {
        this.recResults = res
      })
    },
  }
}
</script>

<style scoped>
.demo-upload-list {
  display: inline-block;
  width: 60px;
  height: 60px;
  text-align: center;
  line-height: 60px;
  border: 1px solid transparent;
  border-radius: 4px;
  overflow: hidden;
  background: #fff;
  position: relative;
  box-shadow: 0 1px 1px rgba(0, 0, 0, .2);
  margin-right: 4px;
}

.demo-upload-list img {
  width: 100%;
  height: 100%;
}

.demo-upload-list-cover {
  display: none;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, .6);
}

.demo-upload-list:hover .demo-upload-list-cover {
  display: block;
}

.demo-upload-list-cover i {
  color: #fff;
  font-size: 20px;
  cursor: pointer;
  margin: 0 2px;
}
</style>