<script setup>
import TheWelcome from '../components/TheWelcome.vue'
</script>

<template>
  <main>
    <TheWelcome />
    <button class ="logout" @click = "logout">Log out</button>
  </main>
</template>
<script>
export default{
  name:"Home",
  setup(){

  },
  mounted(){
        if (localStorage.getItem('user') !== 'true'){
            this.$router.push('/login');
        }
    },
  methods:{
    logout(){
      localStorage.setItem("user","false")
      this.$router.push('/login');
    }
  }

}
</script>
