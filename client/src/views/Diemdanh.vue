<template>
  <main>
    <NavBar class="navbar" />
    <div
      class="body"
      @click="
        (event) => {
          const a = document.getElementsById('more');
          if (event.targer === a) {
            console.log('a');
            return;
          } else {
            console.log(state.choose);
            state.choose = -1;
          }
        }
      "
    >
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
          Lớp DD01
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
          Lớp DD02
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
          Lớp DD03
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
          Lớp DD04
        </div>
      </div>
      <input class="search" placeholder="Tìm kiếm" />
      <span class="material-symbols-outlined" id="search"> search </span>
      <div class="bar">
        <div class="name">Sinh viên</div>
        <div class="id">Mã số</div>
      </div>
      <div class="table">
        <div
          v-for="item in state.student_list"
          :key="item"
          class="student"
          :class="{
            first: state.student_list.indexOf(item) === 0,
            choose: state.pop === state.student_list.indexOf(item),
          }"
        >
          <div class="ava"></div>
          <div class="info">
            <div class="ten">
              <div class="fname">{{ item.firstname }}</div>
              <div class="lname">{{ item.lastname }}</div>
            </div>
            <div class="mail">huy.nguyen123@hcmut.edu.vn</div>
          </div>
          <div class="stid">{{ item.studentid }}</div>
          <div class="last">
            <span
              class="material-symbols-outlined"
              id="more"
              @click="choose(state.student_list.indexOf(item))"
            >
              more_vert
            </span>
            <table-more-vue
              class="tablemore"
              v-if="state.pop === state.student_list.indexOf(item)"
              v-on:view="handleView(item)"
            />
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
<script>
import NavBar from "../components/NavBar.vue";
import { reactive } from "vue";
import TableMoreVue from "../components/TableMore.vue";
import axios from "axios";
export default {
  components: { NavBar, TableMoreVue },
  setup() {
    const state = reactive({
      student_list: [{ firstname: "", lastname: "", studentid: "" }],
      class: 0,
      pop: -1,
    });
    const choose = (i) => {
      if (state.pop === i) {
        state.pop = -1;
        return;
      }
      state.pop = i;
    };
    const listener = (e) => {
      if (e.target.id !== "more" && state.pop !== -1) {
        state.pop = -1;
        // let a = document.querySelector(".drop_down");
        // if (!a.contains(e.target)) {
        //   let b = document.querySelector(".container");
        //   if (b === null || !b.contains(e.target))
        //   state.pop = -1;
        // }
      }
    };
    return { state, choose, listener };
  },
  async created() {
    try {
      axios.defaults.withCredentials = true;
      const response = await axios.get("http://localhost:8080/getAllStudent");
      console.log(response.data.data);
      const lst = response.data.data;
      for (var i = 0; i < 6; i++) {
        var obj = {
          firstname: lst[i].firstname,
          lastname: lst[i].lastname,
          studentid: lst[i].studentID,
        };
        this.state.student_list.push(obj);
      }
      this.state.student_list.splice(0, 1);
    } catch (e) {
      console.log(e);
    }
  },
  mounted() {
    window.addEventListener("click", this.listener);
  },
  beforeUnmount() {
    window.removeEventListener("click", this.listener);
  },
  methods: {
    handleView(item) {
      localStorage.setItem("ten", item.firstname + " " + item.lastname);
      this.$router.push(`/xem/${item.studentid}`);
    },
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
}
#search {
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
.student.first {
  border-top: 5px solid black;
}
.student.choose {
  background-color: rgb(240, 240, 240);
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
  border-radius: 35px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.id {
  margin-left: 55px;
}
.last {
  margin-left: 250px;
  cursor: pointer;
  position: fixed;
  margin-left: 750px;
  margin-top: 500px;
  width: 100px;
  height: 100px;
  z-index: 0;
}
#more {
  position: fixed;
  margin-top: -213px;
  z-index: 20;
  font-size: 30px;
  margin-left: 25px;
}
.tablemore {
  position: fixed;
}
</style>
