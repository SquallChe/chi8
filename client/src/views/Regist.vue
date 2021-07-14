<template>
  <v-container>
    <v-card class="pa-4" flat>
      <v-text-field label="邮件地址" v-model="mail"></v-text-field>
      <v-row justify="end">
        <v-btn class="mr-4 primary text-subtitle-1" @click="sendVerificationMail" rounded>
          <v-icon class="mr-1">mdi-email</v-icon>
          发送验证邮件
        </v-btn>
      </v-row>
      <v-text-field label="请填写邮件中的验证码" v-model="verificationCode"></v-text-field>
      <v-row justify="end">
        <v-btn class="mr-4 success text-subtitle-1" rounded @click="regist">
          <v-icon class="mr-1">mdi-account</v-icon>
          注册
        </v-btn>
      </v-row>
    </v-card>

    <v-overlay :value="overlay">
      <v-progress-circular indeterminate size="64" color="amber">
      </v-progress-circular>
    </v-overlay>
    <!-- <Snackbar ref="snackbar" message="验证邮件已发送，请前往邮箱查看。" timeout=-1></Snackbar> -->
    <v-snackbar v-model="snackbar" timeout="4000" multi-line="true" color="success" outlined centered>
      <v-row justify="center" class="title">
        <v-icon class="mr-1" color="success">mdi-checkbox-marked-circle</v-icon>
        验证邮件已发送,请前往邮箱查看
      </v-row>
    </v-snackbar>
    <v-snackbar v-model="snackbarError1" timeout="2000" multi-line="true" outlined color="red lighten-1"  centered>
      <v-row justify="center" class="title">
        <v-icon class="mr-1" color="red lighten-1">mdi-alert</v-icon>
        请输入邮箱地址
      </v-row>
    </v-snackbar>
    <v-snackbar v-model="snackbarError2" timeout="2000" multi-line="true" outlined color="red lighten-1" centered>
      <v-row justify="center" class="title">
        <v-icon class="mr-1" color="red lighten-1">mdi-alert</v-icon>
        请输入验证码
      </v-row>
    </v-snackbar>
  </v-container>
</template>

<script>
//import Snackbar from "@/components/Snackbar"

export default {
  data() {
    return {
      mail: '',
      verificationCode: '',
      overlay: false,
      snackbar: false,
      snackbarError1: false,
      snackbarError2: false
    }
  },
  methods: {
    sendVerificationMail() {
      if(this.mail == '') {
        this.snackbarError1 = true
        return
      }
      
      this.overlay = true
      this.$http.get('/send-verification-mail', {
        params: {
          mail: this.mail
        }
      })
      .then(response => {
        if(response.data.result == 'true') {
          this.overlay = false
          this.snackbar = true
        }
      })
      .catch(error => {
        window.console.log(error)
        this.overlay = false
      })
    },
    regist() {
      if(this.verificationCode == '') {
        this.snackbarError2 = true
        return
      }
    }
  },
  components: {
    //Snackbar
  }
}
</script>

<style scoped lang="scss">
</style>
