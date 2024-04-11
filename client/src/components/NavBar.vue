<template>
  <div class="navbar">
    <div
      class="item"
      v-for="item in state.navbar"
      :key="item.id"
      @mouseover="handleOver(item.id)"
      @mouseleave="handleLeave(item.id)"
    >
      {{ item.title }}
      <div class="list" v-if="item.isHover">
        <div
          class="i"
          v-for="i in item.lst"
          :key="i"
          @click="this.handleClick(i)"
        >
          {{ i }}
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { reactive } from "vue";
export default {
  name: "NavBar",
  setup() {
    const state = reactive({
      navbar: [
        {
          id: 0,
          title: "Trang chủ",
          lst: ["Hồ sơ", "Lịch sử", "Báo lỗi", "Đăng xuất"],
          isHover: false,
        },
        {
          id: 1,
          title: "Quản lý",
          lst: ["Danh sách lớp", "Điểm danh"],
          isHover: false,
        },
        {
          id: 2,
          title: "Nghiên cứu",
          lst: ["Báo lỗi", "API", "Tài liệu"],
          isHover: false,
        },
        {
          id: 3,
          title: "Về chúng tôi",
          lst: ["Lịch sử", "Thành viên"],
          isHover: false,
        },
      ],
    });
    const handleOver = (id) => {
      state.navbar.forEach((i) => {
        if (i.id !== id) i.isHover = false;
      });
      state.navbar[id].isHover = true;
    };
    const handleLeave = (id) => {
      state.navbar[id].isHover = false;
    };
    return { state, handleOver, handleLeave };
  },
  methods: {
    handleClick(i) {
      if (i !== "Đăng xuất") return;
      localStorage.setItem("user", "false");
      this.$router.push("/login");
    },
  },
};
</script>
<style scoped>
.navbar {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  width: 100vw;
  height: 100px;
  /* border: 1px solid black; */
  background-color: aquamarine;
}
.item {
  cursor: pointer;
  height: 100px;
  width: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.item:hover {
  background-color: aqua;
  color: brown;
}
.list {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  position: fixed;
  top: 100px;
  flex-direction: column;
  width: 200px;
}
.i {
  height: 50px;
  width: 200px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-bottom: 1px solid lightblue;
  border-left: 1px solid lightblue;
  border-right: 1px solid lightblue;
}
</style>