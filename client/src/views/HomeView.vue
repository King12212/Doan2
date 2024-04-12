<script setup>
import TheWelcome from "../components/TheWelcome.vue";
import NavBar from "../components/NavBar.vue";
</script>

<template>
  <main>
    <NavBar id="navbar" />
    <!-- <button class="logout" @click="logout">Log out</button> -->
    <div>Đây là trang chủ</div>
  </main>
</template>
<script>
export default {
  name: "Home",
  setup() {},
  mounted() {
    if (localStorage.getItem("user") !== "true") {
      this.$router.push("/login");
    }
  },
  methods: {
    logout() {
      localStorage.setItem("user", "false");
      this.$router.push("/login");
    },
  },
};
</script>
<style scoped>
#navbar {
  /* margin-left: -150px; */
  position: fixed;
  top: 0;
  left: 0;
}
</style>
