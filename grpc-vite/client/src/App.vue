<template>
  <el-container class="container">
    <el-row class="main" justify="center" align="middle">
      <el-row class="word" justify="center">Hello Kratos</el-row>
      <el-space>
        <el-input class="input" placeholder="输入你的名字" v-model="input"></el-input>
        <el-button type="primary" @click="send(input)">发送</el-button>
      </el-space>
      <el-row class="response">{{ response }}</el-row>
    </el-row>
  </el-container>
</template>

<script setup>
import {ref} from "vue"
import {GrpcWebFetchTransport} from "@protobuf-ts/grpcweb-transport"
import {HelloRequest} from "../proto/greeter.ts"
import {GreeterClient} from "../proto/greeter.client.ts"

const input = ref("")
const response = ref("")
const transport = new GrpcWebFetchTransport({
  baseUrl: "http://localhost:8080"
});
const client = new GreeterClient(transport)

function send(input) {
  client.sayHello(HelloRequest.create({name: input})).then(function (reply) {
    response.value = reply.response.message
    console.log(reply)
  })
}
</script>

<style setup lang=scss>
.container {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;

  .main {
    width: 100%;
    height: 100%;
    flex-direction: column;

    .word {
      font-size: 80px;
      margin-bottom: 50px;
    }

    .input {
      width: 300px;
    }

    .response {
      margin-top: 50px;
    }
  }
}
</style>
