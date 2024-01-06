<template>
  <div>
    <!-- <div>
      <h1>Hello, world!</h1>
    </div> -->
    <div>
      <h3>add todo list</h3>
    </div>
    <v-btn color="info" @click="submit()">submit!</v-btn>
    <v-row>
      <v-col cols="12" sm="4">
        Name
        <v-text-field name="件名" label="label" v-model="createTodoForm.name"></v-text-field>
      </v-col>
    </v-row>
  </div>
  </template>
  
  <script lang="ts">
  import Vue from 'vue';
  import { TODOType } from '~/types/TODOType';
  // import { api } from '~/apis/api';
  
  export default Vue.extend({
    data() {
      return {
        todos: [] as TODOType[],
        createTodoForm: {} as TODOType,
        updateTodoType: {}as TODOType,
      };
    },
    methods: {
      // ボタン押下を検知できるかを確認するテストmethod
      submit() {
        console.log("submit!");
        console.log(this.getRecordDate());
      },
      // todoを登録する際の記録日付 YYYY/MM/DD hh:mm:dd を取得するmethod
      getRecordDate(): string {
        const date = new Date();
        date.setDate(date.getDate());
        const year = date.getFullYear();
        const month = this.paddingOfTwoDigit(date.getMonth() + 1);
        const day = this.paddingOfTwoDigit(date.getDate());
        const hour = this.paddingOfTwoDigit(date.getHours());
        const minute = this.paddingOfTwoDigit(date.getMinutes());
        const second = this.paddingOfTwoDigit(date.getSeconds());
        // hh:mm:ddも設定するように回収
        return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
      },
      // 記録日付を取得する際に、要素が2桁になるように0で詰めていくmethod
      paddingOfTwoDigit(num: number): string {
        return num.toString().padStart(2, '0');
      }
      // API叩くテスト用method
      // testWord(word: string) {
      //   api.testAPI(word).then((response: string) => {
      //     console.log(response);
      //   })
      //   .catch((err) => {
      //     console.error(err);
      //   });
      // }
    }
  })
  </script>
  