<template>
  <main>
    <NavBar class="navbar" />
    <div class="body">
      <div class="title">Danh sách lớp</div>
      <div class="classlist">
        <div
          class="item"
          :class="{ '-mark': state.class === 0 }"
          @click="
            () => {
              state.class = 0;
            }
          "
        >
          Lớp CC01
        </div>
        <div
          class="item"
          :class="{ '-mark': state.class === 1 }"
          @click="
            () => {
              state.class = 1;
            }
          "
        >
          Lớp CC02
        </div>
        <div
          class="item"
          :class="{ '-mark': state.class === 2 }"
          @click="
            () => {
              state.class = 2;
            }
          "
        >
          Lớp CC03
        </div>
        <div
          class="item"
          :class="{ '-mark': state.class === 3 }"
          @click="
            () => {
              state.class = 3;
            }
          "
        >
          Lớp CC04
        </div>
      </div>
      <input class="search" placeholder="Tìm kiếm" />
      <span class="material-symbols-outlined"> search </span>
      <div class="bar">
        <div class="name">Sinh viên</div>
        <div class="id">Mã số</div>
      </div>
      <div class="table">
        <div v-for="item in state.student_list" :key="item" class="student">
          <div class="ava"></div>
          <div class="info">
            <div class="ten">
              <div class="fname">{{ item.firstname }}</div>
              <div class="lname">{{ item.lastname }}</div>
            </div>
            <div class="mail">huy.nguyen123@hcmut.edu.vn</div>
          </div>
          <div class="stid">{{ item.studentid }}</div>
          <div class="btn"></div>
        </div>
      </div>
    </div>
  </main>
</template>
<script>
import NavBar from "../components/NavBar.vue";
import { reactive } from "vue";
export default {
  components: { NavBar },
  setup() {
    const state = reactive({
      student_list: [
        { firstname: "Le", lastname: "Khoi", studentid: "2053151" },
        { firstname: "Le", lastname: "Huy", studentid: "2053152" },
        { firstname: "Le", lastname: "Minh", studentid: "2053153" },
      ],
      class: 0,
    });
    return { state };
  },
};
</script>
<style scoped>
.navbar {
  /* margin-left: -150px; */
  position: fixed;
  top: 0;
  left: 0;
  z-index: 100;
}
.body {
  position: fixed;
  top: 200px;
  left: 0;
  z-index: 0;
}
.classlist {
  position: fixed;
  top: 200px;
  left: 50px;
  display: flex;
  flex-direction: column;
  width: 100px;
  height: 190px;
  justify-content: space-between;
  align-items: center;
}
.title {
  position: fixed;
  top: 130px;
  left: 50px;
  font-size: 20px;
  font-weight: 700;
}
.item {
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 20px;
  width: 100px;
  padding: 10px;
  color: rgb(72, 61, 139);
  text-align: left;
  cursor: pointer;
  border-radius: 1px;
  font-size: 14px;
}
.-mark {
  color: rgb(247, 84, 25);
  background-color: azure;
}
.item:hover {
  color: rgb(247, 84, 25);
}
.search {
  position: fixed;
  top: 130px;
  left: 400px;
  width: 800px;
  height: 60px;
  padding: 10px;
  font-size: 20px;
  border-radius: 5px;
  border: 1px solid lightgrey;
}
.search:focus {
  outline: none;
}
.material-symbols-outlined {
  font-variation-settings: "FILL" 0, "wght" 600, "GRAD" 40, "opsz" 48;
  font-size: 40px;
  position: fixed;
  top: 140px;
  left: 1150px;
}
.bar {
  position: fixed;
  top: 220px;
  left: 400px;
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 800px;
  font-size: 17px;
  font-weight: 600;
}
.name {
  padding-left: 30px;
}
.id {
  padding-left: 250px;
}
.table {
  position: fixed;
  top: 250px;
  left: 400px;
  display: flex;
  flex-direction: column;
  width: 800px;
  align-items: center;
}
.student {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 800px;
  border-top: 1px solid lightgrey;
  height: 70px;
}
.ava {
  border-radius: 50%;
  width: 60px;
  height: 60px;
  border: 1px solid lightgray;
}
.info {
  padding-left: 15px;
  display: flex;
  flex-direction: column;
  width: 200px;
  font-size: 15px;
  justify-content: space-between;
  /* align-items: center; */
}
.ten {
  display: flex;
  flex-direction: row;
  width: 70px;
  justify-content: space-between;
  align-items: center;
}
.stid {
  margin-left: 95px;
  width: 140px;
  font-size: 16px;
  height: 35px;
  border: 2px solid lightgrey;
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.id {
  margin-left: 55px;
}
</style>
